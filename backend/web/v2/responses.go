package v2

import (
	"github.com/google/uuid"
	"ppo/web"
)

type UserResponse struct {
	User web.User `json:"user"`
}

type SaladsResponse struct {
	Pages  int         `json:"num_pages"`
	Salads []web.Salad `json:"salads"`
}

type SaladResponse struct {
	Salad web.Salad `json:"salad"`
}

type RecipeResponse struct {
	Recipe web.Recipe `json:"recipe"`
}

type RecipeStepsResponse struct {
	Steps []web.RecipeStep `json:"step"`
}

type IngredientsResponse struct {
	Ingredients []web.Ingredient `json:"step"`
}

type LinkIngredientSaladResponse struct {
	Id uuid.UUID `json:"link_id"`
}

type LinkTypeSaladResponse struct {
	Id uuid.UUID `json:"link_id"`
}

type SaladTypesResponse struct {
	Types []web.SaladType `json:"types"`
}

type AllSaladTypesResponse struct {
	Pages  int         `json:"num_pages"`
	Salads []web.Salad `json:"salads"`
}

type MeasurementsResponse struct {
	Measurements []web.Measurement `json:"measurements"`
}

type MeasurementResponse struct {
	Measurements web.Measurement `json:"measurement"`
	Count        int             `json:"count"`
}

type CommentsResponse struct {
	Pages    int           `json:"num_pages"`
	Comments []web.Comment `json:"comments"`
}
