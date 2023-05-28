<template>
  <section id="help_queries" class="flex flex-col gap-4">
    <h1 class="text-3xl p-6 lg:p-0 lg:text-6xl mb-4">My tasks</h1>
    <my-task
      v-for="item in myTasks"
      :key="item.id"
      :task="item"
      @onFinishTask="onFinishTask"
      @onDismissTask="onDismissTask"
    />
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
    ...mapActions(useTasksStore, ['getMyTasks', 'finishTask', 'dismissTask']),
    async initialLoad() {
      try {
        await this.getMyTasks();
      } catch (e) {
        toast.error(e?.messages);
      }
    },
    async onFinishTask(values) {
      const payload = {
        id: values.id,
        hours: values.hours,
      };
      try {
        await this.finishTask(payload);
        await this.getMyTasks();
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async onDismissTask(values) {
      const payload = {
        id: values.id,
      };
      try {
        await this.dismissTask(payload);
        await this.getMyTasks();
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
