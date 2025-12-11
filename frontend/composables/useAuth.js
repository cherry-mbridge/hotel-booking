export default function useAuth() {
  const isLoggedIn = useState('isLoggedIn', () => false)  
  const userType = useState('userType', () => null)

  async function fetchSession() {
    try {
      const data = await $fetch('http://localhost:8080/api/admin/me', { credentials: 'include' })
      isLoggedIn.value = !!data?.type
      userType.value = data.type ?? null
    } catch {
      isLoggedIn.value = false
      userType.value = null
    }
  }

  async function login(email, password) {
    try {
      await $fetch('http://localhost:8080/api/admin/login', {
        method: 'POST',
        body: { email, password },
        credentials: 'include',
      })
      return navigateTo('/admin/bookings')
    } catch (e) {
      throw new Error('Login failed')
    }
  }

  async function logout() {
    try {
      await $fetch('http://localhost:8080/api/admin/logout', { credentials: 'include' })
    } finally {
      isLoggedIn.value = false
      userType.value = null
    }
  }

  return { isLoggedIn, userType, fetchSession, login, logout }
}
