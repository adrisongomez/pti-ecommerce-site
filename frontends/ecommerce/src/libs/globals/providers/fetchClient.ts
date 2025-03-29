/**
 * @description we neet define the OpenAPI client and set the baseUrl
 *
 * @todo Implement the client generation and hook it with tanstack
 * @see https://github.com/7nohe/openapi-react-query-codegen/blob/main/examples/tanstack-router-app/src/fetchClient.ts
 */

import { OpenAPI } from "../generated/requests";
import { getCreds } from "../utilities/auth";

OpenAPI.BASE = "http://localhost:3030";
OpenAPI.TOKEN = async () => {
  const creds = getCreds();
  if (!creds) {
    return "";
  }
  return creds.accessToken;
};
