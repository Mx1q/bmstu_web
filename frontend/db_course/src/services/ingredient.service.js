import axios from 'axios';

const API_URL = 'http://localhost:8081/ingredients';

class IngredientService {
    getRecipeIngredients(recipeId) {
        return axios.get(API_URL + `/${recipeId}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    getAllByPage(page) {
        return axios.get(API_URL + `/?page=${page}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    link(recipeId, saladId, ingredientId, measurementId, amount) {
        return axios.post( API_URL + `/link`,
            {
                salad_id: saladId,
                recipe_id: recipeId,
                ingredient_id: ingredientId,
                measurement_id: measurementId,
                amount: amount,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            })
    }

    unlink(recipeId, saladId, ingredientId) {
        return axios.patch( API_URL + `/unlink`,
            {
                salad_id: saladId,
                recipe_id: recipeId,
                ingredient_id: ingredientId,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            })
    }
}

export default new IngredientService();