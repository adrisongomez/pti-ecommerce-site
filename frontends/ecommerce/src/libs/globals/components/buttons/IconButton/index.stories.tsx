import type { Meta, StoryObj } from "@storybook/react";
import IconButton from ".";
import { Heart } from "react-feather";

type IconButtonType = typeof IconButton;
type Story = StoryObj<IconButtonType>;

const meta: Meta<IconButtonType> = {
  component: IconButton,
};

export const Default: Story = {
  args: {
    children: <Heart size="18" />,
  },
};

export default meta;
