import { FC, Fragment } from "react";
type BreadcrumpsProps = {
  crumps: {
    label: string;
    path: string;
  }[];
  currentCrump: string;
};

const Breadcrumps: FC<BreadcrumpsProps> = ({ crumps, currentCrump }) => {
  return (
    <div className="flex w-full flex-row gap-8 select-none">
      {crumps.map((v, i) => (
        <Fragment key={`crumps-${i}`}>
          <a
            href={v.path}
            className="font-medium tracking-normal text-gray-500 hover:underline"
          >
            {v.label}
          </a>
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
