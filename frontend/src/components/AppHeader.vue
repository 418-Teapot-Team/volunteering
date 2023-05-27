<template>
  <header
    class="relative w-full flex flex-row justify-between px-8 h-14 shadow-sm bg-primary text-white"
  >
    <div class="flex flex-row gap-16">
      <div class="flex flex-col justify-between items-start h-full">
        <span class="text-sm font-bold">hand</span>
        <span class="text-xs">to</span>
        <span class="text-sm font-bold">hand</span>
      </div>
      <div class="flex flex-row justify-start gap-6 items-center h-full">
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/search-for-tasks"
          >Search for tasks</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/projects"
          >Projects</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/my-tasks"
          >My Tasks</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/stats"
          >Statistics</RouterLink
        >
      </div>
    </div>
    <div class="flex flex-row self-end justify-end items-center h-full gap-2">
      <span>Welcome, {{ fullname }}!</span>
      <div class="h-full flex justify-center items-center cursor-pointer" @click="logout">
        <LogoutIcon />
      </div>
    </div>
  </header>
</template>

<script>
import LogoutIcon from '@/components/icons/LogoutIcon.vue';

export default {
  name: 'AppHeader',
  components: {
    LogoutIcon,
  },
  props: ['fullname'],
  methods: {
    logout() {
      localStorage.removeItem('auth-token');
      this.$router.push('/sign-in');
    },
  },
};
</script>
