import { FC } from "react";
import NavigationLink from "./NavigationLink";
import { Home } from "react-feather";

const Navigation: FC = () => {
  return (
    <nav className="flex gap-6">
      <NavigationLink
        to="/"
        label={
          <div className="flex items-center gap-4 font-bold">
            <Home size={16} /> <p className="pt-1">Home</p>
          </div>
        }
      />
    </nav>
  );
};

export default Navigation;
