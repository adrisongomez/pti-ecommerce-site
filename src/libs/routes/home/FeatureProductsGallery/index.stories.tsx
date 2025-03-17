import type { Meta, StoryObj } from "@storybook/react";
import FeatureProductsGallery from ".";

type FeatureProductsGalleryType = typeof FeatureProductsGallery;
type Story = StoryObj<FeatureProductsGalleryType>;

const meta: Meta<FeatureProductsGalleryType> = {
  component: FeatureProductsGallery,
};

export const Default: Story = {
  args: {},
};

export default meta;
