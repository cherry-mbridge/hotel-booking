export default defineNuxtRouteMiddleware(async () => {
  const { fetchSession, isAdminLoggedIn, adminType } = useAuth()

  if (!isAdminLoggedIn.value || !adminType.value) {
    await fetchSession()
  }

  if (!isAdminLoggedIn.value || adminType.value !== 'admin') {
    return navigateTo('/admin')
  }
})
