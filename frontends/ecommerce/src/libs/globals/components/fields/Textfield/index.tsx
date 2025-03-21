import { joinClass } from "@/libs/globals/utilities/joinClass";
import { FC, InputHTMLAttributes } from "react";

type TextfieldProps = {
  label: string;
  error?: boolean;
  helperText?: string;
} & Omit<InputHTMLAttributes<HTMLInputElement>, "className">;

const Textfield: FC<TextfieldProps> = ({
  label,
  error = false,
  helperText,
  ...props
}) => {
  return (
    <label
      className={joinClass(
        "flex flex-col gap-1.5 text-sm",
        error ? "text-red-600" : "text-(--text-primary)",
      )}
    >
      <span
        className={joinClass(
          "font-semibold",
          error ? "text-red-600" : "text-inherit",
        )}
      >
        {label}{" "}
        {props.required && <span className="font-bold text-red-600">*</span>}
      </span>
      <input
        className={joinClass(
          error ? "border-red-600" : "",
          "text-md rounded-md border p-2 outline-none",
        )}
        {...props}
      />
      {helperText && (
        <p
          className={joinClass(
            "text-sm font-light",
            error ? "text-red-600" : "",
          )}
        >
          {helperText}
        </p>
      )}
    </label>
  );
};

export default Textfield;
