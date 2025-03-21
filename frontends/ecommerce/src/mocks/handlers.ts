import { PRODUCTS } from "@/assets/data";
import { http, HttpResponse } from "msw";

export const handlers = [
  http.get("/products/", () => {
    return HttpResponse.json(PRODUCTS, { status: 200 });
  }),

  http.get("/products/:productId", ({ params }) => {
    const product = PRODUCTS.find((p) => p.id.toString() === params.productID);
    if (!product) {
      return HttpResponse.json(null, { status: 404 });
    }
    return HttpResponse.json(product, { status: 200 });
  }),
];
