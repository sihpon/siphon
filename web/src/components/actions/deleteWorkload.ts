'use server'

import APIClient from "@/util/api-client"
import { redirect } from "next/navigation"

export async function deleteWorkload(id: string) {
  await new APIClient().Workload().Delete(id)
  redirect('/workloads')
}
