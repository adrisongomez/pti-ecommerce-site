import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ReactNode } from "@tanstack/react-router";
import { FC, HTMLAttributes } from "react";

type CardProps = {
  children?: ReactNode | ReactNode[];
} & HTMLAttributes<HTMLDivElement>;

const Card: FC<CardProps> = ({ children, ...props }) => {
  return (
    <div
      {...props}
      className={joinClass(props.className ?? "", "rounded-lg p-3 shadow-lg")}
    >
      {children}
    </div>
  );
};

export default Card;
