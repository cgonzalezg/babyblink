#API

##FAMILY
###create
<b>POST```/famlily/create```

<b>body
```json
{
  "id":"",
  "name":"Kokoska"
}
```
<b>Response

```json
{
    "data":
      {
        "id":"5444e91ff360c5f41d000002",
        "name":"test56",
        "c":"2014-10-20T12:51:11.496+02:00",
        "u":"2014-10-20T12:51:11.496+02:00"
      }
}
```

###update
<b>POST</b> ```/famlily/update```

<b>body
```json
{
  "id":"5444e91ff360c5f41d000002", //needs to be set
  "name":"Kokoska",
}
```
<b>Response

```json
{
    "data":
      {
        "id":"5444e91ff360c5f41d000002",
        "name":"test56",
        "c":"2014-10-20T12:51:11.496+02:00",
        "u":"2014-11-20T14:04:14.778+02:00"
      }
}
```
####all

<b>GET</b> ```/famlily/all```

<b>Response
```json
{
  "data": [
    {
      "id": "5443be6631e79f8a835042f3",
      "name": "item1",
      "c": "0001-01-01T00:00:00Z",
      "u": "0001-01-01T00:00:00Z"
    },
    {
      "id": "5443be7f31e79f8a835042f4",
      "name": "item2",
      "c": "0001-01-01T00:00:00Z",
      "u": "0001-01-01T00:00:00Z"
    }
  ]
}
```
###child
- create
- update
- delete
- picture
 - upload
 - get
- all
- get

###parent
- create
- update
- delete
- picture
 - upload
 - get
- all
- get

###family-member
- create
- update
- delete
- picture
 - upload
 - get
- all
- get

###event
- create
- update
- delete
- picture
 - upload
 - get
- all
- get

##Picture

###Upload

###Delete

###update
- child
 - add
 - delete
- comment
 - delete
 - add
- event
 - delete
 - add
- people
 - add
 - delete

###share

###get-info
- event
- all
- child
- people
- comment
