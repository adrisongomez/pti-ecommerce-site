import { ComponentProps, FC } from "react";
import ProductStyledCard from "./ProductStyledCard";
import { Typography } from "@mui/material";
import ProductMediaDropzone from "./ProductMediaDropzone";

const ProductMediasFormSection: FC<
  ComponentProps<typeof ProductMediaDropzone>
> = (props) => (
  <ProductStyledCard raised variant="elevation">
    <Typography variant="h6" fontWeight="500">
      Medias
    </Typography>
    <ProductMediaDropzone {...props} />
  </ProductStyledCard>
);

export default ProductMediasFormSection;
