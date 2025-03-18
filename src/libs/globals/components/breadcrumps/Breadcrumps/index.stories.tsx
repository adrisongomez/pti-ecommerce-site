import type { Meta, StoryObj } from "@storybook/react";
import Breadcrumps from ".";

type BreadcrumpsType = typeof Breadcrumps;
type Story = StoryObj<BreadcrumpsType>;

const meta: Meta<BreadcrumpsType> = {
  component: Breadcrumps,
};

export const DefaultBreadcrumps: Story = {
  args: {
    crumps: [{ label: "Previous", path: "/previous" }],
    currentCrump: "Current",
  },
};

export default meta;
