import { createApp } from 'vue';
import './style.css';
import App from './App.vue';
import { createWebHistory, createRouter } from 'vue-router';

import WorkloadIndex from './components/workloads/Index.vue';
import WorkloadCreate from './components/workloads/Create.vue';
import WorkloadGet from './components/workloads/Get.vue';

const routes = [
  { path: '/', component: WorkloadIndex },
  { path: '/workloads', component: WorkloadIndex },
  { path: '/workloads/new', component: WorkloadCreate },
  { path: '/workloads/:id', component: WorkloadGet },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

createApp(App).use(router).mount('#app')
