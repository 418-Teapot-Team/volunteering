<template>
  <div
    class="flex flex-row justify-between items-start w-full border rounded-xl p-4"
    :class="task?.is_finished ? 'bg-green-50' : 'bg-white'"
  >
    <div class="flex flex-col">
      <div class="flex flex-row gap-4">
        <div class="w-8 h-full cursor-pointer" @click="toggleDoneStatus">
          <DoneIcon v-if="task?.is_finished" />
          <NotDoneIcon v-else />
        </div>
        <span class="md:text-2xl"
          >{{ task?.title }} |
          <span class="text-primary">{{ task?.estimate_time }} hrs.</span></span
        >
      </div>
      <div class="flex flex-col md:flex-row gap-4 pt-2 md:pl-12 text-sm md:text-base">
        <div class="flex flex-col gap-2">
          <div>{{ getFormattedDate(task?.createdAt) }}</div>
        </div>
        <div class="border-l-2 pl-3 md:pr-10">
          {{ task?.description }}
        </div>
      </div>
    </div>
    <div class="flex flex-row justify-between items-center gap-4">
      <div class="w-8 h-8 cursor-pointer" @click="toggleShareTask">
        <share-icon v-if="!task?.shared" />
        <archive-task-icon v-else />
      </div>
      <div class="w-8 h-8 cursor-pointer" @click="showConfirmPopup = true">
        <delete-icon />
      </div>
    </div>
    <confirm-action-popup
      @onConfirm="confirmDeletion"
      @onClose="closeConfirmPopup"
      v-if="showConfirmPopup"
    />
  </div>
</template>

<script>
import ShareIcon from '@/components/icons/ShareIcon.vue';
import DeleteIcon from '@/components/icons/DeleteIcon.vue';
import DoneIcon from '@/components/icons/DoneIcon.vue';
import NotDoneIcon from '@/components/icons/NotDoneIcon.vue';
import ConfirmActionPopup from '@/components/ConfirmActionPopup.vue';
import ArchiveTaskIcon from '@/components/icons/ArchiveTaskIcon.vue';

export default {
  name: 'ProjectTask',
  props: ['task'],
  components: {
    ShareIcon,
    DoneIcon,
    DeleteIcon,
    NotDoneIcon,
    ConfirmActionPopup,
    ArchiveTaskIcon,
  },
  data() {
    return {
      isOffersOpened: false,
      showConfirmPopup: false,
    };
  },
  methods: {
    closeConfirmPopup() {
      this.showConfirmPopup = false;
    },
    confirmDeletion() {
      this.$emit('deleteTask', this.task?.id);
      this.showConfirmPopup = false;
    },
    toggleShareTask() {
      this.$emit('onToggleShareTask', { id: this.task?.id, shared: !this.task?.shared });
    },
    toggleDoneStatus() {
      console.log(!this.task?.is_finished);
      this.$emit('onChangeDoneStatus', {
        id: this.task?.id,
        loggedHours: this.task?.estimate_time,
        done: !this.task?.is_finished,
      });
    },
    getFormattedDate(dateRow) {
      const date = new Date(dateRow);
      return `${date.getDate()}.${
        date.getMonth() < 10 ? '0' + (date.getMonth() + 1) : date.getMonth() + 1
      }.${date.getFullYear()}`;
    },
  },
};
</script>
