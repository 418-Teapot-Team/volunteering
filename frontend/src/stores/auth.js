import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';
export default defineStore('auth', {
  state: () => ({
    user: {},
    token: '',
  }),
  getters: {
    isAuthorized() {
      return !!localStorage.getItem('auth-token') || this.token?.trim() !== '';
    },
  },
  actions: {
    async login(payload) {
      const { data } = await HttpClient.post('auth/v1/sign-in', {
        email: payload.email,
        password: payload.password,
      });
      localStorage.setItem('auth-token', data?.token);
      this.token = data?.token;
    },
    async register(payload) {
      await HttpClient.post('auth/v1/sign-up', {
        firstName: payload.firstName,
        lastName: payload.lastName,
        email: payload.email,
        password: payload.password,
      });
    },

    async whoAmI() {
      const { data } = await HttpClient.get('api/v1/who-am-i');
      this.user = data;
    },
  },
});
