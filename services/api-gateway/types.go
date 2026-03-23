package main

import "ride-sharing/shared/types"

type PreviewTripRequest struct {
	UserId      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}
