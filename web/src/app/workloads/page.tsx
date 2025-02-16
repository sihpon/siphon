import DetailButton from '@/components/DetailButton'
import Table from '@/components/Table'
import { ListResponse } from '@/generated/workload/v1/workload_pb'
import APIClient from '@/util/api-client'
import { timestampDate } from '@bufbuild/protobuf/wkt'

export default async function Page() {
  const response: ListResponse = await new APIClient().Workload().List()
  return (
    <Table columns={['Name', 'SERVER', 'MASTER', 'Created At', 'Updated At', 'Owner', '']}>
      {response.workloads.map((workload) => (
        <tr key={workload.id}>
          <td>
            <span className="font-bold text-lg">{workload.version} / {workload.name}</span><br />
            <span className="italic text-sm">{workload.description}</span>
          </td>
          <td>
            <span className="font-bold">rc/v2.1.0</span>
          </td>
          <td>
            <span className="font-bold">rc/v2.1.0</span>
          </td>
          <td>
            <span>{timestampDate(workload.createdAt!).toISOString()}</span>
          </td>
          <td>
            <span>{timestampDate(workload.updatedAt!).toISOString()}</span>
          </td>
          <td>
            <span>atsuya.siphon@example.com</span>
          </td>
          <td>
            <DetailButton href={`/workloads/${workload.id}`} />
          </td>
        </tr>
      ))}
    </Table>
  )
}
