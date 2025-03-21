import type { Meta, StoryObj } from "@storybook/react";
import Skeleton from ".";

type SkeletonType = typeof Skeleton;
type Story = StoryObj<SkeletonType>;

const meta: Meta<SkeletonType> = {
  component: Skeleton,
};

export const Rectangle: Story = {
  args: {
    variant: "rectangle",
    className: "w-xl h-18",
  },
};

export const Circle: Story = {
  args: {
    variant: "circle",
    className: "w-xl h-18",
    size: 24,
  },
};

export const Composition: Story = {
  render() {
    return (
      <div className="flex w-full flex-row items-center gap-4">
        <Skeleton variant="circle" size={32} />
        <Skeleton variant="rectangle" className="h-10 w-xl" />
      </div>
    );
  },
};

export default meta;
