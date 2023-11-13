package gateways

import (
	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"github.com/fedeveron01/golang-base/cmd/core"
	"github.com/fedeveron01/golang-base/cmd/core/entities"
	_ "github.com/fedeveron01/golang-base/cmd/core/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user gateway_entities.User) (gateway_entities.User, error)
	CreateCompleteUserWithEmployee(user gateway_entities.User, chargeID uint, employee gateway_entities.Employee) (gateway_entities.User, error)
	FindUserById(id int64) (gateway_entities.User, error)
	FindUserByUsernameAndPassword(username string, password string) (gateway_entities.User, error)
	FindUserByUsername(username string) gateway_entities.User
	UpdateUser(user gateway_entities.User) (gateway_entities.User, error)
	DeleteUser(id string) error
}

type UserGatewayImpl struct {
	userRepository UserRepository
}

func NewUserGateway(userRepository UserRepository) *UserGatewayImpl {
	return &UserGatewayImpl{
		userRepository: userRepository,
	}
}

func (i *UserGatewayImpl) CreateCompleteUserWithEmployee(user entities.User, employee entities.Employee) (entities.User, error) {
	userDB := i.ToServiceEntity(user)
	employeeDB := gateway_entities.Employee{
		Name:     employee.Name,
		LastName: employee.LastName,
		DNI:      employee.DNI,
		User: gateway_entities.User{
			UserName: employee.User.UserName,
			Password: employee.User.Password,
		},
	}
	chargeID := employee.Charge.ID

	created, err := i.userRepository.CreateCompleteUserWithEmployee(userDB, chargeID, employeeDB)
	if err != nil {
		return entities.User{}, err
	}
	user = i.ToBusinessEntity(created)
	return user, nil
}

func (i *UserGatewayImpl) CreateUser(user entities.User) (entities.User, error) {
	userDB := i.ToServiceEntity(user)
	created, err := i.userRepository.CreateUser(userDB)
	if err != nil {
		return entities.User{}, err
	}
	user = i.ToBusinessEntity(created)
	return user, nil
}

func (i *UserGatewayImpl) FindUserById(id int64) (entities.User, error) {
	userDB, err := i.userRepository.FindUserById(id)
	if err != nil {
		return entities.User{}, err
	}
	user := i.ToBusinessEntity(userDB)
	return user, nil
}

func (i *UserGatewayImpl) FindUserByUsernameAndPassword(username string, password string) (entities.User, error) {
	userDB, err := i.userRepository.FindUserByUsernameAndPassword(username, password)
	if err != nil {
		return entities.User{}, err
	}
	user := i.ToBusinessEntity(userDB)
	return user, nil
}

func (i *UserGatewayImpl) FindUserByUsername(username string) entities.User {
	userDB := i.userRepository.FindUserByUsername(username)
	user := i.ToBusinessEntity(userDB)
	return user
}

func (i *UserGatewayImpl) UpdateUser(user entities.User) (entities.User, error) {
	userDB := i.ToServiceEntity(user)
	userResponse, err := i.userRepository.UpdateUser(userDB)
	if err != nil {
		return entities.User{}, err
	}
	user = i.ToBusinessEntity(userResponse)
	return user, nil
}

func (i *UserGatewayImpl) DeleteUser(id string) error {
	return i.userRepository.DeleteUser(id)
}

func (i *UserGatewayImpl) ToBusinessEntity(userDB gateway_entities.User) entities.User {
	user := entities.User{
		EntitiesBase: core.EntitiesBase{
			ID: userDB.ID,
		},
		UserName: userDB.UserName,
		Password: userDB.Password,
		Inactive: userDB.Inactive,
	}
	return user
}

func (i *UserGatewayImpl) ToServiceEntity(user entities.User) gateway_entities.User {
	userDB := gateway_entities.User{
		Model: gorm.Model{
			ID: user.ID,
		},
		UserName: user.UserName,
		Password: user.Password,
		Inactive: user.Inactive,
	}
	return userDB
}
