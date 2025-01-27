import axios from 'axios';

const API_URL = 'http://localhost:8081/salads';

class SaladService {
    getSalads(page, ingredients, saladTypes, rate) {
        return axios.get(API_URL + `?page=${page}&minRate=${rate}&ingredients=${ingredients.join("&ingredients=")}&types=${saladTypes.join("&types=")}`,
        {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    getSaladsByStatus(page, ingredients, saladTypes, rate, status) {
        return axios.get(API_URL + `/status?page=${page}&minRate=${rate}&ingredients=${ingredients.join("&ingredients=")}&types=${saladTypes.join("&types=")}&status=${status}`,
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            })
    }

    getUserSalads(page, ingredients, saladTypes) {
        return axios.get(API_URL + `/user?page=${page}&ingredients=${ingredients.join("&ingredients=")}&types=${saladTypes.join("&types=")}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
            }
        })
    }

    getUserRated(page) {
        return axios.get(API_URL + `/userRated?page=${page}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
            }
        })
    }

    getSaladDetails(id) {
        return axios.get(API_URL + `/${id}`, {
            headers: {
                'Authorization': `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    createSalad(salad) {
        return axios.post(
            API_URL + `/create`,
            {
                name: salad.name,
                description: salad.description
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    updateSalad(id, salad) {
        return axios.patch(
            API_URL + `/${id}/update`,
            {
                name: salad.name,
                description: salad.description,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    deleteSalad(id) {
        return axios.delete(API_URL + `/${id}/delete`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
            }
        })
    }
}

export default new SaladService();