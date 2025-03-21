import type { Meta, StoryObj } from "@storybook/react";
import SimpleCenterLayout from "@/libs/globals/components/layouts/SimpleCenterLayout";

type SimpleCenterLayoutType = typeof SimpleCenterLayout;
type Story = StoryObj<SimpleCenterLayoutType>;

const meta: Meta<SimpleCenterLayoutType> = {
  component: SimpleCenterLayout,
};

export const DefaultSimpleCenterLayout: Story = {
  args: {
    children: "Simple Center Layout children",
  },
};

export default meta;
