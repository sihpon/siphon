import { createVersion } from "./action/create";

export default async function Page() {
  return (
    <div className="card bg-base-100 shadow-sm">
      <div className="card-body">
        <form className="form" action={createVersion}>
          <label className="select">
            <span className="label">* バージョン</span>
            <select name="version">
              <option>dev</option>
              <option>v2.1.0</option>
              <option>v2.2.0</option>
            </select>
          </label>

          <label className="floating-label">
            <span>* サーバー名</span>
            <input name="name" className="input" />
          </label>

          <label className="floating-label">
            <span>サーバー説明文</span>
            <input name="description" type="text" placeholder="Your name" className="input" />
          </label>


          <div className="divider"></div>

          <button type="submit" className="btn btn-primary">作成する</button>
        </form>
      </div>
    </div >
  )
}

