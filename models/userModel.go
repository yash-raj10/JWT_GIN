package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id"`
	First_name   *string            `json:"first_name" validation:"required, min=2, max=100"`
	Last_name    *string            `json:"last_name" validation:"required, min=2, max=10"`
	Password     *string            `json:"password" validation:"required, min=6"`
	Email        *string            `json:"email" validation:"email, required"`
	Phone        *string            `json:"phone" validation:"required"`
	Token        *string            `json:"token"`
	User_type    *string            `json:"user_Type" validation:"required, eq=ADMIN|eq=USER"` // same as enum in js
	Refesh_token *string            `json:"refresh_token"`
	Created_at   time.Time          `json:"created_at"`
	Updated_at   time.Time          `json:"updated_at"`
	User_id      string             `json:"user_id" `
}
