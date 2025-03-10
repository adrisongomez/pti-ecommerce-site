import { joinClass } from "@/globals/utilities/joinClass";
import { FC, MouseEventHandler } from "react";
import { ShoppingCart } from "react-feather";
import { FormattedNumber } from "react-intl";

type ProductCardProps = {
  onAddToCarcClick?: MouseEventHandler<HTMLButtonElement>;
  title: string;
  variants: {
    imageUrl: string;
    colorSwatch: string;
    price: number;
  }[];
  label?: {
    color: string;
    labelTitle: string;
  };
};

const ProductCard: FC<ProductCardProps> = ({
  variants,
  title,
  onAddToCarcClick,
}) => {
  const [lower, max] = [
    Math.min(...variants.map((v) => v.price)),
    Math.max(...variants.map((v) => v.price)),
  ];
  const imageUrl = variants.at(0)?.imageUrl;
  return (
    <article className="w-4xs flex flex-col gap-3">
      <img
        width="100%"
        className="h-[240px] object-contain object-top"
        src={imageUrl}
      />
      <div className="flex items-center justify-between">
        <div className="flex flex-col items-start gap-0.5">
          <span
            className={joinClass(
              "w-[20ch] overflow-hidden text-sm font-bold",
              "text-ellipsis whitespace-nowrap",
              "text-(--text-primary) uppercase",
            )}
          >
            {title}
          </span>
          <span
            className={joinClass(
              "w-[20ch] overflow-hidden text-sm",
              "font-medium text-ellipsis whitespace-nowrap",
              "text-(--text-accent)",
            )}
          >
            ${" "}
            <FormattedNumber
              value={lower}
              maximumSignificantDigits={2}
              minimumFractionDigits={2}
            />{" "}
            -{" "}
            <FormattedNumber
              value={max}
              maximumSignificantDigits={2}
              minimumFractionDigits={2}
            />
          </span>
        </div>
        <button
          onClick={onAddToCarcClick}
          className="cursor-pointer rounded-md border-none bg-white p-1 shadow shadow-black outline-none"
        >
          <ShoppingCart className="color-(--bg-main) fill-(--bg-main)" />
        </button>
      </div>
    </article>
  );
};

export default ProductCard;
