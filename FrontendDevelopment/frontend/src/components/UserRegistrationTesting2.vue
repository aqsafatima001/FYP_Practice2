<template>
    <div>
      <h1>User Registration</h1>
      <form @submit.prevent="goToOtpVerification">
        <div>
          <label for="username">Username:</label>
          <input type="text" id="username" v-model="formData.username" required />
        </div>
        <div>
          <label for="email">Email:</label>
          <input type="email" id="email" v-model="formData.email" required />
        </div>
        <div>
          <label for="password">Password:</label>
          <input type="password" id="password" v-model="formData.password" required />
        </div>
        <div>
          <button type="submit">Continue to OTP Verification</button>
        </div>
      </form>
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
      };
    },
    methods: {
      goToOtpVerification() {
        // Redirect to OTP verification page and pass email as a route parameter
        this.$router.push({ name: 'OtpVerification', params: { email: this.formData.email } });
      },
    },
  };
  </script>
  