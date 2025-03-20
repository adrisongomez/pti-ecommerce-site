import type { Meta, StoryObj } from "@storybook/react";
import CollectionsSection from ".";

type CollectionsSectionType = typeof CollectionsSection;
type Story = StoryObj<CollectionsSectionType>;

const meta: Meta<CollectionsSectionType> = {
  component: (props) => (
    <main className="m-auto mb-12 flex h-full w-full max-w-7xl flex-1 bg-transparent px-1.5">
      <CollectionsSection {...props} />
    </main>
  ),
};

export const Default: Story = {
  args: {},
};

export default meta;
