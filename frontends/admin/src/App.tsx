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
import ProductList from "./pages/products/list";
import CreateProductForm from "./pages/products/create";
import { APP_RESOURCES, APP_DATA_PROVIDER } from "./config/constants";
import EditProduct from "./pages/products/edit";
import { Container, Typography } from "@mui/material";
import UserList from "./pages/users/list";
import CreateUserPage from "./pages/users/create";
import EditUserPage from "./pages/users/edit";

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
                dataProvider={APP_DATA_PROVIDER}
                notificationProvider={useNotificationProvider}
                routerProvider={routerBindings}
                resources={APP_RESOURCES}
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
                      <ThemedLayoutV2
                        Title={() => <Typography>CMS</Typography>}
                        Header={() => <Header sticky />}
                      >
                        <Container maxWidth="lg" sx={{ width: "100%" }}>
                          <Outlet />
                        </Container>
                      </ThemedLayoutV2>
                    }
                  >
                    <Route
                      index
                      element={<NavigateToResource resource="products" />}
                    />
                    <Route path="/files">
                      <Route index element={<FileLists />} />
                    </Route>
                    <Route path="/users">
                      <Route index element={<UserList />} />
                      <Route
                        path="/users/create"
                        element={<CreateUserPage />}
                      />
                      <Route path="/users/:id" element={<EditUserPage />} />
                    </Route>
                    <Route path="/products">
                      <Route index element={<ProductList />} />
                      <Route
                        path="/products/create"
                        element={<CreateProductForm />}
                      />
                      <Route path="/products/:id" element={<EditProduct />} />
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
