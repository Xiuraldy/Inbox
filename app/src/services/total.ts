import { apiCall } from '@/services/utils'

export async function loadTotal(filters: Record<string, string>) {
  const query = new URLSearchParams(filters).toString()
  const baseURL = query ? `/total?${query}` : '/total'
  const data = await apiCall(baseURL)
  if (data && data.length <= 0) {
    return -1
  }
  return data[0]['count(*)'] || -1
}
