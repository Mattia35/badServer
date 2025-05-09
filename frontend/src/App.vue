<template>
  <div class="app-container">
    <LoginView />
  </div>
</template>

<script setup>
import LoginView from './views/loginView.vue'
</script>

<style scoped>
.app-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f9f9f9;
}
</style>
