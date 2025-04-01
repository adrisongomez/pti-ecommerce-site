import { Link } from "@tanstack/react-router";
import { ComponentProps, FC, ReactNode } from "react";

type NavigationLinkProps = { label: ReactNode } & Omit<
  ComponentProps<typeof Link>,
  "className"
>;

const NavigationLink: FC<NavigationLinkProps> = ({ label, ...props }) => {
  return <Link {...props}> {label} </Link>;
};

export default NavigationLink;
