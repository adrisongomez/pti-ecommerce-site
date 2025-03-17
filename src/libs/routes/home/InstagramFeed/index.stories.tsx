import type { Meta, StoryObj } from "@storybook/react";
import InstagramFeed from ".";

type InstagramFeedType = typeof InstagramFeed;
type Story = StoryObj<InstagramFeedType>;

const meta: Meta<InstagramFeedType> = {
  component: InstagramFeed,
};

export const Default: Story = {
  args: {},
};

export default meta;
