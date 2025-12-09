export default defineNuxtRouteMiddleware(async () => {
  const { fetchSession, isLoggedIn, userType } = useAuth()

  if (!isLoggedIn.value || !userType.value) {
    await fetchSession()
  }

  if (!isLoggedIn.value || userType.value !== 'admin') {
    return navigateTo('/admin')
  }
})
