// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Meal struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Rating *int    `json:"rating,omitempty"`
	Image  *string `json:"image,omitempty"`
}

type Mutation struct {
}

type NewMeal struct {
	Name   string  `json:"name"`
	Rating *int    `json:"rating,omitempty"`
	Image  *string `json:"image,omitempty"`
}

type Query struct {
}

type RatingValue struct {
	Value *int `json:"value,omitempty"`
}

type UpdateMeal struct {
	ID     string       `json:"id"`
	Name   *string      `json:"name,omitempty"`
	Rating *RatingValue `json:"rating,omitempty"`
	Image  *string      `json:"image,omitempty"`
}
