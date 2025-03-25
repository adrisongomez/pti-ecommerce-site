import React from "react";

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
import MediaList from "./pages/files/list";
import FileApiProvider from "./dataProviders/FileApiProvider";

function App() {
  return (
    <BrowserRouter>
      <RefineKbarProvider>
        <ColorModeContextProvider>
          <CssBaseline />
          <GlobalStyles styles={{ html: { WebkitFontSmoothing: "auto" } }} />
          <RefineSnackbarProvider>
            <DevtoolsProvider>
              <Refine
                dataProvider={{
                  default: FileApiProvider("http://localhost:3030"),
                  files: FileApiProvider("http://localhost:3030"),
                }}
                notificationProvider={useNotificationProvider}
                routerProvider={routerBindings}
                resources={[
                  {
                    name: "files",
                    list: "/files",
                    show: "/files/:id",
                    create: "/files/create",
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
                      element={<NavigateToResource resource="blog_posts" />}
                    />
                    <Route path="/files">
                      <Route index element={<MediaList />} />
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
}

export default App;
