import { HttpError, useForm } from "@refinedev/core";
import { FC } from "react";
import { Product, ProductInput } from "../../generated";
import { Formik } from "formik";
import { Box, Stack } from "@mui/material";
import { Create } from "@refinedev/mui";
import ProductFormGeneralSection from "./ProductFormGeneralSection";
import ProductMediasFormSection from "./ProductMediasFormSection";
import ProductVariantsFormSection from "./ProductVariantsFormSection";

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
      }}
      onSubmit={(data) => {
        form.onFinish(data);
      }}
    >
      {(formik) => (
        <Box
          component="form"
          onSubmit={formik.handleSubmit}
          onReset={formik.handleReset}
        >
          <Create
            resource="products"
            saveButtonProps={{
              type: "submit",
            }}
          >
            <Stack spacing={6}>
              <ProductFormGeneralSection
                value={{
                  title: formik.values.title,
                  description: formik.values.description,
                  tags: formik.values.tags,
                }}
                onSave={(d) => {
                  formik.setValues({ ...formik.values, ...d });
                }}
              />
              <ProductMediasFormSection
                mapFromMediaToT={(m, i) => ({
                  mediaId: m.id,
                  alt: "",
                  sortNumber: i,
                })}
                values={formik.values.medias ?? []}
                onChange={(m) => {
                  formik.setFieldValue("medias", m);
                }}
              />
              <ProductVariantsFormSection
                value={formik.values.variants}
                onSave={(d) => {
                  formik.setFieldValue("variants", d);
                }}
              />
            </Stack>
          </Create>
        </Box>
      )}
    </Formik>
  );
};

export default CreateProductForm;
