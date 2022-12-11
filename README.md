# transactions

API to save and list user transactions.

## How to start API

You need to have `docker` and `docker-compose` installed before execute `transactions` api

In a terminal shell:

- Clone this repository
```
git clone https://github.com/lawmatsuyama/transactions.git
cd transactions
```

- In transactions directory, run the command below to start the API
```
docker-compose up -d
```
It will build the API and create containers to mongodb, rabbitmq and transactions API

Observation: It will use the net ports 15672, 5672, 27017 and 8080. So be attention to don't lock those ports before run the command above.

## How to test API

With the API up and running, you can access the swagger in some browser:
```
http://localhost:8080/transactions/swagger/index.html
```

It contains two operations:
- /v1/get: List transactions by giving filter. All fields you send will be considered for filtering transactions. For example if you send `user_id` and `date_from`, it will return transactions by user_id and from the date you sent. For the fields `date_from` and `date_to`, you have to send in RFC3339 format, like `2022-12-31T23:59:59-03:00`. You can omit or just send zero values in the fields that you want to ignore in filter.

Request example:
```json
{
  "_id": "",
  "amount_greater": 0,
  "amount_less": 0,
  "description": "",
  "operation_type": "",
  "origin": "",
  "date_from":"2022-12-10T23:19:42-03:00",
  "user_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe"
}
```

It will return 20 transactions for each page. So if you want to view the next page of transactions, you can check response object `paging`, get the `next_page` and then set `page` in the new request.
If response don't return `next_page` in `paging`, it means there are no more transactions to return.


```json
{
  "_id": "",
  "amount_greater": 0,
  "amount_less": 0,
  "description": "",
  "operation_type": "",
  "origin": "",
  "date_from":"2022-12-10T23:19:42-03:00",
  "user_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe",
  "paging":{
        "page": 3
    }
}
```

- /v1/save: Receives transactions data, registed it in application and finish notifying other applications. You can send a list of transactions to register it. 

Request example:
```json
{
    "user_id": "74211f47-9c6f-4648-88a3-cb7d0614b5fe",
    "origin_channel": "desktop-web",
    "transactions": [
        {
            "description": "test21",
            "amount": 5678.34,
            "operation": "debit"
        },
        {
            "description": "test2",
            "amount": 14.1,
            "operation": "credit"
        },
        {
            "description": "test3",
            "amount": 14.1,
            "operation": "debit"
        }
    ]
}
```

You can send up to 20 transactions for each request. More than that it will return an error. 

The `origin_channel` field accept the follow values:
- "desktop-web"
- "mobile-android"
- "mobile-ios"

`operation` field accept the follow values:
- debit
- credit

