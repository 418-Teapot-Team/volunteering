<template>
  <dev
    class="fixed top-0 left-0 right-0 bottom-0 bg-black bg-opacity-40 flex justify-center items-center z-10"
  >
    <div class="w-96 h-fit bg-white shadow-sm rounded-2xl p-4">
      <vee-form :validation-schema="schema" @submit="onSubmit">
        <div class="text-2xl mb-4 flex flex-row justify-between items-center">
          <span>Create new project</span>
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
        <div class="w-full h-12"><app-button type="submit" text="Craete" /></div>
      </vee-form>
    </div>
  </dev>
</template>

<script>
import CloseIcon from '@/components/icons/CloseIcon.vue';
export default {
  name: 'AddProjectPopup',
  components: {
    CloseIcon,
  },
  data() {
    return {
      schema: {
        title: 'required|min:3|max:100|alpha_spaces',
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
