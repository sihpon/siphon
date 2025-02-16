"use client";

import { useState } from "react";

interface DeleteButtonProps {
  deleter: Promise<void>;
  itemName?: string;
}

export default function DeleteButton({ deleter, itemName }: DeleteButtonProps) {
  const [loading, setLoading] = useState(false);

  const handleDelete = async () => {
    if (!confirm(`本当に削除しますか？ ${itemName ? `(${itemName})` : ""}`)) {
      return;
    }

    setLoading(true);
    await deleter;
    setLoading(false);
  };

  return (
    <button onClick={handleDelete} className="btn btn-error" disabled={loading}>
      {loading ? "削除中..." : "削除"}
    </button>
  );
}
