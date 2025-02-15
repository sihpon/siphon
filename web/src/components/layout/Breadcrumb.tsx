'use client'

import { useSelectedLayoutSegments } from 'next/navigation'

export function BreadCrumb() {
  const segments = useSelectedLayoutSegments()

  return (
    <div className="breadcrumbs text-sm mb-4">
      <ul>
        {segments.map((segment, index) => (
          <li key={index}>
            <a>
              <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24"
                className="h-4 w-4 stroke-current">
                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                  d="M3 7v10a2 2 0 002 2h14a2 2 0 002-2V9a2 2 0 00-2-2h-6l-2-2H5a2 2 0 00-2 2z"></path>
              </svg>
              {segment}
            </a>
          </li>
        ))}
      </ul>
    </div>
  )
}
