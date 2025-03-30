import type { StoryObj, Meta } from "@storybook/react";
import Spinner from ".";

type SpinnerType = typeof Spinner;
type Story = StoryObj<SpinnerType>;

const meta: Meta<SpinnerType> = {
  component: Spinner,
};

export const DefaultSpinner: Story = {};

export default meta;
