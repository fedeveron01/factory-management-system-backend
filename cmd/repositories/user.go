package repositories

import (
	"errors"

	gateway_entities "github.com/fedeveron01/golang-base/cmd/adapters/gateways/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(database *gorm.DB) *UserRepository {
	return &UserRepository{
		db: database,
	}
}

func (r *UserRepository) GetAllUser() ([]gateway_entities.User, error) {
	var users []gateway_entities.User
	res := r.db.Find(&users)
	if res.Error != nil {
		return nil, res.Error
	}
	return users, nil
}

func (r UserRepository) CreateCompleteUserWithEmployee(
	user gateway_entities.User,
	charge gateway_entities.Charge,
	employee gateway_entities.Employee) (gateway_entities.User, error) {
	r.db.Transaction(func(tx *gorm.DB) error {
		tx.Create(&user)
		// find charge
		res := tx.Where("name = ?", charge.Name).Find(&charge)
		if res.Error != nil {
			return res.Error
		}
		if res.RowsAffected == 0 {
			res := tx.Create(&charge)
			if res.Error != nil {
				return res.Error
			}
		}

		employee.UserId = user.ID
		employee.ChargeId = charge.ID

		res = tx.Create(&employee)
		if res.Error != nil {
			return res.Error
		}
		return nil

	})

	user = r.FindUserByUsername(user.UserName)
	return user, nil

}

func (r UserRepository) CreateUser(user gateway_entities.User) (gateway_entities.User, error) {
	var userDB gateway_entities.User
	id := r.db.Create(&user)
	if id.RowsAffected == 0 {
		return gateway_entities.User{}, errors.New("user not created")
	}
	r.db.Where("user_name = ?", user.UserName).First(&userDB)

	return userDB, nil
}

func (r *UserRepository) FindUserByUsername(username string) gateway_entities.User {
	var user gateway_entities.User
	res := r.db.Where("user_name = ?", username).First(&user)
	if res.Error != nil {
		return gateway_entities.User{}
	}
	return user
}

func (r *UserRepository) FindUserByUsernameAndPassword(username string, password string) (gateway_entities.User, error) {
	var user gateway_entities.User
	user.UserName = username
	user.Password = password
	res := r.db.Where("user_name = ? AND password = ?", username, password).First(&user)
	if res.Error != nil {
		return gateway_entities.User{}, res.Error
	}
	if res.RowsAffected == 0 {
		return gateway_entities.User{}, errors.New("user not found")
	}
	return user, nil
}

func (r *UserRepository) UpdateUser(user gateway_entities.User) error {
	r.db.Save(&user)
	return nil
}

func (r *UserRepository) DeleteUser(id string) error {
	r.db.Delete(&gateway_entities.User{}, id)
	return nil
}
