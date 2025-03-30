import type { Meta, StoryFn, StoryObj } from "@storybook/react";

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
    variant: "outlined",
    className: "",
  },
};

export const TextButton: Story = {
  args: {
    children: "Primary Button",
    variant: "text",
    className: "",
  },
};

export const LoadingButton: StoryFn = () => {
  return (
    <div className="flex w-sm flex-col gap-5">
      <div className="flex items-center gap-5">
        <p className="w-[10ch] font-bold">Text</p>
        <Button variant="text" className="flex-1" loading></Button>
      </div>
      <div className="flex items-center gap-5">
        <p className="w-[10ch] font-bold">Outlined</p>
        <Button variant="outlined" className="flex-1" loading></Button>
      </div>
      <div className="flex items-center gap-5">
        <p className="w-[10ch] font-bold">Contained</p>
        <Button variant="contained" className="flex-1" loading></Button>
      </div>
      <div className="flex items-center gap-5">
        <p className="w-[10ch] font-bold">Disabled</p>
        <Button
          variant="contained"
          disabled
          className="flex-1"
          loading
        ></Button>
      </div>
    </div>
  );
};
export default meta;
