import { FC, Fragment, ImgHTMLAttributes, useState } from "react";
import Skeleton from "../../loaders/Skeleton";
import { joinClass } from "@/libs/globals/utilities/joinClass";

type ImageProps = {} & ImgHTMLAttributes<HTMLImageElement>;

const Image: FC<ImageProps> = ({ className = "", ...props }) => {
  const [loaded, setLoaded] = useState(false);
  return (
    <Fragment>
      <img
        className={joinClass(loaded ? "flex" : "hidden", className)}
        onLoad={() => {
          setLoaded(true);
        }}
        {...props}
      />
      {!loaded && (
        <Skeleton
          variant="rectangle"
          className={className}
          style={{ width: props.width, height: props.height }}
        />
      )}
    </Fragment>
  );
};

export default Image;
