package v2

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"net/http"
	"net/mail"
	"ppo/domain"
	"ppo/internal/app"
	"ppo/pkg/base"
	errs "ppo/pkg/errors"
	"ppo/services/dto"
	"ppo/web"
	"strconv"
	"time"
)

// LoginHandler godoc
//
//	@Summary		Аутентификация
//	@ID				login
//	@Produce json
//	@Description	Метод для получения bearer-токена
//	@Param data body web.LoginReq true "Login data"
//	@Tags			users
//	@Success		200	{object} web.SuccessResponseStruct
//	@Failure		401	{object} web.ErrorResponseStruct
//	@Router			/login [post]
func LoginHandler(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "Authentication"

		var req web.LoginReq

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ua := &domain.UserAuth{Username: req.Login, Password: req.Password}
		token, err := app.AuthService.Login(r.Context(), ua)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusUnauthorized)
			return
		}

		_, err = base.VerifyAuthToken(token, app.Config.Jwt.Key)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: проверка JWT-токена: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:    "access_token",
			Value:   token,
			Path:    "/",
			Secure:  true,
			Expires: time.Now().Add(3600 * 24 * time.Second),
		}

		http.SetCookie(w, &cookie)
		web.SuccessResponse(w, http.StatusOK, map[string]string{"token": token})
	}
}

// RegisterHandler godoc
//
//	@Summary		Регистрация
//	@ID				signup
//	@Produce json
//	@Description	Метод для регистрации
//	@Param data body web.RegisterReq true "Register data"
//	@Tags			users
//	@Success		200	{object} web.SuccessResponseStruct
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Router			/register [post]
func RegisterHandler(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "Registration"

		var req web.RegisterReq

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		var mailAddr mail.Address
		mailAddr.Address = req.Email

		ua := &domain.User{
			Name:     req.Name,
			Username: req.Username,
			Password: req.Password,
			Email:    mailAddr,
		}
		token, err := app.AuthService.Register(r.Context(), ua)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		_, err = base.VerifyAuthToken(token, app.Config.Jwt.Key)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: проверка JWT-токена: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		cookie := http.Cookie{
			Name:    "access_token",
			Value:   token,
			Path:    "/",
			Secure:  true,
			Expires: time.Now().Add(3600 * 24 * time.Second),
		}

		http.SetCookie(w, &cookie)

		web.SuccessResponse(w, http.StatusOK, map[string]string{"token": token})
	}
}

// GetUser godoc
//
//	@Summary		Информация о пользователе
//	@ID				getUser
//	@Produce json
//	@Description	Метод для получения информации о пользователе
//	@Tags			users
//	@Param id path string true "User`s id"
//	@Success		200	{object} UserResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/users/{id} [get]
func GetUser(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "Getting user info"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		user, err := app.UserService.GetById(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"user": web.ToUserTransport(user)})
	}
}

func GetSalads(app *app.App, w http.ResponseWriter, r *http.Request) {
	prompt := "getting salads"
	r.ParseForm()

	page := r.URL.Query().Get("page")
	if page == "" {
		web.ErrorResponse(w, fmt.Errorf("%s: empty page number", prompt).Error(), http.StatusBadRequest)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: converting page number to int: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}

	minRate := r.URL.Query().Get("minRate")
	if minRate == "" {
		minRate = "0.0"
	}
	minRateFloat, err := strconv.ParseFloat(minRate, 64)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: converting minRAte to float: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}

	ingredients := r.Form["ingredients"]
	if len(ingredients) == 0 || len(ingredients[0]) == 0 {
		ingredients = make([]string, 0)
	}
	ingredientUuids := make([]uuid.UUID, len(ingredients))
	for i := 0; i < len(ingredients); i++ {
		ingredientUuids[i], err = uuid.Parse(ingredients[i])
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting ingredient id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
	}

	saladTypes := r.Form["types"]
	if len(saladTypes) == 0 || len(saladTypes[0]) == 0 {
		saladTypes = make([]string, 0)
	}
	saladUuids := make([]uuid.UUID, len(saladTypes))
	for i := 0; i < len(saladTypes); i++ {
		saladUuids[i], err = uuid.Parse(saladTypes[i])
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting salad type id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
	}

	filter := new(dto.RecipeFilter)

	filter.MinRate = minRateFloat
	filter.Status = dto.PublishedSaladStatus
	filter.SaladTypes = saladUuids
	filter.AvailableIngredients = ingredientUuids

	salads, numPages, err := app.SaladService.GetAll(r.Context(), filter, pageInt)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
		return
	}

	saladsTransport := make([]web.Salad, len(salads))
	for i, salad := range salads {
		saladsTransport[i] = web.ToSaladTransport(salad)
	}

	web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salads": saladsTransport})
}

