# Smart Contracts

## Address Factory
A Smart contract that generate addresses can be used to withdraw without any ether, using a common address pay the transaction fee.

### Compilation
```bash
$ solc --abi AddressFactory.sol -o build
$ solc --bin AddressFactory.sol -o build
$ abigen --bin=./build/AddressFactory.bin --abi=./build/AddressFactory.abi --pkg=contracts --type=AddressFactory --out=addressFactory.go
$ solc --abi ERC20.sol -o build
$ abigen --abi=./build/ERC20.abi --pkg=contracts --type=ERC20 --out=ERC20.go
```

### Deploy
```bash
$ INFURA_ENDPOINT= INFURA_ID= go run deploy/*
```
