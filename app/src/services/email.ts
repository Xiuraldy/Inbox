import type { Inbox } from '@/types'
import { apiCall } from '@/services/utils'
import { ref } from 'vue'

export function useEmail() {
  const inbox = ref<Inbox[]>([])

  async function loadEmail(filters: Record<string, string>) {
    const query = new URLSearchParams(filters).toString()
    const baseURL = query ? `/email?${query}` : '/email'
    const data = await apiCall(baseURL)
    inbox.value = data
  }

  return { loadEmail, inbox }
}

export async function loadEmail(filters: Record<string, string>) {
  const query = new URLSearchParams(filters).toString()
  const baseURL = query ? `/email?${query}` : '/email'
  const data = await apiCall(baseURL)
  return data
}
