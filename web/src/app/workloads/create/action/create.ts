'use server'

import { CreateRequestSchema } from "@/generated/workload/v1/workload_pb"
import APIClient from "@/util/api-client"
import { create } from "@bufbuild/protobuf"
import { redirect } from "next/navigation"

export async function createWorkload(formData: FormData) {
  console.log(formData)
  const request = create(CreateRequestSchema)
  request.versionId = formData.get('version') as string
  request.name = formData.get('name') as string
  request.description = formData.get('description') as string
  await new APIClient().Workload().Create(request)
  redirect('/workloads')
}
