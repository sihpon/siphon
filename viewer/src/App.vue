<script setup lang="ts">
import Header from './components/Header.vue';
import Footer from './components/Footer.vue';
import Sidebar from './components/Sidebar.vue';
import WorkloadIndex from './components/workload/WorkloadIndex.vue';
</script>

<template>
<div class="flex flex-col h-screen">
  <Header />

  <div class="flex flex-1">
    <Sidebar />
    <div class="flex-1 p-6">
    <WorkloadIndex />
    </div>
  </div>

  <Footer />
</div>
</template>
