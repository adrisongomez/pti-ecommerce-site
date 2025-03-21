import type { Meta, StoryObj } from "@storybook/react";
import HeroComponent from "@/libs/globals/components/heros/Hero";

type Story = StoryObj<typeof HeroComponent>;

const meta: Meta<typeof HeroComponent> = {
  component: HeroComponent,
};

export const Hero: Story = {
  args: {
    actionText: "CTA Button",
    title: "Hero Title",
    captionText: "Caption of hero, this is just a random long text",
    imageUrl:
      "https://media.istockphoto.com/id/1496111932/photo/the-antique-oil-lamp-on-the-old-wooden-floor-in-the-forest-at-night-camping-atmosphere.jpg?s=612x612&w=0&k=20&c=pxEG7f-1hq9cgvQ7wOH6WrrJgW9s7NeG0rOEWLyHzTA=",
    onActionClick: console.log,
  },
};

export default meta;
