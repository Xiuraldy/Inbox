<script setup lang="ts">
import { RouterLink, RouterView } from 'vue-router'
import Linkedin from './components/Linkedin.vue';
</script>

<template>
  <nav class="w-full text-xs text-center flex items-center justify-between h-[8vh] bg-primary text-white px-8 transition-all duration-300">
    <h1 class="text-xl">Email</h1>
    <RouterLink 
      to="/email" 
      class="inline-block px-4 border-l border-border text-lg last:border-0 router-link-exact-active:text-white router-link-exact-active:bg-secondary router-link-exact-active:rounded-md router-link-exact-active:px-5 router-link-exact-active:hover:bg-transparent">
      Inbox
    </RouterLink>
  </nav>
  <RouterView />
  <footer class="h-[6vh] flex justify-center items-center bg-tertiary">
    <h3 class="mr-2">Code developed by </h3>
    <Linkedin/>
  </footer>
</template>

