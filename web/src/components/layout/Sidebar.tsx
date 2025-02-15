import Link from 'next/link'
import React from 'react'

const Sidebar = () => {
  return (
    <div id="sidebar" className="z-10 w-64 bg-base-200 p-4 hidden lg:block fixed lg:relative h-full">
      <h2 className="text-xl font-bold mb-4">Menu</h2>
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
