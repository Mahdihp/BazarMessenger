package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	UserName    string             `bson:"user_name"`
	CellNumber  string             `bson:"cell_number"`
	AreaCode    string             `bson:"area_code"`
	Bio         string             `bson:"bio"`
	Avatar      string             `bson:"avatar"`
	Acive       bool               `bson:"acive"`
	ContactList []Contact          `bson:"contact_list"`
	Store       Store              `bson:"store"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
type Store struct {
	ID          string    `bson:"_id"`
	Title       string    `bson:"title"`
	Description string    `bson:"description"`
	Address     []string  `bson:"address"`
	PhoneNumber []string  `bson:"phone_number"`
	Products    []Product `bson:"products"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
type Product struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Code        string             `bson:"code"`
	Price       int32              `bson:"price"`
	Description string             `bson:"description"`

	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

type Contact struct {
	UserID   primitive.ObjectID `bson:"_id"`
	UserName string             `bson:"user_name"`
}

type UserRepository interface {
	Create(user *User, context context.Context) error
	Update(user *User, context context.Context) error
	GetById(filter interface{}, context context.Context) (*User, error)
	GetByUsername(filter interface{}, context context.Context) (*User, error)
}

type UserUsecase interface {
	Create(user *User, context context.Context) error
	Update(user *User, context context.Context) error
	GetById(id string, context context.Context) (*User, error)
	GetByUsername(userName string, context context.Context) (*User, error)
}
