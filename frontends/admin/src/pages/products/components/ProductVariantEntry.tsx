import {
  Card,
  Grid2,
  IconButton,
  InputAdornment,
  TextField,
} from "@mui/material";
import DeleteIcon from "@mui/icons-material/Delete";
import { FC, MouseEventHandler } from "react";
import { useFormik } from "formik";
import { ProductVariantCreateInput } from "../../../generated";
import { useDebounce } from "react-use";

type ProductVariantEntryProps = {
  colorName: string;
  colorHex?: string;
  price: number;
  onRemove: MouseEventHandler<HTMLButtonElement>;
  onSave(v: ProductVariantCreateInput): void;
};

const ProductVariantEntry: FC<ProductVariantEntryProps> = ({
  colorName,
  colorHex,
  price,
  onSave,
  onRemove,
}) => {
  const formik = useFormik<ProductVariantCreateInput>({
    initialValues: {
      colorName,
      colorHex,
      price,
    },
    onSubmit(d) {
      onSave(d);
    },
  });

  useDebounce(
    () => {
      formik.submitForm();
    },
    500,
    [formik.values],
  );

  return (
    <Card raised elevation={3} sx={{ p: 2 }} variant="elevation">
      <Grid2 container spacing={2}>
        <Grid2 size={{ xs: 11, md: 6 }}>
          <TextField
            fullWidth
            variant="filled"
            label="Option name"
            name="colorName"
            value={formik.values.colorName}
            onChange={formik.handleChange}
            slotProps={{
              inputLabel: {
                shrink: true,
              },
            }}
          />
        </Grid2>
        <Grid2 size={{ xs: 12, md: 4 }}>
          <TextField
            fullWidth
            variant="filled"
            value={formik.values.price}
            onChange={formik.handleChange}
            slotProps={{
              input: {
                startAdornment: (
                  <InputAdornment position="start">$</InputAdornment>
                ),
              },
            }}
            type="number"
            label="Option price"
            name="price"
          />
        </Grid2>
        <Grid2
          size={{ xs: 12, md: 1 }}
          sx={{ display: "flex", alignItems: "center" }}
        >
          <input
            value={formik.values.colorHex ?? 0}
            style={{ flex: 1 }}
            name="colorHex"
            type="color"
            onChange={formik.handleChange}
          />
        </Grid2>
        <Grid2 size={1} sx={{ display: "flex", alignItems: "center" }}>
          <IconButton onClick={onRemove}>
            <DeleteIcon />
          </IconButton>
        </Grid2>
      </Grid2>
    </Card>
  );
};

export default ProductVariantEntry;
