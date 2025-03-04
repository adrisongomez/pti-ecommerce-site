import { FC } from "react";
import NavigationLink from "./NavigationLink";

const Navigation: FC = () => {
  return (
    <nav className="flex gap-6">
      <NavigationLink to="/about" label="About us" />
      <NavigationLink to="/about" label="Contact" />
      <NavigationLink to="/about" label="Blog" />
    </nav>
  );
};

export default Navigation;
