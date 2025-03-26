// This file is auto-generated by @hey-api/openapi-ts

export type AddMediaRequestBody = {
  payload?: ProductMediaInput;
};

export type AddVariantRequestBody = {
  payload?: ProductVariantCreateInput;
};

export type CreateMediaResponse = {
  media: Media;
  uploadUrl: string;
};

export type _Error = {
  /**
   * Is the error a server-side fault?
   */
  fault: boolean;
  /**
   * ID is a unique identifier for this particular occurrence of the problem.
   */
  id: string;
  /**
   * Message is a human-readable explanation specific to this occurrence of the problem.
   */
  message: string;
  /**
   * Name is the name of this class of errors.
   */
  name: string;
  /**
   * Is the error temporary?
   */
  temporary: boolean;
  /**
   * Is the error a timeout?
   */
  timeout: boolean;
};

export type HealthcheckResponse = {
  status?: string;
};

export type Media = {
  bucket: string;
  /**
   * Datetime
   */
  createdAt: string;
  filename: string;
  /**
   * Key ID
   */
  id: number;
  key: string;
  /**
   * Type of the media
   */
  mediaType: "IMAGE" | "VIDEO" | "UNKNWON";
  mimeType: string;
  size: number;
  /**
   * Datetime
   */
  updatedAt?: string;
  url: string;
};

export type MediaInput = {
  bucket: string;
  filename: string;
  key: string;
  mimeType: string;
  size: number;
};

/**
 * Paginated results
 */
export type MediaList = {
  /**
   * Data
   */
  data: Array<Media>;
  pageInfo: PageInfo;
};

/**
 * Pagination information
 */
export type PageInfo = {
  /**
   * The ending cursor for pagination
   */
  endCursor: number;
  /**
   * Indicates if there are more results available
   */
  hasMore: boolean;
  /**
   * The starting cursor for pagination
   */
  startCursor: number;
  /**
   * Total number of resources available
   */
  totalResource: number;
};

/**
 * Product information
 */
export type Product = {
  /**
   * Datetime
   */
  createdAt: string;
  /**
   * Product description
   */
  description: string;
  /**
   * Handle
   */
  handle: string;
  /**
   * Key ID
   */
  id: number;
  medias: Array<ProductMedia>;
  /**
   * Define the status of product on the site
   */
  status: "ACTIVE" | "DRAFT";
  /**
   * Product tags
   */
  tags?: Array<string>;
  /**
   * Title
   */
  title: string;
  /**
   * Datetime
   */
  updatedAt?: string;
  variants: Array<ProductVariant>;
  vendor?: Vendor;
};

export type ProductInput = {
  /**
   * Product description
   */
  description: string;
  /**
   * Last part of the url which use to idepntify the user
   */
  handle?: string;
  medias?: Array<ProductMediaInput>;
  /**
   * Define the status of product on the site
   */
  status?: "ACTIVE" | "DRAFT";
  /**
   * Product tags
   */
  tags: Array<string>;
  /**
   * Title's product
   */
  title: string;
  /**
   * Product variants
   */
  variants: Array<ProductVariantCreateInput>;
  /**
   * Vendor's product
   */
  vendorId: number;
};

/**
 * Instance of media in a product
 */
export type ProductMedia = {
  /**
   * Alt text that would show in case the image does not render
   */
  alt?: string;
  /**
   * Datetime
   */
  createdAt: string;
  /**
   * Key ID
   */
  id: number;
  /**
   * ID of the media record where the resource has being upload
   */
  mediaId: number;
  /**
   * Type of the media
   */
  mediaType: "IMAGE" | "VIDEO" | "UNKNWON";
  /**
   * Position on the images of the product
   */
  sortNumber: number;
  /**
   * Datetime
   */
  updatedAt?: string;
  /**
   * URL to the media
   */
  url: string;
};

export type ProductMediaInput = {
  /**
   * Alt text that would show in case the image does not render
   */
  alt?: string;
  /**
   * ID of the media record where the resource has being upload
   */
  mediaId: number;
  /**
   * Position on the images of the product
   */
  sortNumber: number;
};

export type ProductUpdateInput = {
  /**
   * Product description
   */
  description: string;
  /**
   * Last part of the url which use to idepntify the user
   */
  handle?: string;
  /**
   * Define the status of product on the site
   */
  status?: "ACTIVE" | "DRAFT";
  /**
   * Product tags
   */
  tags: Array<string>;
  /**
   * Title's product
   */
  title: string;
  /**
   * Vendor's product
   */
  vendorId: number;
};

