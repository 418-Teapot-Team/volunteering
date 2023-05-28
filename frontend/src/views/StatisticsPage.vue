<template>
  <section id="statistics h-full">
    <h1 class="text-6xl mb-4">Stats</h1>
    <div class="w-full shadow-md border rounded-xl mb-4 h-80">
      <hours-chart
        :values="hoursData.map((item) => item.Value)"
        :bottomLabels="hoursData.map((item) => item.Date)"
      />
    </div>
    <div class="flex flex-row justify-between w-full gap-8 h-96">
      <projects-chart
        :values="projectsData.map((item) => item.Value)"
        :rightLabels="projectsData.map((item) => item.Title)"
        class="rounded-xl shadow-md border"
      />
      <div class="w-full border rounded-xl shadow-md p-4 flex flex-col justify-center items-center">
        <span class="text-2xl">
          Score: <span class="text-primary">{{ generalData.Score }}</span>
        </span>
        <span class="text-2xl">
          Completed tasks: <span class="text-primary">{{ generalData.TaskAmount }}</span>
        </span>
        <span class="text-2xl">
          Hours spent: <span class="text-primary">{{ generalData.HoursAmount }}</span>
        </span>
      </div>
    </div>
  </section>
</template>

<script>
import HoursChart from '@/components/HoursChart.vue';
import ProjectsChart from '@/components/ProjectsChart.vue';
import useStatisticsStore from '@/stores/statistic';
import { mapActions, mapState } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'StatisticsPage',
  components: {
    HoursChart,
    ProjectsChart,
  },
  computed: {
    ...mapState(useStatisticsStore, ['hoursData', 'projectsData', 'generalData']),
  },
  methods: {
    ...mapActions(useStatisticsStore, ['getStats', 'getProjectStats', 'getGeneralStats']),
    async initialLoad() {
      try {
        await this.getStats();
        await this.getProjectStats();
        await this.getGeneralStats();
      } catch (e) {
        toast.error(e?.message);
      }
    },
  },
  mounted() {
    this.initialLoad();
  },
};
</script>
