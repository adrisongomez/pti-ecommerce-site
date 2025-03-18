import { FC, Fragment } from "react";
import { RouterProvider, createRouter } from "@tanstack/react-router";
import { IntlProvider } from "react-intl";
// Import the generated route tree
import { routeTree } from "@/routeTree.gen";

// Create a new router instance
const router = createRouter({
  routeTree,
  basepath:
    process.env.NODE_ENV !== "development"
      ? "/programming-the-internet-tarea-2/"
      : undefined, // this is just for enable github pages
});

// Register the router instance for type safety
declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const RootProvider: FC = () => {
  return (
    <Fragment>
      <IntlProvider locale="en" defaultLocale="en">
        <RouterProvider router={router} />
      </IntlProvider>
    </Fragment>
  );
};

export default RootProvider;
