import { Box, IconButton, styled, Typography, useTheme } from "@mui/material";
import isPropValid from "@emotion/is-prop-valid";
import { useDropzone } from "react-dropzone";
import { Media, MediaInput, svcMediaCreate } from "../../generated";
import { useState } from "react";
import axios from "axios";
import SortableList, { SortableItem } from "react-easy-sort";
import DeleteIcon from "@mui/icons-material/Delete";
import Image from "../images/Images";
import AddIcon from "@mui/icons-material/Add";

type DropdownZoneProps<T> = {
  values: T[];
  onChange: (medias: T[]) => void | Promise<void>;
  mapFromMediaToT: (m: Media, i: number) => T;
};

export function DropZone<
  T extends { mediaId: number; sortNumber: number; url?: string | undefined },
>({ values, onChange, mapFromMediaToT }: DropdownZoneProps<T>): JSX.Element {
  const theme = useTheme();
  const [loadingFiles, setLoadingFiles] = useState<
    Record<string, Pick<Media, "url">>
  >({
    "": {
      url: "",
    },
  });
  const { getRootProps, getInputProps, isDragActive } = useDropzone({
    multiple: true,
    accept: {
      "image/*": [],
    },
    async onDrop(files) {
      const response = await Promise.all(
        files.map(async (value) => {
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
          setLoadingFiles((value) => ({
            ...value,
            [data.id]: {
              url: data.url,
            },
          }));
          await axios.put(uploadUrl, value, {
            headers: { "Content-Type": value.type },
          });
          return data;
        }),
      );
      onChange(
        [...values, ...response.map(mapFromMediaToT)].toSorted(
          (a, b) => a.sortNumber - b.sortNumber,
        ),
      );
    },
  });
  return (
    <DropzoneContainer $isDragging={isDragActive} {...getRootProps()}>
      <input type="file" {...getInputProps()} />
      {!values.length && (
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
      )}
      {!!values.length && (
        <SortableList
          className="list"
          style={{
            display: "flex",
            gap: 12,
            flexWrap: "wrap",
          }}
          draggedItemClassName="item"
          onSortEnd={(oldI, newI) => {
            const newMedias = values.map((v) => {
              if (![oldI, newI].includes(v.sortNumber)) {
                return v;
              }
              return {
                ...v,
                sortNumber: oldI === v.sortNumber ? newI : oldI,
              };
            });
            onChange(newMedias.toSorted((a, b) => a.sortNumber - b.sortNumber));
          }}
        >
          {values.map((v) => (
            <SortableItem key={v.mediaId}>
              <Box
                className="items"
                position="relative"
                onDragStart={(e) => {
                  e.preventDefault();
                }}
              >
                <IconButton
                  sx={{
                    position: "absolute",
                    top: 6,
                    right: 3,
                    zIndex: 10,
                  }}
                  onClick={(e) => {
                    e.stopPropagation();
                    onChange(
                      values
                        .filter((m) => v.mediaId !== m.mediaId)
                        .map((v, i) => ({
                          ...v,
                          sortNumber: i,
                        })),
                    );
                  }}
                >
                  <DeleteIcon fontSize="small" />
                </IconButton>
                <Image
                  sx={{ cursor: "pointer" }}
                  onDragStart={(e) => {
                    e.preventDefault();
                  }}
                  width={"166px"}
                  height={"208px"}
                  src={v.url ?? loadingFiles[v.mediaId]?.url}
                />
              </Box>
            </SortableItem>
          ))}
          <Box
            width="166px"
            height="208px"
            display="flex"
            alignItems="center"
            justifyContent="center"
            borderRadius={theme.shape.borderRadius}
            bgcolor={theme.palette.background.default}
          >
            <IconButton>
              <AddIcon fontSize="large" />
            </IconButton>
          </Box>
        </SortableList>
      )}
    </DropzoneContainer>
  );
}

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
