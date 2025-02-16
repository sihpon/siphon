import { GetResponse } from "@/generated/workload/v1/workload_pb"
import APIClient from "@/util/api-client"
import { notFound } from "next/navigation"
import Table from "@/components/Table"
import { timestampDate } from "@bufbuild/protobuf/wkt"
import DeleteButton from "@/components/DeleteButton"
import { deleteWorkload } from "@/components/actions/deleteWorkload"

export default async function Page({
  params,
}: {
  params: Promise<{ id: string }>
}) {
  const id = (await params).id
  let response: GetResponse

  try {
    response = await new APIClient().Workload().Get(id)
  } catch (error) {
    return notFound()
  }

  return (
    <Table columns={['Name', 'SERVER', 'MASTER', 'Created At', 'Updated At', 'Owner', '']}>
      <tr>
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
          <span>{timestampDate(response.workload?.createdAt!).toISOString()}</span>
        </td>
        <td>
          <span>{timestampDate(response.workload?.updatedAt!).toISOString()}</span>
        </td>
        <td>
          <span>atsuya.siphon@example.com</span>
        </td>
        <td>
          <DeleteButton
            deleter={deleteWorkload(id)}
            itemName={id}
          />
        </td>
      </tr>
    </Table>
  )
}
