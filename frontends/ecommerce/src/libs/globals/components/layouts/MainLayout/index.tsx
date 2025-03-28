import { Link, ReactNode } from "@tanstack/react-router";
import { FC } from "react";
import Navigation from "../../navigations/Navigation";
import Footer from "../../footers/Footer";
import { joinClass } from "@/libs/globals/utilities/joinClass";
import MainHeader from "../../headers/MainHeader";
import { useAppDispatch, useAppSelector } from "@/libs/globals/hooks/redux";
import Button from "../../buttons/Button";
import authSlice from "@/libs/globals/redux/AuthReducer";

type MainLayoutProps = {
  children: ReactNode | ReactNode[];
};

const MainLayout: FC<MainLayoutProps> = ({ children }) => {
  const dispatch = useAppDispatch();
  const auth = useAppSelector((state) => state.auth);
  console.log(auth);
  return (
    <div
      className={joinClass(
        "font-monserat flex h-full min-h-screen",
        "flex-col justify-between gap-4",
      )}
    >
      <MainHeader
        navigation={<Navigation />}
        logo={<div />}
        actions={
          <div className="flex flex-row items-center gap-4">
            {auth.user ? (
              <div className="flex gap-3">
                <span className="cursor-pointer select-none">
                  Hi, {auth.user.firstName}.
                </span>
                <Button
                  variant="outline"
                  onClick={() => {
                    dispatch(authSlice.actions.logout());
                  }}
                >
                  Logout
                </Button>
              </div>
            ) : (
              <div className="flex gap-3">
                <Link to="/auth/sign-up">
                  <Button>Sign up</Button>
                </Link>
                <Link to="/auth/login">
                  <Button variant="outline">Login</Button>
                </Link>
              </div>
            )}
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

// <div className="no-wrap flex flex-row items-center gap-2">
//   <Smartphone />
//   <span>123 456 7891</span>
// </div>

export default MainLayout;
