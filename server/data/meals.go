package data

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kor-ren/babyfood/graph/model"
)

type scanInterface interface {
	Scan(dest ...any) error
}

const mealFields = "id, name, rating, image, created_at, updated_at"

func (data *DataContext) GetMeals() ([]*model.Meal, error) {

	var result []*model.Meal

	rows, err := data.db.Query(fmt.Sprintf("SELECT %s FROM meals ORDER BY name, id", mealFields))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		meal, err := scanAsMeal(rows)

		if err != nil {
			return nil, err
		}

		result = append(result, meal)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return result, nil

}

func (data *DataContext) GetMealById(id string) (*model.Meal, error) {
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return nil, fmt.Errorf("id is not allowed")
	}

	row := data.db.QueryRow(fmt.Sprintf("SELECT %s FROM meals WHERE id = ?", mealFields), idInt)

	if err := row.Err(); err != nil {
		return nil, err
	}

	meal, err := scanAsMeal(row)

	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (data *DataContext) CreateMeal(input model.NewMeal) (*model.Meal, error) {
	row := data.db.QueryRow(
		fmt.Sprintf("INSERT INTO meals (name, rating, image) VALUES (?, ?, ?) RETURNING %s", mealFields),
		input.Name,
		input.Rating,
		input.Image,
	)

	if err := row.Err(); err != nil {
		return nil, err
	}

	meal, err := scanAsMeal(row)
	if err != nil {
		return nil, err
	}

	return meal, nil
}

func (data *DataContext) UpdateMeal(input model.UpdateMeal) (*model.Meal, error) {
	return nil, fmt.Errorf("not implemented")
}

func scanAsMeal(rows scanInterface) (*model.Meal, error) {
	var id int
	var rating *int
	var name string
	var image *string
	var created_at, updated_at time.Time

	if err := rows.Scan(&id, &name, &rating, &image, &created_at, &updated_at); err != nil {
		return nil, err
	}

	return &model.Meal{
		ID:        fmt.Sprintf("%d", id),
		Name:      name,
		Rating:    rating,
		Image:     image,
		CreatedAt: created_at,
		UpdatedAt: updated_at,
	}, nil
}
