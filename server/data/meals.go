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

func (data *DataContext) GetMeals(name *string) ([]*model.Meal, error) {

	var result []*model.Meal

	var args []interface{}
	where := ""

	if name != nil {
		where = " WHERE name LIKE ?"
		args = append(args, fmt.Sprintf("%%%s%%", *name))
	}

	rows, err := data.db.Query(fmt.Sprintf("SELECT %s FROM meals %s ORDER BY name, id", mealFields, where), args...)
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

	var args []interface{}
	atLeastOne := false

	sql := "UPDATE meals SET "

	addIfSet := func(value any, col string, checkEmpty bool) {
		if checkEmpty && value == nil {
			return
		}

		if atLeastOne {
			sql += ", "
		}

		sql = fmt.Sprintf("%s %s = ? ", sql, col)
		args = append(args, value)
		atLeastOne = true
	}

	addIfSet(input.Name, "name", true)
	addIfSet(input.Image, "image", true)

	if input.Rating != nil {
		addIfSet(input.Rating.Value, "rating", false)
	}

	if !atLeastOne {
		return nil, fmt.Errorf("no changes")
	}

	sql = fmt.Sprintf("%s, updated_at = CURRENT_TIMESTAMP WHERE id = ? RETURNING %s", sql, mealFields)
	args = append(args, input.ID)

	row := data.db.QueryRow(sql, args...)

	if err := row.Err(); err != nil {
		return nil, err
	}

	meal, err := scanAsMeal(row)
	if err != nil {
		return nil, err
	}

	return meal, nil
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
