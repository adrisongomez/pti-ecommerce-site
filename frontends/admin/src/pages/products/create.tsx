import { HttpError, useForm } from "@refinedev/core";
import { FC } from "react";
import { Product, ProductInput } from "../../generated";
import { Formik } from "formik";
import {
  Autocomplete,
  Box,
  Card,
  Container,
  FormLabel,
  IconButton,
  Stack,
  styled,
  TextField,
  Typography,
} from "@mui/material";
import MDEditor from "@uiw/react-md-editor";
import ProductMediaDropzone from "./ProductMediaDropzone";
import ProductVariantEntry from "./ProductVariantEntry";
import AddIcon from "@mui/icons-material/Add";

const CreateProductForm: FC = () => {
  const form = useForm<Product, HttpError, ProductInput>({
    resource: "products",
    action: "create",
  });
  return (
    <Formik<ProductInput>
      initialValues={{
        description: "",
        tags: [],
        title: "",
        variants: [{ colorName: "Default", price: 0 }],
        medias: [],
        vendorId: 1,
      }}
      onSubmit={(data) => {
        form.onFinish(data);
      }}
    >
      {(formik) => (
        <Container
          maxWidth={"md"}
          component="form"
          onSubmit={formik.handleSubmit}
          onReset={formik.handleReset}
        >
          <Stack spacing={6}>
            <StyledCard variant="elevation">
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
                slotProps={{
                  inputLabel: {
                    shrink: true,
                  },
                }}
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
                  <TextField
                    label="Tags"
                    name="tags"
                    variant="filled"
                    slotProps={{
                      inputLabel: {
                        shrink: true,
                      },
                    }}
                    {...props}
                  />
                )}
              />
              <Stack spacing={1}>
                <FormLabel>Description</FormLabel>
                <MDEditor
                  value={formik.values.description}
                  onChange={(value) =>
                    formik.setFieldValue("description", value)
                  }
                />
              </Stack>
            </StyledCard>
            <StyledCard variant="elevation">
              <Typography variant="h6" fontWeight="500">
                Medias
              </Typography>
              <ProductMediaDropzone />
            </StyledCard>
            <StyledCard>
              <Typography variant="h6" fontWeight="500">
                Variants
              </Typography>
              <Stack spacing={2}>
                {formik.values.variants.map((v, i) => (
                  <ProductVariantEntry
                    key={`${v.colorName}-${i}`}
                    colorName={v.colorName}
                    price={v.price}
                    colorHex={v.colorHex}
                    onSave={(d) => {
                      formik.setFieldValue(
                        "varians",
                        formik.values.variants.map((v, idx) =>
                          i === idx ? d : v,
                        ),
                      );
                    }}
                    onRemove={() => {
                      if (formik.values.variants.length <= 1) {
                        return;
                      }
                      formik.setFieldValue(
                        "varians",
                        formik.values.variants.filter((_, idx) => i !== idx),
                      );
                    }}
                  />
                ))}
                <Box display="flex" alignItems="center" justifyContent="center">
                  <IconButton
                    onClick={() => {
                      formik.setFieldValue("variants", [
                        ...formik.values.variants,
                        { colorName: "", price: 0, colorHex: undefined },
                      ]);
                    }}
                  >
                    <AddIcon fontSize="large" />
                  </IconButton>
                </Box>
              </Stack>
            </StyledCard>
          </Stack>
        </Container>
      )}
    </Formik>
  );
};

const StyledCard = styled(Card)(({ theme }) => ({
  padding: theme.spacing(4),
  display: "flex",
  flexDirection: "column",
  gap: theme.spacing(3),
  borderRadius: theme.shape.borderRadius,
}));

export default CreateProductForm;
