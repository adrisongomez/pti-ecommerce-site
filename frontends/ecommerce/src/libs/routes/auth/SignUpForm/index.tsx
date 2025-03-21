import Button from "@/libs/globals/components/buttons/Button";
import Card from "@/libs/globals/components/cards/Card";
import Textfield from "@/libs/globals/components/fields/Textfield";
import { useFormik } from "formik";
import { FC } from "react";
import * as Yup from "yup";

const SignUpDataSchema = Yup.object({
  firstName: Yup.string().required("First name is required"),
  lastName: Yup.string().optional(),
  email: Yup.string()
    .email("Email isn't valid")
    .strict()
    .required("Email is required"),
  password: Yup.string()
    .required("Password is required")
    .min(8, "Password must be at least 8 characters"),
  confirmPassword: Yup.string()
    .required("Please confirm your password")
    .oneOf([Yup.ref("password")], "Password must match"),
});

type SignUpDataFormState = Yup.InferType<typeof SignUpDataSchema>;

type SignUpFormProps = {
  onSignUpForm?: (data: SignUpDataFormState) => void | Promise<void>;
};

const SignUpForm: FC<SignUpFormProps> = ({ onSignUpForm }) => {
  const formik = useFormik<SignUpDataFormState>({
    validationSchema: SignUpDataSchema,
    initialValues: {
      firstName: "",
      password: "",
      confirmPassword: "",
      email: "",
      lastName: "",
    },
    onSubmit(d) {
      onSignUpForm?.(d);
    },
  });
  return (
    <Card className="w-full max-w-lg bg-(--bg-light) px-10">
      <form className="flex w-full flex-col gap-4 px-4 py-6">
        <h3 className="mb-3 text-center text-3xl font-medium">
          Sign to the <strong className="text-(--bg-dark)">Store</strong>
        </h3>
        <Textfield
          required
          label="First name"
          name="firstName"
          error={formik.touched.firstName && Boolean(formik.errors.firstName)}
          value={formik.values.firstName}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          helperText={
            formik.touched.firstName ? formik.errors.firstName : undefined
          }
        />
        <Textfield
          label="Last name"
          name="lastName"
          value={formik.values.lastName}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
        />
        <Textfield
          required
          label="Email"
          name="email"
          value={formik.values.email}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.email && Boolean(formik.errors.email)}
          helperText={formik.touched.email ? formik.errors.email : undefined}
        />
        <Textfield
          required
          label="Password"
          type="password"
          name="password"
          value={formik.values.password}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.password && Boolean(formik.errors.password)}
          helperText={
            formik.touched.password ? formik.errors.password : undefined
          }
        />
        <Textfield
          required
          label="Confirm password"
          type="password"
          name="confirmPassword"
          value={formik.values.confirmPassword}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={
            formik.touched.confirmPassword &&
            Boolean(formik.errors.confirmPassword)
          }
          helperText={
            formik.touched.confirmPassword
              ? formik.errors.confirmPassword
              : undefined
          }
        />
        <Button className="mt-2 py-2" variant="contained">
          Submit
        </Button>
      </form>
    </Card>
  );
};

export default SignUpForm;
