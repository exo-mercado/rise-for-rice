
## SMART PARK API Reference

### CLIENT API
#### Get all items end point

```http
  GET /client
```

| Query Parameter | Description                |
| :-------- |  :------------------------- |
| `keyword` | **Optional**. Keyword to search |


### Consumer API
#### Register consumer end point

```http
  POST /consumer
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `first_name` | `string` | **Required**. |
| `last_name` | `string` | **Required**. |
| `phone_number` | `string` | **Required**. |
| `password` | `string` | **Required**. |

#### Get consumer by ID

```http
  GET /consumer/{:id}
```

| Query Parameter | Description                |
| :-------- | :------------------------- |
| `join`| **optional**.enum(*vehicle*) |


#### Register consumer end point

```http
  POST /consumer/login
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `phone_number` | `string` | **Required**. |
| `password` | `string` | **Required**. |


### Vehicle API

#### Create Vehicle end point
```http
  GET /vehicle/consumer/{:id}
```

#### Create Vehicle end point
###### when you create a new vehicle, this will be set as default vehicle for the consumer

```http
  POST /vehicle
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `make` | `string` | **Required**. |
| `model` | `string` | **Required**. |
| `year` | `string` | **Required**. |
| `color` | `string` | **Required**. |
| `plate` | `string` | **Required**. |
| `consumer_id` | `string` | **Required**. Logged in consumer|

### Reservation API
#### Book reservation end point

```http
  POST /reservation
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `from` | `DateTime unix` | **Required**. |
| `to` | `DateTime unix` | **Required**. |
| `grid_number` | `int` | **Required**. |
| `vehicle_id` | `int` | **Required**. |

