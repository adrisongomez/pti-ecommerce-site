import { FC, InputHTMLAttributes } from "react";

type TextfieldProps = {
  label: string;
  error?: boolean;
  helperText?: string;
} & Omit<InputHTMLAttributes<HTMLInputElement>, "className">;

const Textfield: FC<TextfieldProps> = ({ label, ...props }) => {
  return (
    <label className="flex flex-col gap-1.5 text-sm text-(--text-primary)">
      <span>
        {label}{" "}
        {props.required && <span className="font-bold text-red-900">*</span>}
      </span>
      <input
        className={"text-md rounded-sm border p-2 outline-none"}
        {...props}
      />
    </label>
  );
};

export default Textfield;
