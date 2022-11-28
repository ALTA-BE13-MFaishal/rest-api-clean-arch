package repository

import (
	"14-api-clean-arch/features/user"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *userRepository) Create(input user.Core) (row int, err error) {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *userRepository) GetAll() (data []user.Core, err error) {
	var users []User

	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(users)
	return dataCore, nil
}

// GetById implements user.RepositoryInterface
func (repo *userRepository) GetById(id int) (data user.Core, err error) {
	var user User

	tx := repo.db.First(&user, id)
	if tx.Error != nil {
		return data, tx.Error
	}
	var dataCore = user.toCore()
	return dataCore, nil
}

// Update implements user.Repository
func (repo *userRepository) Update(input user.Core, id int) (row int, err error) {
	userGorm := fromCore(input)
	var user User
	tx := repo.db.Model(&user).Where("ID = ?", id).Updates(&userGorm) // proses update
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("update failed")
	}
	return int(tx.RowsAffected), nil
}

// Delete implements user.Repository
func (repo *userRepository) Delete(id int) (row int, err error) {
	var user User
	tx := repo.db.Delete(&user, id) // proses delete
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("delete failed")
	}
	return int(tx.RowsAffected), nil
}
