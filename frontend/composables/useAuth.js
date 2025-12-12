export default function useAuth() {
  const isAdminLoggedIn = useState('isAdminLoggedIn', () => false)  
  const adminType = useState('adminType', () => null)
  const errorMsg = ref("")

  async function fetchSession() {
    try {
      const data = await $fetch('http://localhost:8080/api/admin/me', { credentials: 'include' })
      isAdminLoggedIn.value = data?.type === 'admin'
      adminType.value = data?.type === 'admin' ? 'admin' : null
    } catch {
      isAdminLoggedIn.value = false
      adminType.value = null
    }
  }

  async function login(email, password) {
    try {
      await $fetch('http://localhost:8080/api/admin/login', {
        method: 'POST',
        body: { email, password },
        credentials: 'include',
      })
      return navigateTo('/admin/dashboard')
    } catch (e) {
      errorMsg.value = e.data?.error || 'Login failed'
    }
  }

  async function logout() {
    try {
      await $fetch('http://localhost:8080/api/admin/logout', { credentials: 'include' })
    } finally {
      isAdminLoggedIn.value = false
      adminType.value = null
    }
  }

  return { isAdminLoggedIn, adminType, errorMsg, fetchSession, login, logout }
}
