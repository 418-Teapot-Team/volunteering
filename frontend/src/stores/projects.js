import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';

export default defineStore('projects', {
  state: () => ({
    projects: [],
  }),
  actions: {
    async createProject(payload) {
      await HttpClient.post('api/v1/projects', {
        title: payload.title,
        description: payload.description,
      });
    },
    async getProjects() {
      const { data } = await HttpClient.get('api/v1/projects');
      this.projects = data?.result;
    },
    async addTaskToProject(payload) {
      await HttpClient.post('api/v1/tasks', {
        projectId: payload.projectId,
        title: payload.title,
        estimate_time: +payload.estimate_time,
        description: payload.description,
      });
    },
    async markProjectTaskAsDone(payload) {
      await HttpClient.post('api/v1/done-employer?done=' + payload.done, {
        id: payload.id,
        loggedHours: payload.loggedHours,
      });
    },
    async toogleTaskSharing(payload) {
      await HttpClient.post('api/v1/tasks/share', {
        id: payload.id,
        share: payload.shared,
      });
    },
    async deleteTaskFromProject(id) {
      await HttpClient.delete('api/v1/projects/task', {
        data: {
          id,
        },
      });
    },
    async deleteProject(id) {
      await HttpClient.delete('api/v1/projects', {
        data: {
          id,
        },
      });
    },
  },
});
