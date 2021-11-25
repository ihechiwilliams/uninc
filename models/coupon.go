package models

type Coupon struct {
	DocumentUrl string `json:"documenturl" bson:"documenturl"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	CouponCode  string `json:"couponcode" bson:"couponcode"`
}
