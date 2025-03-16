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
    <div className="mb-12 flex w-full flex-col items-center justify-center gap-12 sm:flex-col md:flex-row">
      <img
        src={imageUrl}
        alt="Hero Image"
        className="w-full bg-black md:w-md lg:w-xl"
      />
      <section className="flex flex-col items-center gap-6 md:items-start">
        <h4 className="text-center text-7xl font-bold text-(--bg-dark) md:text-left">
          {title}
        </h4>
        <p className="text-center text-lg text-(--text-accent) md:text-left">
          {captionText}
        </p>
        <Button className="px-8 py-2 text-lg" onClick={onActionClick}>
          {actionText}
        </Button>
      </section>
    </div>
  );
};

export default Hero;
