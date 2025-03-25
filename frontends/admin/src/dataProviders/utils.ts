import { Client, createClient } from "@hey-api/client-axios";

export function getBaseApiClients(baseUrl: string): Client {
  return createClient({
    baseURL: baseUrl,
  });
}
