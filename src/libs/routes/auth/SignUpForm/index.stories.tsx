import type { Meta, StoryObj } from "@storybook/react";
import SignUpForm from "@/libs/routes/auth/SignUpForm";

type SignUpFormType = typeof SignUpForm;
type Story = StoryObj<SignUpFormType>;

const meta: Meta<SignUpFormType> = {
  component: (props) => (
    <div className="flex w-full items-center justify-center">
      <SignUpForm {...props} />
    </div>
  ),
};

export const DefaultSignUpForm: Story = {};

export default meta;
