<template>
  <div>
    <template v-if="user.name">
      <h1>Ваш профиль</h1>
      <div>
        <h2>Nickname: {{ user.name }}</h2>
        <p>Роль: {{ user.role }}</p>
        <p>email: {{ user.email.Address }}</p>
      </div>
    </template>
    <template v-else>
      <h1>Ваш профиль не заполнен. Обратитесь к администратору.</h1>
    </template>
  </div>
</template>

<script lang="ts">
import UserService from "../services/user.service.ts";
import ButtonGroup from 'primevue/buttongroup';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Utils from '../services/auth-header.ts';

export default {
  name: 'ProfilePage',
  components: {
    Button,
    Column,
    ButtonGroup
  },
  data() {
    return {
      user: {},
      isAuthValue: null,
      role: 'guest'
    }
  },
  created() {
    if (Utils.isAuth()) {
      this.role = Utils.getUserRoleJWT();
    }
    this.fetchUserDetails()
  },
  methods: {
    fetchUserDetails() {
      const userId = Utils.getUserIdJWT();
      UserService.getUserDetails(userId)
          .then(response => {
            const user = response.data.data.user
            this.user = user
          })
          .catch(error => {
            console.error('Error fetching user details:', error)
          })
    },
    isAuth() {
      this.isAuthValue = Utils.isAuth();
    },
  }
}
</script>