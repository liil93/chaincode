# Hyperledger Fabric - Pet Sitting

## 개요
펫시팅 체인코드(PS.go)는 사용자의 정보, 집과 펫에 대한 자산을 등록하여 펫시터와
이용자 사이에 오가는 거래를 하나의 블록체인 내에서 관리하기 위한 코드입니다.


체인 코드는 다음과 같은 기능을 제공합니다.  

#### **chaincode - invoke**
1. user_insert:처음 유저를 등록하는 기능이며 userID, password만을 받아서 블록에 저장합니다.
2. home_insert:집을 등록하는 기능이며 args로 집에 대한 정보를 입력받으며 user의 citycode를 업데이트합니다.
3. pet_insert:펫을 등록하는 기능이며 펫에 대한 정보를 저장함과 동시에 user의 정보에 펫의 수를 업데이트합니다.
4. user_change:유저의 password와 펫시팅 여부를 변경하여주는 기능을 합니다.
5. ~~home_change:집은 insert와 delete가 있으므로 변경은 필요하지 않다고 생각합니다.~~
6. pet_change:유저가 소유한 펫에 대한 사이즈, 중성화수술 여부, 백신 여부를 변경하는 기능을 합니다.
7. home_delete:유저가 소유하고 있는 집에 대한 정보를 지워줍니다.
8. pet_delete:유저가 소유하고 있는 펫에 대한 정보를 지워줍니다.
9. trade_insert:펫시터와 고객간의 거래 정보를 블록에 저장하는 기능을 합니다.

#### **chaincode - query**
1. user_read:유저의 현재 상태를 조회합니다.
2. home_read:유저의 집에 대한 자산정보를 조회합니다.
3. pet_read:유저의 펫에 대한 자산정보를 조회합니다.
4. city_search:입력한 지역에 펫시팅 여부가 1인 유저들의 ID를 조회합니다.
5. trade_search:펫시터와 고객간의 거래 정보를 조회합니다.

### *membersrvc*
##### localhost:7050/registrar

    {
        "enrollId": "admin",
        "enrollSecret": "Xurw3yU9zI0l"
    }

admin으로 로그인하여 아래와 같은 문장이 뜨면 됩니다.

    {
      "OK": "Login successful for user 'admin'."
    }

### *init(deploy)*
##### localhost:7050/chaincode

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

mycc라는 PS체인코드를 deploy하여 줍니다. 이 때 args로는 아무것도 입력하지 않습니다.  

### *user_insert(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["user_insert", "userID", "password"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

args로는 키값으로 사용될 유저의 ID와 패스워드를 입력받습니다. 나머지의 유저정보는 0으로 초기화됩니다.  

### *home_insert(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["home_insert", "userID", "citycode", "address", "hometype", "room", "area", "elevator", "parking"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

유저의 집 자산을 등록하는 기능을 하며 각각의 args에는 다음과 같이 들어갑니다.  이 때 엘리베이터와 주차는 Y 또는 N으로 표시됩니다.

### *pet_insert(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["pet_insert", "userID", "name", "birth", "gender", "kind", "size", "ns", "vac"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

유저의 펫 자산을 등록하는 기능을 하며 각각의 args에는 다음과 같이 들어갑니다.  
이 중 Size와 NS, Vac만 수정이 가능합니다.

### *user_change(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["user_change", "userID", "new password", "AP"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

유저의 정보를 수정하며 첫번 째 값에는 유저의 이메일이 들어가고 두번 째에는 변경하고자 하는 패스워드,  
세번 째에는 펫시팅 여부가 들어갑니다. 이 때, 패스워드를 변경하고 싶지 않을 경우에는 '0'을 입력받고  
AP를 변경할 경우, 집 자산이 꼭 있어야 하며 city_search에 조회가 되게끔 저장됩니다.

### *pet_change(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["pet_change", "userID", "size", "ns", "vac"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

유저의 펫 정보를 수정하는 기능을 하며 첫번 째에는 펫의 사이즈를 변경하고 두번 째에는 펫의 중성화 수술  
여부를, 세번 째에는 펫의 백신정보를 변경하여 줍니다.

### *home_delete(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["home_delete", "userID"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

유저의 집 자산을 삭제하는 역할을 합니다.

### *pet_delete(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["pet_delete", "userID"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

유저의 펫 자산을 삭제하는 역할을 합니다.

### *trade_insert(invoke)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "invoke",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["trade_insert", "petsitterID", "consumerID", "start time", "end time", "transaction complete time", "transaction amount", "transaction history"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }

펫시터와 이용자간의 거래를 블록에 저장하는 기능을 하며, 거래의 결제가 종료되는 시간에 기록이 됩니다.

### *user_read(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["user_read","userID"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
~~~~
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"PW\":\"password\",\"PN\":\"0\",\"CC\":\"0\",\"AP\":\"0\"}"
      },
      "id": 3
    }
~~~~
유저의 정보를 조회하는 기능입니다.

### *home_read(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["home_read","userID"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
~~~~
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"Address\":\"address\",\"HomeType\":\"hometype\",\"Room\":\"room\",\"Area\":\"area\",\"Elevator\":\"elevator\",\"Parking\":\"parking\"}"
      },
      "id": 3
    }
~~~~
유저의 ID로 집 자산에 대한 정보를 조회하는 기능입니다.

### *pet_read(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["pet_read","userID"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
~~~~
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"Name\":\"name\",\"Birth\":\"birth\",\"Gender\":\"gender\",\"Kind\":\"kind\",\"Size\":\"size\",\"NS\":\"ns\",\"Vac\":\"vac\"}"
      },
      "id": 3
    }
~~~~
유저의 ID로 펫 자산에 대한 정보를 조회하는 기능입니다.

### *city_search(query)*
##### localhost:7050/chaincode


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
~~~~
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "/user_1/user_2/user_3/"
      },
      "id": 3
    }
~~~~
City code를 가지고 그 지역에 있는 펫시팅이 가능한 유저들의 이메일 값을 보여줍니다.

### *trade_search(query)*
##### localhost:7050/chaincode

    {
      "jsonrpc": "2.0",
      "method": "query",
      "params": {
          "type": 1,
          "chaincodeID":{
              "name":"mycc"
          },
          "ctorMsg": {
             "args":["trade_search", "petsitterID", "consumerID", "transaction complete time"]
          },
          "secureContext": "admin"
      },
      "id": 3
    }
~~~~
    {
      "jsonrpc": "2.0",
      "result": {
        "status": "OK",
        "message": "{\"PSID\":\"petsitterID\",\"CSID\":\"consumerID\",\"TS\":\"start time\",\"TE\":\"end time\",\"TC\":\"transaction complete time\",\"TA\":\"transaction amount\",\"TH\":\"transaction history\"}"
      },
      "id": 3
    }
~~~~
거래 기록을 보여주며 펫시터, 이용자의 ID와 결제 완료 시간을 가지고 검색을 하게 됩니다.
