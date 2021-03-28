<template>
  <div>
    <div class="content">
      <span class="content-item" style="padding-bottom: 20px"
        ><h2>File List</h2>
        <button @click="getFileList()" class="btn btn-primary">Refresh</button>
        <ul id="example-1">
          <li v-for="item in items" :key="item">
            {{ item }}
          </li>
        </ul>
      </span>
      <span class="content-item">
        <form method="post" action="#" id="" enctype="multipart/form-data" />
        <div class="form-group files text-center" ref="fileform">
          <input type="file" ref="file" multiple="multiple" />
          <span id="val"></span>
          <button @click="submitFiles()" class="btn btn-primary"  id="button">Upload Photo</button>
        </div>
      </span>
      <span class="content-item"
        ><p>
          "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do
          eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad
          minim veniam, quis nostrud exercitation ullamco laboris nisi ut
          aliquip ex ea commodo consequat. Duis aute irure dolor in
          reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla
          pariatur. Excepteur sint occaecat cupidatat non proident, sunt in
          culpa qui officia deserunt mollit anim id est laborum. Sed ut
          perspiciatis unde omnis iste natus error sit voluptatem accusantium
          doloremque laudantium, totam rem aperiam, eaque ipsa quae ab illo
          inventore veritatis et quasi architecto beatae vitae dicta sunt
          explicabo. Nemo enim ipsam voluptatem quia voluptas sit aspernatur aut
          odit aut fugit, sed quia consequuntur magni dolores eos qui ratione
          voluptatem sequi nesciunt. Neque porro quisquam est, qui dolorem ipsum
          quia dolor sit amet, consectetur, adipisci velit, sed quia non numquam
          eius modi tempora incidunt ut labore et dolore magnam aliquam quaerat
          voluptatem. Ut enim ad minima veniam, quis nostrum exercitationem
          ullam corporis suscipit laboriosam, nisi ut aliquid ex ea commodi
          consequatur? Quis autem vel eum iure reprehenderit qui in ea voluptate
          velit esse quam nihil molestiae consequatur, vel illum qui dolorem eum
          fugiat quo voluptas nulla pariatur?"
        </p></span
      >
      <span class="content-item">Hello</span>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Files",
  setup() {},
  data() {
    return {
      envVar: "Test",
      items: ["TestFile1.txt", "TestFile2.txt"],
    };
  },
  created() {
    this.envVar = process.env.VUE_APP_IP;
  },
  methods: {
    submitFiles() {
      let formData = new FormData();

      for (var i = 0; i < this.$refs.file.files.length; i++) {
        let file = this.$refs.file.files[i];
        formData.append("multiplefiles", file);
      }

      axios
        .post("http://" + this.envVar + ":8080/api/upload", formData, {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        })
        .then(function() {})
        .catch(function() {});
    },
    getFileList() {
         axios
        .post("http://" + this.envVar + ":8080/api/getfiles", {})
        .then((response) => {
          this.items = response.data.files;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
    },
  },
};
</script>

<style scoped></style>