/**
 * Definition of product variants
 */
export type ProductVariant = {
  /**
   * Color in HEX value that would be used on the variant picker
   */
  colorHex?: string;
  /**
   * Color variant option
   */
  colorName: string;
  /**
   * Datetime
   */
  createdAt: string;
  /**
   * ProductMedia which would be focus when a variant is picked by the user
   */
  featureMediaId?: number;
  /**
   * Key ID
   */
  id: number;
  /**
   * Price on cents
   */
  price: number;
  productId: string;
  /**
   * Datetime
   */
  updatedAt?: string;
};

export type ProductVariantCreateInput = {
  /**
   * Color in HEX value that would be used on the variant picker
   */
  colorHex?: string;
  /**
   * Color variant option
   */
  colorName: string;
  /**
   * Price on cents
   */
  price: number;
};

/**
 * Paginated results
 */
export type ProductsList = {
  /**
   * Data
   */
  data: Array<Product>;
  pageInfo: PageInfo;
};

export type TypeFooter = {
  /**
   * Datetime
   */
  createdAt: string;
  /**
   * Datetime
   */
  updatedAt?: string;
};

export type UpdateProductByIdRequestBody = {
  payload?: ProductUpdateInput;
};

export type Vendor = {
  /**
   * Datetime
   */
  createdAt: string;
  /**
   * Key ID
   */
  id?: number;
  name: string;
  /**
   * Datetime
   */
  updatedAt?: string;
};

export type VendorInput = {
  name: string;
};

/**
 * Paginated results
 */
export type VendorList = {
  /**
   * Data
   */
  data: Array<Vendor>;
  pageInfo: PageInfo;
};

export type SvcHealthcheckCheckData = {
  body?: never;
  path?: never;
  query?: never;
  url: "/api/healthcheck";
};

export type SvcHealthcheckCheckResponses = {
  /**
   * OK response.
   */
  200: HealthcheckResponse;
};

export type SvcHealthcheckCheckResponse =
  SvcHealthcheckCheckResponses[keyof SvcHealthcheckCheckResponses];

export type SvcMediaListData = {
  body?: never;
  path?: never;
  query?: {
    /**
     * Record per page
     */
    pageSize?: number;
    /**
     * Start listing after this resource
     */
    after?: number;
    /**
     * S3 bucket where data is store
     */
    bucket?: string;
  };
  url: "/api/medias";
};

export type SvcMediaListErrors = {
  /**
   * BadRequest: Bad Gateway response.
   */
  502: _Error;
};

export type SvcMediaListError = SvcMediaListErrors[keyof SvcMediaListErrors];

export type SvcMediaListResponses = {
  /**
   * OK response.
   */
  200: MediaList;
};

export type SvcMediaListResponse =
  SvcMediaListResponses[keyof SvcMediaListResponses];

export type SvcMediaCreateData = {
  body: MediaInput;
  path?: never;
  query?: never;
  url: "/api/medias";
};

export type SvcMediaCreateErrors = {
  /**
   * BadRequest: Bad Request response.
   */
  400: _Error;
};

export type SvcMediaCreateError =
  SvcMediaCreateErrors[keyof SvcMediaCreateErrors];

export type SvcMediaCreateResponses = {
  /**
   * Created response.
   */
  201: CreateMediaResponse;
};

export type SvcMediaCreateResponse =
  SvcMediaCreateResponses[keyof SvcMediaCreateResponses];

export type SvcMediaDeleteByIdData = {
  body?: never;
  path: {
    mediaId: number;
  };
  query?: never;
  url: "/api/medias/{mediaId}";
};

export type SvcMediaDeleteByIdErrors = {
  /**
   * NotFound: Not Found response.
   */
  404: _Error;
};

export type SvcMediaDeleteByIdError =
  SvcMediaDeleteByIdErrors[keyof SvcMediaDeleteByIdErrors];

export type SvcMediaDeleteByIdResponses = {
  /**
   * Created response.
   */
  201: boolean;
};

export type SvcMediaDeleteByIdResponse =
  SvcMediaDeleteByIdResponses[keyof SvcMediaDeleteByIdResponses];

export type SvcMediaGetByIdData = {
  body?: never;
  path: {
    mediaId: number;
  };
  query?: never;
  url: "/api/medias/{mediaId}";
};

export type SvcMediaGetByIdErrors = {
  /**
   * NotFound: Not Found response.
   */
  404: _Error;
};

export type SvcMediaGetByIdError =
  SvcMediaGetByIdErrors[keyof SvcMediaGetByIdErrors];

