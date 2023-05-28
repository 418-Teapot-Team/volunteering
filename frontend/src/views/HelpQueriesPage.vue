<template>
  <main>Help queries</main>
</template>

<script>
import useTasksStore from '@/stores/tasks';
import { mapState, mapActions } from 'pinia';
export default {
  name: 'HelpQueriesPage',
  computed: {
    ...mapState(useTasksStore),
  },
};
</script>
