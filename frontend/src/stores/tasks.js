import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';

export default defineStore('tasks', {
  state: () => ({
    sharedTasks: [],
    myTasks: [],
    applies: [],
    pendingTasks: [],
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
    async finishTask(payload) {
      await HttpClient.post('api/v1/done-volunteer', {
        id: payload.id,
        loggedHours: Number(payload.hours),
      });
    },
    async dismissTask(payload) {
      await HttpClient.delete('api/v1/tasks', {
        data: {
          id: payload.id,
        },
      });
    },
    async getPendingTasks() {
      const { data } = await HttpClient.get('api/v1/tasks/pending');
      this.pendingTasks = data.result;
    },
    async approveTask(payload) {
      await HttpClient.post('api/v1/done-employer?done=true', {
        id: payload.id,
        loggedHours: payload.loggedHours,
      });
    },
    async denyTask(payload) {
      await HttpClient.delete('api/v1/tasks', {
        data: {
          id: payload.id,
        },
      });
    },
  },
});
