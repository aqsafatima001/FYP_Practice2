import { createRouter , createWebHistory } from "vue-router";
import Body from "./../components/Body.vue"
import UserLogin from "./../components/UserLogin.vue"
import AdminLogin from "./../components/AdminLogin.vue"
import NextPage from "./../components/NextPage.vue"
import HomePage from "./../components/HomePage.vue"
import UserRegistration from "./../components/UserRegistration.vue"
import UserRegistrationTesting from "./../components/UserRegistrationTesting.vue"
import OtpVerification from "./../components/OtpVerification.vue"
import UserDashboard from "./../components/UserDashboard.vue"
import AdminDashboard from "./../components/AdminDashboard.vue"
import RegistrationSuccess from "./../components/RegistrationSuccess.vue"
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
            path: "/UserRegistration",
            name: "UserRegistration",
            component: UserRegistration
        },
        {
            path: "/UserRegistrationTesting",
            name: "UserRegistrationTesting",
            component: UserRegistrationTesting
        },
        {
            path: "/OtpVerification",
            name: "OtpVerification",
            component: OtpVerification
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
        },
        {
            path:"/userdashboard",
            name:"UserDashboard",
            component: UserDashboard
        },
        {
            path:"/admindashboard",
            name:"AdminDashboard",
            component: AdminDashboard
        },
        {
            path:"/registration-success",
            name:"RegistrationSuccess",
            component:RegistrationSuccess
        }
    ]
})

export default router