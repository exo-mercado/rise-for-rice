# rise-for-rice
API for parking application

## API Reference

### CLIENT API
#### Get all items

```http
  GET /client
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `keyword` | `string` | **Optional**. Keyword to search |


### Consumer API
#### Register consumer

```http
  POST /consumer
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `first_name` | `string` | **Required**. |
| `last_name` | `string` | **Required**. |
| `phone_number` | `string` | **Required** **Unique**. |
| `password` | `string` | **Required**. |