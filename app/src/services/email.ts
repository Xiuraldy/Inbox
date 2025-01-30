import { apiCall } from '@/services/utils'

export async function loadEmail(filters: Record<string, string>) {
  const query = new URLSearchParams(filters).toString() // Convierte los filtros en cadena de consulta
  const baseURL = query ? `/email?${query}` : '/email'
  const data = await apiCall(baseURL) // Respuesta del API
  return data
}