func GetSaladsWithStatus(app *app.App, w http.ResponseWriter, r *http.Request) { // FIXME
	prompt := "getting salads (by status)"
	r.ParseForm()

	page := r.URL.Query().Get("page")
	if page == "" {
		web.ErrorResponse(w, fmt.Errorf("%s: empty page number", prompt).Error(), http.StatusBadRequest)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: converting page number to int: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}

	minRate := r.URL.Query().Get("minRate")
	if minRate == "" {
		minRate = "0.0"
	}
	minRateFloat, err := strconv.ParseFloat(minRate, 64)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: comverting minRate to float: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}

	ingredients := r.Form["ingredients"]
	if len(ingredients) == 0 || len(ingredients[0]) == 0 {
		ingredients = make([]string, 0)
	}
	ingredientUuids := make([]uuid.UUID, len(ingredients))
	for i := 0; i < len(ingredients); i++ {
		ingredientUuids[i], err = uuid.Parse(ingredients[i])
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting ingredient id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
	}

	saladTypes := r.Form["types"]
	if len(saladTypes) == 0 || len(saladTypes[0]) == 0 {
		saladTypes = make([]string, 0)
	}
	saladUuids := make([]uuid.UUID, len(saladTypes))
	for i := 0; i < len(saladTypes); i++ {
		saladUuids[i], err = uuid.Parse(saladTypes[i])
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting salad type id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
	}

	filter := new(dto.RecipeFilter)

	status := r.URL.Query().Get("status")
	statusInt := dto.PublishedSaladStatus
	if status != "" {
		statusInt, err = strconv.Atoi(status)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting status to int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
	}

	filter.Status = statusInt
	filter.MinRate = minRateFloat

	//filter.Status = dto.PublishedSaladStatus
	filter.SaladTypes = saladUuids
	filter.AvailableIngredients = ingredientUuids

	salads, numPages, err := app.SaladService.GetAll(r.Context(), filter, pageInt)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
		return
	}

	saladsTransport := make([]web.Salad, len(salads))
	for i, salad := range salads {
		saladsTransport[i] = web.ToSaladTransport(salad)
	}

	web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salads": saladsTransport})
}

//	@Param minRate query number false "Minimal rating of salad to filter"
//	@Param ingredients query []string false "Available ingredients to filter"
//	@Param types query []string false "Available salad types to filter"

// GetUserSaladsByFlag godoc
//
//	@Summary		Получение салатов пользователя
//	@ID				getUserSalads
//	@Produce json
//	@Description	Получение списка салатов авторизованного пользователя
//	@Tags			salads
//	@Param page query integer true "Page number"
//	@Param onlyRated query boolean false "Show salads rated by user"
//	@Success		200	{object} SaladsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/users/me/salads [get]
func GetUserSaladsByFlag(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if onlyRated := r.URL.Query().Get("onlyRated"); onlyRated == "true" {
			GetUserRatedSalads(app, w, r)
		} else {
			GetUserSalads(app, w, r)
		}
	}
}

// GetSaladsFunc godoc
//
//	@Summary		Получение салатов
//	@ID				getSalads
//	@Produce json
//	@Description	Получение списка салатов по указанным фильтрам
//	@Tags			salads
//	@Param page query integer true "Page number"
//	@Param minRate query number false "Minimal rating of salad to filter"
//	@Param ingredients query []string false "Available ingredients to filter"
//	@Param types query []string false "Available salad types to filter"
//	@Param status query integer false "Status of salad (only for admins)"
//	@Success		200	{object} SaladsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads [get]
func GetSaladsFunc(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		role, err := web.GetStringClaimFromJWT(r.Context(), "role")
		if err != nil || role != "admin" {
			GetSalads(app, w, r)
		} else {
			GetSaladsWithStatus(app, w, r)
		}
	}

}

