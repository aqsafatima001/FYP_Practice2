<template>
    <div class="container">
      <div class="row">
          <div class="column">
              <h1 class="mt-5">OTP Verification</h1>
              <hr />
              <form @submit.prevent="verifyOtp">
                <div>
                  <div>
                    <!-- <label>Your OTP has been sent to : {{ email }} and the password is {{ password }} and the username is {{
                      username }} </label> -->
                      <label>Your OTP has been sent to : <b>{{ email }} </b></label>
          
                  </div>
                  <label for="otp">Enter OTP:</label>
                  <input type="text" id="otp" v-model="formData.otp" required />
                </div>
                <div>
                  <button type="submit" class="VerifyButton">Verify OTP</button>
                </div>
              </form>
  
              
          </div>
      </div>
  </div>
  
  </template>
    
  <script>
  export default {
    name: 'OtpVerificationWithAdditionalInformation',
    data() {
  
      return {
        formData: {
          otp: '',
          email: '',
          password: '',
          username: '',
          dateOfBirth: '',
                nationality: '',
                nicNumber: '',
                passportNumber: '',
                title: '',
                firstName: '',
                middleName: '',
                lastName: '',
                universityCentre: '',
                thesisTitle: '',
                startDate: '',
                endDate: '',
                pictureProof: ''
        },
      };
    },
    created() {
      // Access the email from the route's query parameters
      this.email = this.$route.query.email;
      this.password = this.$route.query.password;
      this.username = this.$route.query.username;
      // Access additional parameters if available
    this.formData.dateOfBirth = this.$route.query.dateOfBirth || '';
    this.formData.nationality = this.$route.query.nationality || '';
    this.formData.nicNumber = this.$route.query.nicNumber || '';
    this.formData.passportNumber = this.$route.query.passportNumber || '';
    this.formData.title = this.$route.query.title || '';
    this.formData.firstName = this.$route.query.firstName || '';
    this.formData.middleName = this.$route.query.middleName || '';
    this.formData.lastName = this.$route.query.lastName || '';
    this.formData.universityCentre = this.$route.query.universityCentre || '';
    this.formData.thesisTitle = this.$route.query.thesisTitle || '';
    this.formData.startDate = this.$route.query.startDate || '';
    this.formData.endDate = this.$route.query.endDate || '';
    this.formData.pictureProof = this.$route.query.pictureProof || '';
    },
    methods: {
      async verifyOtp() {
        // Send OTP and email to the server for verification
        const response = await fetch('http://localhost:8090/api/verify-otp-user-additonal-info', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: this.email,
            password: this.password,
            username: this.username,
            otp: this.formData.otp,
            dateOfBirth: this.formData.dateOfBirth,
            nationality: this.formData.nationality,
            nicNumber: this.formData.nicNumber,
            passportNumber: this.formData.passportNumber,
            title: this.formData.title,
            firstName: this.formData.firstName,
            middleName: this.formData.middleName,
            lastName: this.formData.lastName,
            universityCentre: this.formData.universityCentre,
            thesisTitle: this.formData.thesisTitle,
            startDate: this.formData.startDate,
            endDate: this.formData.endDate,
            pictureProof: this.formData.pictureProof
          }),
        });
  
        if (response.status === 200) {
          // OTP verification successful, proceed with data storage
          this.$router.push('/registration-success');
  
        } else {
          alert('OTP verification failed');
        }
      },
    },
  };
  </script>
   
  
  <style scoped>
  
  
  label {
    display: block;
    margin-top: 5px;
    font-family: sans-serif;
  }
  
  input[type="text"] {
    padding: 5px;
      border-radius: 5px;
      width: 20%;
      box-sizing: border-box;
  }
  .VerifyButton{
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
    width: 20%
  }
  .VerifyButton:hover{
    background-color: #0056b3;
    /* Change to your preferred hover background color */
  }
  </style>