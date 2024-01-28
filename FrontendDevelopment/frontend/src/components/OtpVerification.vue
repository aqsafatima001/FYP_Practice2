<template>
    <div>
      <h1>OTP Verification</h1>
      <form @submit.prevent="verifyOtp">
        <div>
          <label for="otp">Enter OTP:</label>
          <input type="text" id="otp" v-model="otp" required />
        </div>
        <div>
          <button type="submit">Verify OTP</button>
        </div>
      </form>
    </div>
  </template>
  
  <script>
  export default {
    name: 'OtpVerification',
    data() {
      return {
        otp: '',
      };
    },
    methods: {
      async verifyOtp() {
        // Send OTP and email to the server for verification
        const response = await fetch('http://localhost:8090/api/verify-otp', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            email: this.$route.params.email,
            otp: this.otp,
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
  