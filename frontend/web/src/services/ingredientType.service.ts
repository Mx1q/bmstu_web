import axios from 'axios';

const API_URL = 'http://localhost:8081/ingredientTypes';

class IngredientTypeService {
    getType(id) {
        return axios.get(API_URL + `/${id}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }
}

export default new IngredientTypeService();