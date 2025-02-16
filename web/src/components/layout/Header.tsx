import Link from "next/link"

interface HeaderProps {
  title: string
}

const Header = ({ title }: HeaderProps) => {
  return (
    <div className="z-50 navbar shadow glass">
      <div className="lg:hidden">
        <button className="btn btn-square btn-ghost">
          hoge
        </button>
      </div>
      <div className="flex-1">
        <Link href="/workloads" className="btn btn-ghost text-xl">{title}</Link>
      </div>
      <div className="flex-none">
        <ul className="menu menu-horizontal px-1">
          <li><Link href="/workloads">Workloads</Link></li>
          <li>
            <details>
              <summary>siphonist@example.com</summary>
              <ul className="bg-base-100 rounded-t-none p-2">
                <li><a>Link 1</a></li>
                <li><a>Link 2</a></li>
              </ul>
            </details>
          </li>
        </ul>
      </div>
    </div>
  )
}

export default Header
