package dto

type Destination struct{
	DestinationID string `json:"destination_id"`
	DestinationNmae string `json:"destination_name"`
	DestinationLocation string `json:"destination_location"`
	DestinationDescription string `json:"destination_description"`
	DestinationReviews string `json:"destination_reviews"`
	DestinationPrice string `json:"destination_price"`
	DestinationGimmickPrice string `json:"destination_gimmick_price"`
	DestinationCategory Category `json:"destination_category"`
}