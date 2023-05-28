<template>
  <div class="w-full rounded-2xl shadow-md border p-4 bg-gray-50">
    <div class="flex flex-row justify-between items-center h-20">
      <div class="flex flex-row justify-start gap-4">
        <span class="text-3xl">{{ project?.project?.title }}</span>
        <div v-if="isTasksOpened" @click="showConfirmPopup = true" class="w-8 h-8">
          <delete-icon />
        </div>
      </div>

      <div
        class="w-14 h-14 cursor-pointer transition duration-150"
        :class="isTasksOpened ? '' : '-rotate-90'"
        @click="isTasksOpened = !isTasksOpened"
      >
        <arrow-icon />
      </div>
    </div>
    <div :class="isTasksOpened ? 'h-fit visible' : 'h-0 hidden'">
      <p class="mb-6 pr-10">
        {{ project?.project?.description }}
      </p>
      <div class="flex flex-col justify-start items-start w-full gap-4">
        <div class="flex flex-row justify-between w-full">
          <span class="text-3xl mb-6">Tasks</span>
          <div class="h-10 w-36">
            <app-button type="button" text="Add task" @onClick="showAddTaskPopup = true" />
          </div>
        </div>
        <project-task
          v-for="task in project.tasks"
          :key="task.id"
          :task="task"
          @onToggleShareTask="toggleShareTask"
          @deleteTask="deleteTask"
          @onChangeDoneStatus="toggleDoneStatus"
        />
      </div>
    </div>
    <add-task-to-project-popup
      @onClose="closePopup"
      @onSubmit="addTaskToProject"
      v-if="showAddTaskPopup"
    />
    <confirm-action-popup
      @onConfirm="confirmDeletion"
      @onClose="closeConfirmPopup"
      v-if="showConfirmPopup"
    />
  </div>
</template>

<script>
import ArrowIcon from '@/components/icons/ArrowIcon.vue';
import DeleteIcon from '@/components/icons/DeleteIcon.vue';
import ProjectTask from '@/components/ProjectTask.vue';
import AddTaskToProjectPopup from '@/components/AddTaskToProjectPopup.vue';
import ConfirmActionPopup from '@/components/ConfirmActionPopup.vue';

export default {
  name: 'Project',
  props: ['project'],
  components: {
    ArrowIcon,
    ProjectTask,
    AddTaskToProjectPopup,
    ConfirmActionPopup,
    DeleteIcon,
  },
  data() {
    return {
      isTasksOpened: false,
      showAddTaskPopup: false,
      showConfirmPopup: false,
    };
  },
  methods: {
    addTaskToProject(values) {
      this.$emit('onAddTaskToProject', { ...values, projectId: this.project.project.id });
      this.showAddTaskPopup = false;
    },
    toggleShareTask(values) {
      this.$emit('onToggleShareTask', values);
    },
    closePopup() {
      this.showAddTaskPopup = false;
    },
    deleteTask(id) {
      this.$emit('onDeleteTask', id);
    },
    toggleDoneStatus(values) {
      this.$emit('onChangeDoneStatus', values);
    },
    async confirmDeletion() {
      this.$emit('onDeleteProject', this.project?.project?.id);
      this.showConfirmPopup = false;
    },
    closeConfirmPopup() {
      this.showConfirmPopup = false;
    },
  },
};
</script>
