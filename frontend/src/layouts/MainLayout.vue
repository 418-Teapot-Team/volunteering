<template>
  <app-header :fullname="`${user.firstName} ${user.lastName}`" class="mb-10" />
  <main class="container m-auto pb-6">
    <router-view />
  </main>
</template>

<script>
import useAuthStore from '@/stores/auth';
import { mapActions, mapState } from 'pinia';
import AppHeader from '@/components/AppHeader.vue';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'MainLayout',
  components: {
    AppHeader,
  },
  watch: {
    isAuthorized(val) {
      if (val === false) {
        this.$router.push('/sign-in');
      }
    },
    '$route.path'() {
      if (this.isAuthorized === false) {
        this.$router.push('/sign-in');
      }
    },
  },
  computed: {
    ...mapState(useAuthStore, ['isAuthorized', 'user']),
  },
  methods: {
    ...mapActions(useAuthStore, ['whoAmI']),
    async getCurrentUser() {
      try {
        await this.whoAmI();
      } catch (e) {
        toast.error(e?.message);
        console.log(e);
        if (e?.code === 403) {
          localStorage.removeItem('auth-token');
          this.$router.push('/sign-in');
        }
      }
    },
  },
  mounted() {
    if (!this.isAuthorized) {
      this.$router.push('/sign-in');
    } else {
      this.getCurrentUser();
    }
  },
};
</script>
