import './assets/main.css';

import { createApp } from 'vue';
import { createPinia } from 'pinia';

import App from './App.vue';
import router from './router';

import VeeValidatePlugin from './includes/validation';
import Components from './includes/components';
import Toast, { POSITION } from 'vue-toastification';
import 'vue-toastification/dist/index.css';

const app = createApp(App);

app.use(createPinia());
app.use(router);

app.use(VeeValidatePlugin);
app.use(Components);
app.use(Toast, {
  position: POSITION.TOP_CENTER,
});

app.mount('#app');
