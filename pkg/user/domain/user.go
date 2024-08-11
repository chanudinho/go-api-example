package domain

import (
	"time"

	"github.com/chanudinho/go-api-example/pkg/user/models"
)

type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) FromModel(m models.User) {
	u.ID = m.ID
	u.Name = m.Name
	u.Email = m.Email
	u.Role = m.Role
	u.CreatedAt = m.CreatedAt
	u.UpdatedAt = m.UpdatedAt
}

func (u *User) ToModel() models.User {
	return models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Role:     u.Role,
	}
}
