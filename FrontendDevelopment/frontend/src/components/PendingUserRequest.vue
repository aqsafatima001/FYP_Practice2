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
    <!-- Modal for displaying additional user information -->
    <div v-if="showModal" class="modal">
      <div class="modal-content">
        <span @click="closeModal" class="close">&times;</span>
        <h3>User Information</h3>
        <p><strong>Username:</strong> {{ selectedUserInfo.username }}</p>
        <p><strong>Email:</strong> {{ selectedUserInfo.email }}</p>
        <p><strong>NCP ID:</strong> {{ selectedUserInfo.ncpId }}</p>
        <p><strong>Date of Birth:</strong> {{ selectedUserInfo.dateOfBirth }}</p>
        <p><strong>Nationality:</strong> {{ selectedUserInfo.nationality }}</p>
        <p><strong>NIC Number:</strong> {{ selectedUserInfo.nicNumber }}</p>
        <p><strong>Passport Number:</strong> {{ selectedUserInfo.passportNumber }}</p>
        <p><strong>Title:</strong> {{ selectedUserInfo.title }}</p>
        <p><strong>First Name:</strong> {{ selectedUserInfo.firstName }}</p>
        <p><strong>Middle Name:</strong> {{ selectedUserInfo.middleName }}</p>
        <p><strong>Last Name:</strong> {{ selectedUserInfo.lastName }}</p>
        <p><strong>University Centre:</strong> {{ selectedUserInfo.universityCentre }}</p>
        <p><strong>Thesis Title:</strong> {{ selectedUserInfo.thesisTitle }}</p>
        <p><strong>Start Date:</strong> {{ selectedUserInfo.startDate }}</p>
        <p><strong>End Date:</strong> {{ selectedUserInfo.endDate }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      pendingUsers: [],
      selectedUser: null,
      selectedUserInfo: null,
      showModal: false
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
    },
    async fetchUserInfo(email) {
  try {
    const response = await fetch(`http://localhost:8090/api/getUserInfo?email=${email}`);
    if (!response.ok) {
      throw new Error('Network response was not ok');
    }
    const userInfo = await response.json();
    console.log(userInfo); // Check the received user information

    // Populate selectedUserInfo with the received user information
    this.selectedUserInfo = {
      username: userInfo.Username,
      email: userInfo.Email,
      ncpId: userInfo.NCP_ID,
      dateOfBirth: userInfo.DateOfBirth,
      nationality: userInfo.Nationality,
      nicNumber: userInfo.NIC_Number,
      passportNumber: userInfo.Passport_Number,
      title: userInfo.Title,
      firstName: userInfo.FirstName,
      middleName: userInfo.MiddleName,
      lastName: userInfo.LastName,
      universityCentre: userInfo.University_Centre,
      thesisTitle: userInfo.ThesisTitle,
      startDate: userInfo.StartDate,
      endDate: userInfo.EndDate,
      // Ensure you handle PictureProof according to your requirements
    };

    this.showModal = true; // Show the modal after fetching user info
  } catch (error) {
    console.error('Error fetching user information:', error);
  }
},

    closeModal() {
      this.showModal = false;
    }
  },
  watch: {
    selectedUser(newValue) {
      if (newValue) {
        this.fetchUserInfo(newValue);
      }
    }
  }
};
</script>

<style>
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

#pageheading {
  font-size: 40px;
  font-weight: bold;
  font-family: 'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif;
}

/* CSS for the modal */
.modal {
  display: block;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.5);
  /* Black background with transparency */
}

.modal-content {
  background-color: #fefefe;
  margin: 15% auto;
  padding: 20px;
  border: 1px solid #888;
  width: 80%;
  max-width: 600px;
  border-radius: 5px;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  color: black;
  text-decoration: none;
  cursor: pointer;
}
</style>




<!-- <template>
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
   -->