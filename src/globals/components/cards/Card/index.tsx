import { joinClass } from "@/globals/utilities/joinClass";
import { ReactNode } from "@tanstack/react-router";
import { FC, HTMLAttributes } from "react";

type CardProps = {
  children?: ReactNode | ReactNode[];
} & HTMLAttributes<HTMLDivElement>;

const Card: FC<CardProps> = ({ children, ...props }) => {
  return (
    <div
      className={joinClass(props.className ?? "", "rounded-lg p-3 shadow")}
      {...props}
    >
      {children}
    </div>
  );
};

export default Card;
