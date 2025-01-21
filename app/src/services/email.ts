import { apiCall } from '@/services/utils'

export async function loadEmail(filters: Record<string, string>) {
  const query = new URLSearchParams(filters).toString()
  const baseURL = query ? `/email?${query}` : '/email'
  const data = await apiCall(baseURL)
  return data
}
