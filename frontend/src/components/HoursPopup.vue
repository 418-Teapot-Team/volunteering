<template>
  <div
    class="fixed top-0 left-0 right-0 bottom-0 bg-black bg-opacity-40 flex justify-center items-center z-10"
  >
    <div class="w-96 h-fit bg-white shadow-sm rounded-2xl p-4 flex flex-col gap-2">
      <div class="w-6 h-6 cursor-pointer self-end" @click="onClose">
        <close-icon :isPopup="true" />
      </div>
      <vee-form :validation-schema="schema" @submit="onSubmit">
        <div class="w-full h-10 mb-8">
          <vee-field
            name="hours"
            type="number"
            placeholder="Hours spent*"
            class="p-2 rounded-lg shadow-md outline-none w-full h-10 border resize-none"
          ></vee-field>
          <ErrorMessage name="hours" class="text-red-500 text-xs" />
        </div>
        <div class="w-full h-12"><app-button type="submit" text="Continue" /></div>
      </vee-form>
    </div>
  </div>
</template>

<script>
import CloseIcon from '@/components/icons/CloseIcon.vue';
export default {
  name: 'HoursPopup',
  components: { CloseIcon },
  data() {
    return {
      schema: {
        hours: 'required|min_value:1|max_value:360',
      },
    };
  },
  methods: {
    onClose() {
      this.$emit('onClose');
    },
    onSubmit(values) {
      this.$emit('onSubmit', values);
    },
  },
};
</script>
