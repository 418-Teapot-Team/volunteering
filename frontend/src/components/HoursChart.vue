<template>
  <Line :data="dataset" :options="options" class="pb-2 scale-75 xl:scale-100" />
</template>

<script>
import { Line } from 'vue-chartjs';
export default {
  name: 'HoursChart',
  components: {
    // eslint-disable-next-line vue/no-reserved-component-names
    Line,
  },
  props: {
    bottomLabels: Array,
    values: Array,
  },
  computed: {
    dataset() {
      return {
        labels: this.bottomLabels,
        datasets: [
          {
            borderColor: '#1c1c21',
            backgroundColor: '#1c1c211A',
            pointerBackgroundColor: '#1c1c21',
            fill: true,
            data: this.values,
            tension: 0.2,
            borderWidth: 2,
            pointRadius: 3,
            pointBorderColor: '#1c1c21',
            pointBackgroundColor: '#1c1c21',
          },
        ],
      };
    },
  },
  data() {
    return {
      options: {
        responsive: true,
        maintainAspectRatio: false,
        plugins: {
          legend: null,
        },
      },
    };
  },
};
</script>
