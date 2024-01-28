<template>
    <div class="container">
      <div class="success-message">
        <p>CONGRATULATIONS!!!</p>
        <p> Your email registration is successful</p>
      </div>
      <button type="submit" class="request_resources-button">Request Resources</button>
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
    flex-direction: column;
    /*justify-content: center; */
    align-items: center; 
    height: 100vh; /* Make the container full height of the viewport */
  }

  .request_resources-button{
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
  .request_resources-button:hover{
    background-color: #0056b3;
  }
  .success-message {
    background-color: #4CAF50;
    color: white;
    padding: 20px;
    text-align: center;
    border-radius: 8px;
  }
  </style>
  