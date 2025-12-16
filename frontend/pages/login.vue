<template>
  <div class="container">
    <h2>Login</h2>

    <form @submit.prevent="submit">
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />

      <button type="submit">Login</button>
    </form>

    <p v-if="userErrMsg" class="text-red-500">{{ userErrMsg }}</p>

    <p>
      No account?
      <NuxtLink to="/register">Register</NuxtLink>
    </p>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'guest-user' 
})

const email = ref("")
const password = ref("")

const { login, userErrMsg } = useUserAuth()

const submit = async () => {
  try {
    await login(email.value, password.value)
  } catch (e) {
    alert('Invalid login')
  }
}
</script>
