<script setup lang="ts">
import SearchBar from '@/components/SearchBar.vue'
import EmailTable from '@/components/EmailTable.vue'
import Pagination from '@/components/Pagination.vue'
import EmailModal from '@/components/EmailModal.vue'
import { useEmail } from '@/composables/useEmail';

// Extrae variables y funciones de useEmail() para usarlas en la vista
const {
  inbox,
  total,
  loading,
  loadingDirection,
  pag,
  search,
  selectedEmail,
  isModalOpen,
  loadData,
  changePages,
  openModal,
  closeModal,
  currentPage,
  totalPages
} = useEmail()
</script>

<template>
  <section>
    <SearchBar 
      v-model="search" 
      :loading="loading" 
      @searchTriggered="loadData('0')" 
    />
    <EmailTable :inbox="inbox" :openModal="openModal" />
    <Pagination 
      :inbox="inbox" 
      :pag="pag" 
      :total="total" 
      :loading="loading" 
      :loadingDirection="loadingDirection"
      :changePages="changePages"
      :currentPage="currentPage"
      :totalPages="totalPages"
    />
    <EmailModal
      :isModalOpen="isModalOpen" 
      :selectedEmail="selectedEmail" 
      :closeModal="closeModal" 
    />
  </section>
</template>