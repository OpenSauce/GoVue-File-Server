<template>
  <div style="padding-bottom: 20px">
    <div style="padding-bottom: 20px">
      <h2>Execute Command</h2>
    </div>
    <div class="container-fluid" style="display:flex;">
      <input
        style="margin-right: 20px;"
        v-model="commandToExecute"
        type="text"
        class="form-control"
        id="exampleInputEmail1"
        placeholder=""
      />
      <button type="button" class="btn btn-primary" @click="SendCommand()">
        Execute
      </button>
    </div>
    <div v-if="result" class="container-fluid" style="margin-top:20px;">
      <textarea class="form-control" id="exampleFormControlTextarea1" rows="3" v-model="result"> </textarea>
    </div>
    <div v-if="error" class="container-fluid" style="margin-top:20px;">
      <textarea
        class="form-control"
        id="exampleFormControlTextarea1"
        rows="3"
        v-model="error"
        ></textarea>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "ExecuteCommand",
  data() {
    return {
      commandToExecute: "",
      result: "",
      error: "",
      envVar: "Test",
    };
  },
  created() {
    this.envVar = process.env.VUE_APP_IP;
  },
  methods: {
    SendCommand() {
      axios
        .post("http://" + this.envVar + ":8080/api/executecommand", {
          command: this.commandToExecute,
        })
        .then((response) => {
          console.log(response.data.output);
          this.result = response.data.output;
          console.log(response.data.error);
          this.error = response.data.error;
        })
        .catch((error) => {
          window.alert(`The API returned an error: ${error}`);
        });
      this.commandToExecute = "";
    },
  },
};
</script>

<style></style>
