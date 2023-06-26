package dto

type Review struct{
	ReviewID string `json:"review_id"`
	ReviewUserName string `json:"review_user_name"`
	ReviewUserIcon string `json:"review_user_icon"`
	ReviewUserLocation string `json:"review_user_location"`
	ReviewDescription string `json:"review_description"`
	ReviewLocation string `json:"review_location"`
	ReviewDate string `json:"review_date"`
	ReviewStar int `json:"review_star"`
}