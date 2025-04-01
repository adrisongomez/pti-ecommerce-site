import { ReactNode } from "@tanstack/react-router";
import { JSX } from "react";

function Table<T extends Record<string, ReactNode>>({
  heads,
  rows,
}: {
  heads: { field: keyof T; label: ReactNode }[];
  rows: T[];
}): JSX.Element {
  return (
    <div className="relative overflow-x-auto">
      <table className="w-full bg-(--bg-light) text-left rounded text-sm text-(--bg-dark) rtl:text-right">
        <thead className="bg-(--bg-main) text-xs text-gray-900 uppercase">
          <tr>
            {heads.map((v, i) => (
              <th key={`header-${i}-table`} scope="col" className="px-6 py-3">
                {v.label}
              </th>
            ))}
          </tr>
        </thead>
        <tbody>
          {rows.map((v, i) => (
            <tr key={`key-row-${i}-table`} className="">
              {heads.map((h, i) =>
                i === 0 ? (
                  <th
                    key={`cell-value-table-${i}`}
                    scope="row"
                    className="px-6 py-4 font-medium whitespace-nowrap text-(--text-accent)"
                  >
                    {v[h.field]}
                  </th>
                ) : (
                  <td key={`cell-value-table-${i}`} className="px-6 py-4">
                    {v[h.field]}
                  </td>
                ),
              )}
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}

export default Table;
