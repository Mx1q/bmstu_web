<template>
  <Toast />

  <div id="app">
    <nav class="navbar navbar-expand navbar-dark bg-dark">
      <a href="/legacy" class="navbar-brand">Salads</a>

      <div class="navbar-nav mr-auto">
        <li class="nav-item">
          <router-link to="/legacy/salads" class="nav-link">
            <font-awesome-icon icon="utensils" /> Салаты
          </router-link>
        </li>
      </div>

      <div v-if="!currentUser" class="navbar-nav ml-auto">
        <li class="nav-item">
          <router-link to="/legacy/register" class="nav-link">
            <font-awesome-icon icon="user-plus" /> Зарегистрироваться
          </router-link>
        </li>
        <li class="nav-item">
          <router-link to="/legacy/login" class="nav-link">
            <font-awesome-icon icon="sign-in-alt" /> Войти
          </router-link>
        </li>
      </div>

      <div v-if="currentUser" class="navbar-nav ml-auto">
        <li class="nav-item">
          <router-link to="/legacy/ratedSalads" class="nav-link">
            <i class="pi pi-heart"></i> Оцененные
          </router-link>
        </li>

        <li class="nav-item">
          <router-link to="/legacy/mySalads" class="nav-link">
            <font-awesome-icon icon="utensils" /> Мои салаты
          </router-link>
        </li>

        <li class="nav-item">
          <button class="nav-link" @click.prevent="createEmptySalad">
            <font-awesome-icon icon="plus" /> Создать салат
          </button>
        </li>

        <li class="nav-item" v-if="user !== null">
          <router-link to="/legacy/profile" class="nav-link">
            <font-awesome-icon icon="user" /> {{ user.name }}
          </router-link>
        </li>

        <li class="nav-item">
          <button class="nav-link" @click.prevent="logOut">
            <font-awesome-icon icon="sign-out-alt" /> Выход
          </button>
        </li>
      </div>
    </nav>

    <div class="container">
      <router-view />
    </div>
  </div>
</template>

<script>
import Toast from 'primevue/toast';
import Utils from "@/services/auth-header.js";
import UserService from "@/services/user.service.js";
import SaladService from "@/services/salad.service.js";
import SaladModel from "@/models/SaladModel.js";
import { useToast } from "primevue/usetoast";

export default {
  name: 'Nav',
  components: {
    Toast,
  },
  data() {
    return {
      toast: useToast(),

      user: null,
      isAuthValue: null,
      role: 'guest'
    }
  },
  created() {
    console.log("created")
    if (Utils.isAuth()) {
      this.role = Utils.getUserRoleJWT();
      this.fetchUserDetails()
    }
  },
  // updated() {
  //   if (Utils.isAuth()) {
  //     this.role = Utils.getUserRoleJWT();
  //     this.fetchUserDetails()
  //   }
  // },

  computed: {
    currentUser() {
      return this.$store.state.auth.user;
    },
  },
  methods: {
    logOut() {
      this.$store.dispatch('auth/logout');
      this.$router.push('/legacy/login');
    },

    fetchUserDetails() {
      if (Utils.isAuth()) {
        const userId = Utils.getUserIdJWT();
        UserService.getUserDetails(userId)
            .then(response => {
              this.user = response.data.data.user
            })
            .catch(error => {
              console.error('Error fetching user details:', error)
            })
      }
    },

    createEmptySalad() {
      let isAuth = Utils.isAuth()

      if (isAuth) {
        const salad = new SaladModel("Название салата", "Описание салата")
        SaladService.createSalad(salad)
            .then(response => {
              this.toast.add({ severity: 'success', summary: 'Успех', detail: 'Салат успешно создан', life: 3000 });
              const saladId = response.data.data.salad.id
              this.$router.push(`/legacy/salads/${saladId}`)
                  .then(() => { this.$router.go(0) })
            })
            .catch(error => {
              this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `Произошла ошибка при создании салата: ${error.response.data.error}`, life: 3000 });
            })

      } else {
        console.log("Auth required to create salad")
      }
    }
  }
};
</script>