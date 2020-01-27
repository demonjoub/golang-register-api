# Register
**build**

    go mod tindy
    go mod vender
    go build

**run**

    ./register-service

**URL**

    #GET /personas
    [{
        "id": 1,
        "UserId": "xxxx",
        "FirstName": "xxx",
        "LastName": "xxx",
        ...,
    },
    ...,
    ]
    
    #POST /register
    {
	    "id": 1,
        "UserId": "xxxx",
        "FirstName": "xxx",
        "LastName": "xxx",
        ...,
    }
    
    #PUT /register/:id
    {
	    "id": 1,
        "UserId": "xxxx",
        "FirstName": "xxx",
        "LastName": "xxx",
        ...,
    }
    

  