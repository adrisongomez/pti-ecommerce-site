import { joinClass } from "@/libs/globals/utilities/joinClass";
import { CSSProperties, FC } from "react";

type SkeletonProps = {
  className?: string;
  style?: CSSProperties;
} & (
  | {
      variant: "rectangle";
      size?: CSSProperties["width"];
    }
  | {
      variant: "circle";
      size: CSSProperties["width"];
    }
);

const Skeleton: FC<SkeletonProps> = ({
  variant = "rectangle",
  style = {},
  className = "",
  ...props
}) => {
  return (
    <div
      className={joinClass(
        className,
        variant === "rectangle" ? "rounded-md" : "rounded-full",
        "animate-pulse bg-slate-300",
        "shadow-sm",
      )}
      style={{
        ...style,
        ...(variant == "circle"
          ? {
              height: props.size,
              width: props.size,
            }
          : {}),
      }}
    />
  );
};

export default Skeleton;
