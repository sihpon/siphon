<script setup lang="ts">
import { Timestamp } from "google-protobuf/google/protobuf/timestamp_pb.ts";
import { ref, onMounted } from 'vue'
import { timestampDate } from "@bufbuild/protobuf/wkt";
import { Client } from '../../shared/client.ts'

const response = ref({})
onMounted(async () => {
  response.value = await new Client().Workload().List()
})

</script>
<template>
    <h1 class="text-5xl font-extrabold mb-4">🏗️ Workloads</h1>
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
                    Index
                </span>
            </li>
        </ul>
    </div>
    <div class="overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
        <table class="table">
            <thead>
                <tr class="text-xs uppercase italic font-stretch-extra-condensed">
                    <th>STATUS</th>
                    <th>NAME</th>
                    <th>SERVER</th>
                    <th>MASTER</th>
                    <th>CREATED AT</th>
                    <th>UPDATED AT</th>
                    <th>CREATED BY</th>
                    <th>DETAIL</th>
                </tr>
            </thead>
            <tbody>
                <!-- row 1 -->
                <tr v-for="(workload, index) in response.workloads">
                    <td>
                        <div class="inline-grid *:[grid-area:1/1]">
                            <div class="status status-error animate-ping"></div>
                            <div class="status status-error"></div>
                        </div>
                    </td>
                    <td>
                        <span class="font-bold text-lg">{{ workload.version }} - {{ workload.name }}</span><br>
                        <span class="italic text-sm">{{ workload.description }}</span>
                    </td>
                    <td>
                        <span class="font-bold">rc/v2.1.0</span>
                    </td>
                    <td>
                        <span class="font-bold">rc/v2.1.0</span>
                    </td>
                    <td>
                        <span>{{ timestampDate(workload.createdAt) }}</span>
                    </td>
                    <td>
                        <span>{{ timestampDate(workload.updatedAt) }}</span>
                    </td>
                    <td>
                        <span>atsuya.siphon@example.com</span>
                    </td>
                    <td>
                    <router-link :to="`/workloads/${workload.id}`">
                      <button class="btn btn-neutral">
                          <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5"
                              stroke="currentColor" class="size-[1.2em]">
                              <path stroke-linecap="round" stroke-linejoin="round"
                                  d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z" />
                          </svg>
                          詳細
                      </button>
                    </router-link>
                    </td>
                </tr>
            </tbody>
        </table>
    </div>
</template>
