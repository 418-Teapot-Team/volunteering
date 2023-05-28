<template>
  <section id="applies" class="flex flex-col gap-4">
    <h1 class="text-6xl mb-4">Applies</h1>
    <apply-item v-for="item in applies" :key="item.id" :apply="item" @onApproval="approve" />
  </section>
</template>

<script>
import ApplyItem from '@/components/Apply.vue';
import useTasksStore from '@/stores/tasks';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'AppliesPage',
  components: { ApplyItem },
  computed: {
    ...mapState(useTasksStore, ['applies']),
  },
  methods: {
    ...mapActions(useTasksStore, ['getApplies', 'approveApply']),
    async approve(values) {
      try {
        this.approveApply(values);
        this.getApplies();
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async initialLoad() {
      try {
        await this.getApplies();
      } catch (e) {
        toast.error(e?.messages);
      }
    },
  },
  mounted() {
    this.initialLoad();
  },
};
</script>
