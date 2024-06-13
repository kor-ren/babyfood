package data

import (
	"encoding/json"
	"fmt"

	"github.com/dgraph-io/badger/v4"
	"github.com/google/uuid"
	"github.com/kor-ren/babyfood/graph/model"
)

var (
	mealsPrefix = []byte("meals:")
)

func (data *DataContext) GetMeals() ([]*model.Meal, error) {

	result := make([]*model.Meal, 0)

	err := data.db.View(func(txn *badger.Txn) error {
		itr := txn.NewIterator(badger.IteratorOptions{
			Prefix:       mealsPrefix,
			PrefetchSize: 10,
		})
		defer itr.Close()

		for itr.Rewind(); itr.Valid(); itr.Next() {
			item := itr.Item()
			err := item.Value(func(val []byte) error {
				meal := model.Meal{}

				if parseErr := json.Unmarshal(val, &meal); parseErr != nil {
					return parseErr
				}

				result = append(result, &meal)
				return nil
			})

			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (data *DataContext) CreateMeal(newMeal model.NewMeal) (*model.Meal, error) {

	id, err := uuid.NewV7()

	if err != nil {
		return nil, fmt.Errorf("could not create id: %w", err)
	}

	result := &model.Meal{
		ID:     id.String(),
		Name:   newMeal.Name,
		Rating: newMeal.Rating,
		Image:  newMeal.Image,
	}

	key := append(mealsPrefix, []byte(result.ID)...)

	value, err := json.Marshal(result)

	if err != nil {
		return nil, fmt.Errorf("could not convert to json: %w", err)
	}

	err = data.db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value)
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (data *DataContext) UpdateMeal(input *model.UpdateMeal) (*model.Meal, error) {
	key := append(mealsPrefix, []byte(input.ID)...)

	meal := &model.Meal{}

	err := data.db.Update(func(txn *badger.Txn) error {
		item, itemErr := txn.Get(key)
		if itemErr != nil {
			return fmt.Errorf("not found: %w", itemErr)
		}

		itemErr = item.Value(func(val []byte) error {
			return json.Unmarshal(val, meal)
		})

		if itemErr != nil {
			return fmt.Errorf("could read data from db: %w", itemErr)
		}

		if input.Name != nil {
			meal.Name = *input.Name
		}

		if input.Image != nil {
			meal.Image = input.Image
		}

		if input.Rating != nil {
			meal.Rating = input.Rating.Value
		}

		jsonData, jsonErr := json.Marshal(meal)

		if jsonErr != nil {
			return fmt.Errorf("could not serialize new data: %w", jsonErr)
		}

		itemErr = txn.Set(key, jsonData)

		if itemErr != nil {
			return fmt.Errorf("could not store new data: %w", itemErr)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return meal, nil
}
