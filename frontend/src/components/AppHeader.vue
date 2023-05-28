<template>
  <header
    class="hidden w-full lg:flex flex-row justify-between px-8 h-14 shadow-sm bg-primary text-white"
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
          active-class="border-b-2 border-b-white"
          exact
          >Search for tasks</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/projects"
          active-class="border-b-2 border-b-white"
          exact
          >Projects</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/my-tasks"
          active-class="border-b-2 border-b-white"
          exact
          >My Tasks</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/applies"
          active-class="border-b-2 border-b-white"
          exact
          >Applies</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/waiting"
          active-class="border-b-2 border-b-white"
          exact
          >Waiting for approval</RouterLink
        >
        <RouterLink
          class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
          to="/stats"
          active-class="border-b-2 border-b-white"
          exact
          >Statistics</RouterLink
        >
      </div>
    </div>
    <div class="flex flex-row self-end justify-end items-center h-full gap-2">
      <span :class="verified ? 'text-blue-300' : ''">Welcome, {{ fullname }}!</span>
      <div class="h-full flex justify-center items-center cursor-pointer" @click="logout">
        <LogoutIcon />
      </div>
    </div>
  </header>
  <header
    class="flex w-full lg:hidden flex-row justify-between px-8 h-14 shadow-sm bg-primary text-white"
  >
    <div class="flex flex-row gap-16">
      <div class="flex flex-col justify-between items-start h-full">
        <span class="text-sm font-bold">hand</span>
        <span class="text-xs">to</span>
        <span class="text-sm font-bold">hand</span>
      </div>
    </div>
    <div class="w-12 h-full cursor-pointer" @click="isMobileHeaderVisible = !isMobileHeaderVisible">
      <burger-icon />
    </div>
  </header>
  <div
    class="fixed top-14 left-0 right-0 bottom-0 lg:hidden bg-primary z-20"
    v-if="isMobileHeaderVisible"
  >
    <div class="flex flex-col justify-start gap-6 items-center h-full pt-6 text-white">
      <RouterLink
        class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
        to="/search-for-tasks"
        active-class="border-b-2 border-b-white"
        exact
        >Search for tasks</RouterLink
      >
      <RouterLink
        class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
        to="/projects"
        active-class="border-b-2 border-b-white"
        exact
        >Projects</RouterLink
      >
      <RouterLink
        class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
        to="/my-tasks"
        active-class="border-b-2 border-b-white"
        exact
        >My Tasks</RouterLink
      >
      <RouterLink
        class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
        to="/applies"
        active-class="border-b-2 border-b-white"
        exact
        >Applies</RouterLink
      >
      <RouterLink
        class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
        to="/waiting"
        active-class="border-b-2 border-b-white"
        exact
        >Waiting for approval</RouterLink
      >
      <RouterLink
        class="border-b-2 border-transparent hover:border-b-white cursor-pointer transition duration-300"
        to="/stats"
        active-class="border-b-2 border-b-white"
        exact
        >Statistics</RouterLink
      >
      <div class="flex flex-row items-center gap-2">
        <span :class="verified ? 'text-blue-300' : ''">Welcome, {{ fullname }}!</span>
        <div class="h-full flex justify-center items-center cursor-pointer" @click="logout">
          <LogoutIcon />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import LogoutIcon from '@/components/icons/LogoutIcon.vue';
import BurgerIcon from './icons/BurgerIcon.vue';

export default {
  name: 'AppHeader',
  components: {
    LogoutIcon,
    BurgerIcon,
  },
  props: ['fullname', 'verified'],
  watch: {
    '$route.path'() {
      this.isMobileHeaderVisible = false;
    },
  },
  data() {
    return {
      isMobileHeaderVisible: false,
    };
  },
  methods: {
    logout() {
      localStorage.removeItem('auth-token');
      this.$router.push('/sign-in');
    },
  },
};
</script>
