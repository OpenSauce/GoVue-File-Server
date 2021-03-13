<template>
  <div>{{ avaliablespace }}</div>
</template>

<script>
import axios from "axios";

export default {
  name: "App",
  components: {},
  data() {
    return {
      avaliablespace: "Hello",
    };
  },
  created() {
    this.makeWebsiteThumbnail();
  },
  methods: {
    makeWebsiteThumbnail() {
      axios
        .post("http://localhost:3000/api/avaliablespace", {
        })
        .then((response) => {
          this.avaliablespace = response.data.avaliablespace;
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
