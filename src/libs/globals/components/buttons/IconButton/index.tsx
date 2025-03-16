import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ReactNode } from "@tanstack/react-router";
import { ButtonHTMLAttributes, FC } from "react";

type IconButtonProps = {
  children: ReactNode;
} & ButtonHTMLAttributes<HTMLButtonElement>;

const IconButton: FC<IconButtonProps> = ({
  children,
  className = "",
  ...props
}) => {
  return (
    <button
      className={joinClass(
        className,
        "cursor-pointer",
        "rounded-full p-2 hover:bg-slate-200 hover:shadow-md",
        "transition-all duration-500 ease-in-out",
      )}
      {...props}
    >
      {children}
    </button>
  );
};

export default IconButton;
