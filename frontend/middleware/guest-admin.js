export default defineNuxtRouteMiddleware(async () => {
  const { fetchSession, isLoggedIn, userType } = useAuth()

  if (!isLoggedIn.value || !userType.value) {
    await fetchSession()
  }
console.log(isLoggedIn.value)
console.log(userType.value)
  if (isLoggedIn.value && userType.value === 'admin') {
    return navigateTo('/admin/dashboard')
  }
})
