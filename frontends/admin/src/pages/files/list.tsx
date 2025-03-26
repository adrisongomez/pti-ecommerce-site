import { FC, Fragment, useMemo } from "react";
import { DeleteButton, List, useDataGrid } from "@refinedev/mui";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { Media } from "../../generated";
import ImageIcon from "@mui/icons-material/Image";
import MovieIcon from "@mui/icons-material/Movie";
import BrowserNotSupportedIcon from "@mui/icons-material/BrowserNotSupported";
import { IconButton } from "@mui/material";
import VisibilityIcon from "@mui/icons-material/Visibility";

const FileLists: FC = () => {
  const { dataGridProps } = useDataGrid({
    resource: "medias",
    filters: {
      mode: "off",
    },
    sorters: {
      mode: "off",
    },
  });
  const columns = useMemo<GridColDef<Media>[]>(
    () => [
      {
        field: "mediaType",
        headerName: "",
        type: "string",
        align: "center",
        headerAlign: "center",
        renderCell: ({ row }) => (
          <Fragment>
            {row.mediaType === "IMAGE" && <ImageIcon fontSize="large" />}
            {row.mediaType === "VIDEO" && <MovieIcon fontSize="large" />}
            {row.mediaType === "UNKNWON" && (
              <BrowserNotSupportedIcon fontSize="large" />
            )}
          </Fragment>
        ),
      },
      {
        field: "bucket",
        headerName: "Bucket",
        type: "string",
      },
      { field: "key", headerName: "Key", type: "string", flex: 1 },
      { field: "mimeType", headerName: "File type", type: "string", flex: 1 },
      {
        type: "actions",
        field: "actions",
        headerName: "",
        align: "center",
        headerAlign: "center",
        sortable: false,
        renderCell: ({ row }) => {
          return (
            <Fragment>
              <DeleteButton hideText recordItemId={row.id} />
              <IconButton
                color="primary"
                onClick={() => {
                  window.open(row.url, "__blank");
                }}
              >
                <VisibilityIcon fontSize="small" />
              </IconButton>
            </Fragment>
          );
        },
      },
    ],
    [],
  );
  return (
    <List canCreate={false}>
      <DataGrid {...dataGridProps} columns={columns} />
    </List>
  );
};

export default FileLists;
