export default function useUserAuth() {
  const isUserLoggedIn = useState('isUserLoggedIn', () => false)  
  const userType = useState('userType', () => null)
  const userErrMsg = ref("")

  async function fetchSession() {
    try {
      const data = await $fetch('http://localhost:8080/api/users/me', { credentials: 'include' })

      isUserLoggedIn.value = data?.type === 'user'
      userType.value = data?.type === 'user' ? 'user' : null
    } catch {
      isUserLoggedIn.value = false
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
        if (e.data?.errors) {
          userErrMsg.value = Object.values(e.data.errors).join(', ')
        } else {
          userErrMsg.value = e.data?.error || 'Registration failed'
        }
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
      userErrMsg.value = e.data?.error || 'Login failed'
    }
  }

  async function logout() {
    try {
      await $fetch('http://localhost:8080/api/users/logout', { 
        method: 'POST',
        credentials: 'include',
      })
    } finally {
      isUserLoggedIn.value = false
      userType.value = null
    }
  }

  return { isUserLoggedIn, userType, userErrMsg, fetchSession, register, login, logout }
}
