import { ReactNode } from "react";

interface TableProps {
  columns: string[];
  children: ReactNode;
}

export default function Table({ columns, children }: TableProps) {
  return (
    <div className="overflow-x-auto bg-base-100 shadow">
      <table className="table glass">
        <thead>
          <tr className="text-xs uppercase italic font-stretch-extra-condensed bg-base-200">
            {columns.map((col, index) => (
              <th key={index}>{col}</th>
            ))}
          </tr>
        </thead>
        <tbody>
          {children}
        </tbody>
      </table>
    </div>
  );
}
