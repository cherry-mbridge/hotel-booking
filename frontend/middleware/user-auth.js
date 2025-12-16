export default defineNuxtRouteMiddleware(async () => {
  const { fetchSession, isUserLoggedIn, userType } = useUserAuth()

  if (!isUserLoggedIn.value || !userType.value) {
    await fetchSession()
  }

  if (!isUserLoggedIn.value || userType.value !== 'user') {
    return navigateTo('/login')
  }
})