// GetCommentsFunc godoc
//
//	@Summary		Получение комментариев
//	@ID				getComments
//	@Produce json
//	@Description	Получение списка комментариев
//	@Tags			comments
//	@Param id path string true "Salad's id"
//	@Param page query integer false "Page"
//	@Param user_id query string false "User's id"
//
// @Success		200	{object} CommentsResponse
// @Failure		400	{object} web.ErrorResponseStruct
// @Failure 		500 {object} web.ErrorResponseStruct
// @Router			/salads/{id}/comments [get]
func GetCommentsFunc(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.URL.Query().Get("user_id")
		if userId == "" {
			GetCommentsBySalad(app, w, r)
		} else {
			userUuid, err := uuid.Parse(userId)
			if err != nil {
				web.ErrorResponse(w, fmt.Errorf("converting user id to uuid: %w", err).Error(), http.StatusBadRequest)
				return
			}
			GetUserComment(app, userUuid, w, r)
		}
	}

}

func GetUserSalads(app *app.App, w http.ResponseWriter, r *http.Request) {
	prompt := "получение списка салатов"
	r.ParseForm()

	userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
		return
	}
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		web.ErrorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}

	ingredients := r.Form["ingredients"]
	if len(ingredients) == 0 || len(ingredients[0]) == 0 {
		ingredients = make([]string, 0)
	}
	ingredientUuids := make([]uuid.UUID, len(ingredients))
	for i := 0; i < len(ingredients); i++ {
		ingredientUuids[i], err = uuid.Parse(ingredients[i])
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id ингредиента к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
	}

	saladTypes := r.Form["types"]
	if len(saladTypes) == 0 || len(saladTypes[0]) == 0 {
		saladTypes = make([]string, 0)
	}
	saladUuids := make([]uuid.UUID, len(saladTypes))
	for i := 0; i < len(saladTypes); i++ {
		saladUuids[i], err = uuid.Parse(saladTypes[i])
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id типа салата к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
	}

	salads, err := app.SaladService.GetAllByUserId(r.Context(), userUuid)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
		return
	}

	saladsTransport := make([]web.Salad, len(salads))
	for i, salad := range salads {
		saladsTransport[i] = web.ToSaladTransport(salad)
	}

	web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"num_pages": pageInt, "salads": saladsTransport})
}

func GetUserRatedSalads(app *app.App, w http.ResponseWriter, r *http.Request) {
	prompt := "получение списка салатов, оцененных пользователем"

	userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
		return
	}
	userUuid, err := uuid.Parse(userId)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
		return
	}

	page := r.URL.Query().Get("page")
	if page == "" {
		web.ErrorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}

	salads, numPages, err := app.SaladService.GetAllRatedByUser(r.Context(), userUuid, pageInt)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
		return
	}

	saladsTransport := make([]web.Salad, len(salads))
	for i, salad := range salads {
		saladsTransport[i] = web.ToSaladTransport(salad)
	}

	web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salads": saladsTransport})
}

// GetSaladById godoc
//
//	@Summary		Получение информации о салате
//	@ID				getSalad
//	@Produce json
//	@Tags			salads
//	@Param id path string true "Salad`s id"
//	@Success		200	{object} SaladResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/salads/{id} [get]
func GetSaladById(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting salad by id"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: empty salad id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		app.Logger.Infof("!!!!! ID %s", idUuid.String())

		salad, err := app.SaladService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"salad": web.ToSaladTransport(salad)})
	}
}

func GetSaladRating(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение рейтинга салата"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipe, err := app.RecipeService.GetBySaladId(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusOK, map[string]float32{"rating": recipe.Rating})
	}
}

// GetSaladRecipe godoc
//
//	@Summary		Получение информации о салате
//	@ID				getRecipe
//	@Produce json
//	@Tags			recipes
//	@Param id path string true "Salad`s id"
//	@Success		200	{object} RecipeResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/recipes/{id} [get]
func GetSaladRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting salad recipe"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipe, err := app.RecipeService.GetBySaladId(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"recipe": web.ToRecipeTransport(recipe),
		})
	}
}

