import { GetResponse } from "@/generated/workload/v1/workload_pb"
import APIClient from "@/util/api-client"
import Button from "./components/button"
import { notFound } from "next/navigation"


export default async function Page(
  { params }: { params: { id: string } },
) {
  const id = params.id
  let response: GetResponse
  try {
    response = await new APIClient().Workload().Get(id)
  } catch (error) {
    return notFound()
  }

  return (
    <table className="table">
      <thead>
        <tr className="text-xs uppercase italic font-stretch-extra-condensed">
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
        <tr>
          <td>
            <div className="inline-grid *:[grid-area:1/1]">
              <div className="status status-error animate-ping"></div>
              <div className="status status-error"></div>
            </div>
          </td>
          <td>
            <span className="font-bold text-lg">{response.workload?.version} / {response.workload?.name}</span><br />
            <span className="italic text-sm">{response.workload?.description}</span>
          </td>
          <td>
            <span className="font-bold">rc/v2.1.0</span>
          </td>
          <td>
            <span className="font-bold">rc/v2.1.0</span>
          </td>
          <td>
            <span>2025/02/15 19:00:00</span>
          </td>
          <td>
            <span>2025/02/15 19:00:00</span>
          </td>
          <td>
            <span>atsuya.siphon@example.com</span>
          </td>
          <td>
            <Button id="{id}" />
          </td>
        </tr>
      </tbody>
    </table>
  )
}
