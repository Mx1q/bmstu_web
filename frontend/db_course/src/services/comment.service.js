import axios from 'axios';

import VueJwtDecode from 'vue-jwt-decode'; // FIXME: remove unused

const API_URL = 'http://localhost:8081/comments';

class CommentService {
    getAllBySaladId(saladId, page) {
        return axios.get(API_URL + `/?salad=${saladId}&page=${page}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    getUserComment(saladId, userId) {
        return axios.get(API_URL + `/userSalad?salad=${saladId}&user=${userId}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    deleteComment(commentId) {
        return axios.delete(
            API_URL + `/${commentId}/delete`, {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    getById(id) {
        return axios.get(API_URL + `/${id}`, {
            headers: {
                Authorization: `Bearer ${localStorage.getItem('user')}`
            }
        })
    }

    updateComment(id, comment) {
        return axios.patch(
            API_URL + `/${id}/update`,
            {
                text: comment.text,
                rating: comment.rating,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }

    createComment(comment) {
        return axios.post(
            API_URL + `/create`,
            {
                salad_id: comment.salad_id,
                text: comment.text,
                rating: comment.rating,
            },
            {
                headers: {
                    Authorization: `Bearer ${localStorage.getItem('user').replace(/"/g, '')}`
                }
            }
        )
    }
}

export default new CommentService();