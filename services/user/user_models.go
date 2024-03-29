package models

import (
	"github.com/golang-jwt/jwt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

type Session struct {
	Email string `json:"email"`
	Token string `json:"token"`
}

type UserType string

const (
	AdminUser   UserType = "admin"
	DefaultUser UserType = "default"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email     string  `json:"email" validate:"required,email" unique:"true"`
	Password  string  `json:"password" validate:"required"`
	FirstName *string `json:"firstName" omitempty:"true"`
	LastName  *string `json:"lastName" omitempty:"true"`
}

// ---------------------------------------------------------------------------------------------------
// ------------------------------------------ MONGO OBJECTS ------------------------------------------
// ---------------------------------------------------------------------------------------------------
type User struct {
	ID              primitive.ObjectID  `json:"id,omitempty" bson:"_id,omitempty" validate:"required"`
	Username        string              `json:"username,omitempty" bson:"username,omitempty" validate:"required"`
	Password        string              `json:"-" bson:"password,omitempty" validate:"required"`
	Email           string              `json:"email" bson:"email" validate:"required,email"`
	FirstName       *string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName        *string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Phone           *string             `json:"phone,omitempty" bson:"phone,omitempty"`
	ShippingAddress *primitive.ObjectID `json:"shippingAddress,omitempty" bson:"shippingAddress,omitempty"`
	Shop            *primitive.ObjectID `json:"shop,omitempty" bson:"shop,omitempty"`
	CreatedAt       primitive.DateTime  `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt       primitive.DateTime  `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
