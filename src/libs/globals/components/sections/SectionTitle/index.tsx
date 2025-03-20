import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ReactNode } from "@tanstack/react-router";
import { FC } from "react";
const SectionTitle: FC<{ children: ReactNode; className?: string }> = ({
  children,
  className = "",
}) => (
  <h3
    className={joinClass(
      className,
      "mb-12 text-2xl font-medium text-(--text-accent) uppercase",
    )}
  >
    {children}
  </h3>
);

export default SectionTitle;
