<template>
    <div class="container">
      <div class="success-message">
        <p>CONGRATULATIONS!!!</p>
        <p> Your registration is successful</p>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'RegistrationSuccess',
    // Add any component logic here if needed
  }
  </script>
  
  <style scoped>
  .container {
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh; /* Make the container full height of the viewport */
  }
  
  .success-message {
    background-color: #4CAF50;
    color: white;
    padding: 20px;
    text-align: center;
    border-radius: 8px;
  }
  </style>
  