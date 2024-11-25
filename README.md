# golang-crud

## PRODUCT API SPEC

### Create Product

Endpoint: POST /api/products/create

Request Body:

```json
{
  "product_name": "example",
  "price": 18000000.0,
  "description": "example"
}
```

Response Body (Success):

```json
{
  "id": 1,
  "product_name": "example",
  "price": 18000000.0,
  "description": "example",
  "created_at": "2024-11-25T14:54:05.9382381+07:00",
  "updated_at": "2024-11-25T14:54:05.9382381+07:00"
}
```

Response Body (Failed):

```json
{
  "error": "Failed to create product - Internal server error"
}
```

### Get Product By ID

Endpoint: GET /api/products/:id

Example Endpoint: GET /api/products/1

Response Body (Success):

```json
{
  "id": 1,
  "product_name": "example",
  "price": 18000000.0,
  "description": "example",
  "created_at": "2024-11-25T14:54:05.9382381+07:00",
  "updated_at": "2024-11-25T14:54:05.9382381+07:00"
}
```

Response Body (Failed):

```json
{
  "error": "Product not found"
}
```

### Get All Product

Endpoint: GET /api/products

Response Body (Success):

```json
[
  {
    "id": 1,
    "product_name": "example",
    "price": 18000000.0,
    "description": "example",
    "created_at": "2024-11-25T14:54:05.9382381+07:00",
    "updated_at": "2024-11-25T14:54:05.9382381+07:00"
  },
  {
    "id": 2,
    "product_name": "example",
    "price": 18000000.0,
    "description": "example",
    "created_at": "2024-11-25T14:54:05.9382381+07:00",
    "updated_at": "2024-11-25T14:54:05.9382381+07:00"
  },
  {
    "id": 3,
    "product_name": "example",
    "price": 18000000.0,
    "description": "example",
    "created_at": "2024-11-25T14:54:05.9382381+07:00",
    "updated_at": "2024-11-25T14:54:05.9382381+07:00"
  }
]
```

Response Body (Failed):

```json
{
  "error": "Failed to get all product - Internal server error"
}
```

### Update Product

Endpoint: PUT /api/products/update-product/:id

Request param: id

Request Body:

```json
{
  "product_name": "example update",
  "price": 18000000.0,
  "description": "example update"
}
```

Response Body (Success):

```json
{
  "id": 1,
  "product_name": "example update",
  "price": 18000000.0,
  "description": "example update",
  "created_at": "2024-11-25T14:54:05.9382381+07:00",
  "updated_at": "2024-11-25T14:54:05.9382381+07:00"
}
```

Response Body (Failed):

```json
{
  "error": "Product not found", "Failed to update product - Internal server error"
}
```

### Delete Product

Endpoint: Delete /api/products/delete

Request param: id

Response Body (Success):

```json
{
  "message": "Delete product success"
}
```

Response Body (Failed):

```json
{
  "error": "Product not found", "Failed to delete product - Internal server error"
}
```
