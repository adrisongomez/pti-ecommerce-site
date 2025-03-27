import { ComponentProps, FC } from "react";
import { ProductMediaInput } from "../../generated";
import { DropZone } from "../../components/dropZones/DropZone";

const ProductMediaDropzone: FC<
  Pick<
    ComponentProps<typeof DropZone<ProductMediaInput>>,
    "onChange" | "values" | "mapFromMediaToT"
  >
> = ({ onChange, values, mapFromMediaToT }) => {
  return (
    <DropZone<ProductMediaInput>
      values={values}
      onChange={(m) => {
        onChange(m);
      }}
      mapFromMediaToT={mapFromMediaToT}
    />
  );
};

export default ProductMediaDropzone;
