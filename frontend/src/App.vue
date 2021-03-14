<template>
  <div
    class="header h1 start"
    style="display:flex; justify-content: flex-start; align-items: center;"
  >
    GoVue File Server
  </div>

  <div
    class="info-pane start"
    style="display:flex; justify-content: flex-start; align-items: center;"
  >
    <span class="info-item"> Server Name: {{ pcName }} </span>
    <span class="info-item">{{ freeSpace }}/{{ totalSpace }}</span>
    <span class="info-item">{{ percentageAvaliable }}%</span>
  </div>
  <div class="row">
    <div class="column side" style="display: flex; flex-direction:column;">
      <span class="menu-item h6">Files</span>
      <span class="menu-item h6">Settings</span>
    </div>
    <div class="column.middle">
      <div class="content">
        <span class="content-item"
          ><p>
            "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim
            ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
            aliquip ex ea commodo consequat. Duis aute irure dolor in
            reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
            pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
            culpa qui officia deserunt mollit anim id est laborum."
          </p></span
        >
        <span class="content-item"
          ><p>
            "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim
            ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
            aliquip ex ea commodo consequat. Duis aute irure dolor in
            reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
            pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
            culpa qui officia deserunt mollit anim id est laborum."
          </p></span
        >
        <span class="content-item"
          ><p>
            "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
            eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim
            ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut
            aliquip ex ea commodo consequat. Duis aute irure dolor in
            reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
            pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
            culpa qui officia deserunt mollit anim id est laborum. Sed ut
            perspiciatis unde omnis iste natus error sit voluptatem accusantium
            doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo
            inventore veritatis et quasi architecto beatae vitae dicta sunt
            explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur
            aut odit aut fugit, sed quia consequuntur magni dolores eos qui
            ratione voluptatem sequi nesciunt. Neque porro quisquam est, qui
            dolorem ipsum quia dolor sit amet, consectetur, adipisci velit, sed
            quia non numquam eius modi tempora incidunt ut labore et dolore
            magnam aliquam quaerat voluptatem. Ut enim ad minima veniam, quis
            nostrum exercitationem ullam corporis suscipit laboriosam, nisi ut
            aliquid ex ea commodi consequatur? Quis autem vel eum iure
            reprehenderit qui in ea voluptate velit esse quam nihil molestiae
            consequatur, vel illum qui dolorem eum fugiat quo voluptas nulla
            pariatur?"
          </p></span
        >
        <span class="content-item">Hello</span>
        <span class="content-item">Hello</span>
        <span class="content-item">Hello</span>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
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
    // this.makeWebsiteThumbnail();
  },
  computed: {
    percentageAvaliable() {
      return parseFloat(this.avaliablespace).toFixed(2);
    },
  },
  methods: {
    makeWebsiteThumbnail() {
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
.h6 {
  font-family: "Oswald", sans-serif;
  text-transform: uppercase;
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

.header {
  font-size: 48px;
  padding-left: 20px;
  background-color: #131313;
  color: white;
  height: 100px;
  word-spacing: 20px;
}

.start {
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
.row:after {
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
