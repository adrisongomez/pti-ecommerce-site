import { ReactNode } from "@tanstack/react-router";
import { FC } from "react";
import Navigation from "../navigations/Navigation";
import Footer from "../footers/Footer";
import { joinClass } from "@/globals/utilities/joinClass";
import MainHeader from "../headers/MainHeader";
import { Smartphone } from "react-feather";

type MainLayoutProps = {
  children: ReactNode[];
};

const MainLayout: FC<MainLayoutProps> = ({ children }) => {
  return (
    <div
      className={joinClass(
        "font-monserat flex h-full min-h-screen",
        "flex-col justify-between",
      )}
    >
      <MainHeader
        navigation={<Navigation />}
        logo={<div />}
        actions={
          <div className="no-wrap flex flex-row items-center gap-2">
            <Smartphone />
            <span>123 456 7891</span>
          </div>
        }
      />
      <main className="flex flex-1 flex-col">{children}</main>
      <Footer />
    </div>
  );
};

export default MainLayout;
