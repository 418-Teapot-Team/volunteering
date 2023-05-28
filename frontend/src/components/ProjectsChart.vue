<template>
  <div
    class="h-full bg-white p-2 shadow-sm rounded-2xl flex flex-col justify-center items-center gap-5"
  >
    <!-- <span class="text-black text-3xl first-letter:capitalize">{{ title }}</span> -->
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
    // title: String,
  },
  computed: {
    dataset() {
      console.log(this.rightLabels, this.values);
      return {
        labels: this.rightLabels,
        datasets: [
          {
            backgroundColor: this.values.map((item) => item.hash),
            data: this.values.map((item) => item.value),
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
        },
      };
    },
  },
  data() {
    return {};
  },
};
</script>
