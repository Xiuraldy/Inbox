<template>
    <div 
      v-if="loading" 
      class="absolute inset-0 flex items-center justify-center"
    >
      <svg 
        class="animate-spin h-5 w-5 text-white" 
        xmlns="http://www.w3.org/2000/svg" 
        fill="none" 
        viewBox="0 0 24 24"
      >
        <circle 
          class="opacity-25" 
          cx="12" cy="12" r="10" stroke="currentColor" stroke-width="4"
        ></circle>
        <path 
          class="opacity-75" 
          fill="currentColor" 
          d="M4 12a8 8 0 018-8v4a4 4 0 00-4 4H4z"
        ></path>
      </svg>
    </div>
  </template>
  
  <script setup lang="ts">
  defineProps({
    loading: {
      type: Boolean,
      required: true,
    },
  });
  </script>
  