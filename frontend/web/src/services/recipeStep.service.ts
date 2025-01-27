import axios from 'axios';

const API_URL = 'http://localhost:8081/recipeSteps';

class RecipeStepService {
    getRecipeSteps(recipeId) {
        return axios.get(API_URL + `/${recipeId}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    create(recipeStep) {
        return axios.post(
            API_URL + `/create`,
            {
                recipe_id: recipeStep.recipe_id,
                name: recipeStep.name,
                description: recipeStep.description,
                step_num: recipeStep.step_num,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    deleteById(id) {
        return axios.delete(
            API_URL + `/${id}/delete`, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    update(id, step) {
        return axios.patch(
            API_URL + `/${id}/update`,
            {
                name: step.name,
                description: step.description,
                step_num: step.step_num,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    // TODO
    // getSaladsSentToModeration(page) {
    //     return axios.get(API_URL + `/empty?page=${page}`, {
    //         headers: {
    //             Authorization: `Bearer ${localStorage.getItem('user')}`
    //         }
    //     })
    // }

    // getRecipeDetails(saladId) {
    //     return axios.get(API_URL + `/${saladId}`, {
    //         headers: {
    //             'Authorization': `Bearer ${localStorage.getItem('user')}`
    //         }
    //     })
    // }

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

export default new RecipeStepService();