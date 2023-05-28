<template>
  <div class="w-full shadow-md border rouned-xl p-4 flex flex-col rounded-2xl">
    <div class="flex flex-row justify-between mb-4">
      <div class="flex flex-col gap-2 items-center justify-center">
        <span class="text-xl md:text-3xl"
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
    <div
      class="flex flex-row justify-start gap-8 md:gap-4 self-center md:self-end md:mt-0 mt-4 items-center"
    >
      <div class="h-10 w-20 cursor-pointer">
        <app-button type="button" text="Finish" @onClick="showHoursPopup = true" />
      </div>
      <div class="h-10 w-10 cursor-pointer" @click="onDismissTask">
        <delete-icon />
      </div>
    </div>
  </div>
  <hours-popup @onSubmit="finishTask" @onClose="closePopup" v-if="showHoursPopup" />
</template>

<script>
import DeleteIcon from '@/components/icons/DeleteIcon.vue';
import HoursPopup from '@/components/HoursPopup.vue';
export default {
  name: 'MyTask',
  components: {
    DeleteIcon,
    HoursPopup,
  },
  props: ['task'],
  data() {
    return {
      showHoursPopup: false,
    };
  },
  methods: {
    getFormattedDate(dateRow) {
      const date = new Date(dateRow);
      return `${date.getDate()}.${
        date.getMonth() < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1
      }.${date.getFullYear()}`;
    },
    closePopup() {
      this.showHoursPopup = false;
    },
    finishTask(values) {
      this.$emit('onFinishTask', { ...values, id: this.task?.id });
      this.showHoursPopup = false;
    },
    onDismissTask() {
      this.$emit('onDismissTask', { id: this.task.id });
    },
  },
};
</script>