// GetRecipeSteps godoc
//
//	@Summary		Получение шагов рецепта
//	@ID				getRecipeSteps
//	@Produce json
//	@Tags			recipeSteps
//	@Param id path string true "Recipe`s id"
//	@Success		200	{object} RecipeStepsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/recipeSteps/{id} [get]
func GetRecipeSteps(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting recipe steps"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: empty recipe id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		steps, err := app.RecipeStepService.GetAllByRecipeID(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		stepsTransport := make([]web.RecipeStep, len(steps))
		for i, step := range steps {
			stepsTransport[i] = web.ToRecipeStepTransport(step)
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"steps": stepsTransport,
		})
	}
}

// GetRecipeIngredients godoc
//
//	@Summary		Получение ингредиентов, использующихся в рецепте
//	@ID				getRecipeIngredients
//	@Produce json
//	@Tags			ingredients
//	@Param id path string true "Recipe`s id"
//	@Success		200	{object} IngredientsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/ingredients/{id} [get]
func GetRecipeIngredients(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting recipe ingredients"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredients, err := app.IngredientService.GetAllByRecipeId(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		ingredientsTransport := make([]web.Ingredient, len(ingredients))
		for i, ingredient := range ingredients {
			ingredientsTransport[i] = web.ToIngredientTransport(ingredient)
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"ingredients": ingredientsTransport,
		})
	}
}

// GetSaladTypes godoc
//
//	@Summary		Получение типов салата
//	@ID				getSaladTypes
//	@Produce json
//	@Tags			types
//	@Param id path string true "Salad`s id"
//	@Success		200	{object} IngredientsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/salads/{id}/types [get]
func GetSaladTypes(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting types of salad"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		types, err := app.SaladTypeService.GetAllBySaladId(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		typesTransport := make([]web.SaladType, len(types))
		for i, saladType := range types {
			typesTransport[i] = web.ToSaladTypeTransport(saladType)
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"types": typesTransport,
		})
	}
}

func GetIngredientType(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение типа ингредиента"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredientType, err := app.IngredientTypeService.GetById(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"ingredientType": web.ToIngredientTypeTransport(ingredientType),
		})
	}
}

// GetMeasurementByRecipe godoc
//
//	@Summary		Получение единицы измерения ингредиента
//	@ID				getIngredientMeasurement
//	@Produce json
//	@Description	Получение единицы измерения ингредиента, использующегося в рецепте
//	@Tags			measurements
//	@Param id path string true "Recipe`s id"
//	@Param ingredient query string true "Ingredient`s id"
//	@Success		200	{object} MeasurementResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/measurements/{id} [get]
func GetMeasurementByRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting measurement of ingredient"

		ingredientId := r.URL.Query().Get("ingredient")
		if ingredientId == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id ингредиента", prompt).Error(), http.StatusBadRequest)
			return
		}
		ingredientUuid, err := uuid.Parse(ingredientId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeId := chi.URLParam(r, "id")
		if recipeId == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: empty id", prompt).Error(), http.StatusBadRequest)
			return
		}
		recipeUuid, err := uuid.Parse(recipeId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		//recipeId := r.URL.Query().Get("recipe")
		//if ingredientId == "" {
		//	web.ErrorResponse(w, fmt.Errorf("%s: пустой id рецепта", prompt).Error(), http.StatusBadRequest)
		//	return
		//}
		//recipeUuid, err := uuid.Parse(recipeId)
		//if err != nil {
		//	web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
		//	return
		//}

		measurement, count, err := app.MeasurementService.GetByRecipeId(r.Context(), ingredientUuid, recipeUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"measurement": web.ToMeasurementTransport(measurement),
			"count":       count,
		})
	}
}

// GetIngredientsByPage godoc
//
//	@Summary		Получение ингредиентов
//	@ID				getIngredientsByPage
//	@Produce json
//	@Description	Получение списка ингредиентов
//	@Tags			ingredients
//	@Param page query integer true "Page number"
//	@Success		200	{object} IngredientsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/ingredients [get]
func GetIngredientsByPage(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение списка ингредиентов"

		page := r.URL.Query().Get("page")
		if page == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой номер страницы", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование номера страницы к int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		ingredients, numPages, err := app.IngredientService.GetAll(r.Context(), pageInt)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		ingredientsTransport := make([]web.Ingredient, len(ingredients))
		for i, ingredient := range ingredients {
			ingredientsTransport[i] = web.ToIngredientTransport(ingredient)
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "ingredients": ingredientsTransport})
	}
}

