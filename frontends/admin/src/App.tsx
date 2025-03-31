import * as React from "react";
import { Authenticated, AuthPage, Refine } from "@refinedev/core";
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
  CatchAllNavigate,
} from "@refinedev/react-router";
import { ColorModeContextProvider } from "./contexts/color-mode";
import { Header } from "./components/header";
import FileLists from "./pages/files/list";
import ProductList from "./pages/products/list";
import CreateProductForm from "./pages/products/create";
import {
  APP_RESOURCES,
  APP_DATA_PROVIDER,
  APP_AUTH_PROVIDER,
} from "./config/constants";
import EditProduct from "./pages/products/edit";
import { Card, Container, Typography } from "@mui/material";
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
                authProvider={APP_AUTH_PROVIDER}
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
                    path="/login"
                    element={
                      <AuthPage
                        renderContent={(content) => (
                          <Container
                            maxWidth="sm"
                            sx={{
                              display: "flex",
                              justifyContent: "center",
                              alignItems: "center",
                              minHeight: "100vh",
                            }}
                          >
                            <Card
                              raised
                              sx={{ p: 2, borderRadius: 3, width: "100%" }}
                            >
                              {content}
                            </Card>
                          </Container>
                        )}
                        type="login"
                        rememberMe={false}
                      />
                    }
                  />
                  <Route
                    element={
                      <Authenticated
                        key="catch-all-auth"
                        fallback={<CatchAllNavigate to="/login" />}
                      >
                        <ThemedLayoutV2>
                          <Container maxWidth="lg" sx={{ width: "100%" }}>
                            <Outlet />
                          </Container>
                        </ThemedLayoutV2>
                      </Authenticated>
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
