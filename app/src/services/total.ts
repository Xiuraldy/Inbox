import { apiCall } from '@/services/utils'

export async function loadTotal(filters: Record<string, string>) {
  const query = new URLSearchParams(filters).toString() // Convierte los filtros en cadena de consulta
  const baseURL = query ? `/total?${query}` : '/total'
  const data = await apiCall(baseURL) // Respuesta del API
  if (data && data.length <= 0) {
    return -1 // Indica que no hay resultados
  }
  return data[0]['count(*)'] || -1
}
