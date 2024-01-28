<template>
  <div>
    <h1>User Registration</h1>
    <form @submit.prevent="registerUser">
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
        <button type="submit">Send OTP</button>
      </div>
    </form>
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
          // this.$router.push('/OtpVerification');
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
  