import { Link, ReactNode, useNavigate } from "@tanstack/react-router";
import { FC, useState } from "react";
import Navigation from "../../navigations/Navigation";
import Footer from "../../footers/Footer";
import { joinClass } from "@/libs/globals/utilities/joinClass";
import MainHeader from "../../headers/MainHeader";
import { useAppDispatch, useAppSelector } from "@/libs/globals/hooks/redux";
import Button from "../../buttons/Button";
import authSlice from "@/libs/globals/redux/AuthReducer";
import { ShoppingCart } from "react-feather";
import IconButton from "../../buttons/IconButton";
import AssitantChat from "@/libs/routes/ChatModal";
import { useChatServicePostApiChats } from "@/libs/globals/generated/queries";

type MainLayoutProps = {
  children: ReactNode | ReactNode[];
};

const MainLayout: FC<MainLayoutProps> = ({ children }) => {
  const [open, setOpen] = useState(false);
  const [sessionId, setSessionId] = useState(0);
  const session = useChatServicePostApiChats({
    mutationKey: ["MainLayout_CreateChatSession"],
  });
  const dispatch = useAppDispatch();
  const cartQty = useAppSelector((state) =>
    state.cart.data.flatMap((v) => v.quantity).reduce((acc, v) => acc + v, 0),
  );
  const nav = useNavigate();
  const auth = useAppSelector((state) => state.auth);
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
              <div className="flex items-center gap-3">
                <Link to="/profile">
                  <span className="cursor-pointer select-none">
                    <Button variant="text">Profile</Button>
                  </span>
                </Link>
                <Button
                  variant="text"
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
                  <Button variant="text">Login</Button>
                </Link>
              </div>
            )}
            <div className="relative">
              {!!cartQty && (
                <span className="absolute -top-1 -right-1.5 rounded-full bg-(--bg-dark) p-1 px-2 text-[10px] font-medium text-white">
                  {cartQty}
                </span>
              )}
              <IconButton
                className="p-2"
                onClick={() => {
                  nav({ to: "/carts" });
                }}
              >
                <ShoppingCart className="text-(--bg-dark)" />
              </IconButton>
            </div>
          </div>
        }
      />
      <main
        className={joinClass(
          "m-auto mb-12 flex h-full",
          "w-full max-w-7xl flex-1 bg-transparent px-1.5",
        )}
      >
        {children}
      </main>
      <Footer />
      {auth.status === "logged" && (
        <AssitantChat
          open={open}
          sessionId={sessionId}
          onOpen={async () => {
            setOpen(true);
            const sessionId = await session.mutateAsync();
            setSessionId(sessionId);
          }}
          onClose={async () => {
            setOpen(false);
          }}
        />
      )}
    </div>
  );
};

export default MainLayout;
