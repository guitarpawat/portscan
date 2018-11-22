# Portscan

## SampleInput

### POST /api/scan

Input:
```json
{
  "targets": [
    {
      "address": "www.google.com",
      "ports": [
        80,
        433,
        443,
        27017
      ]
    },
    {
      "address": "127.0.0.1",
      "ports": [
        10,
        20,
        30,
        40,
        50,
        60,
        70,
        80,
        90,
        100,
        200,
        300,
        400,
        433,
        443,
        500,
        4433,
        8000,
        27017
      ]
    }
  ]
}
```

Output:
```json
{
  "token":"1542889524688374400-2241993678076296288",
  "ip":[
    "172.217.26.132",
    "127.0.0.1"
  ]
}
```

Input:
```json
{
  "targets": [
    {
      "address": "::1",
      "ports": [
        80,
        433,
        443,
        27017
      ]
    }
  ]
}
```

Output:
```json
{
  "msg":"ipv6 is not currently support"
}
```

### GET /api/token/<token_id>
Output:
```json
{
  "results": [
    {
      "ip":"172.217.166.132",
      "ports": [
        {
          "port":80,
          "description":"HyperText Transfer Protocol (HTTP)"
        },
        {
          "port":443,
          "description":"HTTP with Secure Sockets Layer (SSL)"
        }
      ],
      "finished":true
    },
    {
      "ip":"127.0.0.1",
      "ports": [
        {
          "port":80,
          "description":"HyperText Transfer Protocol (HTTP)"
        },
        {
          "port":27017,
          "description":"MongoDB daemon process"
        }
      ],
      "finished":true
    }
  ],
  "last_update":"2018-11-22T19:49:06.3678471+07:00"
}
```
or
```json
{
  "msg":"token id not found"
}
```

### DELETE /api/token/<token_id>
Output: `200 OK` or
```json
{
  "msg":"token id not found"
}
```