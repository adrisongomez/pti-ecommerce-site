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
    <div className="flex w-full flex-row">
      <img
        src={imageUrl}
        alt="Hero Image"
        width="40%"
        className="flex-1 object-contain"
      />
      <section className="flex flex-1 flex-col items-start justify-center">
        <h4 className="text-lg font-bold">{title}</h4>
        <caption className="font-medium text-(--text-accent)">
          {captionText}
        </caption>
        <Button className="px-8 py-6" onClick={onActionClick}>
          {actionText}
        </Button>
      </section>
    </div>
  );
};

export default Hero;
