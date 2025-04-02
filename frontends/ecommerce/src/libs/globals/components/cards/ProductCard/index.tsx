import { joinClass } from "@/libs/globals/utilities/joinClass";
import { FC, MouseEventHandler, useState } from "react";
import { Heart, ShoppingCart } from "react-feather";
import { FormattedNumber } from "react-intl";
import IconButton from "../../buttons/IconButton";

type Variant = {
    imageUrl: string;
    colorSwatch: string;
    price: string;
};

type ProductCardProps = {
  onClick?: MouseEventHandler<HTMLDivElement>;
  onAddToCarcClick?: MouseEventHandler<HTMLButtonElement>;
  title: string;
  variants: Variant[];
  label?: {
    color: string;
    labelTitle: string;
  };
};

const ProductCard: FC<ProductCardProps> = ({
  variants,
  title,
  onAddToCarcClick,
  label,
  onClick,
}) => {
  const [currentPosition, setCurrentPosition] = useState<number | null>();
  const [lower, max] = [
    Math.min(...variants.map((v) => Number(v.price))),
    Math.max(...variants.map((v) => Number(v.price))),
  ];
  const currentVariant = variants.at(currentPosition ?? 0);
  const imageUrl = currentVariant?.imageUrl ?? "placeholder";
  return (
    <article
      className="relative mb-6 flex w-fit flex-col gap-3"
      onClick={onClick}
    >
      {label && (
        <span
          className="absolute top-4 -left-3 w-[10ch] text-center font-bold tracking-widest text-white uppercase"
          style={{ background: label.color }}
        >
          {label.labelTitle}
        </span>
      )}
      <div className="absolute top-3 right-2">
        <IconButton>
          <Heart size={16} />
        </IconButton>
      </div>
      <img
        className="h-[256px] w-[256px] rounded object-cover object-top sm:w-xs md:w-2xs lg:w-3xs"
        src={imageUrl}
      />
      <div className="flex flex-row gap-4">
        {variants.map((v, i) => (
          <div
            key={"variant-" + i}
            onClick={() => setCurrentPosition(i)}
            style={{ background: v.colorSwatch }}
            className={joinClass(`size-4 rounded-full`, "cursor-pointer")}
          />
        ))}
      </div>
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
            {currentVariant && currentPosition !== null ? (
              currentVariant.price
            ) : (
              <>
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
                />{" "}
              </>
            )}
          </span>
        </div>
        <button
          onClick={onAddToCarcClick}
          className="cursor-pointer rounded-md border-none bg-white p-2 shadow outline-none"
        >
          <ShoppingCart size={18} className="fill-(--bg-main)" />
        </button>
      </div>
    </article>
  );
};

export default ProductCard;
