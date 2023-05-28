import { defineStore } from 'pinia';
import { httpClient as HttpClient } from '../utils/HttpClient';

export default defineStore('statistic', {
  state: () => ({
    hoursData: [],
    projectsData: [],
    generalData: {},
  }),
  actions: {
    async getStats() {
      const { data } = await HttpClient.get('api/v1/get-stats');
      this.hoursData = data.result;
    },
    async getProjectStats() {
      const { data } = await HttpClient.get('api/v1/get-project-stats');
      this.projectsData = data.result;
    },
    async getGeneralStats() {
      const { data } = await HttpClient.get('api/v1/get-general-stats');
      this.generalData = data.result;
    },
  },
});
