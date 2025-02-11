<script setup lang="ts">
import { ref } from 'vue'
const message = ref('Hello, Vite!')
import { createClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { WorkloadService, CreateRequest } from '../../generated/workload/v1/workload_pb.ts';
import { onMounted } from 'vue'

async function create() {
    const transport = createConnectTransport({
        baseUrl: "http://localhost:8080",
    });
    const client = createClient(WorkloadService, transport);
    const response = await client.create(request)
    console.log(response)
}

const request: CreateRequest = {
  versionId : "v2.1.0",
  name : "test",
  description : "test",
};
</script>

<template>
    <h1 class="text-5xl font-extrabold mb-4">🏗️ New Workload</h1>
    <div class="breadcrumbs text-sm mb-4">
        <ul>
            <li>
                <a>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                        class="h-4 w-4 stroke-current">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"></path>
                    </svg>
                    Home
                </a>
            </li>
            <li>
                <a>
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                        class="h-4 w-4 stroke-current">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"></path>
                    </svg>
                    Workloads
                </a>
            </li>
            <li>
                <span class="inline-flex items-center gap-2 italic underline">
                    <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                        class="h-4 w-4 stroke-current">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                            d="M9 13h6m-3-3v6m5 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z">
                        </path>
                    </svg>
                    New
                </span>
            </li>
        </ul>
    </div>
    <div class="card bg-base-100 shadow-sm">
        <div class="card-body">
            <p>Message is: {{ message }}</p>
            <input v-model="message" placeholder="edit me" />
            <label class="select">
                <span class="label">バージョン</span>
                <select>
                    <option>dev</option>
                    <option>v2.1.0</option>
                    <option>v2.2.0</option>
                </select>
            </label>
            <label class="floating-label">
                <span>サーバー名</span>
                <input type="text" placeholder="Your name" class="input" />
            </label>
            <label class="floating-label">
                <span>サーバー説明文</span>
                <input type="text" placeholder="Your name" class="input" />
            </label>
            <label class="floating-label">
                <span>追従サーバーブランチ</span>
                <input type="text" placeholder="Your name" class="input" />
            </label>
            <div class="divider"></div>
            <label class="floating-label">
                <span>追従マスターブランチ</span>
                <input type="text" placeholder="Your name" class="input" />
            </label>
            <label class="floating-label">
                <span>表示優先度</span>
                <input type="text" placeholder="Your name" class="input" />
            </label>
            <div class="divider"></div>
            <button class="btn btn-primary" @click="create">作成する</button>
        </div>
    </div>
</template>
