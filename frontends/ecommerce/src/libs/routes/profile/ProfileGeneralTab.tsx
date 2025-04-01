import Button from "@/libs/globals/components/buttons/Button";
import Textfield from "@/libs/globals/components/fields/Textfield";
import { useAuthServiceGetApiAuthMeSuspense } from "@/libs/globals/generated/queries/suspense";
import { useFormik } from "formik";
import { FC } from "react";

const ProfileGeneralTab: FC = () => {
  const user = useAuthServiceGetApiAuthMeSuspense([
    "ProfileGeneratTab_Me_Query",
  ]);
  console.log(user.data);
  const formik = useFormik({
    initialValues: {
      firstName: user.data.firstName,
      lastName: user.data.lastName,
      email: user.data.email,
    },
    onSubmit(d) {
      console.log(d);
    },
  });
  return (
    <form
      onSubmit={formik.handleSubmit}
      onReset={formik.handleReset}
      className="flex flex-col gap-6"
    >
      <h3 className="text-xl font-bold">General</h3>
      <Textfield
        fullWidth
        label="First name"
        name="firstName"
        value={formik.values.firstName}
        onChange={formik.handleChange}
        onBlur={formik.handleBlur}
      />
      <Textfield
        fullWidth
        label="Last name"
        name="lastName"
        value={formik.values.lastName}
        onChange={formik.handleChange}
        onBlur={formik.handleBlur}
      />
      <Textfield
        fullWidth
        label="Email"
        name="email"
        value={formik.values.email}
        onChange={formik.handleChange}
        onBlur={formik.handleBlur}
      />
      <div className="flex items-center justify-end gap-3">
        <Button
          disabled={!formik.isValid || formik.isSubmitting}
          loading={formik.isSubmitting}
          className="w-22"
          type="submit"
        >
          Save
        </Button>
        <Button className="w-22" variant="outlined" type="reset">
          Cancel
        </Button>
      </div>
    </form>
  );
};

export default ProfileGeneralTab;
