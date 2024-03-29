<template>
    <div>
      <label id="pageheading" for="pendingRequests">PENDING REQUESTS:</label>
      <div v-if="pendingUsers.length === 0" class="no-requests">No pending requests</div>
      <div v-else>
        <select id="pendingRequests" v-model="selectedUser" class="select-box">
            <option v-for="user in pendingUsers" :key="user.email" :value="user.email" class="select-option">
              {{ user.username }} - {{ user.email }}
            </option>
          </select>
        <button @click="acceptRequest" class="action-button">Accept</button>
        <button @click="declineRequest" class="action-button">Decline</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    data() {
      return {
        pendingUsers: [],
        selectedUser: null
      };
    },
    mounted() {
      this.fetchPendingRequests();
    },
    methods: {
      async fetchPendingRequests() {
        try {
          const response = await fetch('http://localhost:8090/api/getPendingRequests');
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          const data = await response.json();
          this.pendingUsers = data;
        } catch (error) {
          console.error('Error fetching pending requests:', error);
        }
      },
      async acceptRequest() {
        if (!this.selectedUser) return;
        try {
          const response = await fetch('http://localhost:8090/api/acceptRequest', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email: this.selectedUser })
          });
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          // Refresh pending requests after accepting
          this.fetchPendingRequests();
        } catch (error) {
          console.error('Error accepting request:', error);
        }
      },
      async declineRequest() {
        if (!this.selectedUser) return;
        try {
          const response = await fetch('http://localhost:8090/api/declineRequest', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json'
            },
            body: JSON.stringify({ email: this.selectedUser })
          });
          if (!response.ok) {
            throw new Error('Network response was not ok');
          }
          // Refresh pending requests after declining
          this.fetchPendingRequests();
        } catch (error) {
          console.error('Error declining request:', error);
        }
      }
    }
  };
  </script>
  
  <style>

/* Styling for select box */
.select-box {
    width: 200px;
    padding: 8px;
    border: 1px solid #ccc;
    border-radius: 4px;
    font-size: 16px;
  }
  
  /* Styling for select options */
  .select-option {
    font-size: 14px;
  }

  /* Add your CSS styles here */
  .no-requests {
    color: red;
    font-weight: bold;
  }
  
  .action-button {
    margin-left: 10px;
    padding: 5px 10px;
    background-color: #007bff;
    color: white;
    border: none;
    border-radius: 5px;
    cursor: pointer;
  }
  
  .action-button:hover {
    background-color: #0056b3;
  }

  #pageheading{
    font-size: 40px;
    font-weight: bold;
    font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
  }
  </style>
  