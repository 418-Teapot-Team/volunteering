<template>
  <div class="w-full shadow-md border rouned-xl p-4 flex flex-col rounded-2xl">
    <div class="flex flex-row justify-between mb-4">
      <div class="flex flex-col gap-2 items-center justify-center">
        <span class="text-xl md:text-3xl"
          >{{ apply?.task?.title }} |
          <span class="text-primary">{{ apply?.task?.estimate_time }} hrs.</span></span
        >
        <span class="md:text-xl text-sm self-start">{{ apply?.task?.project?.title }}</span>
      </div>
      <div class="flex flex-col gap-2 items-end justify-center">
        <span>{{ getFormattedDate(apply?.createdAt) }}</span>
        <span>{{ `${apply?.user?.firstName} ${apply?.user?.lastName}` }}</span>
      </div>
    </div>
    <div class="text-gray-500">
      {{ apply?.task?.description }}
    </div>
    <div class="h-10 w-40 self-center md:self-end md:mt-0 mt-4">
      <app-button type="button" text="Approve" @onClick="approve" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'Apply',
  props: ['apply'],
  methods: {
    getFormattedDate(dateRow) {
      const date = new Date(dateRow);
      return `${date.getDate()}.${
        date.getMonth() < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1
      }.${date.getFullYear()}`;
    },
    approve() {
      this.$emit('onApproval', { id: this.apply.id, applied_id: this.apply.user.userId });
    },
  },
};
</script>
