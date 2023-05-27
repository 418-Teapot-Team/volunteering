<template>
  <section id="auth" class="flex flex-row items-center justify-center h-screen w-full">
    <vee-form
      class="flex flex-col justify-center items-center p-6 gap-2 border w-96"
      :validation-schema="schema"
      @submit="onSubmit"
    >
      <div class="flex flex-col justify-start iems-start self-start mb-4">
        <span class="text-4xl text-black">Login</span>
        <span class="text-sm text-grey-light"
          >Do not have an account?
          <RouterLink to="/sign-up" class="text-grey">Sign Up</RouterLink></span
        >
      </div>
      <div class="w-full h-20">
        <vee-field
          name="email"
          type="email"
          placeholder="Email*"
          class="p-2 rounded-lg shadow-md outline-none w-full h-12 border"
        ></vee-field>
        <ErrorMessage name="email" class="text-red-500 text-xs" />
      </div>
      <div class="w-full h-20">
        <vee-field
          name="password"
          type="password"
          placeholder="Password*"
          class="p-2 rounded-lg shadow-md outline-none w-full h-12 border"
        ></vee-field>
        <ErrorMessage name="password" class="text-red-500 text-xs" />
      </div>
      <div class="w-full h-12"><app-button type="submit" text="Continue" /></div>
    </vee-form>
  </section>
</template>
<script>
import useAuthStore from '@/stores/auth';
import { mapActions, mapState } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'AuthPage',
  data() {
    return {
      schema: {
        email: 'required|email|min:3|max:100',
        password: 'required|min:6|max:100',
      },
    };
  },
  computed: {
    ...mapState(useAuthStore, ['isAuthorized']),
  },
  methods: {
    ...mapActions(useAuthStore, ['login']),
    async onSubmit(values) {
      const payload = {
        email: values?.email,
        password: values?.password,
      };
      try {
        await this.login(payload);
        toast.success('Welcome!');
        console.log(this.isAuthorized);
        this.$router.push('/');
      } catch (e) {
        toast.error(e?.message);
      }
    },
  },
};
</script>
