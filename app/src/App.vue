<script setup lang="ts">
import { RouterView } from 'vue-router'
import NavBar from './components/NavBar.vue';
import FooterEmail from './components/Footer.vue';
</script>

<template>
  <!-- Barra de navegación -->
  <NavBar
    title="Email"
    :links="[{ name: 'Inbox', path: '/email' }]"  
  /> 
  <RouterView /> <!-- Renderiza el componente correspondiente a la ruta actual -->
  <FooterEmail />
</template>

