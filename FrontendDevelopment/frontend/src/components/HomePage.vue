<script setup>
import Header from "./../components/Header.vue"
import Footer from "./../components/Footer.vue"
</script>

<template>
   
<div>
    <Header/>
    <div class="container1">
        <h1>Welcome to NCP Big Data Facility</h1>
        
        <div class="Button-Container1">
            <router-link to="/userlogin" class="button">Login as User</router-link>
            <router-link to="/adminlogin" class="button">Login as Admin</router-link>
        </div>

        <p>Don't have an account?</p>
        <div class="Button-Container2">
            <router-link to="/UserRegistrationTesting" class="button">Register as User</router-link>
            <router-link to="/admin-register" class="button">Register as Admin</router-link>
        </div>
        
    </div>
    <Footer/>
</div>
</template>
  
<style>
h1{
    font-family: Arial, Helvetica, sans-serif;
    font-weight: 800;
    font-size: 35px;
}
.container1 {
    display: flex;
    flex-direction: column;
    align-items: center;
    align-content: center;
    padding-top: 50px;
    margin-top: auto;
    height: 85vh;
    

}
.Button-Container1{
    padding-top:50px ;
    padding-bottom:30px ;
}
.Button-Container2{
    margin-top: -10;

   
}
p{
    margin: 0;
}
.button {
    display: inline-block;
    padding: 10px 20px;
    margin: 5px;
    text-decoration: none;
    color: #fff;
    background-color: #27496D;
    /* You can choose your preferred color */
    background-color: #27496D;
    border-radius: 5px;
    cursor: pointer;
}

.button:hover {
    background-color: #0056b3;
    /* Change the color on hover if you like */
}
</style>
  