<template>
    <div class="container">
      <div class="row">
          <div class="column">
              <h1 class="mt-5"> Admin Registration</h1>
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
  
                        <!-- CAPTCHA CODE -->
                        <div>
                          <label>Enter Captcha:</label>
                      
                          <!-- <p>Enter Captcha:</p> -->
                          <div id="Catcha_Boxes">
                                <input type="text" id="mainCaptcha" readonly="readonly" />
  
                                <i id="refresh" class="fa fa-refresh fa-spin" style="font-size:24px"  @click="generateCaptcha"></i>
  
                                <input type="text" id="txtinput" v-model="captchaInput" />
  
                                <!-- <input type="button" id="Button1" value="Check" @click="checkValidCaptcha" /> -->
                                <span id="error" style="color:red">{{ captchaError }}</span>
                                <span id="success" style="color:green">{{ captchaSuccess }}</span>
                              </div>
                        </div>
                        <!-- CAPTCHA CODE  -->
                
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
    data() {
      return {
        formData: {
          username: '',
          email: '',
          password: '',
        },
        captchaInput: '',
        captchaError: '',
        captchaSuccess: '',
      };
    },
    methods: {
      async registerUser() {
        // Check if captcha is valid before proceeding
        if (!this.checkValidCaptcha()) {
          return; // Stop execution if captcha is not valid
        }
  
        try {
          const response = await fetch('http://localhost:8090/api/admin-otp-verfication', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
            },
            body: JSON.stringify(this.formData),
          });
  
          if (response.status === 201) {
            // Registration successful, redirect to email verification page
            this.$router.push({
              name: 'AdminOtpVerification',
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
  
      generateCaptcha() {
        const cap = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz';
        let code = '';
  
        for (let i = 0; i < 4; i++) {
          code += cap[Math.floor(Math.random() * cap.length)];
        }
  
        document.getElementById("mainCaptcha").value = code;
      },
  
      checkValidCaptcha() {
        const string1 = this.removeSpaces(document.getElementById('mainCaptcha').value);
        const string2 = this.removeSpaces(this.captchaInput);
  
        if (string1 === string2) {
          this.captchaSuccess = "Form is validated Successfully";
          // alert("Form is validated Successfully");
          // window.location = "https://www.youtube.com/c/buzz2daytech/";
          return true;
        } else {
          this.captchaError = "Please Enter a Valid Captcha";
          return false;
        }
      },
  
      removeSpaces(string) {
        return string.split(' ').join('');
      },
    },
  
    mounted() {
      // Call generateCaptcha on component mount
      this.generateCaptcha();
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
  
  
  
  /* Styling for Captcha Text */
  
  #Catcha_Boxes{
    display: flex;
  }
  #mainCaptcha {
    font-family: 'Monotype Corsiva', cursive;
    font-size: 26px;
    color: #009688;
    width: 150px;
    height: 35px;
    text-align: center;
  }
  
  #txtinput{
    width: 150px;
    height: 35px;
    text-align: center;
    font-size: 26px;
  }
  
  #refresh{
    margin-left: 10px;
    margin-right: 10px;
  }
  
  
  
  
  .input-captcha {
    padding: 5px;
    border-radius: 5px;
    width: 100%;
    box-sizing: border-box;
    font-family: 'Monotype Corsiva', cursive; /* Use your preferred font family */
    font-size: 18px;
  }
  
  
  #error {
    color: red;
  }
  
  #success {
    color: green;
  }
  
  </style>
  