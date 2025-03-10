import { FC } from "react";
import { ShoppingCart } from "react-feather";
import { FormattedNumber } from "react-intl";

type ProductCardProps = {
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

const ProductCard: FC<ProductCardProps> = ({ variants, title }) => {
  const [lower, max] = [
    Math.min(...variants.map((v) => v.price)),
    Math.max(...variants.map((v) => v.price)),
  ];
  const imageUrl = variants.at(0)?.imageUrl;
  return (
    <article className="flex w-md flex-col gap-6">
      <img
        width="100%"
        className="h-[240px] object-cover object-center"
        src={imageUrl}
      />
      <div>
        <div>
          <span className="name">{title}</span>
          <span className="price">
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
        <button>
          <ShoppingCart />
        </button>
      </div>
    </article>
  );
};

export default ProductCard;
