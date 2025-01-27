import axios from 'axios';

const API_URL = 'http://localhost:8081/recipe';

class RecipeService {
    // TODO
    // getSaladsSentToModeration(page) {
    //     return axios.get(API_URL + `/empty?page=${page}`, {
    //         headers: {
    //             Authorization: `Bearer ${localStorage.getItem('user')}`
    //         }
    //     })
    // }

    getRecipeRating(saladId) {
        return axios.get(API_URL + `/${saladId}/rating`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    getRecipeDetails(saladId) {
        return axios.get(API_URL + `/${saladId}`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    createRecipe(recipe) {
        return axios.post(
            API_URL + `/create`,
            {
                salad_id: recipe.salad_id,
                number_of_servings: recipe.number_of_servings,
                time_to_cook: recipe.time_to_cook
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    updateRecipe(id, recipe) {
        return axios.patch(
            API_URL + `/${id}/update`,
            {
                number_of_servings: recipe.number_of_servings,
                time_to_cook: recipe.time_to_cook,
                status: recipe.status
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    // TODO
    // updateSalad(id, user) {
    //     return axios.patch(
    //         API_URL + `/${id}/update`,
    //         {
    //             id: user.id,
    //             username: user.username,
    //             full_name: user.fullName,
    //             birthday: user.birthday,
    //             gender: user.gender,
    //             city: user.city
    //         },
    //         {
    //             headers: {
    //                 Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
    //             }
    //         }
    //     )
    // }
}

export default new RecipeService();