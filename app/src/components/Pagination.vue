<template>
  <div class="flex flex-col items-center mt-2">
    <div class="flex justify-center">
      <!-- ⬅ Botón para retroceder (solo visible si no estamos en la primera página) -->
      <button 
        v-if="parseInt(pag || '0') > 0" 
        class="relative bg-primary text-white m-1 text-3xl px-8 h-8 rounded transition-all duration-300 flex justify-center items-center hover:bg-secondary hover:px-12"
        :disabled="loading"
        @click="changePages('back')"
      >
        <Spinner :loading="loading && loadingDirection === 'back'" />
        <span v-if="!(loading && loadingDirection === 'back')">⬅</span>
      </button>

      <!-- ⮕ Botón para avanzar (solo visible si hay más páginas disponibles) -->
      <button 
        v-if="parseInt(pag || '0') + 5 < (total || 0)"
        class="relative bg-primary text-white m-1 text-3xl px-8 h-8 rounded transition-all duration-300 flex justify-center items-center hover:bg-secondary hover:px-12"
        :disabled="loading"
        @click="changePages('next')"
      >
        <Spinner :loading="loading && loadingDirection === 'next'" />
        <span v-if="!(loading && loadingDirection === 'next')">⮕</span>
      </button>
    </div>

    <!-- Indicador de página -->
    <p v-if="inbox?.length !== 0" class="text-sm text-gray-500">Page {{ currentPage }} of {{ totalPages }}</p>
    <p v-else><b>No emails found matching your search ✘</b></p>
  </div>
</template>

<script setup lang="ts">
import type { Inbox } from '@/types';
import Spinner from '../components/Spinner.vue'

// Extrae props 
defineProps<{
  inbox: Inbox[]; 
  pag: string; 
  total: number; 
  loading: boolean; 
  loadingDirection: string; 
  changePages: (direction: string) => void;
  currentPage: number; 
  totalPages: number; 
}>()
</script>
