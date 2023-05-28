<template>
  <section id="projects">
    <div class="mb-10 flex flex-row justify-between items-center">
      <h1 class="text-6xl">Projects</h1>
      <div class="h-10 w-36">
        <app-button text="Create project" @onClick="showAddPopup" type="button" />
      </div>
    </div>
    <div class="flex flex-col justify-start items-center gap-6">
      <project
        v-for="project in projects"
        :key="project.project.id"
        :project="project"
        @onAddTaskToProject="addTask"
        @onToggleShareTask="toggleShareTask"
        @onDeleteTask="deleteTask"
        @onChangeDoneStatus="changeDoneStatus"
        @onDeleteProject="onDeleteProject"
      />
    </div>
    <add-project-popup
      @onClose="closeAddPopup"
      @onSubmit="createProjectSubmit"
      v-if="showProjectPopup"
    />
  </section>
</template>

<script>
import Project from '@/components/Project.vue';
import AddProjectPopup from '@/components/AddProjectPopup.vue';
import useProjectsStore from '@/stores/projects';
import { mapState, mapActions } from 'pinia';
import { useToast } from 'vue-toastification';

const toast = useToast();

export default {
  name: 'ProjectsPage',
  components: {
    Project,
    AddProjectPopup,
  },
  data() {
    return {
      showProjectPopup: false,
    };
  },
  computed: {
    ...mapState(useProjectsStore, ['projects']),
  },
  methods: {
    ...mapActions(useProjectsStore, [
      'getProjects',
      'createProject',
      'deleteProject',
      'addTaskToProject',
      'toogleTaskSharing',
      'deleteTaskFromProject',
      'markProjectTaskAsDone',
    ]),
    closeAddPopup() {
      this.showProjectPopup = false;
    },
    showAddPopup() {
      this.showProjectPopup = true;
    },
    async createProjectSubmit(values) {
      try {
        const payload = {
          title: values.title,
          description: values.description,
        };
        await this.createProject(payload);
        await this.getProjects();
        toast.success('Successfully created new project!');
        this.showProjectPopup = false;
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async initialLoad() {
      try {
        await this.getProjects();
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async addTask(values) {
      try {
        const payload = {
          projectId: values.projectId,
          title: values.title,
          estimate_time: values.estimate_time,
          description: values.description,
        };
        await this.addTaskToProject(payload);
        await this.getProjects();
        toast.success('Successfully added task');
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async toggleShareTask(values) {
      try {
        const payload = {
          id: values.id,
          shared: values.shared,
        };
        await this.toogleTaskSharing(payload);
        await this.getProjects();
        if (values.shared) {
          toast.success('Task is now shared!');
        } else {
          toast.info('Task is archived!');
        }
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async changeDoneStatus(values) {
      try {
        const payload = {
          id: values.id,
          loggedHours: values.loggedHours,
          done: values.done,
        };
        await this.markProjectTaskAsDone(payload);
        await this.getProjects();
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async deleteTask(id) {
      try {
        await this.deleteTaskFromProject(id);
        await this.getProjects();
      } catch (e) {
        toast.error(e?.message);
      }
    },
    async onDeleteProject(id) {
      try {
        await this.deleteProject(id);
        await this.getProjects();
      } catch (e) {
        toast.error(e?.message);
      }
    },
  },
  mounted() {
    this.initialLoad();
  },
};
</script>
