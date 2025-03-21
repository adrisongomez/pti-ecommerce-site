import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ReactNode } from "@tanstack/react-router";
import { ButtonHTMLAttributes, FC } from "react";

type IconButtonProps = {
  children: ReactNode;
  disabled?: boolean;
} & ButtonHTMLAttributes<HTMLButtonElement>;

const IconButton: FC<IconButtonProps> = ({
  children,
  className = "",
  disabled,
  ...props
}) => {
  return (
    <button
      disabled={disabled}
      className={joinClass(
        className,
        disabled ? "disabled" : "",
        "cursor-pointer",
        "disabled:text-(--text-accent) disabled:hover:bg-inherit disabled:hover:shadow-none",
        "rounded-full p-2 hover:bg-slate-200 hover:shadow-md",
        "transition-all duration-200 ease-in-out",
        "active:bg-(--bg-dark) active:text-white",
      )}
      {...props}
    >
      {children}
    </button>
  );
};

export default IconButton;
