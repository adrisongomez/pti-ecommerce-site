import { Edit } from "@refinedev/mui";
import { useFormik } from "formik";
import { FC } from "react";
import {
  Product,
  ProductMediaUpsertInput,
  ProductUpdateInput,
  ProductVariantUpsertInput,
} from "../../generated";
import { HttpError, useForm } from "@refinedev/core";
import { useUpdateEffect } from "react-use";
import { Stack } from "@mui/material";
import ProductFormGeneralSection from "./ProductFormGeneralSection";
import ProductMediasFormSection from "./ProductMediasFormSection";
import ProductVariantsFormSection from "./ProductVariantsFormSection";

type EditProductFormState = {
  productUpdate: ProductUpdateInput;
  medias: ProductMediaUpsertInput[];
  variants: ProductVariantUpsertInput[];
};

const EditProduct: FC<{ id: string }> = ({ id }) => {
  const { formLoading, query, onFinish } = useForm<
    Product,
    HttpError,
    ProductUpdateInput
  >({ id, resource: "products", action: "edit" });
  const originalProduct = query?.data;
  const formik = useFormik<EditProductFormState>({
    initialValues: {
      medias: originalProduct?.data?.medias ?? [],
      variants: originalProduct?.data?.variants ?? [],
      productUpdate: {
        handle: originalProduct?.data?.handle ?? "",
        title: originalProduct?.data?.title ?? "",
        description: originalProduct?.data?.description ?? "",
        tags: originalProduct?.data?.tags ?? [],
      },
    },
    async onSubmit(d) {
      const originalVariantIds = originalProduct?.data.variants.map(
        (v) => v.id,
      );
      const localMediaIds = d.medias.map((m) => m.mediaId);
      const localVariantIds = d.variants.map((v) => v.id);
      const oldMediaIds =
        originalProduct?.data.medias.map((v) => v.mediaId) ?? [];
      const deleteVariantIds =
        originalVariantIds?.filter(
          (id) => id && localVariantIds?.includes(id) === false,
        ) ?? [];
      // const variants = d.variants.map((v) => v.id);)
      const payload: ProductUpdateInput = {
        ...d.productUpdate,
        removeVariantIds: deleteVariantIds.filter(filterNull),
        medias: d.medias.filter((v) => oldMediaIds.includes(v.mediaId)),
        removeMediaIds:
          originalProduct?.data.medias
            .filter((v) => !localMediaIds.includes(v.mediaId))
            .map((v) => v.id) ?? [],
        variants: d.variants,
      };
      await onFinish(payload);
    },
  });
  useUpdateEffect(() => {
    if (originalProduct && !formik.dirty) {
      formik.setValues({
        medias: originalProduct?.data.medias,
        variants: originalProduct?.data.variants,
        productUpdate: {
          handle: originalProduct?.data.handle ?? "",
          title: originalProduct?.data.title ?? "",
          description: originalProduct?.data.description ?? "",
          tags: originalProduct?.data.tags ?? [],
        },
      });
    }
  }, [originalProduct]);
  return (
    <Edit
      isLoading={formLoading}
      resource="products"
      saveButtonProps={{ onClick: () => formik.submitForm() }}
    >
      <Stack spacing={6}>
        <ProductFormGeneralSection
          value={{
            title: formik.values.productUpdate.title,
            description: formik.values.productUpdate.description,
            tags: formik.values.productUpdate.tags,
          }}
          onSave={(d) => {
            formik.setFieldValue("productUpdate", {
              ...d,
            });
          }}
        />
        <ProductMediasFormSection
          mapFromMediaToT={(m, i) => ({
            mediaId: m.id,
            alt: "",
            sortNumber: i,
          })}
          values={formik.values.medias}
          onChange={(d) => {
            formik.setFieldValue("medias", d);
          }}
        />
        <ProductVariantsFormSection
          value={formik.values.variants}
          onSave={(d) => {
            formik.setFieldValue("variants", d);
          }}
        />
      </Stack>
    </Edit>
  );
};

function filterNull<T>(
  v: T | null | undefined,
): v is Exclude<T, null | undefined> {
  return Boolean(v);
}

export default EditProduct;
