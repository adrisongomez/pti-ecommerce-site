import { DataGrid, GridColDef } from "@mui/x-data-grid";
import { DeleteButton, EditButton, List, useDataGrid } from "@refinedev/mui";
import { FC, Fragment } from "react";
import { User } from "../../generated";
import { filterNull } from "../../utils/utils";
import { format } from "date-fns";

const USER_LIST_COLUMNS: GridColDef<User>[] = [
  {
    field: "id",
    type: "number",
    headerName: "Id",
  },
  {
    field: "name",
    flex: 1,
    type: "string",
    headerName: "Name",
    valueGetter: (_, row) =>
      [row.firstName, row.lastName].filter(filterNull).join(" "),
  },
  {
    field: "email",
    flex: 1,
    headerName: "Email",
    type: "string",
  },
  {
    field: "role",
    headerName: "Type",
    valueGetter: (_, row) =>
      row.role[0].concat(row.role.slice(1).toLowerCase()),
  },
  {
    field: "Created",
    flex: 1,
    headerName: "created_at",
    type: "dateTime",
    valueFormatter: (_, row) => format(row.createdAt, "MMM dd, yyyy"),
  },
  {
    field: "Last Updated",
    flex: 1,
    headerName: "created_at",
    type: "dateTime",
    valueFormatter: (_, row) =>
      row.updatedAt ? format(row.updatedAt, "MMM dd, yyyy") : null,
  },
  {
    field: "actions",
    type: "actions",
    renderCell: ({ row }) => {
      return (
        <Fragment>
          <EditButton hideText recordItemId={row.id} />
          <DeleteButton hideText recordItemId={row.id} />
        </Fragment>
      );
    },
  },
];

const UserList: FC = () => {
  const { dataGridProps } = useDataGrid<User>({
    resource: "users",
  });
  return (
    <List canCreate>
      <DataGrid {...dataGridProps} columns={USER_LIST_COLUMNS} />
    </List>
  );
};

export default UserList;
