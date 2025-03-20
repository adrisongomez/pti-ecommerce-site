import { FC } from "react";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import { IntlProvider } from "react-intl";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

// Import the generated route tree
import { routeTree } from "@/routeTree.gen";

const queryClient = new QueryClient({});

// Create a new router instance
const router = createRouter({
  routeTree,
  basepath:
    process.env.NODE_ENV !== "development"
      ? "/programming-the-internet-tarea-2/"
      : undefined, // this is just for enable github pages
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
    <QueryClientProvider client={queryClient}>
      <IntlProvider locale="en" defaultLocale="en">
        <RouterProvider router={router} />
      </IntlProvider>
    </QueryClientProvider>
  );
};

export default RootProvider;
