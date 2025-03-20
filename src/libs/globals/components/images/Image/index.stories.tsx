import type { Meta, StoryObj } from "@storybook/react";
import Image from ".";

type ImageType = typeof Image;
type Story = StoryObj<ImageType>;

const meta: Meta<ImageType> = {
  component: Image,
};

export const Default: Story = {
  args: {
    src: "https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcQccmP97OzfxH96Ck6MSlZM72AXSdNjekNh6Q&s",
    height: "126px",
    width: "280px",
    className: "rounded",
  },
};

export default meta;
