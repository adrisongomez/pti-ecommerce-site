# Ecommerce API

*An REST API build on golang using goa to generate OpenAPI specification.*

## Endpoints
| Method | Endpoints                                | Description                                                               | Auth       | Erros |
| -----  | ---------------------------------------- | ------------------------------------------------------------------------- | ---------- | ----- |
| GET    | /products/                               | List all the products and also filter products by multiples criterias     | public     | 400 |
| GET    | /products/:productId                     | Get a product specify on the path params                                  | public     | 404 |
| POST   | /products/                 -             | To Create products only allowed by users with ADMIN access                | Admin-only | 422 400 |
| PUT    | /products/:productId                     | To update a product                                                       | Admin-only | 422 400 |
| PATCH  | /products/:productId      -              | Update a field on the products                                            | Admin-only | 422 400 |
| GET | /categories/ | List all the Categories | public
| GET | /categories/:categoryId | Get Category By Id | public
| POST | | | | 

