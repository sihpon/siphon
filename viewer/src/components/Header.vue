<script setup lang="ts">
function toggleSidebar() {
    const sidebar = document.getElementById("sidebar");
    sidebar.classList.toggle("hidden");
}
</script>

<template>
    <div class="z-50 navbar bg-base-100 shadow glass">
        <div class="lg:hidden">
            <button class="btn btn-square btn-ghost" @click="toggleSidebar()">
                <svg xmlns="http://www.w3.org/2000/svg" class="h-6 w-6" fill="none" viewBox="0 0 24 24"
                    stroke="currentColor">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h16m-7 6h7" />
                </svg>
            </button>
        </div>
        <div class="flex-1">
            <a class="btn btn-ghost text-xl">Siphon</a>
        </div>
        <div class="flex-none">
            <ul class="menu menu-horizontal px-1">
                <li><a>Workloads</a></li>
                <li>
                    <details>
                        <summary>Parent</summary>
                        <ul class="bg-base-100 rounded-t-none p-2">
                            <li><a>Link 1</a></li>
                            <li><a>Link 2</a></li>
                        </ul>
                    </details>
                </li>
            </ul>
        </div>
    </div>
</template>
