import Link from 'next/link'
import React from 'react'

const Sidebar = () => {
  return (
    <div className="flex w-256">
      <ul className="menu">
        <li><Link href="/workloads" className="active">Workloads</Link></li>
        <li><Link href="/versions">Version</Link></li>
        <li><Link href="/workloads">Profile</Link></li>
        <li><Link href="/workloads">Logout</Link></li>
      </ul>
    </div>
  )
}

export default Sidebar
