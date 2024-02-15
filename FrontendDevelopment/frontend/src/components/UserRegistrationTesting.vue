<template>
  <div class="container">
    <div class="row">
        <div class="column">
            <h1 class="mt-5"> User Registration</h1>
            <hr />
            <form @submit.prevent="registerUser" class="registration-form">

              <div class="form-group">
                <label for="username">Username:</label>
                    <input 
                        type="text" 
                        id="username"
                        v-model="formData.username" 
                        required 
                    />
             
                <label for="email">Email:</label>
                    <input 
                        type="email" 
                        id="email" 
                        v-model="formData.email" 
                        required 
                        placeholder="Enter your Email here"
                    />
             
                <label for="password">Password:</label>
                    <input 
                        type="password" 
                        id="password" 
                        v-model="formData.password" 
                        required 
                        placeholder="Enter your password"
                    />
              
                <button type="submit" class="submit-button">Send OTP</button>
              </div>
            
          </form>

            
        </div>
    </div>
</div>

</template>

  
<script>
export default {
  name: 'registration',
  components: {

  },
  data() {
    return {
      formData: {
        username: '',
        email: '',
        password: '',
      },
    };
  },
  methods: {
    async registerUser() {
      try {
        const response = await fetch('http://localhost:8090/api/otp-verfication', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify(this.formData),
        });

        if (response.status === 201) {
          // Registration successful, redirect to email verification page
          this.$router.push({
            name: 'OtpVerification',
            query: {
              email: this.formData.email,
              password: this.formData.password,
              username: this.formData.username,
            },
          });
        } else {
          if (response.headers.get('content-type').indexOf('application/json') === -1) {
            // Response is not JSON, handle it accordingly
            alert('Non-JSON response from the server');
          } else {
            try {
              // Attempt to parse the response as JSON
              const data = await response.json();

              // Check if data is valid JSON and contains a message
              if (data && data.message) {
                alert(data.message);
              } else {
                // Handle unexpected JSON response
                alert('Unexpected JSON response from the server');
              }
            } catch (error) {
              // Handle JSON parsing error
              alert('Error parsing JSON response from the server');
            }
          }
        }
      } catch (error) {
        console.error(error);
        alert('An error occurred during registration.');
      }
    },


  },
};
</script>
  


<style scoped>
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
  font-size: 36px;
  text-align: center;
  font-family: sans-serif;
  font-weight: bold;
  color: #27496D;
}

.registration-form {
  max-width: 400px;
  margin: 0 auto;
}

.form-group {
  margin-bottom: 15px;
}

label {
  display: block;
  margin-top: 5px;
  font-family: sans-serif;
}

input[type="text"],
input[type="email"],
input[type="password"] {
  padding: 5px;
    border-radius: 5px;
    width: 100%;
    box-sizing: border-box;
}

.submit-button {
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

.submit-button:hover {
  background-color: #0056b3;
  /* Change to your preferred hover background color */
}
</style>