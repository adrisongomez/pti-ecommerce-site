import { Link } from "@tanstack/react-router";
import { ComponentProps, FC } from "react";

type NavigationLinkProps = { label: string } & Omit<
  ComponentProps<typeof Link>,
  "className"
>;

const NavigationLink: FC<NavigationLinkProps> = ({ label, ...props }) => {
  return <Link {...props}> {label} </Link>;
};

export default NavigationLink;
