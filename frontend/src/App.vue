<template>
  <div class="app-container">
    <router-view />
  </div>
</template>

<script setup>
import LoginView from './views/loginView.vue'
</script>

<style scoped>
.app-container {
  display: flex;
  justify-content: center;
  height: 90vh;
  background-color: white;
}
</style>
