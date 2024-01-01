<!-- <template>
  <nav class="navbar navbar-expand-lg bg-body-tertiary">
  <div class="container-fluid">
    <a class="navbar-brand" href="#">Navbar</a>
    <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
      <span class="navbar-toggler-icon"></span>
    </button>
    <div class="collapse navbar-collapse" id="navbarNav">
      <ul class="navbar-nav">
        <li class="nav-item">
          <router-link class="nav-link active" aria-current="page" to="/">Home</router-link>
        </li>
        <li class="nav-item">
          <router-link class="nav-link" to="/login">Login</router-link>
        </li>
      </ul>
    </div>
  </div>
</nav>
</template> -->

<template>
  <header>
      <nav>
          <ul>
              <li><a href="#">Home</a></li>
              <li><a href="#">About</a></li>
              <li><a href="#">Contact Us</a></li>
          </ul>
      </nav>
  </header>
</template>

<style scoped>
/* Reset some default styles and apply basic styling to the header */
body {
  margin: 0;
  padding: 0;
  font-family: Arial, sans-serif;
}

header {
  background-color: #27496D; /* Change to your preferred header background color */
  color: #fff; /* Change to your preferred text color */
  padding: 10px 0;
}

/* Style the navigation menu */
nav ul {
  list-style: none;
  padding: 0;
  margin: 0;
  display: flex;
}

nav li {
  margin-right: 20px; /* Adjust spacing between menu items */
}

nav a {
  text-decoration: none;
  color: #fff; /* Change to your preferred link color */
  font-weight: bold;
}

/* Style navigation links on hover */
nav a:hover {
  color: #ff6600; /* Change to your preferred hover color */
}

</style>