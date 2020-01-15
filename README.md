# Ally
An ecommerce agent based on [Ethereum](https://ethereum.org/). The business can collect any ERC-20 token through sell products , such as USDT, BAT and so on.

## Quick Start
### Dependencies
`./scripts/dev_setup.sh` will install the following dependencies.
1. Ally implemented in [Golang](https://golang.org/).
2. [beanstalk](https://beanstalkd.github.io/) is in charge of internal communication.
3. Besides [infura](https://infura.io/) and [Etherscan](https://etherscan.io/) provide the Ethereum APIs. Please get the api keys from their platforms.

### Environment
The Ethereum has one mainnet and some testnet. Also it has many token. Ally choose them through the operating system environment variables. Besides other configurations are set by this way. Please set the following environment variables before setup ally.
- `INFURA_ENDPOINT` [the infura endpoint](https://infura.io/docs/gettingStarted/chooseaNetwork.md), default is `kovan`
- `INFURA_ID` the infura project ID from [your dashboard](https://infura.io/dashboard)
- `ETHERSCAN_HOST` the Etherscan host, default is `api-kovan.etherscan.io`
- `ETHERSCAN_APIKEY` the Etherscan APIKEY
- `TOKEN_CONTRACT` the token contract address, default is `0xFab46E002BbF0b4509813474841E0716E6730136`
- `CONTRACT_ADDRESS_FACTORY_ADDRESS` the contract address, get it after [deploy the smart contract](#deploy-smart-contracts)

### Setup
#### start component: api
```bash
$ go run cmd/api/main.go
```

#### get the management address and deposit some ethers as gas fee
```bash
$ curl "http://localhost:8080/api/v1/admin/rootAddress"
```

#### deploy smart contracts
Deploy the contracts under [the contracts directory](contracts/) and get the contract address. Remember set it into the [Environment](#environment).

#### start component: dispatcher and worker
Don't forget start `beanstalk` before this step.
- dispatcher
```bash
$ go run cmd/dispatcher/main.go
```

- worker
```bash
$ go run cmd/worker/main.go
```

### Try
1. create a product
```bash
$ curl -X POST -H "Content-type: application/json" -H "Accept: application/json" -d '{"externalID": "my_id", "price": "0.01"}' http://localhost:8080/api/v1/items
```

2. create an order
```bash
$ curl -X POST -H "Content-type: application/json" -H "Accept: application/json" -d '{"itemID": 1, "sequence": 1, "account": "test"}' http://localhost:8080/api/v1/orders
```

3. get the deposit address and then deposit the exact tokens into the address
> It maybe take a little time for confirmation. You can check the order's transaction status with [the ethereum browser](https://etherscan.io/).
```bash
$ curl "http://localhost:8080/api/v1/orders/1"
```

4. check the balance of Ally agent
```bash
$ curl "http://localhost:8080/api/v1/stat"
```

5. withdraw tokens
```bash
$ curl -X POST -H "Content-type: application/json" -H "Accept: application/json" -d '{"to": "0xAdbf42299d432Db7A70e298d07B7f33ce84Ae095", "amount": "0.01"}' http://localhost:8080/api/v1/admin/withdraw
```

6. check the status of withdraw
```bash
$ curl "http://localhost:8080/api/v1/transactiongroups/1"
```
