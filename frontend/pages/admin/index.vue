<template>
  <div>
    <h1>Login</h1>
    <form @submit.prevent="onLogin">
      <input v-model="email" placeholder="Email" />
      <input v-model="password" type="password" placeholder="Password" />
      <button type="submit">Login</button>
    </form>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: ['guest-admin']
})
import { ref, watchEffect } from 'vue'
import { useRouter } from 'vue-router'
import useAuth from '~/composables/useAuth'

const router = useRouter()
const { login, isLoggedIn } = useAuth()
const email = ref('')
const password = ref('')

// Watch login state to redirect automatically
// watchEffect(() => {
//   if (isLoggedIn.value) {
//     router.push('/admin/dashboard')
//   }
// })

async function onLogin() {
  try {
    await login(email.value, password.value)
    // No need to manually redirect; watchEffect handles it
  } catch (e) {
    alert(e.message)
  }
}
</script>
