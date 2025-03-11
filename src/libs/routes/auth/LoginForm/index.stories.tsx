import type { Meta, StoryObj } from "@storybook/react";
import LoginForm from "@/libs/routes/auth/LoginForm";

type LoginFormType = typeof LoginForm;
type Story = StoryObj<LoginFormType>;

const meta: Meta<LoginFormType> = {
  component: (props) => (
    <div className="flex w-full items-center justify-center">
      <LoginForm {...props} />
    </div>
  ),
};

export const DefaultLoginForm: Story = {
  args: {},
};

export default meta;
