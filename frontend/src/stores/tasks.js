import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';

export default defineStore('tasks', {
  state: () => ({
    sharedTasks: [],
    myTasks: [],
    applies: [],
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
    async getApplies() {
      const { data } = await HttpClient.get('api/v1/applies');
      this.applies = data.result;
    },
    async approveApply(payload) {
      await HttpClient.post('api/v1/applies/approve', {
        id: payload.id,
        applied_id: payload.applied_id,
      });
    },
    async createApply(payload) {
      await HttpClient.post('api/v1/apply', {
        taskId: payload.taskId,
        respondUserId: payload.userId,
      });
    },
  },
});
