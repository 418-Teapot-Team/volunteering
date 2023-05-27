<template>
  <section id="auth" class="flex flex-row items-center justify-center h-screen w-full">
    <vee-form
      class="flex flex-col justify-center items-center p-6 gap-2 border w-96"
      :validation-schema="schema"
      @submit="onSubmit"
    >
      <div class="flex flex-col justify-start iems-start self-start mb-4">
        <span class="text-4xl text-black">New account</span>
        <span class="text-sm text-grey-light"
          >Already registered? <RouterLink to="/sign-in" class="text-grey">Login</RouterLink></span
        >
      </div>
      <div class="flex flex-row justify-between w-full gap-4 h-20">
        <div class="flex flex-col justify-start items-start gap-1 w-full">
          <vee-field
            name="first_name"
            type="text"
            placeholder="First name*"
            class="p-2 rounded-lg shadow-md outline-none w-full h-12 border"
          ></vee-field>
          <ErrorMessage name="first_name" class="text-red-500 text-xs" />
        </div>
        <div class="flex flex-col justify-start items-start gap-1 w-full">
          <vee-field
            name="last_name"
            type="text"
            placeholder="Last name*"
            class="p-2 rounded-lg shadow-md outline-none w-full h-12 border"
          ></vee-field>
          <ErrorMessage name="last_name" class="text-red-500 text-xs" />
        </div>
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
import { mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'RegisterPage',
  data() {
    return {
      schema: {
        first_name: 'required|alpha_spaces|min:3|max:100',
        last_name: 'required|alpha_spaces|min:3|max:100',
        email: 'required|email|min:3|max:100',
        password: 'required|min:6|max:100',
      },
    };
  },
  methods: {
    ...mapActions(useAuthStore, ['register']),
    async onSubmit(values) {
      const payload = {
        firstName: values?.first_name,
        lastName: values?.last_name,
        email: values?.email,
        password: values?.password,
      };
      try {
        await this.register(payload);
        toast.success('Account successfully created! Please login');
        this.$router.push('/sign-in');
      } catch (e) {
        toast.error(e?.message);
      }
    },
  },
};
</script>
