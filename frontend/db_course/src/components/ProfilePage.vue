<template>
  <div>
    <template v-if="user.name">
      <h1>Ваш профиль</h1>
      <div>
        <h2>Nickname: {{ user.name }}</h2>
        <p>Роль: {{ user.role }}</p>
        <p>email: {{ user.email.Address }}</p>
<!--        <p>Дата рождения: {{ formatBirthday(entrepreneur.birthday) }}</p>-->
<!--        <p>Город: {{ entrepreneur.city }}</p>-->
<!--        <p>Пол: {{ formatGender(entrepreneur.gender) }}</p>-->
<!--        <p>Рейтинг: {{ entrepreneur.rating }}</p>-->
      </div>
    </template>
    <template v-else>
      <h1>Ваш профиль не заполнен. Обратитесь к администратору.</h1>
    </template>
  </div>

<!--  <div class="settings">-->
<!--    <ButtonGroup>-->
<!--      <RouterLink :to="`/entrepreneurs/${entrepreneur.id}/contacts`"><Button label="Средства связи" icon="pi pi-address-book"></Button></RouterLink>-->
<!--      <RouterLink :to="`/entrepreneurs/${entrepreneur.id}/skills`"><Button label="Навыки" icon="pi pi-bolt"></Button></RouterLink>-->
<!--      <RouterLink :to="`/entrepreneurs/${entrepreneur.id}/companies`"><Button label="Компании" icon="pi pi-building"></Button></RouterLink>-->
<!--    </ButtonGroup>-->
<!--  </div>-->

<!--  <div v-if="role=='admin'">-->
<!--    <ButtonGroup>-->
<!--      <RouterLink :to="`/activity-fields`"><Button label="Все сферы деятельности" icon="pi pi-address-book"></Button></RouterLink>-->
<!--      <RouterLink :to="`/skills`"><Button label="Все навыки" icon="pi pi-bolt"></Button></RouterLink>-->
<!--    </ButtonGroup>-->
<!--  </div>-->
</template>

<script>
import UserService from "../services/user.service.js";
import ButtonGroup from 'primevue/buttongroup';
import Column from 'primevue/column';
import Button from 'primevue/button';
import Utils from '../services/auth-header';

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