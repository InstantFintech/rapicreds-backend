package domain

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"`
	GoogleID string `json:"google_id,omitempty"`
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
