import { FC } from "react";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import { IntlProvider } from "react-intl";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { Provider } from "react-redux";

// Import the generated route tree
import { routeTree } from "@/routeTree.gen";

import "./fetchClient";
import { store } from "../redux/stores";

const queryClient = new QueryClient({});

// Create a new router instance
const router = createRouter({
  routeTree,
  context: {
    queryClient,
  },
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const RootProvider: FC = () => {
  return (
    <Provider store={store}>
      <QueryClientProvider client={queryClient}>
        <IntlProvider locale="en" defaultLocale="en">
          <RouterProvider router={router} />
        </IntlProvider>
      </QueryClientProvider>
    </Provider>
  );
};

export default RootProvider;
