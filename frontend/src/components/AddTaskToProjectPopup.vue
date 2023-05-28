<template>
  <div
    class="fixed top-0 left-0 right-0 bottom-0 bg-black bg-opacity-40 flex justify-center items-center z-10"
  >
    <div class="w-96 h-fit bg-white shadow-sm rounded-2xl p-4">
      <vee-form :validation-schema="schema" @submit="onSubmit">
        <div class="text-2xl mb-4 flex flex-row justify-between items-center">
          <span>Create new task</span>
          <div class="w-6 h-6 cursor-pointer" @click="onClose">
            <close-icon :isPopup="true" />
          </div>
        </div>
        <div class="w-full h-20">
          <vee-field
            name="title"
            type="text"
            placeholder="Title*"
            class="p-2 rounded-lg shadow-md outline-none w-full h-12 border"
          ></vee-field>
          <ErrorMessage name="title" class="text-red-500 text-xs" />
        </div>
        <div class="w-full h-20">
          <vee-field
            name="estimate_time"
            type="number"
            placeholder="Estimate hours*"
            class="p-2 rounded-lg shadow-md outline-none w-full h-12 border"
          ></vee-field>
          <ErrorMessage name="estimate_time" class="text-red-500 text-xs" />
        </div>
        <div class="w-full h-40 mb-4">
          <vee-field
            name="description"
            as="textarea"
            type="text"
            placeholder="Description"
            class="p-2 rounded-lg shadow-md outline-none w-full h-36 border resize-none"
          ></vee-field>
          <ErrorMessage name="description" class="text-red-500 text-xs" />
        </div>
        <div class="w-full h-12"><app-button type="submit" text="Create" /></div>
      </vee-form>
    </div>
  </div>
</template>

<script>
import CloseIcon from '@/components/icons/CloseIcon.vue';
export default {
  name: 'AddTaskToProjectPopup',
  components: {
    CloseIcon,
  },
  data() {
    return {
      schema: {
        title: 'required|min:3|max:100|alpha_spaces',
        estimate_time: 'required|min_value:1|max_value:360',
        description: 'max:500',
      },
    };
  },
  methods: {
    onSubmit(values) {
      this.$emit('onSubmit', values);
    },
    onClose() {
      this.$emit('onClose');
    },
  },
};
</script>
