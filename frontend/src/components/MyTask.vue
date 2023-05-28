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
    <div class="flex flex-row justify-start gap-4 self-end items-center">
      <div class="h-10 w-10 cursor-pointer">
        <DoneIcon v-if="task?.is_finished" />
        <NotDoneIcon v-else />
      </div>
      <div class="h-10 w-10 cursor-pointer">
        <delete-icon />
      </div>
    </div>
  </div>
</template>

<script>
import DeleteIcon from '@/components/icons/DeleteIcon.vue';
import DoneIcon from '@/components/icons/DoneIcon.vue';
import NotDoneIcon from '@/components/icons/NotDoneIcon.vue';
export default {
  name: 'MyTask',
  components: {
    DeleteIcon,
    DoneIcon,
    NotDoneIcon,
  },
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
