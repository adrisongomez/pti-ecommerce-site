import Button from "@/globals/components/buttons/Button";
import Card from "@/globals/components/cards/Card/index";
import Textfield from "@/globals/components/fields/Textfield";
import { useRouter } from "@tanstack/react-router";
import { useFormik } from "formik";
import { FC } from "react";
import * as Yup from "yup";

const LoginFormSchema = Yup.object({
  email: Yup.string().email("Must be a valid email").required().strict(),
  password: Yup.string().required(),
});
type LoginFormState = Yup.InferType<typeof LoginFormSchema>;
type LoginFormProps = {
  onLogin?: (d: LoginFormState) => void | Promise<void>;
};

const LoginForm: FC<LoginFormProps> = ({ onLogin }) => {
  const router = useRouter();
  const formik = useFormik<LoginFormState>({
    validationSchema: LoginFormSchema,
    initialValues: {
      email: "",
      password: "",
    },
    onSubmit(d) {
      return onLogin?.(d);
    },
    onReset() {
      router.navigate({
        to: "/",
      });
    },
  });
  return (
    <Card className="w-full max-w-lg bg-(--bg-light) px-10">
      <form
        className="flex w-full flex-1 flex-col gap-2 px-4 py-6"
        onSubmit={formik.handleSubmit}
        onReset={formik.handleReset}
      >
        <h3 className="mb-3 text-center text-3xl font-medium">
          Login to the <strong className="text-(--bg-dark)">Store</strong>
        </h3>
        <Textfield
          required
          label="Email"
          id="email"
          type="email"
          name="email"
          value={formik.values.email}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
        />
        <Textfield
          required
          label="Password"
          id="password"
          type="password"
          name="password"
          value={formik.values.password}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
        />
        <div className="mt-6 flex gap-3">
          <Button className="flex-1 py-1.5" type="submit" variant="contained">
            Login
          </Button>
          <Button className="flex-1 py-1.5" type="reset" variant="outline">
            Sign Up
          </Button>
        </div>
      </form>
    </Card>
  );
};

export default LoginForm;
