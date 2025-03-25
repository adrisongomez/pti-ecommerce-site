/* eslint-disable */
import { DataProvider } from "@refinedev/core";
import { svcMediaList } from "../generated";
import { getBaseApiClients } from "./utils";

export default function FileApiProvider(baseUrl: string): DataProvider {
  const client = getBaseApiClients(baseUrl);
  return {
    // @ts-ignore
    async getList({ pagination }) {
      const response = await svcMediaList({
        client,
        query: {
          pageSize: pagination?.pageSize,
          after: ((pagination?.current ?? 0) - 1) * (pagination?.pageSize ?? 0),
        },
      });
      if (response.error) {
        throw response.error;
      }
      return {
        data: response.data.data,
        total: response.data.pageInfo.totalResource,
      };
    },
    getOne() {
      throw new Error("NotImplemented");
    },
    create() {
      throw new Error("NotImplemented");
    },
    update() {
      throw new Error("NotImplemented");
    },
    deleteOne() {
      throw new Error("NotImplemented");
    },
    getApiUrl: () => baseUrl,
  };
}
