<template>
  <section id="help_queries" class="flex flex-col gap-4">
    <h1 class="text-3xl p-6 lg:p-0 lg:text-6xl mb-4">Tasks</h1>
    <help-query
      v-for="item in sharedTasks"
      :myid="user.userId"
      :key="item.id"
      :task="item"
      @onApply="apply"
    />
  </section>
</template>

<script>
import HelpQuery from '@/components/HelpQuery.vue';
import useTasksStore from '@/stores/tasks';
import useAuthStore from '@/stores/auth';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'HelpQueriesPage',
  components: {
    HelpQuery,
  },
  computed: {
    ...mapState(useTasksStore, ['sharedTasks']),
    ...mapState(useAuthStore, ['isAuthorized', 'user']),
  },
  methods: {
    ...mapActions(useTasksStore, ['getSharedTasks', 'createApply']),
    async initialLoad() {
      try {
        if (this.isAuthorized) {
          await this.getSharedTasks();
        }
      } catch (e) {
        toast.error(e?.messages);
      }
    },
    async apply(values) {
      try {
        this.createApply(values);
        this.getSharedTasks();
        toast.success('You have successfully aplied to this task!');
      } catch (e) {
        toast.error(e?.message);
      }
    },
  },
  mounted() {
    console.log(this.user);
    this.initialLoad();
  },
};
</script>
