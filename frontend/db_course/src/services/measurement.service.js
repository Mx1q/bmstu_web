import axios from 'axios';

const API_URL = 'http://localhost:8081/measurements';

class MeasurementService {
    getMeasurementByRecipe(ingredientId, recipeId) {
        return axios.get(API_URL + `?ingredient=${ingredientId}&recipe=${recipeId}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    getAll() {
        return axios.get(API_URL + `/all`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }
}

export default new MeasurementService();