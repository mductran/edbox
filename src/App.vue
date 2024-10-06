<script setup>
import Sandbox from "@/components/Sandbox.vue";
import OutputBox from "@/components/OutputBox.vue";
import TaskList from "@/components/TaskList.vue";
</script>

<template>
  <main>
    <div class="left">
      <TaskList/>
    </div>
    <div class="middle">
      <Sandbox/>
    </div>
    <div class="right">
      <OutputBox/>
    </div>
  </main>
</template>

<style scoped>
  main {
    display: flex;
    height: 100vh;
    align-content: center;
  }

  .left {
    border: 1px solid red;
    width: 10vw;
  }

  .middle {
    border: 1px solid red;
    width: 35vw;
  }

  .right {
    border: 1px solid red;
    width: 35vw;
  }
</style>
