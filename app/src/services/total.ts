import { apiCall } from '@/services/utils'
import { ref } from 'vue'

export function useTotal() {
  const total = ref(0)

  async function loadTotal(filters: Record<string, string>) {
    const query = new URLSearchParams(filters).toString()
    const baseURL = query ? `/total?${query}` : '/total'
    const data = await apiCall(baseURL)
    total.value = data[0]['count(*)']
    console.log('data0', data[0])
  }

  return { loadTotal, total }
}

export async function loadTotal(filters: Record<string, string>) {
  const query = new URLSearchParams(filters).toString()
  const baseURL = query ? `/total?${query}` : '/total'
  const data = await apiCall(baseURL)
  console.log('data0', data[0])
  return data[0]['count(*)']
}
