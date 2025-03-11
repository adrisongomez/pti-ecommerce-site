import type { Meta, StoryObj } from "@storybook/react";
import Textfield from "@/globals/components/fields/Textfield";

type TextfieldType = typeof Textfield;
type Story = StoryObj<TextfieldType>;

export const DefaultTextfield: Story = {
  args: { label: "Field Form" },
};

export const RequiredTextfield: Story = {
  args: { label: "Field Form", required: true },
};

export const ErrorTextfield: Story = {
  args: {
    label: "Field Form",
    error: true,
    helperText: "Field Form error message",
  },
};

const meta: Meta<TextfieldType> = {
  component: Textfield,
};

export default meta;
