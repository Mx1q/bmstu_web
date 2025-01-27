import { createWebHistory, createRouter } from "vue-router";
// import Home from "./components/Home.vue"; // TODO
import Login from "./components/Login.vue";
import Register from "./components/Register.vue";
import ProfilePage from "./components/ProfilePage.vue";
import NotFound from "./components/NotFound.vue";
import Salads from "./components/Salads.vue"; // TODO
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
        path: "/legacy/login",
        component: Login,
    },
    {
        path: "/legacy/register",
        component: Register,
    },
    {
        path: "/legacy/profile",
        name: "ProfilePage",
        component: ProfilePage,
    },
    {
        path: "/legacy",
        component: Salads,
    },
    {
        path: "/legacy/salads",
        component: Salads,
    },
    {
        path: '/legacy/salads/:id',
        name: 'SaladPage',
        component: Salad,
        props: true
    },
    {
        path: '/legacy/mySalads',
        name: 'userSalads',
        component: UserSalads,
        props: true
    },
    {
        path: '/legacy/ratedSalads',
        name: 'ratedSalads',
        component: UserRated,
        props: true
    },

    // TODO: place to add new paths

    {
        path: "/legacy/:pathMatch(.*)*",
        component: NotFound
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    const publicPages = ['/legacy/login', '/legacy/register', '/legacy/salads', '/legacy'];
    const authRequired = !publicPages.includes(to.path) && !to.path.startsWith('/legacy/salads/');
    const loggedIn = localStorage.getItem('user');

    if (authRequired && !loggedIn) {
        next('/legacy/login');
    } else {
        next();
    }
});

export default router;