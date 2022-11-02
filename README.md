# golang-microservice-101
golang-microservice-101

![alt text](https://raw.githubusercontent.com/SerhatSelim/golang-microservice-101/main/go_microservice_101.drawio.png)


## test proxies
//curl localhost:3000/payment

//curl localhost:3000/moneygram

## test gateways:
//curl localhost:30001/api/v1/gateway/healty

//curl localhost:30001/api/v1/gateway/payment-api/healty

//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww"}' localhost:30001/api/v1/gateway/payment

//curl localhost:30001/api/v1/gateway/fastpay-api/healty

//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true}' localhost:30001/api/v1/gateway/fastpay

//curl localhost:30002/api/v1/moneygram-gateway/healty

//curl localhost:30002/api/v1/moneygram-gateway/send-money-api/healty

//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true, "desc":"send money"}' localhost:30002/api/v1/moneygram-gateway/send-money

//curl localhost:30002/api/v1/moneygram-gateway/merchant-api/healty

//curl -X POST -H "Content-Type: application/json"  -d '{"amount": 100,"from": "cobadeff","to": "nabatww","flag": true, "desc":"merchant"}' localhost:30002/api/v1/moneygram-gateway/merchant
