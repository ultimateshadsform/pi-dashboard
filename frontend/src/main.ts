import { createApp } from 'vue';
import '@/scss/main.scss';
import 'primevue/resources/themes/aura-light-green/theme.css';
import App from '@/App.vue';
import router from '@/router';

import PrimeVue from 'primevue/config';

import { createPinia } from 'pinia';

const app = createApp(App);

app.use(createPinia());

app.use(router);

app.use(PrimeVue);

app.mount('#app');