// GetSaladTypesByPage godoc
//
//	@Summary		Получение типов салатов
//	@ID				getSaladTypesByPage
//	@Produce json
//	@Description	Получение списка типов салатов
//	@Tags			types
//	@Param page query integer true "Page number"
//	@Success		200	{object} AllSaladTypesResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/types [get]
func GetSaladTypesByPage(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting all salad types"

		page := r.URL.Query().Get("page")
		if page == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: empty page number", prompt).Error(), http.StatusBadRequest)
			return
		}
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting page to int: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladTypes, numPages, err := app.SaladTypeService.GetAll(r.Context(), pageInt)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		typesTransport := make([]web.SaladType, len(saladTypes))
		for i, saladType := range saladTypes {
			typesTransport[i] = web.ToSaladTypeTransport(saladType)
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "salad_types": typesTransport})
	}
}

func GetCommentsBySalad(app *app.App, w http.ResponseWriter, r *http.Request) {
	prompt := "getting salad comments"

	page := r.URL.Query().Get("page")
	if page == "" {
		web.ErrorResponse(w, fmt.Errorf("%s: empty page number", prompt).Error(), http.StatusBadRequest)
		return
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: converting page number to int: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}

	saladId := chi.URLParam(r, "id")
	if saladId == "" {
		web.ErrorResponse(w, fmt.Errorf("%s: empty salad id", prompt).Error(), http.StatusBadRequest)
		return
	}
	saladUuid, err := uuid.Parse(saladId)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}
	//saladId := r.URL.Query().Get("salad")
	//if saladId == "" {
	//	web.ErrorResponse(w, fmt.Errorf("%s: пустой id салата", prompt).Error(), http.StatusBadRequest)
	//	return
	//}
	//saladUuid, err := uuid.Parse(saladId)
	//if err != nil {
	//	web.ErrorResponse(w, fmt.Errorf("%s: преобразование id салата к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
	//	return
	//}

	comments, numPages, err := app.CommentService.GetAllBySaladID(r.Context(), saladUuid, pageInt)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
		return
	}

	commentsTransport := make([]web.Comment, len(comments))
	for i, comment := range comments {
		commentsTransport[i] = web.ToCommentTransport(comment)
	}

	web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"num_pages": numPages, "comments": commentsTransport})

}

func GetUserComment(app *app.App, userUuid uuid.UUID, w http.ResponseWriter, r *http.Request) {
	prompt := "getting user comment to salad"

	//userId := r.URL.Query().Get("user")
	//if userId == "" {
	//	web.ErrorResponse(w, fmt.Errorf("%s: empty user id", prompt).Error(), http.StatusBadRequest)
	//	return
	//}
	//userUuid, err := uuid.Parse(userId)
	//if err != nil {
	//	web.ErrorResponse(w, fmt.Errorf("%s: converting user id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
	//	return
	//}

	saladId := chi.URLParam(r, "id")
	if saladId == "" {
		web.ErrorResponse(w, fmt.Errorf("%s: empty salad id", prompt).Error(), http.StatusBadRequest)
		return
	}
	saladUuid, err := uuid.Parse(saladId)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
		return
	}
	//saladId := r.URL.Query().Get("salad")
	//if saladId == "" {
	//	web.ErrorResponse(w, fmt.Errorf("%s: empty salad id", prompt).Error(), http.StatusBadRequest)
	//	return
	//}
	//saladUuid, err := uuid.Parse(saladId)
	//if err != nil {
	//	web.ErrorResponse(w, fmt.Errorf("%s: converting salad id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
	//	return
	//}

	comment, err := app.CommentService.GetBySaladAndUser(r.Context(), saladUuid, userUuid)
	if err != nil {
		web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
		return
	}

	web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"comment": web.ToCommentTransport(comment)})
}

