import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import "bootstrap";
import "bootstrap/dist/css/bootstrap.min.css";
import 'primevue/resources/primevue.min.css'
import PrimeVue from 'primevue/config';
import { FontAwesomeIcon } from './plugins/font-awesome'
import "primevue/resources/themes/aura-light-green/theme.css";
import 'primeicons/primeicons.css';
import "primeflex/primeflex.css";
import ToastService from 'primevue/toastservice';
import ConfirmationService from 'primevue/confirmationservice';

const app = createApp(App)
    .use(router)
    .use(store)
    .use(ToastService)
    .use(ConfirmationService)
    .component("font-awesome-icon", FontAwesomeIcon)

app.use(PrimeVue)

app.mount('#app')
