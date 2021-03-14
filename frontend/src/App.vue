<template>
  <div>{{ pcName }}</div>
  <div>{{ avaliablespace }}</div>
  <div>{{ freeSpace }}</div>
  <div>{{ totalSpace }}</div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
  data() {
    return {
      avaliablespace: "Hello",
      freeSpace: "",
      totalSpace: "",
      pcName: "",
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
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
