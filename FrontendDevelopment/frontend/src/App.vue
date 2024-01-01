<template>
  <body>
    
    <Header />
    <div>
      <RouterView/>
    </div>
    
    <Footer />
    
  </body>
</template>

<script setup>
import Header from "./components/Header.vue"
import Body from "./components/Body.vue"
import Footer from "./components/Footer.vue"

// export default {
//   name: 'App',
//   components: {
//     Header,
//     Body,
//     Footer,
//   }
// }
</script>

<style>

</style>

