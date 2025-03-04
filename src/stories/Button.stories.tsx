import type { Meta, StoryObj } from "@storybook/react";

import Button from "@/globals/components/buttons/Button";

type ButtonType = typeof Button;

const meta: Meta<ButtonType> = {
  component: Button,
};

type Story = StoryObj<ButtonType>;

export const ContainButton: Story = {
  args: {
    children: "Primary Button",
    className: "",
    variant: "contain",
  },
};

export const OutlineButton: Story = {
  args: {
    children: "Primary Button",
    variant: "outline",
    className: "",
  },
};
export default meta;
