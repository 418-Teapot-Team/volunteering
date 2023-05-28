<template>
  <div
    class="h-full w-full bg-white p-2 rounded-2xl flex flex-col justify-center items-center gap-5"
  >
    <Pie :options="options" :data="dataset" />
  </div>
</template>

<script>
import { Pie } from 'vue-chartjs';
export default {
  name: 'ProjectsChart',
  components: {
    Pie,
  },
  props: {
    rightLabels: Array,
    values: Array,
  },
  computed: {
    dataset() {
      console.log(this.rightLabels, this.values);
      return {
        labels: this.rightLabels,
        datasets: [
          {
            data: this.values,
          },
        ],
      };
    },
    options() {
      return {
        responsive: true,
        maintainAspectRation: false,
        plugins: {
          legend: {
            position: 'right',
          },
          colors: {
            forceOverride: true,
          },
        },
      };
    },
  },
  data() {
    return {};
  },
};
</script>
