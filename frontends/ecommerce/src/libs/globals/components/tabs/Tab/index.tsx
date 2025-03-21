import { joinClass } from "@/libs/globals/utilities/joinClass";
import { JSX, MouseEvent } from "react";

type TabProps<T = string> = {
  value: T;
  onChange: (e: MouseEvent, v: T) => void;
  options: { label: string; value: T }[];
};

function Tab<T>({ options, value, onChange }: TabProps<T>): JSX.Element {
  return (
    <div
      className="flex flex-row select-none"
      style={{ borderBottom: "2px solid var(--color-gray-300)" }}
    >
      {options.map((v, i) => (
        <span
          key={`tab-option-${i}`}
          className={joinClass(
            value === v.value
              ? "rounded-t-lg text-(--bg-dark)"
              : "rounded-t-lg text-gray-500",
            "cursor-pointer px-6 py-2 font-medium hover:bg-slate-200",
            "transition-all duration-200 ease-in-out",
            "active:bg-(--bg-dark) active:text-white",
          )}
          onClick={(e) => onChange(e, v.value)}
          style={
            value === v.value
              ? { borderBottom: "3px solid var(--bg-dark)" }
              : {}
          }
        >
          {v.label}
        </span>
      ))}
    </div>
  );
}

export default Tab;
