import DetailButton from '@/components/DetailButton'
import Table from '@/components/Table'
import { ListResponse } from '@/generated/version/v1/version_pb'
import APIClient from '@/util/api-client'

export default async function Page() {
  const response: ListResponse = await new APIClient().Version().List()
  return (
    <Table columns={['ID', 'DESCRIPTION', 'DETAIL']}>
      {response.versions.map((version) => (
        <tr key={version.id}>
          <td>
            <span className="font-bold text-lg">{version.id}</span><br />
          </td>
          <td>
            <span className="font-bold">{version.description}</span>
          </td>
          <td>
            <DetailButton href={`/versions/${version.id}`} />
          </td>
        </tr>
      ))}
    </Table>
  )
}
