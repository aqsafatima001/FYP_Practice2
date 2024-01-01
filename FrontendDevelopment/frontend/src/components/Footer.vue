<!-- <template>
    <footer class="d-flex align-items-center justify-content-center mt-5">
        Copyright &copy; {{ new Date().getFullYear()}}
    </footer>
</template>
<style scoped>
footer {
    background: black;
    color: white;
    min-height: 3em;
    font-size: 0.8em;
}
</style> -->

<template>
    <footer>
        <p>&copy; Copyright <b>2023 - NCP.</b> All Rights Reserved</p>
        <!-- <p>Contact: <a href="mailto:contact@example.com">contact@example.com</a></p> -->
    </footer>
</template>

<style scoped>
 /* Styling for the footer */
 footer {
    position: fixed;
    left: 0;
    bottom: 0;
    width: 100%;
    background-color: #27496D;
    color: #fff;
    padding: 20px;
    text-align: center;
}

/* Styling for links in the footer */
footer a {
    color: #fff;
    text-decoration: none;
}
</style>