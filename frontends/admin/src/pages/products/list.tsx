import { Avatar } from "@mui/material";
import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { DeleteButton, EditButton, List, useDataGrid } from "@refinedev/mui";
import { FC, Fragment, useMemo } from "react";
import { Product } from "../../generated";
import ImageIcon from "@mui/icons-material/Image";
import Image from "../../components/images/Images";

const ProductList: FC = () => {
  const { dataGridProps } = useDataGrid({ resource: "products" });
  const columns = useMemo<GridColDef<Product>[]>(
    () => [
      {
        headerName: "",
        type: "custom",
        display: "flex",
        sortable: false,
        field: "medias",
        align: "center",
        renderCell: ({ row }) => {
          const media = row.medias.at(0);
          if (!media) {
            return (
              <Avatar variant="rounded" sx={{ p: 2 }}>
                <ImageIcon />
              </Avatar>
            );
          }
          return (
            <Image
              width="44px"
              height="44px"
              sx={{ borderRadius: 1 }}
              src={media?.url}
              alt={media?.alt}
            />
          );
        },
      },
      {
        headerName: "Title",
        type: "string",
        field: "title",
        flex: 1,
      },
      {
        headerAlign: "center",
        align: "center",
        headerName: "Status",
        type: "string",
        field: "status",
        width: 244,
      },
      {
        headerAlign: "center",
        headerName: "Created",
        align: "center",
        type: "dateTime",
        field: "createdAt",
        width: 244,
        valueGetter: (_, r) => new Date(r.createdAt),
      },
      {
        width: 244,
        headerAlign: "center",
        align: "center",
        headerName: "Updated",
        type: "dateTime",
        valueGetter: (_, r) => (r.updatedAt ? new Date(r.updatedAt) : null),
        field: "updatedAt",
      },
      {
        field: "actions",
        type: "actions",
        renderCell: ({ row }) => (
          <Fragment>
            <EditButton hideText recordItemId={row.id} />
            <DeleteButton hideText recordItemId={row.id} />
          </Fragment>
        ),
      },
    ],
    [],
  );
  return (
    <List canCreate>
      <DataGrid {...dataGridProps} rowSelection columns={columns} />
    </List>
  );
};

export default ProductList;
