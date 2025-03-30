import Button from "@/libs/globals/components/buttons/Button";
import Card from "@/libs/globals/components/cards/Card/index";
import FormHelperText from "@/libs/globals/components/fields/FormHelperText";
import Textfield from "@/libs/globals/components/fields/Textfield";
import { joinClass } from "@/libs/globals/utilities/joinClass";
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
  showErrorCrendetialsMessage?: boolean;
};

const LoginForm: FC<LoginFormProps> = ({
  onLogin,
  showErrorCrendetialsMessage = false,
}) => {
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
        to: "/auth/sign-up",
      });
    },
  });
  return (
    <Card className="w-full max-w-lg bg-(--bg-light) px-10">
      <form
        className="flex w-full flex-col gap-2 px-4 py-6"
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
          error={formik.touched.email && Boolean(formik.errors.email)}
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
          error={formik.touched.password && Boolean(formik.errors.password)}
        />
        {showErrorCrendetialsMessage && (
          <FormHelperText error>Email or password is not valid</FormHelperText>
        )}
        <div className="mt-6 flex gap-3">
          <Button
            disabled={!formik.isValid || formik.isSubmitting}
            className={joinClass(
              "flex-1 py-1.5",
              formik.isSubmitting ? "flex items-center justify-center" : "",
            )}
            type="submit"
            variant="contained"
            loading={formik.isSubmitting}
          >
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
