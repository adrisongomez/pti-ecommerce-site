import type { Meta, StoryObj } from "@storybook/react";
import Gallery from ".";

type GalleryType = typeof Gallery;
type Story = StoryObj<GalleryType>;

const meta: Meta<GalleryType> = {
  component: (props) => {
    return (
      <main className="m-auto mb-12 flex h-full w-full max-w-7xl flex-1 bg-transparent px-1.5">
        <Gallery {...props}>
          <div className="h-96">Step {props.value ?? 1}</div>
        </Gallery>
      </main>
    );
  },
};

export const Default: Story = {
  args: {
    value: 1,
  },
};

export const WithMaxValue: Story = {
  args: {
    value: 1,
    maxValue: 4,
  },
};

export const Disabled: Story = {
  args: {
    disabled: true,
  },
};

export default meta;
