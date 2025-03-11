import { ReactNode } from "@tanstack/react-router";
import { FC } from "react";

type SimpleCenterLayoutProps = {
  children: ReactNode | ReactNode[];
};

const SimpleCenterLayout: FC<SimpleCenterLayoutProps> = ({ children }) => {
  return (
    <div className="flex h-full w-full flex-1 items-center justify-center">
      {children}
    </div>
  );
};

export default SimpleCenterLayout;
