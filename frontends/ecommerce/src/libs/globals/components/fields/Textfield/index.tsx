import { joinClass } from "@/libs/globals/utilities/joinClass";
import { FC, InputHTMLAttributes } from "react";
import FormHelperText from "../FormHelperText";

type TextfieldProps = {
  label: string;
  error?: boolean;
  helperText?: string;
  fullWidth?: boolean;
} & Omit<InputHTMLAttributes<HTMLInputElement>, "className">;

const Textfield: FC<TextfieldProps> = ({
  label,
  error = false,
  helperText,
  fullWidth = false,
  ...props
}) => {
  return (
    <label
      className={joinClass(
        fullWidth ? "w-full flex-1" : "",
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
      {helperText && <FormHelperText>{helperText}</FormHelperText>}
    </label>
  );
};

export default Textfield;
