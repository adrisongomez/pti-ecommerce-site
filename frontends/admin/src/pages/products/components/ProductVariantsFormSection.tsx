import ProductStyledCard from "./ProductStyledCard";
import { Box, IconButton, Stack, Typography } from "@mui/material";
import ProductVariantEntry from "./ProductVariantEntry";
import AddIcon from "@mui/icons-material/Add";
import { ProductVariantCreateInput } from "../../../generated";

type ProductVariantsFormSectionProps<T extends ProductVariantCreateInput> = {
  value: T[];
  onSave(d: (T | ProductVariantCreateInput)[]): void;
};

const ProductVariantsFormSection = <T extends ProductVariantCreateInput>({
  value,
  onSave,
}: ProductVariantsFormSectionProps<T>) => {
  return (
    <ProductStyledCard raised>
      <Typography variant="h6" fontWeight="500">
        Variants
      </Typography>
      <Stack spacing={2}>
        {value.map((v, i) => (
          <ProductVariantEntry
            key={`${v.colorName}-${i}`}
            colorName={v.colorName}
            price={v.price}
            colorHex={v.colorHex}
            onSave={(d) => {
              onSave(value.map((v, idx) => (i === idx ? d : v)));
            }}
            onRemove={() => {
              if (value.length <= 1) {
                return;
              }
              onSave(value.filter((_, idx) => i !== idx));
            }}
          />
        ))}
        <Box display="flex" alignItems="center" justifyContent="center">
          <IconButton
            onClick={() => {
              onSave([
                ...value,
                { colorName: "", price: 0, colorHex: undefined },
              ]);
            }}
          >
            <AddIcon fontSize="large" />
          </IconButton>
        </Box>
      </Stack>
    </ProductStyledCard>
  );
};

export default ProductVariantsFormSection;
