<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { loadEmail } from '@/services/email'
import type { Inbox } from '@/types';
import { loadTotal } from '@/services/total';
import Spinner from '../components/Spinner.vue';

const inbox = ref<Inbox[]>([])
const total = ref(0)
const loading = ref(false)
const loadingDirection = ref("")

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
  loadingDirection.value = direction
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
    <div class="flex items-center ml-5 justify-start gap-2 mb-2 md:justify-center md:ml-0">
      <input 
        type="text" 
        placeholder="Search Email..." 
        class="w-2/3 max-w-md rounded-md p-2 text-sm border border-gray-300 outline-none"
        v-model="search" 
        :disabled='loading'
        v-on:keyup.enter="() => { loadData('0') }"
      >
      <button 
        class="relative bg-primary text-white text-xs w-100px px-8 h-8 rounded transition-all duration-300 flex justify-center items-center hover:bg-secondary md:text-sm md:w-150px"
        @click="() => { loadData('0') }"
      >
        <Spinner :loading="loading" />
        {{ loading ? "" : "Search 🔍︎" }}
      </button>
    </div>
    <div class="overflow-x-auto flex justify-start ml-5 md:justify-center md:ml-0">
      <table class="border-collapse w-90 text-xs">
        <tbody>
          <tr 
            v-for="email in inbox" 
            :key="email.message_id" 
            class="hover:bg-gray-200 cursor-pointer"
            @click="openModal(email)"
          >
            <td class="border p-5 align-middle text-left">
              <span class="text-tertiary">From:</span><br/> {{ email.from }}
              <br/>
              <span class="text-secondary">To:</span><br/> {{ email.to }}
            </td>
            <td class="border p-5 text-left align-middle">
              <span class="overflow-hidden font-bold underline">{{ email.subject }}</span>
              <span class="text-gray-500 px-2 italic">{{ email.date }}</span>
              <span class="line-clamp-2 overflow-hidden text-ellipsis h-full">{{ email.body }}</span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="flex flex-col items-center mt-2">
      <div class="flex justify-center">
        <button 
          class="relative bg-primary text-white text-3xl px-8 h-8 rounded transition-all duration-300 flex justify-center items-center hover:bg-secondary hover:px-12"
          :disabled='loading'
          @click="() =>changePages('back')"
          v-if="parseInt(pag) > 0 && inbox.length !== 0" 
        >
          <Spinner :loading="loading && loadingDirection === 'back'"/>
          {{loading && loadingDirection === 'back' ? "" : "⬅" }}
        </button>
        <button 
          class="relative bg-primary text-white text-3xl px-8 h-8 rounded transition-all duration-300 flex justify-center items-center hover:bg-secondary hover:px-12"
          :disabled="loading"
          @click="changePages('next')"
          v-if="parseInt(pag) + 5 < total && inbox.length !== 0" 
        >
          <Spinner :loading="loading && loadingDirection === 'next'"/>
          {{loading && loadingDirection === 'next' ? "" : "⮕"}}
        </button>

      </div>
      <p v-if="inbox.length !== 0" class="text-sm text-gray-500">Page {{ Math.ceil((parseInt(pag) + 5) / 5) }} of {{ Math.ceil(total/5) }}</p>
      <p v-else><b>No emails found matching your search ✘</b></p>
    </div>
  </section>

  <div v-if="isModalOpen" class="fixed top-0 left-0 w-full h-full bg-black bg-opacity-50 flex justify-center items-center z-50" @click.self="closeModal">
    <div class="bg-white rounded-lg p-10 w-90 h-80 shadow-md">
      <div class="h-90">
        <h6 class="text-gray-500 text-xs mb-5">{{ selectedEmail?._timestamp }} | {{ selectedEmail?.message_id }}</h6>
        <h4 class="italic text-gray-600">{{ selectedEmail?.date }}</h4>
        <h2 class="text-tertiary text-lg font-bold text-xl">{{ selectedEmail?.subject }}</h2>
        <hr class="border-2 border-tertiary">
        <hr class="border-2 border-gray">
        <div class="my-2 text-gray-500 text-sm italic">
          <div><strong>☺ From:</strong> {{ selectedEmail?.from }}</div>
          <div><strong>☺ To:</strong> {{ selectedEmail?.to }}</div>
        </div>
        <div 
        class="py-5 h-70 overflow-y-auto text-sm text-gray-600"
          v-html="selectedEmail?.body
            .replace(/-----Original Message-----/g, '<br/><br/><hr/>')
            .replace(/(From:|Sent:|To:|Cc:|Subject:)/g, '<br/><strong>$1</strong>')"
        ></div>
      </div>
      <div class="h-10">
        <button 
          class="mt-4 p-2 bg-primary text-white rounded hover:bg-secondary transition-all"
          @click="closeModal"
        >Close</button>
      </div>
    </div>
  </div>
</template>
