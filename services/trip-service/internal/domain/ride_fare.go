package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type RideFareModel struct {
	ID                primitive.ObjectID
	UserId            string
	PackageSlug       string
	TotalPriceinCents float64
}
