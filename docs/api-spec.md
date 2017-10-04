# API spec

## GET/products/:product_name

ex:

```
$ curl http://104.154.247.230:8080/products/fly_rods?search=sage
```

#### Query Parameters
```
?search={string}

```

#### Response
```
{
    "metadata": {
        "count": {int},
        "page": {int},
        "page_count": {int},
        "per_page": {int}
    },
    products: [
        {
            "product_id": {string},
            "active": {bool},
            "url": {string},
            "image": {string},
            "type": {string},
            "brand": {string},
            "name": {string},
            "title": {string},
            "price": {string},
            "retailer": {string},
            "details": [
                {string},
                ...
            ],
            "managed": {bool}
        },
        ...
    ]
}
```

## POST/product-type

#### Request Payload
```
{
    "name": {string}
}
```

ex:

```
$ curl -d '{"name":"Fly Rods"}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/product-type
```

#### Response
```
{
    "name": {string}
}
```

## GET/product-types

ex:

```
$ curl http://localhost:8080/product-types
```

#### Response
```
[
    {
        "_id": {bson.ObjectId}
        "name": {string}
    }
]
```

## POST/retailer

#### Request Payload
```
{
    "name": {string},
    "base_url": {string},
    "products": {
        {string}: {string},
        ...
    }
}
```

ex:

```
$ curl -d '{"name":"smapler", "base_url":"s.com", "products": {"sample": "/sample"}}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/retailer
```

#### Response
```
{
    "name": {string},
    "base_url": {string},
    "products": {
        {string}: {string},
        ...
    }
}
```

## POST/:scraper/scrape

ex:

```
$ curl -X POST http://localhost:8080/backcountry/scrape
$ curl -X POST http://localhost:8080/cabelas/scrape
```

#### Response
```
{
    "items_added": {int},
    "items_found": {int}
}
```
