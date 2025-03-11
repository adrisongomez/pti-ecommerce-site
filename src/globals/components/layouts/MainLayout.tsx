import { ReactNode } from "@tanstack/react-router";
import { FC } from "react";
import Navigation from "../navigations/Navigation";
import Footer from "../footers/Footer";
import { joinClass } from "@/globals/utilities/joinClass";
import MainHeader from "../headers/MainHeader";
import { Smartphone, User } from "react-feather";

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
          <div className="flex flex-row items-center gap-4">
            <User />
            <div className="no-wrap flex flex-row items-center gap-2">
              <Smartphone />
              <span>123 456 7891</span>
            </div>
          </div>
        }
      />
      <main className="m-auto mb-12 flex h-full w-full max-w-7xl flex-1 bg-transparent px-1.5">
        {children}
      </main>
      <Footer />
    </div>
  );
};

export default MainLayout;
