import { ReactNode } from "@tanstack/react-router";
import { FC } from "react";

type MainHeaderProps = {
  navigation: ReactNode;
  logo: ReactNode;
  actions: ReactNode;
};

const MainHeader: FC<MainHeaderProps> = ({ navigation, logo, actions }) => {
  return (
    <header className="flex flex-row justify-between px-6 py-2">
      {navigation}
      {logo}
      {actions}
    </header>
  );
};

export default MainHeader;
