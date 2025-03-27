import { FC } from "react";
import ProductStyledCard from "./ProductStyledCard";
import {
  Autocomplete,
  FormLabel,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import MDEditor from "@uiw/react-md-editor";
import { Product } from "../../../generated";

type ProductGeneralField = Pick<Product, "title" | "description" | "tags">;

type ProductFormGeneralSectionProps = {
  value: ProductGeneralField;
  onSave(d: ProductGeneralField): void;
};

const ProductFormGeneralSection: FC<ProductFormGeneralSectionProps> = ({
  value,
  onSave,
}) => {
  return (
    <ProductStyledCard raised variant="elevation">
      <Typography variant="h6" fontWeight="500">
        General
      </Typography>
      <TextField
        fullWidth
        required
        label="Title"
        variant="filled"
        value={value.title}
        onChange={(e) => {
          const text = e.currentTarget.value;
          onSave({ ...value, title: text ?? "" });
        }}
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
          console.log(v);
          onSave({ ...value, tags: v });
        }}
        value={value.tags ?? []}
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
          value={value.description}
          onChange={(v) => onSave({ ...value, description: v ?? "" })}
        />
      </Stack>
    </ProductStyledCard>
  );
};

export default ProductFormGeneralSection;
