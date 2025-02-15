import { GetResponse } from "@/generated/version/v1/version_pb"
import APIClient from "@/util/api-client"
import { notFound } from "next/navigation"

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const id = (await params).id
  let response: GetResponse

  try {
    response = await new APIClient().Version().Get(id)
  } catch (error) {
    return notFound()
  }

  return (
    <table className="table">
      <thead>
        <tr className="text-xs uppercase italic font-stretch-extra-condensed">
          <th>{response.version?.id}</th>
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
        <tr>
          <td>
            <div className="inline-grid *:[grid-area:1/1]">
              <div className="status status-error animate-ping"></div>
              <div className="status status-error"></div>
            </div>
          </td>
        </tr>
      </tbody>
    </table>
  )
}
