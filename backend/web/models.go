package web

import (
	"github.com/google/uuid"
	"net/mail"
	"ppo/domain"
)

type LoginReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterReq struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type User struct {
	ID       uuid.UUID    `json:"id,omitempty"`
	Name     string       `json:"name,omitempty"`
	Username string       `json:"username,omitempty"`
	Email    mail.Address `json:"email,omitempty"`
	Role     string       `json:"role,omitempty"`
}

type SaladCreate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type SaladUpdate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Salad struct {
	ID          uuid.UUID `json:"id,omitempty"`
	AuthorID    uuid.UUID `json:"author_id,omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type RecipeCreate struct {
	SaladID          uuid.UUID `json:"salad_id,omitempty"`
	NumberOfServings int       `json:"number_of_servings,omitempty"`
	TimeToCook       int       `json:"time_to_cook,omitempty"`
}

type RecipeUpdate struct {
	Status           int `json:"status,omitempty"`
	NumberOfServings int `json:"number_of_servings,omitempty"`
	TimeToCook       int `json:"time_to_cook,omitempty"`
}

type Recipe struct {
	ID               uuid.UUID `json:"id,omitempty"`
	SaladID          uuid.UUID `json:"salad_id,omitempty"`
	Status           int       `json:"status,omitempty"`
	NumberOfServings int       `json:"number_of_servings,omitempty"`
	TimeToCook       int       `json:"time_to_cook,omitempty"`
	Rating           float32   `json:"rating,omitempty"`
}

type RecipeStepCreate struct {
	RecipeID    uuid.UUID `json:"recipe_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
}

type RecipeStepUpdate struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	StepNum     int    `json:"step_num,omitempty"`
}

type RecipeStep struct {
	ID          uuid.UUID `json:"id,omitempty"`
	RecipeID    uuid.UUID `json:"recipe_id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
	StepNum     int       `json:"step_num,omitempty"`
}

type Ingredient struct {
	ID       uuid.UUID `json:"id,omitempty"`
	TypeID   uuid.UUID `json:"type_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Calories int       `json:"calories,omitempty"`
}

type SaladType struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
}

type IngredientType struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Description string    `json:"description,omitempty"`
}

type Measurement struct {
	ID    uuid.UUID `json:"id,omitempty"`
	Name  string    `json:"name,omitempty"`
	Grams int       `json:"grams,omitempty"`
}

type CommentUpdate struct {
	Text   string `json:"text,omitempty"`
	Rating int    `json:"rating,omitempty"`
}

type CommentCreate struct {
	SaladID uuid.UUID `json:"salad_id,omitempty"`
	Text    string    `json:"text,omitempty"`
	Rating  int       `json:"rating,omitempty"`
}

type Comment struct {
	ID       uuid.UUID `json:"id,omitempty"`
	AuthorID uuid.UUID `json:"author_id,omitempty"`
	SaladID  uuid.UUID `json:"salad_id,omitempty"`
	Text     string    `json:"text,omitempty"`
	Rating   int       `json:"rating,omitempty"`
}

type LinkSaladType struct {
	SaladId uuid.UUID `json:"salad_id,omitempty"`
	TypeId  uuid.UUID `json:"type_id,omitempty"`
}

type LinkIngredientSalad struct {
	RecipeId      uuid.UUID `json:"recipe_id,omitempty"`
	SaladId       uuid.UUID `json:"salad_id,omitempty"`
	IngredientId  uuid.UUID `json:"ingredient_id,omitempty"`
	MeasurementId uuid.UUID `json:"measurement_id,omitempty"`
	Amount        int       `json:"amount,omitempty"`
}

func ToUserTransport(user *domain.User) User {
	return User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Role:     user.Role,
	}
}

func ToSaladTransport(salad *domain.Salad) Salad {
	return Salad{
		ID:          salad.ID,
		AuthorID:    salad.AuthorID,
		Name:        salad.Name,
		Description: salad.Description,
	}
}

func ToSaladModel(salad *Salad) *domain.Salad {
	return &domain.Salad{
		ID:          salad.ID,
		AuthorID:    salad.AuthorID,
		Name:        salad.Name,
		Description: salad.Description,
	}
}

func ToLinkSaladTypeModel(link *LinkSaladType) (SaladId, TypeId uuid.UUID) {
	return link.SaladId, link.TypeId
}

func ToLinkIngredientModel(link *LinkIngredientSalad) (RecipeId, SaladId, IngredientId, MeasurementId uuid.UUID, amount int) {
	return link.RecipeId, link.SaladId, link.IngredientId, link.MeasurementId, link.Amount
}

func ToRecipeTransport(recipe *domain.Recipe) Recipe {
	return Recipe{
		ID:               recipe.ID,
		SaladID:          recipe.SaladID,
		Status:           recipe.Status,
		NumberOfServings: recipe.NumberOfServings,
		TimeToCook:       recipe.TimeToCook,
		Rating:           recipe.Rating,
	}
}

func ToRecipeModel(recipe *Recipe) *domain.Recipe {
	return &domain.Recipe{
		ID:               recipe.ID,
		SaladID:          recipe.SaladID,
		Status:           recipe.Status,
		NumberOfServings: recipe.NumberOfServings,
		TimeToCook:       recipe.TimeToCook,
		Rating:           recipe.Rating,
	}
}

func ToRecipeStepTransport(step *domain.RecipeStep) RecipeStep {
	return RecipeStep{
		ID:          step.ID,
		RecipeID:    step.RecipeID,
		Name:        step.Name,
		Description: step.Description,
		StepNum:     step.StepNum,
	}
}

func ToRecipeStepModel(step *RecipeStep) *domain.RecipeStep {
	return &domain.RecipeStep{
		ID:          step.ID,
		RecipeID:    step.RecipeID,
		Name:        step.Name,
		Description: step.Description,
		StepNum:     step.StepNum,
	}
}

func ToIngredientTransport(ingredient *domain.Ingredient) Ingredient {
	return Ingredient{
		ID:       ingredient.ID,
		TypeID:   ingredient.TypeID,
		Name:     ingredient.Name,
		Calories: ingredient.Calories,
	}
}

func ToSaladTypeTransport(saladType *domain.SaladType) SaladType {
	return SaladType{
		ID:          saladType.ID,
		Name:        saladType.Name,
		Description: saladType.Description,
	}
}

func ToIngredientTypeTransport(ingredientType *domain.IngredientType) IngredientType {
	return IngredientType{
		ID:          ingredientType.ID,
		Name:        ingredientType.Name,
		Description: ingredientType.Description,
	}
}

func ToMeasurementTransport(measurement *domain.Measurement) Measurement {
	return Measurement{
		ID:    measurement.ID,
		Name:  measurement.Name,
		Grams: measurement.Grams,
	}
}

func ToCommentTransport(comment *domain.Comment) Comment {
	return Comment{
		ID:       comment.ID,
		AuthorID: comment.AuthorID,
		SaladID:  comment.SaladID,
		Text:     comment.Text,
		Rating:   comment.Rating,
	}
}

func ToCommentModel(comment *Comment) *domain.Comment {
	return &domain.Comment{
		ID:       comment.ID,
		AuthorID: comment.AuthorID,
		SaladID:  comment.SaladID,
		Text:     comment.Text,
		Rating:   comment.Rating,
	}
}
