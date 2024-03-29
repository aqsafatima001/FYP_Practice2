<template>
    <div>
        <h1>Create a user to access hadoop cluster</h1>
      <button @click="createUser">Create User</button>
    </div>
  </template>
  
  <script>
  export default {
    methods: {
      async createUser() {
        try {
          const response = await fetch('http://localhost:8090/api/createUserInCentOS', {
            method: 'POST'
          });
          if (response.ok) {
            alert('User created successfully!');
          } else {
            throw new Error('Failed to create user');
          }
        } catch (error) {
          console.error('Error creating user:', error);
          alert('Failed to create user');
        }
      }
    }
  }
  </script>
  