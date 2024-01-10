<template>
    <div class="container">
        <div class="row">
            <div class="column">
                <h1 class="mt-5"> User Login</h1>
                <hr />

                <form-tag @myevent="submitHandler" name="myform" event="myevent">

                        <label for="name">Email: </label>
                            <input 
                                v-model="email" 
                                type="email" 
                                name="email" 
                                required="true" 
                                label="Email"
                                placeholder="abc@xyz.com"
                                :autocomplete="name + '-new'" 
                            >
                        
                        <label for="password">Password: </label>
                            <input 
                                v-model="password" 
                                type="password" 
                                name="password" 
                                required="true" 
                                label="Password"
                                placeholder="Enter your password"
                                :autocomplete="name + '-new'" 
                            >
                    <hr />
                    <!-- <input type="submit" class="btn btn-primary" value="Login" /> -->
                    <div class="button-container">
                        <input type="submit" class="Login-button" value="Login">
                    </div>
                </form-tag>
            </div>
        </div>
    </div>
</template>

<script>
import FormTag from "./forms/FormTag.vue"
import TextInput from './forms/TextInput.vue'

export default {
    name: 'login',
    components: {
        FormTag,
        TextInput,
    },
    data() {
        return {
            email:'',
            password:''
        }

    },
    methods: {
        submitHandler() {
            console.log("Submit Handler Called - Success !!!!")

         
            // Log the email and password before making the API call
            console.log("Email entered:", this.email);
            console.log("Password entered:", this.password);

            const payload = {
                email: this.email,
                password: this.password,
            }
            const requestOptions = {
                method: "POST",
                body: JSON.stringify(payload)
            }

            fetch("http://localhost:8090/api/userlogin", requestOptions)
                .then((response) => response.json())
                .then((data) => {
                    if (data.error) {
                        console.log("ERROR: ", data.message)
                    } else {
                        console.log(data)
                        this.$router.push('/userdashboard');
                    }
                })
                .catch((error) => {
                    console.error("Error fetching data:", error);
                });
        }

    },


}
</script>

<style scoped>
* {
    margin: 0;
    padding: 0;
}

.container {
    margin: 0 auto;
    max-width: 400px;
    padding: 20px;
    border-radius: 5px;
    box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
    position: absolute;
    /* or use 'fixed' if you want it to stay in the center even when scrolling */
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 1000px;
}

h1 {
    font-size: 48px;
    text-align: center;
    font-family: sans-serif;
    font-weight: bold;
    color: #27496D;
}

label {
    display: block;
    margin-top: 5px;
    font-family: sans-serif;
}

input[type="email"],
input[type="password"] {
    padding: 5px;
    border-radius: 5px;
    width: 100%;
    box-sizing: border-box;

}

.button-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    /* Horizontally center content */
    align-items: center;
    /* Vertically center content */
}

.Login-button {
    padding: 10px 20px;
    font-size: 16px;
    border: none;
    border-radius: 5px;
    background-color: #27496D;
    /* Change to your preferred background color */
    color: #fff;
    /* Change to your preferred text color */
    cursor: pointer;
    transition: background-color 0.3s ease, color 0.3s ease;
    margin-top: 25px;
    width: 100%
}

/* Hover effect */
.Login-button:hover {
    background-color: #0056b3;
    /* Change to your preferred hover background color */
}

</style>






                        