export type SvcMediaGetByIdResponses = {
  /**
   * OK response.
   */
  200: Media;
};

export type SvcMediaGetByIdResponse =
  SvcMediaGetByIdResponses[keyof SvcMediaGetByIdResponses];

export type OpenapiApiOpenapiJsonData = {
  body?: never;
  path?: never;
  query?: never;
  url: "/api/openapi.json";
};

export type OpenapiApiOpenapiJsonResponses = {
  /**
   * File downloaded
   */
  200: unknown;
};

export type SvcProductsListProductData = {
  body?: never;
  path?: never;
  query?: {
    /**
     * Record per page
     */
    pageSize?: number;
    /**
     * Start listing after this resource
     */
    after?: number;
  };
  url: "/api/products";
};

export type SvcProductsListProductErrors = {
  /**
   * BadRequest: Bad Request response.
   */
  400: _Error;
};

export type SvcProductsListProductError =
  SvcProductsListProductErrors[keyof SvcProductsListProductErrors];

export type SvcProductsListProductResponses = {
  /**
   * OK response.
   */
  200: ProductsList;
};

export type SvcProductsListProductResponse =
  SvcProductsListProductResponses[keyof SvcProductsListProductResponses];

export type SvcProductsCreateProductData = {
  body: ProductInput;
  path?: never;
  query?: never;
  url: "/api/products";
};

export type SvcProductsCreateProductErrors = {
  /**
   * Conflict: Conflict response.
   */
  409: _Error;
};

export type SvcProductsCreateProductError =
  SvcProductsCreateProductErrors[keyof SvcProductsCreateProductErrors];

export type SvcProductsCreateProductResponses = {
  /**
   * Created response.
   */
  201: Product;
};

export type SvcProductsCreateProductResponse =
  SvcProductsCreateProductResponses[keyof SvcProductsCreateProductResponses];

export type SvcProductsDeleteProductByIdData = {
  body?: never;
  path: {
    /**
     * Unique product identifier
     */
    productId: number;
  };
  query?: never;
  url: "/api/products/{productId}";
};

export type SvcProductsDeleteProductByIdErrors = {
  /**
   * ErrNotFound: Not Found response.
   */
  404: _Error;
};

export type SvcProductsDeleteProductByIdError =
  SvcProductsDeleteProductByIdErrors[keyof SvcProductsDeleteProductByIdErrors];

export type SvcProductsDeleteProductByIdResponses = {
  /**
   * OK response.
   */
  200: boolean;
};

export type SvcProductsDeleteProductByIdResponse =
  SvcProductsDeleteProductByIdResponses[keyof SvcProductsDeleteProductByIdResponses];

export type SvcProductsGetProductByIdData = {
  body?: never;
  path: {
    /**
     * Unique product identifier
     */
    productId: number;
  };
  query?: never;
  url: "/api/products/{productId}";
};

export type SvcProductsGetProductByIdErrors = {
  /**
   * ErrNotFound: Not Found response.
   */
  404: _Error;
};

export type SvcProductsGetProductByIdError =
  SvcProductsGetProductByIdErrors[keyof SvcProductsGetProductByIdErrors];

export type SvcProductsGetProductByIdResponses = {
  /**
   * OK response.
   */
  200: Product;
};

export type SvcProductsGetProductByIdResponse =
  SvcProductsGetProductByIdResponses[keyof SvcProductsGetProductByIdResponses];

export type SvcProductsUpdateProductByIdData = {
  body: UpdateProductByIdRequestBody;
  path: {
    /**
     * Unique product identifier
     */
    productId: number;
  };
  query?: never;
  url: "/api/products/{productId}";
};

export type SvcProductsUpdateProductByIdResponses = {
  /**
   * OK response.
   */
  200: Product;
};

export type SvcProductsUpdateProductByIdResponse =
  SvcProductsUpdateProductByIdResponses[keyof SvcProductsUpdateProductByIdResponses];

export type SvcProductsAddMediaData = {
  body: AddMediaRequestBody;
  path: {
    /**
     * Unique product identifier
     */
    productId: number;
  };
  query?: never;
  url: "/api/products/{productId}/product-medias";
};

export type SvcProductsAddMediaErrors = {
  /**
   * BadRequest: Bad Request response.
   */
  400: _Error;
};

export type SvcProductsAddMediaError =
  SvcProductsAddMediaErrors[keyof SvcProductsAddMediaErrors];

export type SvcProductsAddMediaResponses = {
  /**
   * Created response.
   */
  201: Product;
};

