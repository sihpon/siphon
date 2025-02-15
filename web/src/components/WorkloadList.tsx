import { create } from '@bufbuild/protobuf'
import { ListResponse, ListResponseSchema } from '../generated/workload/v1/workload_pb'
import Client from '../util/client'
import { timestampDate } from '@bufbuild/protobuf/wkt'
import Link from 'next/link'
import React from 'react'

export default function Page() {
  let response: ListResponse = create(ListResponseSchema);
  new Client().Workload().List((error: any, res: ListResponse) => {
    console.log("feching workloads");
    console.log(error);
    console.log(res);
    if (error) {
      console.error(error)
      return
    }
    response = res
  });

  return (
    <table className="table">
      <thead>
        <tr className="text-xs uppercase italic font-stretch-extra-condensed">
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
              <Link className="btn btn-neutral" href={`/workloads/${workload.id}`}>
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
  )
}
