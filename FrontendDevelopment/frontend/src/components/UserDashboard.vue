<script>
import Footer from "./../components/Footer.vue"


export default {
    name: 'userdashboard',
    components: {
        Footer,
    }
}
</script>

<template>
    <div class="grid-container">
        <div class="grid-item header">

            <div class="logo">
                <div class="logo-content">
                    <img class="logo-image" src="@/assets/ncp (1).png" alt="Your Logo" />
                    <span class="logo-text">NATIONAL CENTER FOR PHYSICS</span>
                </div>
            </div>

            <div class="header-items">
                <div class="header-item">Dashboard</div>
                <div class="header-item">Services</div>
                <div class="header-item">Alerts</div>
                <div class="header-item">Stack Version</div>
            </div>
        </div>

        <div class="grid-item sidebar">
            <ul class="sidebar_list">
                <li>HDFS</li>
                <li>YARN</li>
                <li>MapReduce</li>
                <li>Hive</li>
                <li>Spark</li>
                <li>Pig</li>
                <li>Sqoop</li>
                <li>Zookeeper</li>
                <li>Hbase</li>
                <li>Kafka</li>
                <li>Flume</li>
                <li>Storm</li>
            </ul>
        </div>

        <div class="grid-item main-content">
            <div class="content">
                Main Content    
            </div>
            <!-- <div class="content">Main Content</div>
            <div class="content">Main Content</div> -->
        </div>

        <div class="grid-item footer">
            <Footer/>
        </div>
    </div>
</template>
  
<style scoped>
.grid-container {
    display: grid;
    grid-template-columns: 1fr 5fr;
    /* Adjust column widths as needed */
    grid-template-rows: auto 1fr auto;
    /* Adjust row heights as needed */
    gap: 10px;
    height: 100vh;
}

.grid-item {
    padding: 20px;
}

.header {
    grid-column: 1 / span 2;
}

.header {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 20px;
    background-color: #27496D;
}

.logo {
    font-size: 24px;
    /* Adjust the font size for your logo */
    font-weight: bold;
}

.logo {
    display: flex;
    align-items: center;
}

.logo-image {
    max-width: 50px;
    /* Adjust the maximum width as needed */
    height: auto;
    /* Maintain aspect ratio */
    margin-right: 20px;
    /* Adjust the spacing to your preference */
}
.logo-text{
    font-size:16px ;
    font-weight:bold ;
    color: white;
}

.header-items {
    display: flex;
}

.header-item {
    margin-right: 20px;
    /* Adjust the spacing between items */
    font-weight: bold;
    background-color: white;
    padding: 6px;
    text-align: center;
    border-radius: 8px;
    color: rgb(25, 25, 112);
    width: 150px;
}

.header-item:hover {
    margin-right: 40px;
    /* Adjust the spacing between items */
    font-weight: bold;
    background-color: rgb(137, 207, 240);
    padding: 6px;
    text-align: center;
    border-radius: 8px;
    color: black;
    width: 150px;
}

.sidebar {
    grid-row: 2 / span 1;
}

.sidebar_list {
    list-style: none;
    display: flex;
    flex-direction: column;
    align-items: center;
    align-content: center;

}

.sidebar_list li {
    padding: 3px 4px;
    margin: 3px 4px;
    width: 200px;
    text-align: center;
    background-color: #27496D;
    color: white;
    font-weight: bold;
}

.sidebar_list li:hover {
    padding: 3px 4px;
    margin: 3px 4px;
    width: 200px;
    text-align: center;
    background-color: rgb(137, 207, 240);
    color: white;
    font-weight: bold;
}

.main-content {
    grid-row: 2 / span 1;
    grid-column: 2 / span 1;
    display: grid;
    grid-template-columns: 1fr;
    gap: 10px;
}

.content {
    padding: 10px;
    text-align: center;
    background-color: rgb(229, 228, 226);
}

.footer {
    grid-column: 1 / span 2;
}

</style>
  