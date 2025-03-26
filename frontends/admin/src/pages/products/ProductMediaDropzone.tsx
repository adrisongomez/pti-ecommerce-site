import { FC, useState } from "react";
import axios from "axios";
import {
  Media,
  MediaInput,
  ProductMediaInput,
  svcMediaCreate,
} from "../../generated";
import isPropValid from "@emotion/is-prop-valid";
import { useDropzone } from "react-dropzone";
import { Box, styled, Typography } from "@mui/material";
import { Reorder } from "motion/react";
import Image from "../../components/images/Images";

type ProductMediaDropzoneProps = {
  medias: ProductMediaInput[];
  onChangeMedia: (medias: ProductMediaInput[]) => void | Promise<void>;
};

const ProductMediaDropzone: FC<ProductMediaDropzoneProps> = ({
  medias,
  onChangeMedia,
}) => {
  const [loadingFiles, setLoadingFiles] = useState<
    Record<string, Pick<Media, "id" | "url"> & { sortNumber: number }>
  >({});
  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    multiple: true,
    accept: {
      "image/*": [],
    },
    async onDrop(files) {
      const response = await Promise.all(
        files.map(async (value, i) => {
          const mediaInput: MediaInput = {
            bucket: "ecommerce-public",
            filename: value.name,
            key: `/product-medias/${value.name}`,
            mimeType: value.type,
            size: value.size,
          };
          const response = await svcMediaCreate<true>({
            throwOnError: true,
            body: mediaInput,
          });
          const { media: data, uploadUrl } = response.data;
          console.log(response);

          setLoadingFiles((value) => ({
            ...value,
            [data.id]: {
              id: data.id,
              url: data.url,
              sortNumber: i,
            },
          }));
          await axios.put(uploadUrl, value, {
            headers: { "Content-Type": value.type },
          });
          return data;
        }),
      );
      onChangeMedia(
        response.map((d, i) => ({
          mediaId: d.id,
          sortNumber: i,
          alt: d.filename,
        })),
      );
    },
  });
  return (
    <DropzoneContainer $isDragging={isDragActive} {...getRootProps()}>
      <input type="file" {...getInputProps()} />
      {!medias.length ? (
        <Box
          position="absolute"
          top="50%"
          left="50%"
          sx={{
            transform: "translate(-50%,-50%)",
          }}
        >
          <Typography color="textSecondary">Drap file or click over</Typography>
        </Box>
      ) : (
        <Reorder.Group axis="x" values={medias} onReorder={onChangeMedia}>
          {medias.map((v) => (
            <Reorder.Item value={v.mediaId}>
              <Image
                width="166"
                height="208"
                src={loadingFiles[v.mediaId]?.url}
              />
            </Reorder.Item>
          ))}
        </Reorder.Group>
      )}
    </DropzoneContainer>
  );
};

const DropzoneContainer = styled(Box, { shouldForwardProp: isPropValid })<{
  $isDragging: boolean;
}>(({ theme, $isDragging }) => ({
  gap: theme.spacing(2),
  flexWrap: "wrap",
  display: "flex",
  outline: "none",
  transition: "all 1s easa-in-out",
  position: "relative",
  cursor: "pointer",
  userSelect: "none",
  padding: theme.spacing(2),
  width: "100%",
  minHeight: 340,
  borderRadius: theme.shape.borderRadius * 4,
  border: `4px dotted ${$isDragging ? theme.palette.primary.dark : theme.palette.text.secondary}`,
}));

export default ProductMediaDropzone;
