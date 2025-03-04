import { FC } from "react";

type FooterTitleProps = {
  label: string;
};

const FooterTitle: FC<FooterTitleProps> = ({ label }) => {
  return <h4 className="text-[16px] font-bold uppercase">{label}</h4>;
};

export default FooterTitle;
