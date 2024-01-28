<template>
  <div>
    <h1>OTP Verification</h1>
    <form @submit.prevent="verifyOtp">
      <div>
        <div>
          <label>Your OTP has been sent to : {{ email }} and the password is {{ password }} and the username is {{
            username }} </label>

        </div>
        <label for="otp">Enter OTP:</label>
        <input type="text" id="otp" v-model="formData.otp" required />
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
      formData: {
        otp: '',
        email: '',
        password: '',
        username: '',
      },
    };
  },
  created() {
    // Access the email from the route's query parameters
    this.email = this.$route.query.email;
    this.password = this.$route.query.password;
    this.username = this.$route.query.username;
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
          email: this.email,
          password: this.password,
          username: this.username,
          otp: this.formData.otp,
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
  