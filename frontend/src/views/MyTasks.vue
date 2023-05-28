<template>
  <section id="help_queries" class="flex flex-col gap-4">
    <h1 class="text-6xl mb-4">My tasks</h1>
    <my-task v-for="item in myTasks" :key="item.id" :task="item" />
  </section>
</template>

<script>
import MyTask from '@/components/MyTask.vue';
import useTasksStore from '@/stores/tasks';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'MyTasks',
  components: {
    MyTask,
  },
  computed: {
    ...mapState(useTasksStore, ['myTasks']),
  },
  methods: {
    ...mapActions(useTasksStore, ['getMyTasks']),
    async initialLoad() {
      try {
        await this.getMyTasks();
      } catch (e) {
        toast.error(e?.messages);
      }
    },
  },
  mounted() {
    this.initialLoad();
  },
};
</script>
