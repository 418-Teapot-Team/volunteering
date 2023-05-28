import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';

export default defineStore('tasks', {
  state: () => ({
    sharedTasks: [],
    myTasks: [],
  }),
  actions: {
    async getSharedTasks() {
      const { data } = await HttpClient.get('api/v1/tasks/shared');
      this.sharedTasks = data.result;
    },
    async getMyTasks() {
      const { data } = await HttpClient.get('api/v1/tasks');
      this.myTasks = data.result;
    },
  },
});
