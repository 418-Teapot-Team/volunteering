<template>
  <button
    class="w-full h-full bg-primary rounded-lg text-xl text-white hover:bg-gradient-to-r hover:from-primary hover:to-primary-light"
    :type="type"
    @click="onClick"
  >
    {{ text }}
  </button>
</template>

<script>
export default {
  name: 'AppButton',
  props: ['text', 'type'],
  methods: {
    onClick() {
      this.$emit('onClick');
    },
  },
};
</script>
