import { ref, computed, onMounted } from 'vue'
import { loadEmail } from '@/services/email'
import { loadTotal } from '@/services/total'
import type { Inbox } from '@/types'

// Hook encargado de la carga de emails y la paginación
export function useEmail() {
  const inbox = ref<Inbox[]>([]) // Lista de emails
  const total = ref(0) // Número Total de emails
  const loading = ref(false)
  const loadingDirection = ref('') // Dirección de la paginación
  const pag = ref('0') // Página actual
  const search = ref<string>('')
  const error = ref<string | null>(null)

  // Estados para manejar el modal
  const selectedEmail = ref<Inbox | null>(null) // Email seleccionado
  const isModalOpen = ref(false) // Estado del modal

  // Carga los emails y el número total de emails
  async function loadData(p?: string) {
    try {
      loading.value = true // Activa el spiner
      error.value = null // Resetea errores antes de una nueva consulta
      // Realiza ambas peticiones en paralelo
      const [inboxData, totalData] = await Promise.all([
        loadEmail({ from: p || pag.value, search: search.value }),
        loadTotal({ from: p || pag.value, search: search.value })
      ])
      if (p) pag.value = p // Actualiza la página si se pasa un valor nuevo
      inbox.value = inboxData // Asigna la lista de emails
      if (totalData !== -1) total.value = totalData // Asigna el número total de emails
    } catch (err) {
      error.value = 'Failed to load emails. Please try again.'
    } finally {
      loading.value = false // Desactiva el spiner
    }
  }

  // Paginación
  function changePages(direction: string) {
    let pagInt = parseInt(pag.value)

    if (direction === 'back') {
      pagInt = pagInt - 5
      if (pagInt < 0) {
        pagInt = 0
      }
    } else {
      pagInt = pagInt + 5
      if (pagInt > total.value) {
        pagInt = total.value / 5
      }
    }

    pag.value = pagInt.toString()
    loadData()
    loadingDirection.value = direction
  }

  const currentPage = computed(() => Math.ceil((parseInt(pag.value) + 5) / 5))
  const totalPages = computed(() => Math.ceil(total.value / 5))

  function openModal(email: Inbox) {
    selectedEmail.value = email // Guarda el email seleccionado
    isModalOpen.value = true // Abre el modal
  }

  function closeModal() {
    isModalOpen.value = false // Cierra el modal
  }

  onMounted(loadData) // Carga al montar el componente

  // Permite que estos valores sean accesibles desde setup() en los componentes Vue
  return {
    inbox,
    total,
    loading,
    loadingDirection,
    pag,
    search,
    error,
    selectedEmail,
    isModalOpen,
    loadData,
    changePages,
    openModal,
    closeModal,
    currentPage,
    totalPages
  }
}
