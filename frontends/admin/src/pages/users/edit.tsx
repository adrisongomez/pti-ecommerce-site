import { HttpError, useForm } from "@refinedev/core";
import { Edit } from "@refinedev/mui";
import { FC } from "react";
import { User, UserCreateInput } from "../../generated";
import { useFormik } from "formik";
import * as Yup from "yup";
import {
  Box,
  FormControl,
  FormHelperText,
  InputLabel,
  MenuItem,
  Select,
  TextField,
} from "@mui/material";
import ProductStyledCard from "../products/components/ProductStyledCard";
import { useUpdateEffect } from "react-use";

const EditUserPage: FC = () => {
  const { formLoading, onFinish, query } = useForm<
    User,
    HttpError,
    UserCreateInput
  >({
    resource: "users",
    action: "edit",
  });
  const user = query?.data?.data;
  const formik = useFormik<UserCreateInput>({
    initialValues: {
      lastName: user?.lastName ?? "",
      firstName: user?.firstName ?? "",
      email: user?.email ?? "",
      role: user?.role ?? "CUSTOMER",
      password: "",
    },
    validationSchema: Yup.object({
      firstName: Yup.string().required("required"),
      email: Yup.string()
        .email("Please enter a valid email")
        .required("Required")
        .strict(),
      role: Yup.string().oneOf(["CUSTOMER", "ADMIN"], "Choose a valid type"),
    }),
    onSubmit(d) {
      onFinish(d);
    },
  });

  useUpdateEffect(() => {
    const user = query?.data?.data;
    if (user && !formik.dirty) {
      formik.setValues(
        {
          lastName: user?.lastName ?? "",
          firstName: user?.firstName ?? "",
          email: user?.email ?? "",
          role: user?.role ?? "CUSTOMER",
          password: "",
        },
        false,
      );
    }
  }, [query?.data?.data]);
  return (
    <Box
      component="form"
      onSubmit={formik.handleSubmit}
      onReset={formik.handleReset}
    >
      <Edit isLoading={formLoading} saveButtonProps={{ type: "submit" }}>
        <ProductStyledCard>
          <TextField
            fullWidth
            required
            variant="filled"
            label="First name"
            name="firstName"
            value={formik.values.firstName}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.firstName && !!formik.errors.firstName}
            helperText={formik.touched.firstName && formik.errors.firstName}
            slotProps={{
              inputLabel: {
                shrink: true,
              },
            }}
          />
          <TextField
            fullWidth
            variant="filled"
            label="Last name"
            name="lastName"
            value={formik.values.lastName}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.lastName && !!formik.errors.lastName}
            helperText={formik.touched.lastName && formik.errors.lastName}
            slotProps={{
              inputLabel: {
                shrink: true,
              },
            }}
          />
          <TextField
            fullWidth
            required
            variant="filled"
            label="Email"
            name="email"
            value={formik.values.email}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.email && !!formik.errors.email}
            helperText={formik.touched.email && formik.errors.email}
            slotProps={{
              inputLabel: {
                shrink: true,
              },
            }}
          />
          <FormControl variant="filled" fullWidth>
            <InputLabel shrink>Type</InputLabel>
            <Select
              required
              label="Type"
              name="role"
              value={formik.values.role}
              onChange={formik.handleChange}
              onBlur={formik.handleBlur}
              error={formik.touched.role && !!formik.errors.role}
            >
              <MenuItem value="CUSTOMER"> Customer</MenuItem>
              <MenuItem value="ADMIN"> Admin</MenuItem>
            </Select>

            {formik.touched.role && !!formik.errors.role && (
              <FormHelperText>{formik.errors.firstName}</FormHelperText>
            )}
          </FormControl>
        </ProductStyledCard>
      </Edit>
    </Box>
  );
};

export default EditUserPage;
