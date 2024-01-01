import { createRouter , createWebHistory } from "vue-router";
import Body from "./../components/Body.vue"
import UserLogin from "./../components/UserLogin.vue"
import AdminLogin from "./../components/AdminLogin.vue"
import NextPage from "./../components/NextPage.vue"
import HomePage from "./../components/HomePage.vue"
// import HomePage from "./..components/HomePage.vue"
// import Login from "./../components/Login.vue"

const router  = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes:[
        {
            path: "/",
            name: "HomePage",
            component: HomePage
        },
        {
            path: "/userlogin",
            name: "Login",
            component: UserLogin
        },
        {
            path:"/adminlogin",
            name:"AdminLogin",
            component: AdminLogin
        },
        {
            path:"/next-page",
            name:"NextPage",
            component: NextPage
        }
    ]
})

export default router