"use client";

import { GetResponse } from "@/generated/workload/v1/workload_pb"
import Client from "@/util/client"
import { NextResponse } from "next/server"
import { useRouter } from "next/navigation";


export default async function Page(
  { params }: { params: { id: string } },
) {
  async function onDelete() {
    try {
      await new Client().Workload().Delete(params.id)
    } catch (e) {
      console.error(e)
      return new NextResponse("ワークロードの削除に失敗", { status: 500 });
    }

    router.push("/workloads")
  }

  const router = useRouter()
  const response: GetResponse = await new Client().Workload().Get(params.id)

  return (
    <div>
      <h1 className="text-5xl font-extrabold mb-4">🏗️ {response.workload?.id}</h1>
      <div className="breadcrumbs text-sm mb-4">
        <ul>
          <li>
            <a>
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                className="h-4 w-4 stroke-current">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                  d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"></path>
              </svg>
              Home
            </a>
          </li>
          <li>
            <a>
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                className="h-4 w-4 stroke-current">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                  d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"></path>
              </svg>
              Workloads
            </a>
          </li>
          <li>
            <span className="inline-flex items-center gap-2 italic underline">
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                className="h-4 w-4 stroke-current">
                <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2"
                  d="M9 13h6m-3-3v6m5 5H7a2 2 0 01-2-2V5a2 2 0 012-2h5.586a1 1 0 01.707.293l5.414 5.414a1 1 0 01.293.707V19a2 2 0 01-2 2z">
                </path>
              </svg>
              Index
            </span>
          </li>
        </ul>
      </div>
      <div className="overflow-x-auto rounded-box border border-base-content/5 bg-base-100">
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
                <button className="btn btn-error" onClick={onDelete}>
                  <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" strokeWidth="2.5"
                    stroke="currentColor" className="size-[1.2em]">
                    <path strokeLinecap="round" strokeLinejoin="round"
                      d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z" />
                  </svg>
                  削除
                </button>
              </td>
            </tr>
          </tbody>
        </table>
      </div>
    </div >
  )
}
