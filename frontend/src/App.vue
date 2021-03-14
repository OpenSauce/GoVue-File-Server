<template>
  <div class="header" style="display:flex; justify-content: flex-start; align-items: center;">GoVue File Server</div>

  <div
    class="info-pane"
    style="display:flex; justify-content: flex-start; align-items: center;"
  >
    <span class="menu-item"> Server Name: {{ pcName }} </span>
    <span class="menu-item">{{ freeSpace }}/{{ totalSpace }}</span>
    <span class="menu-item">{{ avaliablespace }}</span>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
  data() {
    return {
      avaliablespace: "80%",
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
    this.makeWebsiteThumbnail();
  },
  methods: {
    makeWebsiteThumbnail() {
      axios
        .post("http://" + this.envVar + ":3000/api/avaliablespace", {})
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

#app {
  font-family: Segoe UI, Helvetica, Arial, sans-serif;
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
}

.info-pane {
  background-color: #e7e7e7; 
  height: 50px;
  border-bottom: solid black 1px;
}

.menu-item {
  padding-left: 20px;
  padding-right: 20px;
}
</style>
