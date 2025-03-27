import { Stack, TextField } from "@mui/material";
import DragIndicatorIcon from "@mui/icons-material/DragIndicator";
import { FC } from "react";

const ProductVariantEntry: FC = () => {
  return (
    <Stack flexDirection="row" alignItems="center" useFlexGap spacing={3}>
      <DragIndicatorIcon fontSize="small" />
      <TextField
        variant="filled"
        fullWidth
        label="Option name"
        name="colorName"
      />
      <input name="colorHex" type="color" />
    </Stack>
  );
};

export default ProductVariantEntry;
