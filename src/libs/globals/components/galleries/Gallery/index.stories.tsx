import type { Meta, StoryObj } from "@storybook/react";
import Gallery from ".";

type GalleryType = typeof Gallery;
type Story = StoryObj<GalleryType>;

const meta: Meta<GalleryType> = {
  component: (props) => (
    <main className="m-auto mb-12 flex h-full w-full max-w-7xl flex-1 bg-transparent px-1.5">
      <Gallery {...props} />
    </main>
  ),
};

export const Default: Story = {
  args: {
    children: [
      <div className="h-96">Step 1</div>,
      <div className="h-96">Step 2</div>,
      <div className="h-96">Step 3</div>,
      <div className="h-96">Step 4</div>,
    ],
  },
};

export default meta;
