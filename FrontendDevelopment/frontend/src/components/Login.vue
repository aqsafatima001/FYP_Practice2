<template>
    <div class="container">
        <div class="row">
            <div class="column">
                <h1 class="mt-5">Login</h1>
                <hr />

                <form-tag 
                @myevent="submitHandler" 
                name="myform" 
                event="myevent"> 
                
                    <text-input 
                    v-model="email"
                        label="Email" 
                        type="email" 
                        name="email" 
                        required="true">
                    </text-input>

                    <text-input 
                    v-model="password"
                        label="Password" 
                        type="password" 
                        name="password" 
                        required="true">
                    </text-input>

                    <hr />
                    <input type="submit" class="btn btn-primary" value="Login" />
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
    data(){
       return{
        email:"",
        password:""
       } 

    },
    methods:{
        submitHandler(){
            console.log("Submit Hnadler Called - Success !!!!")

            const payload = {
                email: this.email,
                password : this.password,
            }

            const requestOptions = {
                method: "POST",
                body: JSON.stringify(payload)
            }

            fetch("http://localhost:8081/users/login" , requestOptions)
            .then((response)=>response.json())
            .then((data)=>{
                if(data.error){
                    console.log("ERROR: ",data.message)
                }else{
                    console.log(data)
                }

            })
        }

    },
    // mounted() {
    //     (function () {
    //         'use strict'

    //         // Fetch all the forms we want to apply custom Bootstrap validation styles to
    //         var forms = document.querySelectorAll('.needs-validation')

    //         // Loop over them and prevent submission
    //         Array.prototype.slice.call(forms)
    //             .forEach(function (form) {
    //             form.addEventListener('submit', function (event) {
    //                 if (!form.checkValidity()) {
    //                 event.preventDefault()
    //                 event.stopPropagation()
    //                 }

    //                 form.classList.add('was-validated')
    //             }, false)
    //             })
    //         })()
    // }
}
</script>