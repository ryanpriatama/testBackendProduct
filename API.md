# Application Programming Interface Documentation

## List
| No |     Web Service     | Method | URL |
|----|---------------------|--------|-----|
| 1 | [Create Product](#create-product) | POST | /api/products |
| 2 | [Get Products Sorted](#get-products-sorted) | GET | /api/products?sort= |


## Create Product
### URL : `/api/products`
### Method : `POST`

### Header
    'X-API-Key: api-key-rahasia' \
    'Content-Type: application/json'

### Body Request
```json
{
    "name" : "buku gambar",
    "price" : 12000,
    "description" : "barang baru ready siap antar",
    "quantity" : 20
}
```

### Example cURL
curl --location 'localhost:3000/api/products' \
--header 'x-API-Key: api-key-rahasia' \
--header 'Content-Type: application/json' \
--data '{
    "name" : "abcd123",
    "price" : 10000,
    "description" : "barang baru",
    "quantity" : 50
}'

### Body Response
```json
{
    "code": 200,
    "status": "OK",
    "data": {
        "id": 9,
        "name": "buku gambar",
        "price": 12000,
        "description": "barang baru ready siap antar",
        "quantity": 20
    }
}

{
    "code": 401,
    "status": "UNAUTHORIZED",
    "data": Null
}
```

## Get Products Sorted
### URL : `/api/products`
### Method : `GET`

### Query Params
    sort=ascending_name

### Header
    'X-API-Key: api-key-rahasia'

### Example cURL
curl --location 'localhost:3000/api/products?sort=ascending_name' \
--header 'x-api-key: api-key-rahasia'

### Body Response
```json
{
    "code": 200,
    "status": "OK",
    "data": [
        {
            "Id": 1,
            "Name": "abcd123",
            "Price": 10000,
            "Description": "barang baru",
            "Quantity": 50
        },
        {
            "Id": 2,
            "Name": "baju hitam",
            "Price": 80000,
            "Description": "barang baru bersih",
            "Quantity": 100
        },
        {
            "Id": 9,
            "Name": "buku gambar",
            "Price": 12000,
            "Description": "barang baru ready siap antar",
            "Quantity": 20
        },
        {
            "Id": 3,
            "Name": "celana panjang",
            "Price": 120000,
            "Description": "barang baru",
            "Quantity": 33
        },
        {
            "Id": 4,
            "Name": "kartu yugi",
            "Price": 45000,
            "Description": "barang baru ready",
            "Quantity": 20
        },
        {
            "Id": 6,
            "Name": "test",
            "Price": 100,
            "Description": "test product",
            "Quantity": 10
        },
        {
            "Id": 7,
            "Name": "test",
            "Price": 100,
            "Description": "test product",
            "Quantity": 10
        },
        {
            "Id": 8,
            "Name": "test",
            "Price": 100,
            "Description": "test product",
            "Quantity": 10
        },
        {
            "Id": 5,
            "Name": "vas bunga",
            "Price": 39000,
            "Description": "barang baru ready siap antar",
            "Quantity": 25
        }
    ]
}

{
    "code": 401,
    "status": "UNAUTHORIZED",
    "data": Null
}
```