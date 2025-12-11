<template>
  <div>
    <!-- <Navbar /> -->
    <h1>Welcome {{ user?.name }}</h1>
    <button @click="handleLogout">Logout</button>
  </div>
</template>

<script setup>
definePageMeta({
  middleware: 'user-auth'
})

const { user, logout } = useUserAuth()

const handleLogout = async () => {
    await logout()
    navigateTo('/login')
}
</script>
