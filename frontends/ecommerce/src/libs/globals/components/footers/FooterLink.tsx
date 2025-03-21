import { Link } from "@tanstack/react-router";
import { ComponentProps, FC } from "react";

type FooterLinkProps = { label: string } & ComponentProps<typeof Link>;

const FooterLink: FC<FooterLinkProps> = ({ label, ...props }) => {
  return <Link {...props}>{label}</Link>;
};

export default FooterLink;
