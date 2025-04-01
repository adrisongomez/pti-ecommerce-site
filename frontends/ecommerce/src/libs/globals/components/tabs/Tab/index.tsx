import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ReactNode } from "@tanstack/react-router";
import { JSX, MouseEvent } from "react";

type TabProps<T = string> = {
  variant: "vertical" | "horizontal";
  value: T;
  onChange: (e: MouseEvent, v: T) => void;
  options: { label: string; value: T; icon?: ReactNode }[];
};

function Tab<T>({
  options,
  value,
  onChange,
  variant = "horizontal",
}: TabProps<T>): JSX.Element {
  return (
    <div
      className={joinClass(
        "flex gap-1 select-none",
        variant === "vertical" ? "flex-col" : "flex-row",
      )}
    >
      {options.map((v, i) => (
        <span
          key={`tab-option-${i}`}
          className={joinClass(
            "gap-6 flex items-center",
            value === v.value
              ? "bg-(--bg-dark) text-white"
              : "text-(--bg-dark)",
            "cursor-pointer rounded-lg px-6 py-2 font-medium",
            "hover:bg-slate-400 hover:text-white",
            "transition-all duration-200 ease-in-out",
            "active:bg-(--bg-dark) active:text-white",
          )}
          onClick={(e) => onChange(e, v.value)}
        >
          {v?.icon}
          {v.label}
        </span>
      ))}
    </div>
  );
}

export default Tab;
