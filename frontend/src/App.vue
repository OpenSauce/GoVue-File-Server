<template>
  <div
    class="header-local h1 startlocal"
    style="display:flex; justify-content: flex-start; align-items: center;"
  >
    GoVue File Server
  </div>

  <div
    class="info-pane startlocal"
    style="display:flex; justify-content: flex-start; align-items: center;"
  >
    <span class="info-item"> Server Name: {{ pcName }} </span>
    <span class="info-item">{{ freeSpace }}/{{ totalSpace }}</span>
    <span class="info-item">{{ percentageAvaliable }}%</span>
  </div>
  <div class="row2">
    <div class="column side" style="display: flex; flex-direction:column;">
      <router-link tag="li" to="/"><span class="menu-item h6">Home</span></router-link>
      <router-link to="/files"><span class="menu-item h6">Files</span></router-link>
      <span class="menu-item h6">Settings</span>
    </div>
    <div class="column.middle">
      <router-view />
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  data() {
    return {
      avaliablespace: "80",
      freeSpace: "8GB",
      totalSpace: "10GB",
      pcName: "PC",
      envVar: "Test",
    };
  },
  created() {
    this.envVar = process.env.VUE_APP_IP;
  },
  mounted() {
    //this.getServerStats();
  },
  computed: {
    percentageAvaliable() {
      return parseFloat(this.avaliablespace).toFixed(2);
    },
  },
  methods: {
    getServerStats() {
      axios
        .post("http://" + this.envVar + ":8080/api/avaliablespace", {})
        .then((response) => {
          this.avaliablespace = response.data.avaliablespace;
          this.freeSpace = response.data.freeSpace;
          this.totalSpace = response.data.totalSpace;
          this.pcName = response.data.pcName;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },
  },
};
</script>

<style>
* {
  margin: 0;
  padding: 0;
}

.h1,
h2,
h3,
h4,
h5,
.h6,
h1 {
  font-family: "Oswald", sans-serif;
  text-transform: uppercase;
  margin: 0 !important;
}

h2 {
  padding-bottom: 20px;
}

/* a {
  color: black !important;
} */

.active {
  color:black;
}


body {
  font-family: "Segoe UI", serif;
}

#app {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

.header-local {
  font-size: 48px;
  padding-left: 20px;
  background-color: #131313;
  color: white;
  height: 100px;
  word-spacing: 20px;
}

.startlocal {
  padding-left: 60px;
}

.info-pane {
  background-color: #e7e7e7;
  height: 50px;
  border-bottom: solid black 1px;
}

.info-item {
  padding-right: 40px;
}

.menu-item {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 80px;
}

.menu-item:hover {
  background-color: #848484;
  cursor: pointer;
}

/* Create three equal columns that floats next to each other */
.column {
  float: left;
}

/* Left and right column */
.column.side {
  width: 15%;
  background-color: #343434;
  color: white;
  height: 400px;
}

.content {
  display: flex;
  flex-direction: column;
}

.column.middle {
  width: 85%;
  background-color: red;
}

.content-item {
  text-align: left;
  margin-top: 20px;
  padding: 40px;
  background-color: #e7e7e7;
}

/* Clear floats after the columns */
.row2:after {
  background-color: green;
  content: "";
  display: table;
  clear: both;
}

/* Responsive layout - makes the three columns stack on top of each other instead of next to each other on smaller screens (600px wide or less) */
@media screen and (max-width: 1200px) {
  .column.side {
    width: 33%;
  }

  .column.middle {
    width: 67%;
  }
}

/* Responsive layout - makes the three columns stack on top of each other instead of next to each other on smaller screens (600px wide or less) */
@media screen and (max-width: 600px) {
  .column.side,
  .column.middle {
    width: 100%;
  }
}
</style>
