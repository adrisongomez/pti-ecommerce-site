/* eslint-disable */
import { BaseRecord, DataProvider, GetOneResponse } from "@refinedev/core";
import { getBaseApiClients, mapPaginationArgsToAPIArgs } from "./utils";

export default function ApiDataProvider(baseUrl: string): DataProvider {
  const client = getBaseApiClients(baseUrl);
  return {
    // @ts-ignore
    async getList({ pagination, resource }) {
      const response = await client.get<{
        data: unknown[];
        pageInfo: { totalResource: number };
      }>({
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
      const response = await client.get<BaseRecord>({
        url: `/${resource}/${id}`,
      });
      return { data: response.data };
    },
    // @ts-ignore
    async create({ resource, variables }) {
      const response = await client.post<unknown>({
        url: `/${resource}`,
        body: variables,
      });
      return { data: response.data };
    },
    // @ts-ignore
    async update({ resource, id, variables }) {
      const response = await client.put<unknown>({
        url: `/${resource}/${id}`,
        body: { payload: variables },
      });
      return { data: response.data };
    },
    // @ts-ignore
    async deleteOne({ id, resource }) {
      const response = await client.delete<boolean, true>({
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
