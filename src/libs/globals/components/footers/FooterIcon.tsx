import { joinClass } from "@/libs/globals/utilities/joinClass";
import { FC } from "react";
import { Icon } from "react-feather";
import Button from "../buttons/Button";

type FooterIconProps = {
  Icon: Icon;
};

const FooterIcon: FC<FooterIconProps> = ({ Icon }) => {
  return (
    <Button
      variant="contained"
      className={joinClass(
        "flex",
        "size-6 items-center justify-center rounded-full p-4 shadow transition-all duration-100 ease-in-out hover:bg-white",
        "text-white hover:text-(--bg-dark)",
      )}
    >
      <Icon />
    </Button>
  );
};

export default FooterIcon;
