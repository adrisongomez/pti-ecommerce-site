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
        "w-fit items-center justify-center rounded-full p-2",
      )}
    >
      <Icon
        color="white"
        className="shadow transition-all duration-300 ease-in-out active:invert-100"
        size={18}
      />
    </Button>
  );
};

export default FooterIcon;
