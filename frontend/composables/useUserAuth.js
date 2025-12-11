export default function useUserAuth() {
  const isLoggedIn = useState('isLoggedIn', () => false)  
  const userType = useState('userType', () => null)

  async function fetchSession() {
    try {
      const data = await $fetch('http://localhost:8080/api/users/me', { credentials: 'include' })
      isLoggedIn.value = !!data?.type
      userType.value = data.type ?? null
    } catch {
      isLoggedIn.value = false
      userType.value = null
    }
  }

  async function register(name, email, password) {
    // Keep the $fetch logic as is...
    try {
        await $fetch('http://localhost:8080/api/users/register', {
            method: 'POST',
            body: { name, email, password }
        })
        return navigateTo('/login')
    } catch (e) {
        throw new Error('Registration failed.') 
    }
}

  async function login(email, password) {
    try {
      await $fetch('http://localhost:8080/api/users/login', {
        method: 'POST',
        body: { email, password },
        credentials: 'include',
      })
      return navigateTo('/dashboard')
    } catch (e) {
      throw new Error('Login failed')
    }
  }

  async function logout() {
    try {
      await $fetch('http://localhost:8080/api/users/logout', { 
        method: 'POST',
        credentials: 'include',
      })
    } finally {
      isLoggedIn.value = false
      userType.value = null
    }
  }

  return { isLoggedIn, userType, fetchSession, register, login, logout }
}
