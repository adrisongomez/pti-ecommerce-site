import { HttpError, useForm } from "@refinedev/core";
import { FC } from "react";
import { Product, ProductInput } from "../../generated";
import { useFormik } from "formik";
import {
  Autocomplete,
  Card,
  Container,
  FormLabel,
  Stack,
  styled,
  TextField,
  Typography,
} from "@mui/material";
import MDEditor from "@uiw/react-md-editor";
import ProductMediaDropzone from "./ProductMediaDropzone";

const CreateProductForm: FC = () => {
  const form = useForm<Product, HttpError, ProductInput>({
    resource: "products",
    action: "create",
  });
  const formik = useFormik<ProductInput>({
    initialValues: {
      description: "",
      tags: [],
      title: "",
      variants: [],
      vendorId: 1,
    },
    onSubmit(data) {
      form.onFinish(data);
    },
  });
  return (
    <Container
      maxWidth={"md"}
      component="form"
      onSubmit={formik.handleSubmit}
      onReset={formik.handleReset}
    >
      <Stack spacing={6}>
        <StyledCard raised variant="elevation">
          <Typography variant="h6" fontWeight="500">
            General
          </Typography>
          <TextField
            fullWidth
            required
            label="Title"
            variant="filled"
            value={formik.values.title}
            onChange={formik.handleChange}
            name="title"
          />
          <Autocomplete
            fullWidth
            multiple
            freeSolo
            onChange={(_, v) => {
              formik.setFieldValue("tags", v);
            }}
            value={formik.values.tags}
            options={[]}
            renderInput={(props) => (
              <TextField label="Tags" name="tags" variant="filled" {...props} />
            )}
          />
          <Stack spacing={1}>
            <FormLabel>Description</FormLabel>
            <MDEditor
              value={formik.values.description}
              onChange={(value) => formik.setFieldValue("description", value)}
            />
          </Stack>
        </StyledCard>
        <StyledCard>
          <Typography variant="h6" fontWeight="500">
            Medias
          </Typography>
          <ProductMediaDropzone
            medias={formik.values.medias ?? []}
            onChangeMedia={(medias) => {
              formik.setFieldValue("medias", medias);
            }}
          />
        </StyledCard>
        <StyledCard>
          <Typography variant="h6" fontWeight="500">
            Variants
          </Typography>
        </StyledCard>
      </Stack>
    </Container>
  );
};

const StyledCard = styled(Card)(({ theme }) => ({
  padding: theme.spacing(4),
  display: "flex",
  flexDirection: "column",
  gap: theme.spacing(3),
}));

export default CreateProductForm;
