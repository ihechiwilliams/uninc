package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
	"uninc/database"
	"uninc/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var couponCollection *mongo.Collection = database.OpenCollection(database.Client, "coupon")

func CreateCoupon(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var coupon models.Coupon

	//bind the object that comes in with the declared varaible. thrrow an error if one occurs
	if err := c.BindJSON(&coupon); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validationErr := validate.Struct(coupon)
	if validationErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	count, err := couponCollection.CountDocuments(ctx, bson.M{"couponcode": coupon.CouponCode})
	defer cancel()
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error occured while checking for the coupon code"})
		return
	}

	if count > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "this coupon code already exists"})
		return
	}

	coupon.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	coupon.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	coupon.ID = primitive.NewObjectID()
	coupon.Coupon_id = coupon.ID.Hex()

	_, insertErr := couponCollection.InsertOne(ctx, coupon)
	if insertErr != nil {
		msg := fmt.Sprintf("Coupon item was not created")
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}
	defer cancel()

	var coupons []models.Coupon

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"couponcode": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	cursor, err := couponCollection.Find(ctx, filter, findOptions)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &coupons); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, coupons)
}

func ListCoupons(c *gin.Context) {
	var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	var coupons []models.Coupon

	filter := bson.M{}
	findOptions := options.Find()

	if s := c.Query("s"); s != "" {
		filter = bson.M{
			"$or": []bson.M{
				{
					"couponcode": bson.M{
						"$regex": primitive.Regex{
							Pattern: s,
							Options: "i",
						},
					},
				},
			},
		}
	}

	cursor, err := couponCollection.Find(ctx, filter, findOptions)
	defer cancel()
	if err != nil {
		log.Fatal(err)
	}
	if err = cursor.All(ctx, &coupons); err != nil {
		log.Fatal(err)
	}
	c.JSON(http.StatusOK, coupons)
}
