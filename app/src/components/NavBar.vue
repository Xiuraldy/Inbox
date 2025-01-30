<script setup lang="ts">
import { RouterLink, useRoute } from 'vue-router';

// Recibir props para personalizar el navbar
defineProps<{
  title: string; 
  links: { name: string; path: string }[];
}>();

const route = useRoute();
</script>

<template>
  <nav class="w-full text-xs text-center flex items-center justify-between h-[8vh] bg-primary text-white px-8 transition-all duration-300">
    <!-- Título dinámico -->
    <h1 class="text-xl" aria-label="Navigation">{{ title }}</h1>

    <!-- Links dinámicos -->
    <div class="flex">
      <RouterLink
        v-for="link in links"
        :key="link.path"
        :to="link.path"
        class="inline-block px-4 border-l border-border text-lg last:border-0"
        :class="{ 'text-white bg-tertiary rounded-md px-5 hover:bg-secondary': route.path === link.path }">
        {{ link.name }}
      </RouterLink>
    </div>
  </nav>
</template>
