import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ReactNode } from "@tanstack/react-router";
import { FC } from "react";

const CollectionsSection: FC = () => {
  return (
    <div className="grid w-full grid-cols-1 gap-3 md:grid-cols-3">
      <CollectionContainer
        className={joinClass(
          "bg-[url('https://www.naturalbedcompany.co.uk/wp-content/uploads/LED-lamp-collection-portrait.jpg')]",
          "bg-cover bg-center bg-no-repeat",
        )}
      >
        <CollectionTitle>Table Lamps</CollectionTitle>
      </CollectionContainer>
      <CollectionContainer
        className={joinClass(
          "bg-[url('https://m.media-amazon.com/images/I/71D18qNtYvL._AC_UF894,1000_QL80_.jpg')]",
          "bg-cover bg-center bg-no-repeat",
        )}
      >
        <CollectionTitle>Badside lamps</CollectionTitle>
      </CollectionContainer>
      <CollectionContainer
        className={joinClass(
          "h-full md:row-span-2",
          "bg-[url('https://static.ikea.com.do/assets/images/921/0792159_PE764671_S4.webp')]",
          "bg-cover bg-right bg-no-repeat",
        )}
      >
        <CollectionTitle>Floor Lamps</CollectionTitle>
      </CollectionContainer>
      <CollectionContainer
        className={joinClass(
          "md:col-span-2",
          "bg-[url('https://ae01.alicdn.com/kf/S6b36dd155e994dc7b02ab78f3c1032e3J.jpg_640x640q90.jpg')]",
          "bg-cover bg-bottom bg-no-repeat",
        )}
      >
        <CollectionTitle>Desk lamps</CollectionTitle>
      </CollectionContainer>
    </div>
  );
};

const CollectionContainer: FC<{ children: ReactNode; className?: string }> = ({
  children,
  className,
}) => (
  <div
    className={joinClass(
      className ?? "",
      "flex cursor-pointer items-end justify-end md:min-h-64",
      "transition-all duration-500 ease-in-out hover:shadow-lg",
      "hover:opacity-70",
      "overflow-hidden rounded-sm",
      "select-none",
    )}
  >
    {children}
  </div>
);

const CollectionTitle: FC<{ children: ReactNode }> = ({ children }) => (
  <span className="w-full basis-full bg-[rgba(0,0,0,.44)] px-4 py-2 text-right text-sm font-medium text-white uppercase">
    {children}
  </span>
);

export default CollectionsSection;
