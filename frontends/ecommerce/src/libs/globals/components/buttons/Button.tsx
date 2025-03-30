import { joinClass } from "@/libs/globals/utilities/joinClass";
import { ButtonHTMLAttributes, FC } from "react";
import LoadingIndicator from "../progress/LoadingIndicator";

type ButtonVariantType = "contained" | "outline";

type ButtonProps = {
  variant?: ButtonVariantType;
  disabled?: boolean;
  loading?: boolean;
} & ButtonHTMLAttributes<HTMLButtonElement>;

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
        "box-content cursor-pointer px-4 uppercase",
        "transition-all duration-300 ease-in-out",
        !disabled ? "border border-solid border-(--bg-dark)" : "",
        ...(disabled
          ? ["bg-(--text-accent) text-(--bg-altern)"]
          : variant === "contained"
            ? [
                "bg-(--bg-dark)",
                "text-white hover:bg-(--bg-main) hover:text-(--bg-dark) active:bg-(--bg-light) active:text-(--bg-dark)",
              ]
            : [
                "text-(--bg-dark)",
                "hover:bg-(--bg-dark) hover:text-white active:bg-(--bg-light) active:text-(--bg-dark)",
              ]),
      )}
    />
  );
};

export default Button;
