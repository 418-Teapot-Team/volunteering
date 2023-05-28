<template>
  <section id="help_queries" class="flex flex-col gap-4">
    <h1 class="text-6xl mb-4">Pending tasks</h1>
    <pending-task
      v-for="item in pendingTasks"
      :key="item.id"
      :task="item"
      @onApproveTask="onApproveTask"
      @onDismissTask="onDismissTask"
    />
  </section>
</template>

<script>
import PendingTask from '@/components/PendingTask.vue';
import useTasksStore from '@/stores/tasks';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'WaitingForAproval',
  components: {
    PendingTask,
  },
  computed: {
    ...mapState(useTasksStore, ['pendingTasks']),
  },
  methods: {
    ...mapActions(useTasksStore, ['getPendingTasks', 'approveTask', 'denyTask']),
    async initialLoad() {
      try {
        await this.getPendingTasks();
      } catch (e) {
        toast.error(e?.messages);
      }
    },
    async onApproveTask(values) {
      const payload = {
        id: values.id,
        hours: values.hours,
      };
      try {
        await this.approveTask(payload);
        await this.getPendingTasks();
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async onDismissTask(values) {
      const payload = {
        id: values.id,
      };
      try {
        await this.denyTask(payload);
        await this.getPendingTasks();
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
