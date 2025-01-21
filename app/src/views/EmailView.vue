<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { loadEmail } from '@/services/email'
import type { Inbox } from '@/types';
import { loadTotal } from '@/services/total';
import Spinner from '../components/Spinner.vue';

const inbox = ref<Inbox[]>([])
const total = ref(0)
const loading = ref(false)

let pag = '0';
const search = ref<string>('');

function changePages(direction: string) {
  var pagInt = parseInt(pag)

  if(direction == "back") {
    pagInt = pagInt - 5
    if (pagInt < 0) {
      pagInt = 0
    }
  } else {
    pagInt = pagInt + 5
    if (pagInt > total.value) {
      pagInt = total.value/5
    }
  }

  
  pag = pagInt.toString()
  loadData()
}

const selectedEmail = ref<Inbox | null>(null);
const isModalOpen = ref(false);

function openModal(email: any) {
  selectedEmail.value = email;
  isModalOpen.value = true; 
}

function closeModal() {
  isModalOpen.value = false; 
}

onMounted(loadData);

async function loadData(p?: string) {
    loading.value = true
    const [inboxData, totalData] = await Promise.all([loadEmail({ from: p || pag, search: search.value }), loadTotal({ from: p || pag, search: search.value })])
    if(p) {
      pag = p
    }
    inbox.value = inboxData
    if(totalData !== -1) {
      total.value = totalData
    } 
    loading.value = false
}

</script>

<template>
  <section>
    <div class="flex items-center justify-center gap-2 mb-4">
      <input 
        type="text" 
        placeholder="Search Email..." 
        class="w-2/3 max-w-md rounded-md p-2 text-sm border border-gray-300 outline-none"
        v-model="search" 
        :disabled='loading'
      >
      <button 
        class="relative bg-primary text-white text-sm w-32 h-10 rounded transition-all duration-300 hover:bg-secondary flex items-center justify-center"
        @click="() => { loadData('0') }"
      >
        <Spinner :loading="loading" />
        {{ loading ? "" : "Search 🔍︎" }}
      </button>
    </div>
    <div class="overflow-x-auto ml-5">
      <table class="border-collapse text-base w-max">
        <tbody>
          <tr 
            v-for="email in inbox" 
            :key="email.message_id" 
            class="hover:bg-gray-200 cursor-pointer"
            @click="openModal(email)"
          >
            <td class="border p-1 text-center align-middle">
              <span class="font-semibold">{{ email.date }}</span>
              <br/>
              <span class="bg-tertiary text-white px-1 text-sm py-0.5 rounded">From:</span> {{ email.from }}
              <br/>
              <span class="bg-secondary text-white px-1 text-sm py-0.5 rounded">To:</span> {{ email.to }}
            </td>
            <td class="border p-2 text-center align-middle w-[800px] whitespace-nowrap overflow-hidden text-ellipsis flex flex-col">
              <span class="bg-primary text-white px-2 -mt-2 -ml-2 -mr-2">{{ email.subject }}</span>
              <br/><span class="overflow-hidden">{{ email.body }}</span>
            </td>
            <td class="border p-2 text-center align-middle italic">
              {{ email._timestamp }}<br/>{{ email.message_id }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="flex flex-col items-center">
      <div class="flex justify-center">
        <button 
          class="relative bg-primary text-white text-3xl w-36 h-10 rounded transition-all duration-300 hover:bg-secondary hover:w-48"
          :disabled='loading'
          @click="changePages('back')"
          v-if="parseInt(pag) > 0 && inbox.length !== 0" 
        >
          <Spinner :loading="loading"/>
          {{loading ? "" : "⬅" }}
        </button>
        <button 
          class="relative bg-primary text-white text-3xl w-36 h-10 rounded transition-all duration-300 hover:bg-secondary hover:w-48"
          :disabled="loading"
          @click="changePages('next')"
          v-if="parseInt(pag) + 5 < total && inbox.length !== 0" 
        >
          <Spinner :loading="loading"/>
          {{loading ? "" : "⮕"}}
        </button>

      </div>
      <p v-if="inbox.length !== 0">Page {{ Math.ceil((parseInt(pag) + 5) / 5) }} of {{ Math.ceil(total/5) }}</p>
      <p v-else><b>No emails found matching your search ✘</b></p>
    </div>
  </section>

  <div v-if="isModalOpen" class="fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex justify-center items-center z-50" @click.self="closeModal">
    <div class="bg-white rounded-lg p-5 shadow-md w-4/5 max-w-lg h-4/5 overflow-auto flex flex-col justify-between">
      <h6 class="text-gray-500 mb-2 text-xs">{{ selectedEmail?._timestamp }} | {{ selectedEmail?.message_id }}</h6>
      <h4 class="italic text-gray-600 -mb-2">{{ selectedEmail?.date }}</h4>
      <h2 class="text-tertiary text-lg font-bold text-xl -mb-2">{{ selectedEmail?.subject }}</h2>
      <hr class="border-2 border-tertiary">
      <div class="mt-2 mb-2 text-gray-500 italic">
        <div><strong>☺ From:</strong> {{ selectedEmail?.from }}</div>
        <div><strong>☺ To:</strong> {{ selectedEmail?.to }}</div>
      </div>
      <p class="p-2 border-2 border-dashed border-gray-500 rounded-md"><strong>Body:</strong> {{ selectedEmail?.body }}</p>
      <button 
        class="mt-4 p-2 bg-primary text-white rounded hover:bg-secondary transition-all"
        @click="closeModal"
      >Close</button>
    </div>
  </div>
</template>
