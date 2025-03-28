// This file is auto-generated by @hey-api/openapi-ts

export type CreateMediaResponse = {
  media: Media;
  uploadUrl: string;
};

export type Creds = {
  /**
   * Access JWT Token
   */
  accessToken: string;
  /**
   * Refresh JWT Token
   */
  refreshToken: string;
  user?: string;
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

export type ProductMediaUpsertInput = {
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
  medias?: Array<ProductMediaUpsertInput>;
  removeMediaIds?: Array<number>;
  removeVariantIds?: Array<number>;
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
  variants?: Array<ProductVariantUpsertInput>;
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
  featureMediaLoc?: number;
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
   * ProductMedia which would be focus when a variant is picked by the user
   */
  featureMediaLoc?: number;
  /**
   * Price on cents
   */
  price: number;
};

export type ProductVariantUpsertInput = {
  /**
   * Color in HEX value that would be used on the variant picker
   */
  colorHex?: string;
  /**
   * Color variant option
   */
  colorName: string;
  /**
   * ProductMedia which would be focus when a variant is picked by the user
   */
  featureMediaLoc?: number;
  /**
   * Key ID
   */
  id?: number;
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

export type UpdateRequestBody = {
  payload: UserCreateInput;
};

export type User = {
  /**
   * Datetime
   */
  createdAt: string;
  email: string;
  firstName: string;
  /**
   * Key ID
   */
  id: number;
  lastName?: string;
  role: "CUSTOMER" | "ADMIN";
  /**
   * Datetime
   */
  updatedAt?: string;
};

export type UserCreateInput = {
  email: string;
  firstName: string;
  lastName?: string;
  password: string;
  role?: "CUSTOMER" | "ADMIN";
};

/**
 * Paginated results
 */
export type UserList = {
  /**
   * Data
   */
  data: Array<User>;
  pageInfo: PageInfo;
};

export type UserRegistrationInput = {
  email: string;
  firstName: string;
  lastName?: string;
  password: string;
};

export type AuthLoginData = {
  body?: never;
  path?: never;
  query?: never;
  url: "/api/auth/login";
};

export type AuthLoginErrors = {
  /**
   * NotValidCrendentials: Unauthorized response.
   */
  401: _Error;
};

export type AuthLoginError = AuthLoginErrors[keyof AuthLoginErrors];

export type AuthLoginResponses = {
  /**
   * OK response.
   */
  200: Creds;
};

export type AuthLoginResponse = AuthLoginResponses[keyof AuthLoginResponses];

export type AuthMeData = {
  body?: never;
  path?: never;
  query?: never;
  url: "/api/auth/me";
};

export type AuthMeErrors = {
  /**
   * NotValidCrendentials: Unauthorized response.
   */
  401: _Error;
  /**
   * Unproccesable: Unprocessable Entity response.
   */
  422: _Error;
};

export type AuthMeError = AuthMeErrors[keyof AuthMeErrors];

export type AuthMeResponses = {
  /**
   * OK response.
   */
  200: User;
};

export type AuthMeResponse = AuthMeResponses[keyof AuthMeResponses];

export type AuthRefreshRefreshData = {
  body?: never;
  path?: never;
  query?: never;
  url: "/api/auth/refresh";
};

export type AuthRefreshRefreshErrors = {
  /**
   * NotValidToken: Conflict response.
   */
  409: _Error;
};

export type AuthRefreshRefreshError =
  AuthRefreshRefreshErrors[keyof AuthRefreshRefreshErrors];

export type AuthRefreshRefreshResponses = {
  /**
   * OK response.
   */
  200: Creds;
};

export type AuthRefreshRefreshResponse =
  AuthRefreshRefreshResponses[keyof AuthRefreshRefreshResponses];

export type AuthSignupData = {
  body: UserRegistrationInput;
  path?: never;
  query?: never;
  url: "/api/auth/signup";
};

export type AuthSignupErrors = {
  /**
   * BadInput: Bad Request response.
   */
  400: _Error;
};

export type AuthSignupError = AuthSignupErrors[keyof AuthSignupErrors];

export type AuthSignupResponses = {
  /**
   * OK response.
   */
  200: Creds;
};

export type AuthSignupResponse = AuthSignupResponses[keyof AuthSignupResponses];

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

export type SvcuserListData = {
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
  url: "/api/users";
};

export type SvcuserListResponses = {
  /**
   * OK response.
   */
  200: UserList;
};

export type SvcuserListResponse =
  SvcuserListResponses[keyof SvcuserListResponses];

export type SvcuserCreateData = {
  body: UserCreateInput;
  path?: never;
  query?: never;
  url: "/api/users";
};

export type SvcuserCreateResponses = {
  /**
   * Created response.
   */
  201: User;
};

export type SvcuserCreateResponse =
  SvcuserCreateResponses[keyof SvcuserCreateResponses];

export type SvcuserDeleteData = {
  body?: never;
  path: {
    userId: number;
  };
  query?: never;
  url: "/api/users/{userId}";
};

export type SvcuserDeleteResponses = {
  /**
   * Accepted response.
   */
  202: boolean;
};

export type SvcuserDeleteResponse =
  SvcuserDeleteResponses[keyof SvcuserDeleteResponses];

export type SvcuserShowData = {
  body?: never;
  path: {
    userId: number;
  };
  query?: never;
  url: "/api/users/{userId}";
};

export type SvcuserShowResponses = {
  /**
   * OK response.
   */
  200: User;
};

export type SvcuserShowResponse =
  SvcuserShowResponses[keyof SvcuserShowResponses];

export type SvcuserUpdateData = {
  body: UpdateRequestBody;
  path: {
    userId: number;
  };
  query?: never;
  url: "/api/users/{userId}";
};

export type SvcuserUpdateResponses = {
  /**
   * OK response.
   */
  200: User;
};

export type SvcuserUpdateResponse =
  SvcuserUpdateResponses[keyof SvcuserUpdateResponses];

export type ClientOptions = {
  baseURL: "http://localhost:3030" | (string & {});
};
