import { joinClass } from "@/libs/globals/utilities/joinClass";
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
    <article className="mb-6 flex w-fit flex-col gap-3">
      <img
        className="w-lg object-cover object-top sm:w-xs md:w-2xs lg:w-3xs"
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
          className="cursor-pointer rounded-md border-none bg-white p-2 shadow outline-none"
        >
          <ShoppingCart size={12} className="fill-(--bg-main)" />
        </button>
      </div>
    </article>
  );
};

export default ProductCard;
