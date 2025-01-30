import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL), // Permite rutas limpias sin hash mode
  routes: [
    {
      path: '/email',
      name: 'email',
      component: () => import('../views/EmailView.vue')
    },
    {
      path: '/',
      redirect: '/email'
    }
  ]
})

export default router
