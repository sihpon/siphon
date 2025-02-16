import { ListResponse } from '@/generated/version/v1/version_pb'
import APIClient from '@/util/api-client'
import Link from 'next/link'

export default async function Page() {
  const response: ListResponse = await new APIClient().Version().List()
  return (
    <div className="overflow-x-auto rounded-box bg-base-100">
      <table className="table">
        <thead>
          <tr className="text-xs uppercase italic font-stretch-extra-condensed">
            <th>ID</th>
            <th>DESCRIPTION</th>
            <th>DETAIL</th>
          </tr>
        </thead>
        <tbody>
          {response.versions.map((version) => (
            <tr key={version.id}>
              <td>
                <span className="font-bold text-lg">{version.id}</span><br />
              </td>
              <td>
                <span className="font-bold">{version.description}</span>
              </td>
              <td>
                <Link className="btn btn-neutral" href={`/version/${version.id}`}>
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="2.5"
                    stroke="currentColor" className="size-[1.2em]">
                    <path strokeLinecap="round" strokeLinejoin="round"
                      d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z" />
                  </svg>
                  詳細
                </Link>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )
}
