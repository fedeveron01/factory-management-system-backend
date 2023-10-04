package user_usecase

import (
	"time"

	"github.com/fedeveron01/golang-base/cmd/core/entities"
	core_errors "github.com/fedeveron01/golang-base/cmd/core/errors"
	internal_jwt "github.com/fedeveron01/golang-base/cmd/internal/jwt"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase interface {
	ActiveDesactiveUser(id int64, activate bool) error
	CreateUser(user entities.User, employee entities.Employee) error
	UpdateUser(user entities.User) error
	DeleteUser(id string) error
	LoginUser(username string, password string) (string, error)
	LogoutUser(sessionId float64) error
}

type UserGateway interface {
	FindUserById(id int64) (entities.User, error)
	CreateCompleteUserWithEmployee(user entities.User,
		employee entities.Employee) (entities.User, error)
	FindUserByUsernameAndPassword(username string, password string) (entities.User, error)
	FindUserByUsername(username string) entities.User
	UpdateUser(user entities.User) error
	DeleteUser(id string) error
}

type SessionGateway interface {
	CreateSession(session entities.Session) (entities.Session, error)
	FindAll() ([]entities.Session, error)
	UpdateSession(session entities.Session) error
	DeleteSession(id float64) error
}

type EmployeeGateway interface {
	CreateEmployee(employee entities.Employee) error
	FindEmployeeByUserId(id uint) (entities.Employee, error)
	FindAll() ([]entities.Employee, error)
	UpdateEmployee(employee entities.Employee) error
	DeleteEmployee(id string) error
}

type ChargeGateway interface {
	FindByName(name string) (uint, error)
	FindById(id uint) (entities.Charge, error)
	CreateCharge(charge entities.Charge) (entities.Charge, error)
}

type Implementation struct {
	userGateway     UserGateway
	sessionGateway  SessionGateway
	employeeGateway EmployeeGateway
	chargeGateway   ChargeGateway
}

func NewUserUseCase(userGateway UserGateway,
	sessionGateway SessionGateway,
	employeeGateway EmployeeGateway,
	chargeGateway ChargeGateway,
) *Implementation {
	return &Implementation{
		userGateway:     userGateway,
		sessionGateway:  sessionGateway,
		employeeGateway: employeeGateway,
		chargeGateway:   chargeGateway,
	}
}
func (i *Implementation) ActiveDesactiveUser(id int64, activate bool) error {
	user, err := i.userGateway.FindUserById(id)
	if err != nil {
		return err
	}
	user.ID = uint(id)
	user.Inactive = !activate
	err = i.userGateway.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (i *Implementation) CreateUser(user entities.User, employee entities.Employee) error {
	if user.UserName == "" || user.Password == "" {
		return core_errors.ErrUsernameOrPasswordIsEmpty
	}

	if len(user.Password) < 5 {
		return core_errors.ErrPasswordTooShort
	}

	user.Password = encryptPassword(user.Password)

	userRepeated := i.userGateway.FindUserByUsername(user.UserName)
	if userRepeated.ID != 0 {
		return core_errors.ErrUsernameAlreadyExists
	}

	//validate charge
	_, err := i.chargeGateway.FindById(employee.Charge.ID)
	if err != nil {
		return err
	}
	user, err = i.userGateway.CreateCompleteUserWithEmployee(user, employee)

	if err != nil {
		return err
	}

	return nil
}

func (i *Implementation) UpdateUser(user entities.User) error {
	return i.userGateway.UpdateUser(user)
}

func encryptPassword(password string) string {
	encriptedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(encriptedPassword)
}

func isCorrectPassword(encryptedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func (i *Implementation) LoginUser(username string, password string) (string, error) {
	if username == "" || password == "" {
		return "", core_errors.ErrUsernameOrPasswordIsEmpty
	}
	user := i.userGateway.FindUserByUsername(username)
	if user.Inactive {
		return "", core_errors.ErrInactiveUser
	}
	if user.ID == 0 {
		return "", core_errors.ErrUserNotFound
	}
	if !isCorrectPassword(user.Password, password) {
		return "", core_errors.ErrUsernameOrPasswordIsIncorrect
	}
	// create session
	session := entities.Session{
		User: user,
	}
	createdSession, err := i.sessionGateway.CreateSession(session)
	if err != nil {
		return "", err
	}
	// get employee
	employee, err := i.employeeGateway.FindEmployeeByUserId(user.ID)
	if err != nil {
		return "", err
	}
	return generateToken(employee, createdSession)
}

func (i *Implementation) DeleteUser(id string) error {
	return i.userGateway.DeleteUser(id)
}

func (i *Implementation) LogoutUser(sessionId float64) error {

	err := i.sessionGateway.DeleteSession(sessionId)
	if err != nil {
		return err
	}

	return nil
}

func generateToken(employee entities.Employee, session entities.Session) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)
	var role string
	if employee.ID == 0 {
		role = "none"
	} else {
		role = employee.Charge.Name
	}
	t.Claims = &internal_jwt.Claims{
		StandardClaims: &jwt.StandardClaims{

			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
		TokenType:  "level1",
		EmployeeId: float64(employee.ID),
		SessionId:  float64(session.ID),
		Role:       role,
	}

	return t.SignedString([]byte("test"))
}
