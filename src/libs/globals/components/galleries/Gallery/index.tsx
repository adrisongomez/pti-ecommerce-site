import { FC, ReactNode } from "react";
import { ArrowLeft, ArrowRight } from "react-feather";
import IconButton from "../../buttons/IconButton";

type GalleryProps = {
  value: number;
  disabled?: boolean;
  onChange?: (value: number) => void | Promise<void>;
  maxValue?: number;
  children?: ReactNode | ReactNode[];
};

const Gallery: FC<GalleryProps> = ({
  children,
  disabled = false,
  value,
  maxValue,
  onChange,
}) => {
  const handlePreviousClick = () => onChange?.(value - 1);
  const handleNextClick = () => onChange?.(value + 1);
  return (
    <div className="relative w-full">
      <div className="flex w-full flex-row">{children}</div>
      <div className="flex items-center justify-end gap-3">
        <IconButton
          disabled={value === 1 || disabled}
          onClick={handlePreviousClick}
        >
          <ArrowLeft />
        </IconButton>
        {maxValue && (
          <>
            {value} / {maxValue}
          </>
        )}
        <IconButton
          disabled={value === maxValue || disabled}
          onClick={handleNextClick}
        >
          <ArrowRight />
        </IconButton>
      </div>
    </div>
  );
};

export default Gallery;
