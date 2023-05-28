<template>
  <component :is="layout">
    <RouterView />
  </component>
</template>

<script>
import MainLayout from '@/layouts/MainLayout.vue';
import EmptyLayout from '@/layouts/EmptyLayout.vue';

// Chart.js config
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Legend,
  Filler,
  ArcElement,
  Colors,
} from 'chart.js';

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  ArcElement,
  Title,
  Tooltip,
  Legend,
  Filler,
  Colors
);

export default {
  name: 'App',
  components: {
    MainLayout,
    EmptyLayout,
  },
  computed: {
    layout() {
      return (this.$route.meta.layout || 'empty') + '-layout';
    },
  },
};
</script>
