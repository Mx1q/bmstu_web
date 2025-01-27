import { createWebHistory, createRouter } from "vue-router";
// import Home from "./components/Home.vue";
import Login from "./components/Login.vue";
import Register from "./components/Register.vue";
import ProfilePage from "./components/ProfilePage.vue";
import NotFound from "./components/NotFound.vue";
import Salads from "./components/Salads.vue";
import Salad from './components/Salad.vue';
import SaladCreate from "@/components/SaladCreateChild.vue";
import SaladUpdateChild from "@/components/SaladUpdateChild.vue";
import UserSalads from "@/components/UserSalads.vue";
import UserRated from "@/components/UserRated.vue";
// import UpdateEntrepreneurPage from './components/UpdateEntrepreneurPage.vue';
// import UpdateCompanyPage from './components/UpdateCompanyPage.vue';
// import CompanyPage from './components/CompanyPage.vue';
// import CompaniesPage from './components/CompaniesPage.vue';
// import CompanyCreatePage from './components/CompanyCreatePage.vue';
// import UserSkillsPage from './components/UserSkillsPage.vue';
// import SkillsPage from './components/SkillsPage.vue';
// import UpdateSkillPage from './components/UpdateSkillPage.vue';
// import CreateSkillPage from "./components/CreateSkillPage.vue";
// import ActivityFieldsPage from "./components/ActivityFieldsPage.vue";
// import CreateActivityFieldPage from './components/CreateActivityFieldPage.vue';
// import UpdateActivityFieldPage from './components/UpdateActivityFieldPage.vue';
// import CreateFinReportPage from './components/CreateFinReportPage.vue';
// import UpdateFinReportPage from './components/UpdateFinReportPage.vue';
// import CompanyFinancialsPage from './components/CompanyFinancialsPage.vue';

const routes = [
    {
        path: "/login",
        component: Login,
    },
    {
        path: "/register",
        component: Register,
    },
    {
        path: "/profile",
        name: "ProfilePage",
        component: ProfilePage,
    },
    {
        path: "/",
        component: Salads,
    },
    {
        path: "/salads",
        component: Salads,
    },
    {
        path: '/salads/:id',
        name: 'SaladPage',
        component: Salad,
        props: true
    },
    {
        path: '/mySalads',
        name: 'userSalads',
        component: UserSalads,
        props: true
    },
    {
        path: '/ratedSalads',
        name: 'ratedSalads',
        component: UserRated,
        props: true
    },

    // TODO: place to add new paths

    {
        path: "/:pathMatch(.*)*",
        component: NotFound
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const publicPages = ['/login', '/register', '/salads', '/'];
    const authRequired = !publicPages.includes(to.path) && !to.path.startsWith('/salads/');
    const loggedIn = localStorage.getItem('user');

    if (authRequired && !loggedIn) {
        next('/login');
    } else {
        next();
    }
});

export default router;