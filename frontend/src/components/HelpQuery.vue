<template>
  <div class="w-full shadow-md border rouned-xl p-4 flex flex-col rounded-2xl">
    <div class="flex flex-row justify-between mb-4">
      <div class="flex flex-col gap-2 items-center justify-center">
        <span class="text-3xl"
          >{{ task?.title }} |
          <span class="text-primary">{{ task?.estimate_time }} hrs.</span></span
        >
        <span class="text-xl self-start">{{ task?.project?.title }}</span>
      </div>
      <div class="flex flex-col gap-2 items-end justify-center">
        <span>{{ getFormattedDate(task?.createdAt) }}</span>
        <span>{{ `${task?.user?.firstName} ${task?.user?.lastName}` }}</span>
      </div>
    </div>
    <div class="text-gray-500">
      {{ task?.description }}
    </div>
    <div class="h-10 w-40 self-end">
      <app-button text="Apply" type="button" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'HelpQuery',
  props: ['task'],
  methods: {
    getFormattedDate(dateRow) {
      const date = new Date(dateRow);
      return `${date.getDate()}.${
        date.getMonth() < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1
      }.${date.getFullYear()}`;
    },
  },
};
</script>
