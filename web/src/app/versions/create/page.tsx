import { createVersion } from "./action/create";

export default async function Page() {
  return (
    <div className="card bg-base-100 shadow-sm">
      <div className="card-body">
        <form className="form" action={createVersion}>
          <label className="floating-label">
            <span>ID</span>
            <input name="version" className="input" />
          </label>

          <label className="floating-label">
            <span>説明</span>
            <input name="description" type="text" placeholder="Your name" className="input" />
          </label>

          <button type="submit" className="btn btn-primary">作成する</button>
        </form>
      </div>
    </div >
  )
}

