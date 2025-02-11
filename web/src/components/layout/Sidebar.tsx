import React from 'react'

const Sidebar = () => {
  return (
    <div id="sidebar" className="z-10 w-64 bg-base-200 p-4 hidden lg:block fixed lg:relative h-full">
      <h2 className="text-xl font-bold mb-4">Menu</h2>
      <ul className="menu">
        <li><a className="active">Dashboard</a></li>
        <li><a>Settings</a></li>
        <li><a>Profile</a></li>
        <li><a>Logout</a></li>
      </ul>
    </div>
  )
}

export default Sidebar
