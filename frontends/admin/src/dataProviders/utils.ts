import { Client, createClient } from "@hey-api/client-axios";
import { Pagination } from "@refinedev/core";

export function getBaseApiClients(baseUrl: string): Client {
  return createClient({
    baseURL: baseUrl,
  });
}

export function mapPaginationArgsToAPIArgs(
  pagination: Pagination | undefined,
): {
  pageSize: number;
  after: number;
} {
  if (!pagination) {
    return { pageSize: 10, after: 0 };
  }

  return {
    pageSize: pagination?.pageSize ?? 10,
    after: ((pagination?.current ?? 0) - 1) * (pagination?.pageSize ?? 0),
  };
}
