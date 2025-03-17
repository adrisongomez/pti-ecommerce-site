import React from "react";

import type { Preview } from "@storybook/react";
import "../src/index.css";
import { IntlProvider } from "react-intl";

const preview: Preview = {
  parameters: {
    controls: {
      matchers: {
        color: /(background|color)$/i,
        date: /Date$/i,
      },
    },
  },
  decorators: [
    (Story) => (
      <IntlProvider locale="en">
        <Story />
      </IntlProvider>
    ),
  ],
};

export default preview;
