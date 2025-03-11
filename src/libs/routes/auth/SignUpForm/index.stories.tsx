import type { Meta, StoryObj } from "@storybook/react";
import SignUpForm from "@/libs/routes/auth/SignUpForm";

type SignUpFormType = typeof SignUpForm;
type Story = StoryObj<SignUpFormType>;

const meta: Meta<SignUpFormType> = {
  component: SignUpForm,
};

export const DefaultSignUpForm: Story = {};

export default meta;
