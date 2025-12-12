<template>
  <div class="container">
    <h2>User Register</h2>

    <form @submit.prevent="submit">
      <input v-model="name" type="text" placeholder="Name" required />
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />

      <button type="submit">Register</button>
    </form>

    <p v-if="userErrMsg" class="text-red-500">{{ userErrMsg }}</p>

    <p>
      Already have an account?
      <NuxtLink to="/login">Login</NuxtLink>
    </p>
  </div>
</template>

<script setup>
const name = ref("")
const email = ref("")
const password = ref("")

const { register, userErrMsg } = useUserAuth()

const submit = async () => {
  try {
    await register(name.value, email.value, password.value)
  } catch (e) {
    alert('Registration failed')
  }
}
</script>
