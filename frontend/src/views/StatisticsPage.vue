<template>
  <section id="statistics h-full">
    <div class="h-1/2">
      <hours-chart :values="[]" :bottomLabels="[]" />
    </div>
    <div>
      <div></div>
      <div></div>
    </div>

    <projects-chart :values="[]" :rightLabels="[]" />
  </section>
</template>

<script>
import HoursChart from '@/components/HoursChart.vue';
import ProjectsChart from '@/components/ProjectsChart.vue';

export default {
  name: 'StatisticsPage',
  components: {
    HoursChart,
    ProjectsChart,
  },
};
</script>
