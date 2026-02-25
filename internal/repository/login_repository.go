package repository

import (
	"github.com/cecepwahyu/rest-api-go/internal/model"

	"gorm.io/gorm"
)

type LoginRepository interface {
	FindByEmail(email string) (*model.Peserta, error)
	Save(user *model.Peserta) error
}

type loginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) LoginRepository {
	return &loginRepository{db}
}

func (r *loginRepository) FindByEmail(email string) (*model.Peserta, error) {
	var user model.Peserta
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *loginRepository) Save(user *model.Peserta) error {
	return r.db.Save(user).Error
}
