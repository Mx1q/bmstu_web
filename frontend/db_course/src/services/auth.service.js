import axios from 'axios';

const API_URL = 'http://localhost:8081/';

class AuthService {
    login(user) {
        return axios
            .post(API_URL + 'login', {
                login: user.username,
                password: user.password
            })
            .then(response => {
                if (response.data.data.token) {
                    axios.defaults.headers.common['Authorization'] = 'Bearer ' + response.data.data.token;
                    localStorage.setItem('user', JSON.stringify(response.data.data.token));
                }

                return response.data.data;
            });
    }

    logout() {
        localStorage.removeItem('user');
    }

    register(user) {
        return axios.post(API_URL + 'signup', {
            name: user.name,
            username: user.username,
            password: user.password,
            email: user.email
        })
        .then(response => {
            if (response.data.data.token) {
                axios.defaults.headers.common['Authorization'] = 'Bearer ' + response.data.data.token;
                localStorage.setItem('user', JSON.stringify(response.data.data.token));
            }

            return response.data.data;
        });
    }
}

export default new AuthService();