package domain

import (
	"rapicreds-backend/src/app/domain/constants"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint               `json:"id" gorm:"primaryKey"`
	Email    string             `json:"email"`
	Name     string             `json:"name"`
	LastName string             `json:"last_name"`
	Password string             `json:"password,omitempty"`
	Document string             `json:"document,omitempty"`
	Sex      string             `json:"sex,omitempty"`
	GoogleID string             `json:"google_id,omitempty"`
	Role     constants.UserRole `json:"role,omitempty,"`
	Loans    []UserLoan         `json:"loans,omitempty" gorm:"foreignKey:UserID"`
}

func (u *User) IsVerified() bool {
	return u.Name != "" && u.LastName != "" && u.Email != "" && u.Document != "" && u.Sex != ""
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// Si tiene contraseña, encriptarla
	if u.Password != "" {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

func (u *User) CheckPassword(password string) bool {
	// Verificar si la contraseña coincide
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