// DeleteComment godoc
//
//	@Summary		Удаление комментария
//	@ID				deleteComment
//	@Produce json
//	@Tags			comments
//	@Param id path string true "Comment`s id"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/comments/{id} [delete]
func DeleteComment(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "deleting comment"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("getting id of authorized user: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("converting user id to uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("getting id")
			web.ErrorResponse(w, fmt.Errorf("%s: empty comment id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		comment, err := app.CommentService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}
		if comment.AuthorID != userUuid { // TODO: mb admin can delete every comment
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		err = app.CommentService.DeleteById(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

func GetCommentById(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "получение комментария по id"

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		comment, err := app.CommentService.GetById(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"comment": web.ToCommentTransport(comment),
		})
	}
}

// UpdateComment godoc
//
//	@Summary		Обновление информации о комментарии
//	@ID				updateComment
//	@Produce json
//	@Tags			comments
//	@Param id path string true "Comment`s id"
//	@Param data body web.CommentUpdate true "Comment data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/comments/{id} [patch]
func UpdateComment(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "updating comment"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("getting id of authorized user: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("converting user id to uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: empty comment id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting comment id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		commentDb, err := app.CommentService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}
		if commentDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		var req web.Comment
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Text != "" {
			commentDb.Text = req.Text
		}
		if req.Rating != 0 {
			commentDb.Rating = req.Rating
		}

		err = app.CommentService.Update(r.Context(), commentDb)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// CreateComment godoc
//
//	@Summary		Создание комментария
//	@ID				createComment
//	@Produce json
//	@Tags			comments
//	@Param data body web.CommentCreate true "Comment data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/comments [post]
func CreateComment(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "creating comment"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: getting id of authorized user: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting user id to uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.Comment
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		comment := web.ToCommentModel(&req)
		comment.AuthorID = userUuid

		err = app.CommentService.Create(r.Context(), comment)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// CreateSalad godoc
//
//	@Summary		Создание салата
//	@ID				creteSalad
//	@Produce json
//	@Tags			salads
//	@Param data body web.SaladCreate true "Salad data"
//	@Success		200	{object} SaladResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads [post]
func CreateSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "creating salad"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: getting id of authorized user: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting user id to uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.Salad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		salad := web.ToSaladModel(&req)
		salad.AuthorID = userUuid

		saladUuid, err := app.SaladService.Create(r.Context(), salad)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		salad.ID = saladUuid

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"salad": web.ToSaladTransport(salad),
		})
	}
}

// UpdateSalad godoc
//
//	@Summary		Обновление информации о салате
//	@ID				updateSalad
//	@Produce json
//	@Tags			salads
//	@Param id path string true "Salad`s id"
//	@Param data body web.SaladUpdate true "Salad data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads/{id} [patch]
func UpdateSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "updating salad"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("getting id of authorized user: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("converting user id to uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: empty salad id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting salad id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladDb, err := app.SaladService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}
		if saladDb.AuthorID != userUuid { // TODO: mb admin can change every salad?
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		var req web.Salad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Name != "" {
			saladDb.Name = req.Name
		}
		if req.Description != "" {
			saladDb.Description = req.Description
		}

		err = app.SaladService.Update(r.Context(), saladDb)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// DeleteSalad godoc
//
//	@Summary		Удаление информации о салате
//	@ID				deleteSalad
//	@Produce json
//	@Tags			salads
//	@Param id path string true "Salad`s id"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads/{id} [delete]
func DeleteSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "deleting salad"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		userRole, err := web.GetStringClaimFromJWT(r.Context(), "role")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("получение роли авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id комментария", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id комментария к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		saladDb, err := app.SaladService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}
		if saladDb.AuthorID != userUuid && userRole != "admin" {
			web.ErrorResponse(w, fmt.Errorf("%s: только автор салата может изменить его", prompt).Error(), http.StatusForbidden)
			return
		}

		err = app.SaladService.DeleteById(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// CreateRecipe godoc
//
//	@Summary		Создание рецепта
//	@ID				createRecipe
//	@Produce json
//	@Tags			recipes
//	@Param data body web.RecipeCreate true "Recipe data"
//	@Success		200	{object} RecipeResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/recipes [post]
func CreateRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "creating recipe"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: getting id of authorized user: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting id to uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.Recipe
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipe := web.ToRecipeModel(&req)
		recipe.Status = dto.EditingSaladStatus

		saladDb, err := app.SaladService.GetById(r.Context(), recipe.SaladID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		recipeUuid, err := app.RecipeService.Create(r.Context(), recipe)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		recipe.ID = recipeUuid

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"recipe": web.ToRecipeTransport(recipe),
		})
	}
}

// UpdateRecipe godoc
//
//	@Summary		Обновление информации о рецепте
//	@ID				updateRecipe
//	@Produce json
//	@Tags			recipes
//	@Param id path string true "Recipe`s id"
//	@Param data body web.RecipeUpdate true "Recipe data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/recipes/{id} [patch]
func UpdateRecipe(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "обновление рецепта"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("gettign id of authorized user: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("converting id to uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		userRole, err := web.GetStringClaimFromJWT(r.Context(), "role")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("getting role of authorized user: %w", err).Error(), http.StatusBadRequest)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: empty recipe id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting recipe id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeDb, err := app.RecipeService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid && userRole != "admin" {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		var req web.Recipe
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.NumberOfServings != 0 {
			recipeDb.NumberOfServings = req.NumberOfServings
		}
		if req.TimeToCook != 0 {
			recipeDb.TimeToCook = req.TimeToCook
		}
		if req.Status != 0 {
			if userRole == domain.DefaultRole {
				if req.Status == dto.EditingSaladStatus ||
					req.Status == dto.ModerationSaladStatus ||
					req.Status == dto.StoredSaladStatus {
					recipeDb.Status = req.Status
				}
			} else if userRole == "admin" {
				recipeDb.Status = req.Status
			}
		}

		err = app.RecipeService.Update(r.Context(), recipeDb)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// LinkTypeToSalad godoc
//
//	@Summary		Добавление типа к салату
//	@ID				linkType
//	@Produce json
//	@Tags			types
//	@Param id path string true "Salad`s id"
//	@Param data body web.LinkSaladType true "Link data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads/{id}/types [post]
func LinkTypeToSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "добавление типа к салату"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: getting id of authorized: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		var req web.LinkSaladType
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		saladUuid, typeUuid := web.ToLinkSaladTypeModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		err = app.SaladTypeService.Link(r.Context(), saladUuid, typeUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// UnlinkTypeFromSalad godoc
//
//	@Summary		Удаление привязки типа к салату
//	@ID				unlinkType
//	@Produce json
//	@Tags			types
//	@Param id path string true "Salad`s id"
//	@Param data body web.LinkSaladType true "Link data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads/{id}/types [delete]
func UnlinkTypeFromSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "unlinking type from salad"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: getting id of authorized user: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: converting id to uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		var req web.LinkSaladType
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		saladUuid, typeUuid := web.ToLinkSaladTypeModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		err = app.SaladTypeService.Unlink(r.Context(), saladUuid, typeUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusOK, nil)
	}
}

// LinkIngredientToSalad godoc
//
//	@Summary		Добавление ингредиента к рецепту
//	@ID				linkIngredient
//	@Produce json
//	@Tags			ingredients
//	@Param id path string true "Salad`s id"
//	@Param data body web.LinkIngredientSalad true "Link data"
//	@Success		200	{object} LinkIngredientSaladResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads/{id}/ingredients [post]
func LinkIngredientToSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "adding ingredient to salad"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.LinkIngredientSalad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipeUuid, saladUuid, ingredientUuid, measurementuuid, amount := web.ToLinkIngredientModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		linkUuid, err := app.IngredientService.Link(r.Context(), recipeUuid, ingredientUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		err = app.MeasurementService.UpdateLink(r.Context(), linkUuid, measurementuuid, amount)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusOK, map[string]string{
			"link_id": linkUuid.String(),
		})
	}
}

