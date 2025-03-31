/* eslint-disable */
import { BaseRecord, DataProvider, GetOneResponse } from "@refinedev/core";
import { getBaseApiClients, mapPaginationArgsToAPIArgs } from "./utils";
import { getCreds } from "../utils/auth";

export default function ApiDataProvider(baseUrl: string): DataProvider {
  const client = getBaseApiClients(baseUrl);
  return {
    // @ts-ignore
    async getList({ pagination, resource }) {
      const creds = getCreds();
      const response = await client.get<{
        data: unknown[];
        pageInfo: { totalResource: number };
      }>({
        headers: creds
          ? {
              Authorization: `Bearer ${creds?.accessToken}`,
            }
          : undefined,
        url: `/${resource}`,
        query: {
          ...mapPaginationArgsToAPIArgs(pagination),
        },
      });
      return {
        data: response.data?.data,
        total: response.data?.pageInfo.totalResource,
      };
    },
    // @ts-ignore
    async getOne({ id, resource }): Promise<GetOneResponse<Media>> {
      const creds = getCreds();
      const response = await client.get<BaseRecord>({
        headers: creds
          ? {
              Authorization: `Bearer ${creds?.accessToken}`,
            }
          : undefined,
        url: `/${resource}/${id}`,
      });
      return { data: response.data };
    },
    // @ts-ignore
    async create({ resource, variables }) {
      const creds = getCreds();
      const response = await client.post<unknown>({
        headers: creds
          ? {
              Authorization: `Bearer ${creds?.accessToken}`,
            }
          : undefined,
        url: `/${resource}`,
        body: variables,
      });
      return { data: response.data };
    },
    // @ts-ignore
    async update({ resource, id, variables }) {
      const creds = getCreds();
      const response = await client.put<unknown>({
        url: `/${resource}/${id}`,
        headers: creds
          ? {
              Authorization: `Bearer ${creds?.accessToken}`,
            }
          : undefined,
        body: variables,
      });
      return { data: response.data };
    },
    // @ts-ignore
    async deleteOne({ id, resource }) {
      const creds = getCreds();
      const response = await client.delete<boolean, true>({
        headers: creds
          ? {
              Authorization: `Bearer ${creds?.accessToken}`,
            }
          : undefined,
        url: `/${resource}/${id}`,
      });
      if (response.error) {
        throw response.error;
      }
      return {
        data: {
          id,
        },
      };
    },
    getApiUrl: () => baseUrl,
  };
}
