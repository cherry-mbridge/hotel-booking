<template>
  <div class="p-10">
    <h1 class="text-2xl mb-4">Admin Dashboard</h1>
    <p>Welcome</p>
    <button class="bg-red-500 text-white px-4 py-2 rounded" @click="onLogout">Logout</button>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: ['admin-auth']
})
import useAuth from '~/composables/useAuth'
const { logout } = useAuth()

async function onLogout() {
  await logout()
  navigateTo('/admin')
}
</script>
