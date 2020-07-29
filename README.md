# FastFoodAPI
Sample Go API to support [mobile app ](https://github.com/alexandredsa/food_ranks)
 built with Flutter. 
### GET - api/types
```
[
    {
        "name": "hamburguer",
        "count": 6
    },
    {
        "name": "pizza",
        "count": 3
    }
]

```
### POST - api/types/suggest - `Not implemented yet`

#### Request
```
{
    "term": "piz"
}
```

#### Response
```
[
   "pizza"
]
```

### GET - /api/types/:id/reviews/
```    
[
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "BK",
        "texture": 7,
        "flavor": 6, 
        "sauce": 5,
    },
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "BK",
        "texture": 7,
        "flavor": 6, 
        "sauce": 5,
    },
    {
        "id": "507f1f77bcf86cd799439011",
        "name": "BK",
        "texture": 7,
        "flavor": 6, 
        "sauce": 5,
    }
]
```

### POST - /api/types/:id/reviews/
#### Request
```
    {
        "name": "BK",
        "texture": 7,
        "flavor": 6, 
        "sauce": 5,
    }
```
