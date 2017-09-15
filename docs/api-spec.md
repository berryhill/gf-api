# API spec

### GET/products/:product_name

ex:

```
$ curl http://localhost:8080/products/fly_rods
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

### POST/retailer

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

### POST/:scraper/scrape

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
