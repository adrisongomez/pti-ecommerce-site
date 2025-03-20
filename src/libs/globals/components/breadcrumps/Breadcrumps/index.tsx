import { Link, LinkProps } from "@tanstack/react-router";
import { FC, Fragment } from "react";
type BreadcrumpsProps = {
  crumps: ({
    label: string;
  } & Omit<LinkProps, "className">)[];
  currentCrump: string;
};

const Breadcrumps: FC<BreadcrumpsProps> = ({ crumps, currentCrump }) => {
  return (
    <div className="flex w-full flex-row gap-8 select-none">
      {crumps.map((v, i) => (
        <Fragment key={`crumps-${i}`}>
          <Link
            className="font-medium tracking-normal text-gray-500 hover:underline"
            {...v}
          >
            {v.label}
          </Link>
          <p className="text-(--bg-dark)">&#x25A0;</p>
        </Fragment>
      ))}
      <span className="font-medium tracking-normal text-gray-300">
        {currentCrump}
      </span>
    </div>
  );
};

export default Breadcrumps;
