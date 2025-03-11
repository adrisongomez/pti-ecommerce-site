import type { Meta, StoryObj } from "@storybook/react";
import Card from "@/globals/components/cards/Card/index";

type CardType = typeof Card;
type Story = StoryObj<CardType>;

const meta: Meta<CardType> = {
  component: Card,
};

export const DefaultCard: Story = {
  args: {
    children: "Card Content",
  },
};

export default meta;
