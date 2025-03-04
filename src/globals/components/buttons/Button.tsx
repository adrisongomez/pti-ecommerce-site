import { joinClass } from "@/globals/utilities/joinClass";
import { ButtonHTMLAttributes, FC } from "react";

type ButtonProps = {} & ButtonHTMLAttributes<HTMLButtonElement>;

const Button: FC<ButtonProps> = ({ ...props }) => {
  return (
    <button
      {...props}
      className={joinClass(
        props.className ?? "p-2",
        "cursor-pointer bg-(--bg-dark) text-white uppercase",
        "transition-all duration-300 ease-in-out active:bg-(--bg-light) active:text-black",
      )}
    ></button>
  );
};

export default Button;
