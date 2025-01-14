import { createRouter, createWebHistory } from 'vue-router'
import EmailView from '../views/EmailView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/email',
      name: 'email',
      component: () => EmailView
    }
  ]
})

export default router
