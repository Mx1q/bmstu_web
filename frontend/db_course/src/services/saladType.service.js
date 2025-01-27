import axios from 'axios';

const API_URL = 'http://localhost:8081/types';

class SaladTypesService {
    getSaladTypes(saladId) {
        return axios.get(API_URL + `/${saladId}`, {
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

    link(saladId, typeId) {
        return axios.post( API_URL + `/link`,
            {
                salad_id: saladId,
                type_id: typeId,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            })
    }

    unlink(saladId, typeId) {
        return axios.patch( API_URL + `/unlink`,
            {
                salad_id: saladId,
                type_id: typeId,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            })
    }
}

export default new SaladTypesService();