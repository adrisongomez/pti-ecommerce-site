import type { Meta, StoryObj } from "@storybook/react";
import Textfield from "@/globals/components/fields/Textfield";

type TextfieldType = typeof Textfield;
type Story = StoryObj<TextfieldType>;

export const DefaultTextfield: Story = {
  args: { label: "Field Form" },
};

const meta: Meta<TextfieldType> = {
  component: Textfield,
};

export default meta;
