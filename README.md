# 현재 완성도
....invoke....  
user_insert 100%  
home_insert 100%  
pet_insert 100%  
user_change 0%  
home_change 0%  
pet_change 0%  
home_delete 100%  
pet_delete 100%  
trade_insert 100%  

....Query....  
user_read 100%  
home_read 100%  
pet_read 100%  
city_search 100%  
trade_search 100%  


-------------
## [membersrvc]
#### localhost:7050/registrar

    {
    "enrollId": "admin",
    "enrollSecret": "Xurw3yU9zI0l"
    }
#### result

    {
      "OK": "Login successful for user 'admin'."
    }
-------------
## [deploy] chaincode
#### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "deploy",
      "params": {
        "type": 1,
        "chaincodeID":{
            "name": "mycc"
        },
        "ctorMsg": {
            "args":[""]
        },
        "secureContext": "admin"
      },
      "id": 1
    }
#### result
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "mycc"
      },
      "id": 1
    }
-------------
## [invoke] User Insert
#### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["user_insert", "key", "pw"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

#### result
    {
    "jsonrpc": "2.0",
    "result": {
    "status": "OK",
    "message": "8d042656-ca1d-4ce2-bb0c-c4cdd15ddd64"
    },
    "id": 3
    }
-------------
## [invoke] Home Insert
#### localhost:7050/chaincode

    {
    "jsonrpc": "2.0",
    "method": "invoke",
    "params": {
      "type": 1,
      "chaincodeID":{
          "name":"mycc"
      },
      "ctorMsg": {
         "args":["home_insert", "key", "R103", "x1", "x2", "x3", "x4", "x5", "x6"]
      },
      "secureContext": "admin"
    },
    "id": 3
    }


#### result
    {
    "jsonrpc": "2.0",
    "result": {
    "status": "OK",
    "message": "8b75fb78-65e7-47af-b9ca-e94c19c7b191"
    },
    "id": 3
    }

-------------
## [invoke] Pet Insert
#### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["pet_insert", "key", "y1", "y2", "y3", "y4", "y5", "y6", "y7"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
#### result
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "7389bc5d-b376-4989-862a-1c466a1b2f81"
      },
      "id": 3
    }
-------------
## [query] User Read
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["user_read","key"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
#### result
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"PW\":\"pw\",\"PN\":\"1\",\"CC\":\"R103\",\"AP\":\"0\"}"
      },
      "id": 3
    }
-------------
## [query] Home Read
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["home_read","key"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
#### result
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"Address\":\"x1\",\"HomeType\":\"x2\",\"Room\":\"x3\",\"Area\":\"x4\",\"Elevator\":\"x5\",\"Parking\":\"x6\"}"
      },
      "id": 3
    }
-------------
## [query] Pet Read
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["pet_read","key"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
#### result
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"Name\":\"y1\",\"Birth\":\"y2\",\"Gender\":\"y3\",\"Kind\":\"y4\",\"Size\":\"y5\",\"NS\":\"y6\",\"Vac\":\"y7\"}"
      },
      "id": 3
    }
-------------
## [query] City Search
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["city_search","R103"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
#### result
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "/key/key2/"
      },
      "id": 3
    }
-------------
## [invoke] Home Delete
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["pet_delete", "key"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

#### result (query한 결과)
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"PW\":\"pw\",\"PN\":\"0\",\"CC\":\"R103\",\"AP\":\"0\"}"
      },
      "id": 3
    }

    {
      "jsonrpc": "2.0",
      "error": {
        "code": -32003,
        "message": "Query failure",
        "data": "Error when querying chaincode: Error:Transaction or query returned with failure: [PET QUERY] Not exist pet information"
      },
      "id": 3
    }


-------------
## [invoke] Pet Delete
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["home_delete", "key"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

#### result
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"PW\":\"pw\",\"PN\":\"0\",\"CC\":\"0\",\"AP\":\"0\"}"
      },
      "id": 3
    }

    {
      "jsonrpc": "2.0",
      "error": {
        "code": -32003,
        "message": "Query failure",
        "data": "Error when querying chaincode: Error:Transaction or query returned with failure: [HOME QUERY] Not exist home information"
      },
      "id": 3
    }


-------------
## [invoke] Trade Insert
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["trade_insert", "petsitter", "consumer", "s", "e", "c", "a", "h"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

-------------
## [query] Trade Search
#### localhost:7050/chaincode
    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["trade_search", "petsitter", "consumer", "c"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
-------------
