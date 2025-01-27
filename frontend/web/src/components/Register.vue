<template>
  <Toast />

  <div class="col-md-12">
    <div class="card card-container">

      <div class="flex flex-row justify-content-center">
        <img
            id="profile-img"
            src="//ssl.gstatic.com/accounts/ui/avatar_2x.png"
            class="profile-img-card"
        />
      </div>

      <Form @submit="handleRegister" :validation-schema="schema">
        <div v-if="!successful">
          <div class="form-group">
            <label for="name">Имя</label>
            <Field name="name" type="text" class="form-control" />
            <ErrorMessage name="name" class="error-feedback" />
          </div>

          <div class="form-group">
            <label for="username">Логин</label>
            <Field name="username" type="text" class="form-control" />
            <ErrorMessage name="username" class="error-feedback" />
          </div>

          <div class="form-group">
            <label for="email">E-mail</label>
            <Field name="email" type="text" class="form-control" />
            <ErrorMessage name="email" class="error-feedback" />
          </div>

          <div class="form-group">
            <label for="password">Пароль</label>
            <Field name="password" type="password" class="form-control" />
            <ErrorMessage name="password" class="error-feedback" />
          </div>
          <div class="form-group">
            <label for="reenter_password">Подтверждение пароля</label>
            <Field name="reenter_password" type="password" class="form-control" />
            <ErrorMessage name="reenter_password" class="error-feedback" />
          </div>

          <div class="form-group flex flex-row justify-content-center p-2">
            <button class="btn btn-primary btn-block" style="background-color: #04AA6D;" :disabled="loading">
              <span
                v-show="loading"
                class="spinner-border spinner-border-sm"
              ></span>
              Зарегистрироваться
            </button>
          </div>
        </div>
      </Form>
    </div>
  </div>
</template>

<script lang="ts">
import { Form, Field, ErrorMessage } from "vee-validate";
import * as yup from "yup";
import { useToast } from "primevue/usetoast";
import Toast from 'primevue/toast';

export default {
  name: "Register",
  components: {
    Form,
    Field,
    ErrorMessage,
    Toast
  },
  data() {
    const schema = yup.object().shape({
      name: yup
          .string()
          .required("Введите имя"),
      username: yup
        .string()
        .required("Введите логин"),
        // .min(3, "Логин должен быть не менее 3 символов!")
        // .max(30, "Логин должно быть не более 30 символов!"),
      email: yup
          .string()
          .required("Введите почту")
          .email("Невалидная почта"),
      password: yup
        .string()
        .required("Введите пароль"),
        // .min(6, "Пароль должен быть не менее 6 символов!")
        // .max(40, "Пароль должен быть не более 40 символов!"),
      reenter_password: yup
        .string()
        .required("Введите подтверждение пароля")
        // .min(6, "Пароль должен быть не менее 6 символов!")
        // .max(40, "Пароль должен быть не более 40 символов!")
        .oneOf([yup.ref('password'), null], 'Пароли не совпадают'),
    });

    return {
      successful: false,
      loading: false,
      message: "",
      schema,
      toast: useToast(),
    };
  },
  computed: {
    loggedIn() {
      return this.$store.state.auth.status.loggedIn;
    },
  },
  mounted() {
    if (this.loggedIn) {
      this.$router.push("/profile");
    }
  },
  methods: {
    handleRegister(user) {
      this.message = "";
      this.successful = false;
      this.loading = true;

      this.$store.dispatch("auth/register", user).then(
        () => {
          // this.message = data.message;
          // this.successful = true;
          // this.loading = false;
          this.$router.push("/salads")
              .then(() => { this.$router.go(0) })

        },
        (error) => {
          this.successful = false;
          this.loading = false;

          let errText = error.response.data.error
          if (errText.includes("user.email") || errText.includes("user_email_key")) {
            this.message = "Введенная почта уже используется."
          } else if (errText.includes("user.login") || errText.includes("user_login_key")) {
            this.message = "Введенный логин уже используется."
          }

          this.toast.add({ severity: 'error', summary: 'Ошибка', detail: `${this.message}`, life: 3000 });
        }
      );
    },
  },
};
</script>

<style scoped>
label {
  display: block;
  margin-top: 10px;
}

.card-container.card {
  max-width: 350px !important;
  padding: 40px 40px;
}

.card {
  background-color: #f7f7f7;
  padding: 20px 25px 30px;
  margin: 0 auto 25px;
  margin-top: 50px;
  -moz-border-radius: 2px;
  -webkit-border-radius: 2px;
  border-radius: 2px;
  -moz-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  -webkit-box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
  box-shadow: 0px 2px 2px rgba(0, 0, 0, 0.3);
}

.profile-img-card {
  width: 96px;
  height: 96px;
  margin: 0 auto 10px;
  display: block;
  -moz-border-radius: 50%;
  -webkit-border-radius: 50%;
  border-radius: 50%;
}

.error-feedback {
  color: red;
}
</style>