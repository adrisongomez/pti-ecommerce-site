import { joinClass } from "@/libs/globals/utilities/joinClass";
import { FC, ReactNode } from "react";

const FormHelperText: FC<{
  children: ReactNode[] | ReactNode;
  error?: boolean;
}> = ({ children, error = false }) => {
  return (
    <p className={joinClass("text-sm font-light", error ? "text-red-600" : "")}>
      {children}
    </p>
  );
};

export default FormHelperText;
