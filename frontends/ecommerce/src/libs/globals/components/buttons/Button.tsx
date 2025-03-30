import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ButtonHTMLAttributes, FC } from "react";
import LoadingIndicator from "../progress/LoadingIndicator";

type ButtonVariantType = "contained" | "outlined" | "text";

type ButtonProps = {
  variant?: ButtonVariantType;
  disabled?: boolean;
  loading?: boolean;
} & ButtonHTMLAttributes<HTMLButtonElement>;

function getClassNameByVariant(variant: ButtonVariantType): string {
  switch (variant) {
    case "text":
      return joinClass`
hover:bg-(--bg-main) active:bg-(--bg-light)
text-(--bg-dark) font-medium
`;
    case "contained": {
      return joinClass(
        "bg-(--bg-dark)",
        "text-white hover:bg-(--bg-main) hover:text-(--bg-dark) active:bg-(--bg-light) active:text-(--bg-dark)",
        "border border-solid border-(--bg-dark)",
      );
    }
    case "outlined":
    default: {
      return joinClass(
        "text-(--bg-dark)",
        "hover:bg-(--bg-dark) hover:text-white active:bg-(--bg-light) active:text-(--bg-dark)",
        "border border-solid border-(--bg-dark)",
      );
    }
  }
}

const Button: FC<ButtonProps> = ({
  variant = "contained",
  loading = false,
  disabled = false,
  ...props
}) => {
  if (loading) {
    props.children = <LoadingIndicator />;
  }
  return (
    <button
      {...props}
      className={joinClass(
        props.className ?? "",
        "flex justify-center",
        "box-content cursor-pointer px-4 py-2 uppercase",
        "transition-all duration-300 ease-in-out",
        !disabled ? "" : "",
        ...(disabled
          ? ["bg-(--text-accent) text-(--bg-altern)"]
          : [getClassNameByVariant(variant)]),
      )}
    />
  );
};

export default Button;
