<script setup>
import json from '../assets/sampleQuestions.json'

export default {
  data() {
    return {
      tasksList: json
    }
  }
}
</script>

<template>
  <li v-for="task in tasksList" :key="task.id">
    @{{ task.question }}
  </li>
</template>

<style scoped>

</style>