import { FC } from "react";
import { ProductMediaInput } from "../../generated";
import { useField } from "formik";
import { DropZone } from "../../components/dropZones/DropZone";

const ProductMediaDropzone: FC = () => {
  const fields = useField<ProductMediaInput[]>("medias");
  const { value } = fields[0];
  const { setValue: onChange } = fields[2];
  return (
    <DropZone<ProductMediaInput>
      values={value}
      onChange={(m) => {
        onChange(m);
      }}
      mapFromMediaToT={(m, i) => ({
        mediaId: m.id,
        alt: "",
        sortNumber: i,
      })}
    />
  );
};

export default ProductMediaDropzone;
