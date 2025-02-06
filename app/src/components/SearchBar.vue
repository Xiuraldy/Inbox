<template>
  <div class="flex items-center ml-5 justify-start gap-2 mb-2 md:justify-center md:ml-0">
    <!-- Campo de entrada para búsqueda -->
    <input 
      type="text" 
      placeholder="Search Email..." 
      class="w-2/3 max-w-md rounded-md p-2 text-sm border border-gray-300 outline-none" 
      :value="modelValue"
      :disabled="loading"
      @input="$emit('update:modelValue', ($event.target as HTMLInputElement).value)"
      @keyup.enter="$emit('searchTriggered')"
    >
     <!-- Botón de búsqueda -->
    <button 
      class="relative bg-primary text-white text-xs w-100px px-8 h-8 rounded transition-all duration-300 flex justify-center items-center hover:bg-secondary md:text-sm md:w-150px"
      :disabled="loading"
      @click="$emit('searchTriggered')"
    >
      <Spinner :loading="loading" />
      {{ loading ? "" : "Search 🔍︎" }}
    </button>
  </div>
</template>

<script setup lang="ts">
import Spinner from '../components/Spinner.vue'

defineProps<{
  modelValue: string,
  loading: boolean
}>()

defineEmits([
  'update:modelValue',  // Se emite cuando el usuario escribe para actualizar `v-model`
  'searchTriggered'  // Se emite cuando el usuario presiona "Enter" o hace clic en "Search"
  ])
</script>
