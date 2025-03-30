import React from "react";

import type { Preview } from "@storybook/react";
import "../src/index.css";
import { IntlProvider } from "react-intl";
import {
  RouterProvider,
  createMemoryHistory,
  createRootRoute,
  createRouter,
} from "@tanstack/react-router";
import { initialize, mswLoader } from "msw-storybook-addon";

initialize();

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
  loaders: [mswLoader],
  decorators: [
    (Story) => {
      return (
        <RouterProvider
          router={createRouter({
            history: createMemoryHistory(),
            routeTree: createRootRoute({
              component: Story,
            }),
          })}
        />
      );
    },
    (Story) => (
      <IntlProvider locale="en">
        <Story />
      </IntlProvider>
    ),
  ],
};

export default preview;
