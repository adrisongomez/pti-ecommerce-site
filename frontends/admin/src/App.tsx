import * as React from "react";
import { Refine } from "@refinedev/core";
import { DevtoolsPanel, DevtoolsProvider } from "@refinedev/devtools";
import { RefineKbar, RefineKbarProvider } from "@refinedev/kbar";

import {
  ErrorComponent,
  useNotificationProvider,
  RefineSnackbarProvider,
  ThemedLayoutV2,
} from "@refinedev/mui";

import CssBaseline from "@mui/material/CssBaseline";
import GlobalStyles from "@mui/material/GlobalStyles";
import { BrowserRouter, Route, Routes, Outlet } from "react-router";
import routerBindings, {
  NavigateToResource,
  UnsavedChangesNotifier,
  DocumentTitleHandler,
} from "@refinedev/react-router";
import { ColorModeContextProvider } from "./contexts/color-mode";
import { Header } from "./components/header";
import FileLists from "./pages/files/list";
import ApiDataProvider from "./dataProviders/ApiDataProvider";
import ProductList from "./pages/products/list";
import CreateProductForm from "./pages/products/create";

const dataProvider = ApiDataProvider("http://localhost:3030/api/");

const App: React.FC = () => {
  return (
    <BrowserRouter>
      <RefineKbarProvider>
        <ColorModeContextProvider>
          <CssBaseline />
          <GlobalStyles styles={{ html: { WebkitFontSmoothing: "auto" } }} />
          <RefineSnackbarProvider>
            <DevtoolsProvider>
              <Refine
                dataProvider={dataProvider}
                notificationProvider={useNotificationProvider}
                routerProvider={routerBindings}
                resources={[
                  {
                    name: "medias",
                    list: "/files/",
                    meta: {
                      canDelete: true,
                      label: "Files",
                    },
                  },
                  {
                    name: "products",
                    list: "/products",
                    show: "/products/:id",
                    create: "/products/create",
                    edit: "/products/:id/edit",
                    meta: {
                      canDelete: true,
                    },
                  },
                ]}
                options={{
                  syncWithLocation: true,
                  warnWhenUnsavedChanges: true,
                  useNewQueryKeys: true,
                  projectId: "vxqfwr-LxUveD-Vbu6Lx",
                }}
              >
                <Routes>
                  <Route
                    element={
                      <ThemedLayoutV2 Header={() => <Header sticky />}>
                        <Outlet />
                      </ThemedLayoutV2>
                    }
                  >
                    <Route
                      index
                      element={<NavigateToResource resource="files" />}
                    />
                    <Route path="/files">
                      <Route index element={<FileLists />} />
                    </Route>
                    <Route path="/products">
                      <Route index element={<ProductList />} />
                      <Route
                        path="/products/create"
                        element={<CreateProductForm />}
                      />
                    </Route>
                    <Route path="*" element={<ErrorComponent />} />
                  </Route>
                </Routes>

                <RefineKbar />
                <UnsavedChangesNotifier />
                <DocumentTitleHandler />
              </Refine>
              <DevtoolsPanel />
            </DevtoolsProvider>
          </RefineSnackbarProvider>
        </ColorModeContextProvider>
      </RefineKbarProvider>
    </BrowserRouter>
  );
};

export default App;
