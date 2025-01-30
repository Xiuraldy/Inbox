<template>
  <div 
    v-if="isModalOpen" 
    class="fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex justify-center items-center z-50"
    @click.self="closeModal"
  >
    <div class="bg-white rounded-lg p-10 w-90 h-80 shadow-md">
      <div class="h-90">
        <!-- Información del email -->
        <h6 class="text-gray-500 text-xs mb-5">
          {{ props.selectedEmail?._timestamp }} | {{ props.selectedEmail?.message_id }}
        </h6>
        <h4 class="italic text-gray-600">{{ props.selectedEmail?.date }}</h4>
        <h2 class="text-tertiary text-lg font-bold text-xl">{{ props.selectedEmail?.subject }}</h2>
        
        <!-- Separadores -->
        <hr class="border-2 border-tertiary">
        <hr class="border-2 border-gray">

        <!-- Información del remitente -->
        <div class="my-2 text-gray-500 text-sm italic">
          <div><strong>☺ From:</strong> {{ props.selectedEmail?.from }}</div>
          <div><strong>☺ To:</strong> {{ props.selectedEmail?.to }}</div>
        </div>

        <!-- Cuerpo del email con formato -->
        <div 
          class="py-5 h-70 overflow-y-auto text-sm text-gray-600"
          v-html="formattedBody"
        ></div>
      </div>

      <!-- Botón para cerrar el modal -->
      <div class="h-10">
        <button 
          class="mt-4 p-2 bg-primary text-white rounded hover:bg-secondary transition-all"
          @click="props.closeModal"
        >
          Close
        </button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

// Extrae props 
const props = defineProps<{
  isModalOpen: boolean,
  selectedEmail: {
    _timestamp?: number;
    message_id?: string;
    date?: string;
    subject?: string;
    from?: string;
    to?: string;
    body?: string;
  } | null,
  closeModal: () => void
}>()

// Formatea el cuerpo del email 
const formattedBody = computed(() => {
  if (!props.selectedEmail?.body) return '';
  return props.selectedEmail.body
    .replace(/-----Original Message-----/g, '<br/><br/><hr/>')
    .replace(/(From:|Sent:|To:|Cc:|Subject:)/g, '<br/><strong>$1</strong>');
});
</script>
