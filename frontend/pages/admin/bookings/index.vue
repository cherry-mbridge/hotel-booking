<template>
  <div class="p-8">
    <h1 class="text-2xl font-bold mb-6">Booking List</h1>

    <div
      class="bg-white dark:bg-neutral-900 shadow rounded-xl overflow-hidden border dark:border-neutral-700"
    >
      <table class="w-full text-left">
        <thead class="bg-gray-100 dark:bg-neutral-800">
          <tr>
            <th class="p-4 border-b dark:border-neutral-700">User ID</th>
            <th class="p-4 border-b dark:border-neutral-700">Room ID</th>
            <th class="p-4 border-b dark:border-neutral-700">Check-in</th>
            <th class="p-4 border-b dark:border-neutral-700">Check-out</th>
            <th class="p-4 border-b dark:border-neutral-700 text-center">Actions</th>
          </tr>
        </thead>

        <tbody>
          <tr
            v-for="booking in bookings"
            :key="booking.id"
            class="hover:bg-gray-50 dark:hover:bg-neutral-800 transition"
          >
            <td class="p-4 border-b dark:border-neutral-700">{{ booking.user_id }}</td>
            <td class="p-4 border-b dark:border-neutral-700">{{ booking.room_id }}</td>
            <td class="p-4 border-b dark:border-neutral-700">
              {{ formatDate(booking.start_date) }}
            </td>
            <td class="p-4 border-b dark:border-neutral-700">
              {{ formatDate(booking.end_date) }}
            </td>

            <td class="p-4 border-b dark:border-neutral-700 text-center">
              <NuxtLink
                :to="`/bookings/${booking.id}`"
                class="px-3 py-1 text-sm bg-blue-600 text-white rounded-lg hover:bg-blue-700"
              >
                Show
              </NuxtLink>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script setup>
  definePageMeta({
  middleware: ['admin-auth']
})
const bookings = ref([])

const { data, error } = await useFetch("http://localhost:8080/api/admin/bookings", {
  credentials: "include"
})

console.log("DATA:", data)
console.log("ERROR:", error)

bookings.value = data.value?.data || []

function formatDate(date) {
  return new Date(date).toLocaleDateString()
}
</script>
