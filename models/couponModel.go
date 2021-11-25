package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Coupon struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       *string            `json:"title" validate:"required,min=2,max=100"`
	Description *string            `json:"description" validate:"required,min=2,max=100"`
	DocumentUrl *string            `json:"documenturl" validate:"required"`
	CouponCode  *string            `json:"couponcode" validate:"required"`
	Created_at  time.Time          `json:"created_at"`
	Updated_at  time.Time          `json:"updated_at"`
	Coupon_id   string             `json:"coupon_id"`
}
