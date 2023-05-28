<template>
  <div class="w-full shadow-md border rouned-xl p-4 flex flex-col rounded-2xl">
    <div class="flex flex-row justify-between mb-4">
      <div class="flex flex-col gap-2 items-center justify-center">
        <span class="text-xl md:text-3xl"
          >{{ task?.title }} |
          <span class="text-primary">{{ task?.estimate_time }} hrs.</span></span
        >
        <span class="md:text-xl text-sm self-start">{{ task?.project?.title }}</span>
      </div>
      <div class="flex flex-col gap-2 items-end justify-center">
        <span>{{ getFormattedDate(task?.createdAt) }}</span>
        <span>{{ `${task?.user?.firstName} ${task?.user?.lastName}` }}</span>
      </div>
    </div>
    <div class="text-gray-500">
      {{ task?.description }}
    </div>
    <div class="h-10 w-40 self-center md:self-end md:mt-0 mt-4">
      <app-button v-if="task?.user?.userId !== myid" text="Apply" type="button" @onClick="apply" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelpQuery',
  props: ['task', 'myid'],
  methods: {
    getFormattedDate(dateRow) {
      const date = new Date(dateRow);
      return `${date.getDate()}.${
        date.getMonth() < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1
      }.${date.getFullYear()}`;
    },
    apply() {
      this.$emit('onApply', { taskId: this.task.id, userId: this.task.user?.userId });
    },
  },
};
</script>
