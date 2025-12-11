export default defineNuxtRouteMiddleware(async () => {
  const { fetchSession, isLoggedIn, userType } = useUserAuth()

  if (!isLoggedIn.value || !userType.value) {
    await fetchSession()
  }

  if (!isLoggedIn.value || userType.value !== 'user') {
    return navigateTo('/login')
  }
})
