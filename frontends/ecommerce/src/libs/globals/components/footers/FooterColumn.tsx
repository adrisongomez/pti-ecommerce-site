import { ComponentProps, FC } from "react";
import FooterTitle from "./FooterTitle";
import { ReactNode } from "@tanstack/react-router";

type FooterColumnProps = {
  label: ComponentProps<typeof FooterTitle>["label"];
  children: ReactNode | ReactNode[];
};

const FooterColumn: FC<FooterColumnProps> = ({ label, children }) => {
  return (
    <div className="flex flex-1 flex-col gap-3">
      <FooterTitle label={label} />
      <div className="flex flex-col gap-2">{children}</div>
    </div>
  );
};

export default FooterColumn;
