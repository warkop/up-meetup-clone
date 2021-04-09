package repository

import (
	"github.com/warkop/up-meetup-clone/api/v1/user/models"
	"github.com/warkop/up-meetup-clone/lib/db"
)

type UserRepositoryProto interface {
	Create(user *models.User) error
	Fetch() ([]*models.User, error)
	FetchByID(id int64) (*models.User, error)
	Update(id int64, user *models.User) error
	Delete(id int64) error
}

type UserRepository struct {
	Orm db.TransactionProto
}

func NewUserRepository(tx db.TransactionProto) UserRepositoryProto {
	return &UserRepository{
		Orm: tx,
	}
}

func (repo *UserRepository) Create(user *models.User) error {
	tx := repo.Orm.(*db.Transaction)

	tx.Begin()

	if tx.GetError() != nil {
		return tx.GetError()
	}

	tx.Create(user)

	if tx.GetError() != nil {
		tx.Rollback()
		return tx.GetError()
	}

	tx.Commit()

	return nil
}

func (repo *UserRepository) Fetch() ([]*models.User, error) {
	users := make([]*models.User, 0)
	tx := repo.Orm.(*db.Transaction)

	tx.Begin()

	if tx.GetError() != nil {
		return nil, tx.GetError()
	}

	tx.Find(&users)

	if tx.GetError() != nil {
		tx.Rollback()
		return nil, tx.GetError()
	}

	tx.Commit()

	return users, nil
}

func (repo *UserRepository) FetchByID(id int64) (*models.User, error) {
	user := new(models.User)
	tx := repo.Orm.(*db.Transaction)

	tx.Begin()

	if tx.GetError() != nil {
		return nil, tx.GetError()
	}

	tx.First(&user, id)

	if tx.GetError() != nil {
		tx.Rollback()
		return nil, tx.GetError()
	}

	tx.Commit()

	return user, nil
}

func (repo *UserRepository) Update(id int64, user *models.User) error {
	tx := repo.Orm.(*db.Transaction)

	tx.Begin()

	if tx.GetError() != nil {
		return tx.GetError()
	}

	ud := make(map[string]interface{}, 0)
	ud["name"] = user.Name
	ud["email"] = user.Email

	tx.Where("id = ?", id).Update(ud)

	if tx.GetError() != nil {
		tx.Rollback()
		return tx.GetError()
	}

	tx.Commit()

	return nil
}

func (repo *UserRepository) Delete(id int64) error {
	tx := repo.Orm.(*db.Transaction)

	tx.Where("id = ?", id).Delete(models.User{})

	if tx.GetError() != nil {
		tx.Rollback()
		return tx.GetError()
	}

	tx.Commit()

	return nil
}
