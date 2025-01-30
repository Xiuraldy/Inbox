const BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080'

interface Options {
  method?: string
  data?: any
  headers?: any
}
type ApiCallOptions = Options | undefined

export async function apiCall(
  path: string,
  { method = 'GET', data, headers }: ApiCallOptions = {}
) {
  // Hace una petición fetch a la API
  const response = await fetch(BASE_URL + path, {
    method,
    mode: 'cors', // Permite solicitudes entre dominios
    headers: {
      'Content-Type': 'application/json',
      ...headers
    },
    body: JSON.stringify(data) // Convierte data a JSON y lo envía en body
  })

  if (method === 'DELETE') {
    return response.ok // Devuelve true o false basado en si la petición fue exitosa
  }

  const jsonObj = await response.json() // Convierte la respuesta en JSON
  return jsonObj
}