// UnlinkIngredientFromSalad godoc
//
//	@Summary		Удаление ингредиента из рецепта
//	@ID				unlinkIngredient
//	@Produce json
//	@Tags			ingredients
//	@Param id path string true "Salad`s id"
//	@Param data body web.LinkIngredientSalad true "Link data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/salads/{id}/ingredients [delete]
func UnlinkIngredientFromSalad(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "открепление ингредиента от салата"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.LinkIngredientSalad
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipeUuid, saladUuid, ingredientUuid, _, _ := web.ToLinkIngredientModel(&req)

		saladDb, err := app.SaladService.GetById(r.Context(), saladUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		err = app.IngredientService.Unlink(r.Context(), recipeUuid, ingredientUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// GetAllMeasurements godoc
//
//	@Summary		Получение списка единиц измерения
//	@ID				getAllMeasurements
//	@Produce json
//	@Description	Получение списка всех доступных единиц измерения
//	@Tags			measurements
//	@Success		200	{object} MeasurementsResponse
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Router			/measurements [get]
func GetAllMeasurements(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "getting all measurement units"

		measurements, err := app.MeasurementService.GetAll(r.Context())
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		measurementsTransport := make([]web.Measurement, len(measurements))
		for i, measurement := range measurements {
			measurementsTransport[i] = web.ToMeasurementTransport(measurement)
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{"measurements": measurementsTransport})
	}
}

// CreateRecipeStep godoc
//
//	@Summary		Создание шага рецепта
//	@ID				createRecipeStep
//	@Produce json
//	@Tags			recipeSteps
//	@Param data body web.RecipeStepCreate true "Recipe step data"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/recipeSteps [post]
func CreateRecipeStep(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "создание шага рецепта"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: получение id авторизованного пользователя: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id пользователя к uuid: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		var req web.RecipeStep
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}

		recipeStep := web.ToRecipeStepModel(&req)

		recipeDb, err := app.RecipeService.GetById(r.Context(), recipeStep.RecipeID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		err = app.RecipeStepService.Create(r.Context(), recipeStep)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// DeleteRecipeStep godoc
//
//	@Summary		Удаление шага рецепта
//	@ID				deleteRecipeStep
//	@Produce json
//	@Tags			recipeSteps
//	@Param id path string true "Recipe step`s id"
//	@Success		204
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/recipeSteps/{id} [delete]
func DeleteRecipeStep(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "удаление шага рецепта"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			app.Logger.Infof("getting id")
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeStep, err := app.RecipeStepService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}
		recipeDb, err := app.RecipeService.GetById(r.Context(), recipeStep.RecipeID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		err = app.RecipeStepService.DeleteById(r.Context(), idUuid)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusNoContent, nil)
	}
}

// UpdateRecipeStep godoc
//
//	@Summary		Обновление информации о шаге рецепта
//	@ID				updateRecipeStep
//	@Produce json
//	@Tags			recipeSteps
//	@Param id path string true "Recipe step`s id"
//	@Param data body web.RecipeStepUpdate true "Recipe data"
//	@Success		200 {object} web.RecipeStep
//	@Failure		400	{object} web.ErrorResponseStruct
//	@Failure		401
//	@Failure		403
//	@Failure		404
//	@Failure 		500 {object} web.ErrorResponseStruct
//	@Security		BearerAuth
//	@Router			/recipeSteps/{id} [patch]
func UpdateRecipeStep(app *app.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prompt := "обновление шага рецепта"

		userId, err := web.GetStringClaimFromJWT(r.Context(), "user_id")
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("получение id авторизованного пользователя: %w", err).Error(), http.StatusBadRequest)
			return
		}
		userUuid, err := uuid.Parse(userId)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("преобразование id пользователя к uuid: %w", err).Error(), http.StatusInternalServerError)
			return
		}

		id := chi.URLParam(r, "id")
		if id == "" {
			web.ErrorResponse(w, fmt.Errorf("%s: пустой id комментария", prompt).Error(), http.StatusBadRequest)
			return
		}
		idUuid, err := uuid.Parse(id)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: преобразование id комментария к uuid: %w", prompt, err).Error(), http.StatusBadRequest)
			return
		}

		recipeStep, err := app.RecipeStepService.GetById(r.Context(), idUuid)
		if err != nil {
			if errors.Is(err, errs.ErrNotFound) {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusNotFound)
			} else {
				web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			}
			return
		}
		recipeDb, err := app.RecipeService.GetById(r.Context(), recipeStep.RecipeID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		saladDb, err := app.SaladService.GetById(r.Context(), recipeDb.SaladID)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}
		if saladDb.AuthorID != userUuid {
			web.ErrorResponse(w, fmt.Errorf("%s: access denied", prompt).Error(), http.StatusForbidden)
			return
		}

		var req web.RecipeStep
		err = json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			web.ErrorResponse(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Name != "" {
			recipeStep.Name = req.Name
		}
		if req.Description != "" {
			recipeStep.Description = req.Description
		}
		// TODO: should i change recipe step nums (?)
		if req.StepNum != 0 {
			recipeStep.StepNum = req.StepNum
		}

		err = app.RecipeStepService.Update(r.Context(), recipeStep)
		if err != nil {
			web.ErrorResponse(w, fmt.Errorf("%s: %w", prompt, err).Error(), http.StatusInternalServerError)
			return
		}

		web.SuccessResponse(w, http.StatusOK, map[string]interface{}{
			"step": web.ToRecipeStepTransport(recipeStep),
		})
	}
}
