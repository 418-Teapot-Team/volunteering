<template>
  <div
    class="fixed top-0 left-0 right-0 bottom-0 bg-black bg-opacity-40 flex justify-center items-center z-10"
  >
    <div class="w-96 h-fit bg-white shadow-sm rounded-2xl p-4">
      <div class="flex flex-row justify-between w-full items-center mb-6">
        <span class="text-xl">Confirm your action?</span>
        <div class="w-6 h-6 cursor-pointer" @click="onClose">
          <close-icon :isPopup="true" />
        </div>
      </div>
      <div class="h-10 w-24 m-auto">
        <app-button type="button" text="Confirm" @onClick="confirmAction" />
      </div>
    </div>
  </div>
</template>

<script>
import CloseIcon from '@/components/icons/CloseIcon.vue';
export default {
  name: 'ConfirmActionPopup',
  components: {
    CloseIcon,
  },
  methods: {
    confirmAction() {
      this.$emit('onConfirm');
    },
    onClose() {
      this.$emit('onClose');
    },
  },
};
</script>
