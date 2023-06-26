package cmd

import (
	"travelo/internal/dto"

	"github.com/google/uuid"
)

// Category
func (app *Application) GetAllCategory() ([]dto.Category, error) {

	rows, err := app.DB.Query("select * from categories")

	var categories []dto.Category

	if err != nil {
		return categories, err
	}

	defer rows.Close()

	for rows.Next() {
		var cat dto.Category
		if err := rows.Scan(&cat.CategoryID, &cat.CategoryIcon, &cat.CategoryName); err != nil {
			return categories, err
		}
		categories = append(categories, cat)
	}

	return categories, nil
}

func (app *Application) GetCategoryByID(id string) (dto.Category, error) {
	//var err error
	category := dto.Category{}
	row := app.DB.QueryRow(`SELECT * FROM categories WHERE category_id=$1;`, id)
	if err := row.Scan(&category.CategoryID, &category.CategoryIcon, &category.CategoryName); err != nil {
		return category, err
	}
	return category, nil
}

func (app *Application) PostCategory(data dto.Category) (dto.Category, error) {

	data.CategoryID = uuid.New().String()
	_, err := app.DB.Exec(`INSERT INTO categories (category_id, category_icon, category_name) VALUES ($1, $2, $3)`, data.CategoryID, data.CategoryIcon, data.CategoryName)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (app *Application) UpdateCategory(data dto.Category, id string) (dto.Category, error) {
	_, err := app.DB.Exec(`UPDATE categories
	SET category_icon=$1, category_name=$2
	WHERE category_id=$3;`, data.CategoryIcon, data.CategoryName, id)
	if err != nil {
		return data, err
	}

	data.CategoryID = id
	return data, nil
}


// Review
func (app *Application) GetAllReviews() ([]dto.Review, error) {

	rows, err := app.DB.Query("select * from reviews")

	var reviews []dto.Review

	if err != nil {
		return reviews, err
	}

	defer rows.Close()

	for rows.Next() {
		var rev dto.Review
		if err := rows.Scan(&rev.ReviewID, &rev.ReviewDate, &rev.ReviewDescription, &rev.ReviewLocation, 
			&rev.ReviewStar, &rev.ReviewUserIcon, &rev.ReviewUserLocation, &rev.ReviewUserName); err != nil {
			return reviews, err
		}
		reviews = append(reviews, rev)
	}

	return reviews, nil
}

func (app *Application) GetReviewByID(id string) (dto.Review, error) {
	//var err error
	rev := dto.Review{}
	row := app.DB.QueryRow(`SELECT * FROM reviews WHERE review_id=$1;`, id)
	if err := row.Scan(&rev.ReviewID, &rev.ReviewDate, &rev.ReviewDescription, &rev.ReviewLocation, 
		&rev.ReviewStar, &rev.ReviewUserIcon, &rev.ReviewUserLocation, &rev.ReviewUserName); err != nil {
		return rev, err
	}
	return rev, nil
}

func (app *Application) PostReview(data dto.Review) (dto.Review, error) {

	data.ReviewID = uuid.New().String()
	_, err := app.DB.Exec(`INSERT INTO reviews (review_id, review_date, review_description, 
		review_location, review_star, review_user_icon, review_user_location, review_user_name) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, 
			data.ReviewID, data.ReviewDate, data.ReviewDescription, data.ReviewLocation, 
			data.ReviewStar, data.ReviewUserIcon, data.ReviewUserLocation, data.ReviewUserName)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (app *Application) UpdateReview(data dto.Review, id string) (dto.Review, error) {
	_, err := app.DB.Exec(`UPDATE reviews
	SET review_date=$1, review_description=$2, review_location=$3, review_star=$4, review_user_icon=$5, review_user_location=$6, review_user_name=$7
	WHERE review_id=$8;`, data.ReviewDate, data.ReviewDescription, data.ReviewLocation, 
	data.ReviewStar, data.ReviewUserIcon, data.ReviewUserLocation, data.ReviewUserName, id)
	if err != nil {
		return data, err
	}

	data.ReviewID = id
	return data, nil
}