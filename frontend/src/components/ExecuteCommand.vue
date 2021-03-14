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
    <div></div>
  </div>
</template>

<script>
export default {
  name: "ExecuteCommand",
  data() {
    return {
      commandToExecute: "",
    };
  },
  methods: {
    SendCommand() {
      this.commandToExecute = "";
    },
  },
};
</script>

<style></style>
