# Fabric develop 

## Prerequisites
### Install docker、golang、node.js...
```
./installFabricEnv_ubuntu.sh
```
## Install crypto tools
### If you meet some problem,you can set proxy like below:
```
export http_proxy="http://127.0.0.1:54660"
export https_proxy="http://127.0.0.1:54660"
```
```
curl -sSL http://bit.ly/2ysbOFE | bash -s 1.2.0
```
## Start blockchain service
```
cd fabric-samples/balance-transfer
./runApp.sh
```

## Test Apis

```
./testApi.sh
```