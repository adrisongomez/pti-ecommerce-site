import { PRODUCTS } from "@/assets/data";
import ProductCard from "@/libs/globals/components/cards/ProductCard";
import { FC } from "react";

const FeatureProductsGallery: FC = () => (
  <div className="flex w-full flex-row gap-12 overflow-x-scroll overflow-y-hidden">
    {PRODUCTS.slice(0, 5).map((p, i) => (
      <ProductCard
        key={`key-${i}-product-card`}
        title={p.name}
        variants={p.colorOptions.map((v) => ({
          imageUrl: v.imageUrl,
          colorSwatch: v.colorSwatch,
          price: "500",
        }))}
      />
    ))}
  </div>
);

export default FeatureProductsGallery;
