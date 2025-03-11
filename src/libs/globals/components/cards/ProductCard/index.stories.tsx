import type { Meta, StoryObj } from "@storybook/react";
import ProductCard from "@/libs/globals/components/cards/ProductCard";
import { PRODUCTS } from "@/assets/data";
import { IntlProvider } from "react-intl";

type ProductCardType = typeof ProductCard;

const meta: Meta<ProductCardType> = {
  component: ProductCard,
  decorators: [
    (Story) => (
      <IntlProvider locale="en" defaultLocale="en">
        <Story />
      </IntlProvider>
    ),
  ],
};

type Story = StoryObj<ProductCardType>;

const PRODUCT = PRODUCTS[0];

export const ProductCardDefault: Story = {
  args: {
    title: PRODUCT.name,
    variants: PRODUCT.colorOptions.map((v) => ({
      imageUrl: v.imageUrl,
      colorSwatch: v.colorSwatch,
      price: Math.random() * 100.0,
    })),
  },
};

export default meta;
