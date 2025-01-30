<template>
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
</template>

<script setup lang="ts">
defineProps(['inbox', 'openModal'])
</script>
