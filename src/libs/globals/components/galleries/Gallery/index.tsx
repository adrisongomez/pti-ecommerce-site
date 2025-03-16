import { FC, ReactNode, useState } from "react";
import { ArrowLeft, ArrowRight } from "react-feather";
import IconButton from "../../buttons/IconButton";

type GalleryProps = {
  children: ReactNode | ReactNode[];
};

const Gallery: FC<GalleryProps> = ({ children }) => {
  const [currentPosition, setCurrentPosition] = useState(0);
  if (!Array.isArray(children)) {
    return children;
  }
  const maxLength = children?.length ?? 0;
  const handlePreviousClick = () =>
    setCurrentPosition((v) => {
      const nextValue = v - 1;
      return nextValue;
    });
  const handleNextClick = () =>
    setCurrentPosition((v) => {
      const nextValue = v + 1;
      return nextValue;
    });
  return (
    <div className="relative w-full">
      <div className="flex w-full flex-row">{children[currentPosition]}</div>
      <div className="flex items-center justify-end gap-3">
        <IconButton
          disabled={currentPosition === 0}
          onClick={handlePreviousClick}
        >
          <ArrowLeft />
        </IconButton>
        {currentPosition + 1} / {maxLength}
        <IconButton
          disabled={currentPosition + 1 === maxLength}
          onClick={handleNextClick}
        >
          <ArrowRight />
        </IconButton>
      </div>
    </div>
  );
};

export default Gallery;
