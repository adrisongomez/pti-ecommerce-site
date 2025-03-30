import type { Meta, StoryObj } from "@storybook/react";

import Button from "@/libs/globals/components/buttons/Button";

type ButtonType = typeof Button;

const meta: Meta<ButtonType> = {
  component: Button,
};

type Story = StoryObj<ButtonType>;

export const ContainButton: Story = {
  args: {
    children: "Primary Button",
    className: "",
    variant: "contained",
  },
};

export const OutlineButton: Story = {
  args: {
    children: "Primary Button",
    variant: "outline",
    className: "",
  },
};

export const LoadingButton: Story = {
  args: {
    children: "Primary Button",
    variant: "outline",
    className: "",
    loading: true,
  },
};
export default meta;
