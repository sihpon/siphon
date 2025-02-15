'use server'

import { CreateRequestSchema, VersionSchema } from "@/generated/version/v1/version_pb"
import APIClient from "@/util/api-client"
import { create } from "@bufbuild/protobuf"
import { redirect } from "next/navigation"

export async function createVersion(formData: FormData) {
  console.log(formData)
  const request = create(CreateRequestSchema)
  request.version = create(VersionSchema)
  request.version.id = formData.get('version') as string
  request.version.description = formData.get('description') as string

  await new APIClient().Version().Create(request)
  redirect('/versions')
}
