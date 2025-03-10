import { FC, MouseEventHandler } from "react";
import Button from "../buttons/Button";

type HeroProps = {
  imageUrl: string;
  title: string;
  captionText: string;
  actionText: string;
  onActionClick: MouseEventHandler<HTMLButtonElement>;
};

const Hero: FC<HeroProps> = ({
  onActionClick,
  actionText,
  captionText,
  title,
  imageUrl,
}) => {
  return (
    <div className="mb-12 flex w-full gap-12 md:flex-col lg:flex-row lg:items-center">
      <img src={imageUrl} alt="Hero Image" className="sm:w-full md:w-xl" />
      <section className="flex flex-col items-start gap-6">
        <h4 className="text-7xl font-bold">{title}</h4>
        <caption className="text-lg text-(--text-accent)">
          {captionText}
        </caption>
        <Button className="px-8 py-2 text-lg" onClick={onActionClick}>
          {actionText}
        </Button>
      </section>
    </div>
  );
};

export default Hero;
