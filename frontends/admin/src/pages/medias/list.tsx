import { FC, useMemo } from "react";
import { List, useDataGrid } from "@refinedev/mui";
import { DataGrid, GridColDef } from "@mui/x-data-grid";

const MediaList: FC = () => {
  const { dataGridProps } = useDataGrid({});
  const columns = useMemo<GridColDef[]>(
    () => [
      { field: "url", headerName: "Media", type: "string" },
      {
        field: "bucket",
        headerName: "Bucket",
        type: "string",
      },
      { field: "key", headerName: "Key", type: "string" },
    ],
    [],
  );
  return (
    <List>
      <DataGrid {...dataGridProps} columns={columns} />
    </List>
  );
};

export default MediaList;
