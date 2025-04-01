import type { Meta, StoryObj } from "@storybook/react";
import Tab from ".";

type TabType = typeof Tab<"damage" | "emotional">;
type Story = StoryObj<TabType>;

const meta: Meta<TabType> = {
  component: Tab,
};

export const DefaultTab: Story = {
  args: {
    variant: "horizontal",
    onChange: console.log,
    options: [
      { label: "Emotional", value: "emotional" },
      { label: "Damage", value: "damage" },
    ],
    value: "damage",
  },
};

export const VerticalTab: Story = {
  args: {
    variant: "vertical",
    onChange: console.log,
    options: [
      { label: "Emotional", value: "emotional" },
      { label: "Damage", value: "damage" },
    ],
    value: "damage",
  },
};

export default meta;