export type SvcProductsAddMediaResponse =
  SvcProductsAddMediaResponses[keyof SvcProductsAddMediaResponses];

export type SvcProductsRemoveMediaData = {
  body?: never;
  path: {
    /**
     * Unique product identifier
     */
    productId: number;
    productMediaId: number;
  };
  query?: never;
  url: "/api/products/{productId}/product-medias/{productMediaId}";
};

export type SvcProductsRemoveMediaErrors = {
  /**
   * ErrNotFound: Not Found response.
   */
  404: _Error;
};

export type SvcProductsRemoveMediaError =
  SvcProductsRemoveMediaErrors[keyof SvcProductsRemoveMediaErrors];

export type SvcProductsRemoveMediaResponses = {
  /**
   * OK response.
   */
  200: Product;
};

export type SvcProductsRemoveMediaResponse =
  SvcProductsRemoveMediaResponses[keyof SvcProductsRemoveMediaResponses];

export type SvcProductsAddVariantData = {
  body: AddVariantRequestBody;
  path: {
    /**
     * Unique product identifier
     */
    productId: number;
  };
  query?: never;
  url: "/api/products/{productId}/variants";
};

export type SvcProductsAddVariantErrors = {
  /**
   * ErrNotFound: Not Found response.
   */
  404: _Error;
};

export type SvcProductsAddVariantError =
  SvcProductsAddVariantErrors[keyof SvcProductsAddVariantErrors];

export type SvcProductsAddVariantResponses = {
  /**
   * Created response.
   */
  201: Product;
};

export type SvcProductsAddVariantResponse =
  SvcProductsAddVariantResponses[keyof SvcProductsAddVariantResponses];

export type SvcProductsRemoveVariantData = {
  body?: never;
  path: {
    /**
     * Product ID
     */
    productId: number;
    /**
     * Product ID
     */
    variantId: number;
  };
  query?: never;
  url: "/api/products/{productId}/variants/{variantId}";
};

export type SvcProductsRemoveVariantErrors = {
  /**
   * ErrNotFound: Not Found response.
   */
  404: _Error;
};

export type SvcProductsRemoveVariantError =
  SvcProductsRemoveVariantErrors[keyof SvcProductsRemoveVariantErrors];

export type SvcProductsRemoveVariantResponses = {
  /**
   * OK response.
   */
  200: Product;
};

export type SvcProductsRemoveVariantResponse =
  SvcProductsRemoveVariantResponses[keyof SvcProductsRemoveVariantResponses];

export type SvcVendorListData = {
  body?: never;
  path?: never;
  query?: {
    /**
     * Record per page
     */
    pageSize?: number;
    /**
     * Start listing after this resource
     */
    after?: number;
  };
  url: "/api/vendors";
};

export type SvcVendorListErrors = {
  /**
   * BadRequest: Bad Request response.
   */
  400: _Error;
};

export type SvcVendorListError = SvcVendorListErrors[keyof SvcVendorListErrors];

export type SvcVendorListResponses = {
  /**
   * OK response.
   */
  200: VendorList;
};

export type SvcVendorListResponse =
  SvcVendorListResponses[keyof SvcVendorListResponses];

export type SvcVendorCreateData = {
  body: VendorInput;
  path?: never;
  query?: never;
  url: "/api/vendors";
};

export type SvcVendorCreateErrors = {
  /**
   * BadRequest: Bad Request response.
   */
  400: _Error;
};

export type SvcVendorCreateError =
  SvcVendorCreateErrors[keyof SvcVendorCreateErrors];

export type SvcVendorCreateResponses = {
  /**
   * Created response.
   */
  201: Vendor;
};

export type SvcVendorCreateResponse =
  SvcVendorCreateResponses[keyof SvcVendorCreateResponses];

export type SvcVendorDeleteByIdData = {
  body?: never;
  path: {
    /**
     * Unique product identifier
     */
    vendorId: number;
  };
  query?: never;
  url: "/api/vendors/{vendorId}";
};

export type SvcVendorDeleteByIdErrors = {
  /**
   * NotFound: Not Found response.
   */
  404: _Error;
};

export type SvcVendorDeleteByIdError =
  SvcVendorDeleteByIdErrors[keyof SvcVendorDeleteByIdErrors];

export type SvcVendorDeleteByIdResponses = {
  /**
   * OK response.
   */
  200: boolean;
};

export type SvcVendorDeleteByIdResponse =
  SvcVendorDeleteByIdResponses[keyof SvcVendorDeleteByIdResponses];

export type ClientOptions = {
  baseURL: "http://localhost:3030" | (string & {});
};
