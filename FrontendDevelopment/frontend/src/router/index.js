import { createRouter , createWebHistory } from "vue-router";
// import Body from "./../components/Body.vue"
import UserLogin from "./../components/UserLogin.vue"
import AdminLogin from "./../components/AdminLogin.vue"
import NextPage from "./../components/NextPage.vue"
import HomePage from "./../components/HomePage.vue"
import UserRegistration from "./../components/UserRegistration.vue"
import UserRegistrationTesting from "./../components/UserRegistrationTesting.vue"
import AdminRegistration from "./../components/AdminRegistration.vue"
import OtpVerification from "./../components/OtpVerification.vue"
import AdminOtpVerification from "./../components/AdminOtpVerification.vue"
import UserDashboard from "./../components/UserDashboard.vue"
import AdminDashboard from "./../components/AdminDashboard.vue"
import RegistrationSuccess from "./../components/RegistrationSuccess.vue"
import PendingUserRequest from "./../components/PendingUserRequest.vue"
// import HomePage from "./..components/HomePage.vue"
// import Login from "./../components/Login.vue"


// USER CREATION IN CLUSTER
import UserCreationInCluster from "./../components/UserCreationInCluster/UserCreationInCluster.vue";

import UserRegistrationTesting2 from "./../components/UserRegistrationTesting2.vue"
import OtpVerificationWithAdditionalInformation from "./../components/OtpVerificationWithAdditionalInformation.vue"

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
            path: "/AdminRegistration",
            name: "AdminRegistration",
            component: AdminRegistration
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
            path: "/OtpVerificationWithAdditionalInformation",
            name: "OtpVerificationWithAdditionalInformation",
            component: OtpVerificationWithAdditionalInformation
        },
        {
            path: "/AdminOtpVerification",
            name: "AdminOtpVerification",
            component: AdminOtpVerification
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
        },
        {
            path:"/PendingUserRequest",
            name:"PendingUserRequest",
            component:PendingUserRequest
        },

        // USER CREATION IN CLUSTER
        {
            path:"/UserCreationInCluster",
            name:"UserCreationInCluster",
            component:UserCreationInCluster
        },



        {
            path: "/UserRegistrationTesting2",
            name: "UserRegistrationTesting2",
            component: UserRegistrationTesting2
        },
    ]
})

export default router