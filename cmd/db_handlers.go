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
	row := app.DB.QueryRow(`SELECT * FROM categories WHERE id=$1;`, id)
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
	WHERE id=$3;`, data.CategoryIcon, data.CategoryName, id)
	if err != nil {
		return data, err
	}

	data.CategoryID = id
	return data, nil
}