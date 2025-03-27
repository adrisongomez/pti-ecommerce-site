import { ResourceProps } from "@refinedev/core";
import StoreIcon from "@mui/icons-material/Store";
import Inventory2Icon from "@mui/icons-material/Inventory2";
import ApiDataProvider from "../dataProviders/ApiDataProvider";
import PersonIcon from "@mui/icons-material/Person";

export const APP_RESOURCES: ResourceProps[] = [
  {
    name: "products",
    list: "/products",
    create: "/products/create",
    edit: "/products/:id",
    meta: {
      icon: <StoreIcon />,
      canDelete: true,
    },
  },
  {
    name: "users",
    list: "/users",
    create: "/users/create",
    edit: "/users/:id",
    meta: {
      icon: <PersonIcon />,
      canDelete: true,
    },
  },
  {
    name: "medias",
    list: "/files/",
    meta: {
      icon: <Inventory2Icon />,
      canDelete: true,
      label: "Files",
    },
  },
];

export const APP_DATA_PROVIDER = ApiDataProvider("http://localhost:3030/api/");
