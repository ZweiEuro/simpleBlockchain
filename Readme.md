# Simple Blockchain

This repository contains the golang code of a simple blockchain implementation.
  
This blockchain consists of three parts:  
- A simple wallet that you can get address, scan utxos, sign transaction.  
- A simple blockchain can sync block from other known nodes, mining new block, send transaction and broadcast to other node.  
- A simple restful server you can query blocks and utxos from blockchain.  
- This project only supports two nodes.

There are many part are not like real blockchain because it's just simple implementation, still
insecure and incomplete. you can learn the basic operation of the blockchain through this project.

## changes from zweieuro (author):
- The peers were saved in an array, but there are only ever 2
  - reduced to a single string, making management easier
  - removed file that wrote the peers down, no reason to have this in this case, looks like it has no other function, removed dead code
- Make the peers' and owns' address with respect to port or address:port arguments
  - Same functionality as before when only the port is specified
  

### Docker compose users:

I re-implemented some of the addressing (really it was mostly arg parsing) so that the node and its peer
can have independent base addresses. This means they are now docker container compatible.

The folder `docker_compose` shows how this might be used in your `docker-compose.yml` file.
The `docker-compose.yml` expects `existingWallets/` to have two files:
- `alice_wallet`
- `bob_wallet`

WARNING: If the wallet files are missing docker will attempt to bind directories instead! This will also crash. Additionally docker might make them write-protected. Remove them at your own discretion.


Of course you can rename the files as you want if you change their occurrences.

These wallets are copied to the individual container before they start up, they will instantly crash without them.


When binding, in order to retain DB information, both DB files are linked back into a folder for each node. This is mostly due to docker volume limitations as dual binding is not really a good idea (binding the same folder of the host to two different sub-containers), especially since we _want_ separate files.


 
#### Docker compose address resolution
Containers that run in docker and are declared inside a `docker-compose.yml` have a built-in hostname resolution that can make them find other containers in the same service file. They are automatically available to any container. This is what this example files uses to route data between the two node instances.


## Running commands without compiling your own version:
Since you already have a node running inside docker containers, it seems a bit backwards to then compile again just to talk to it.
You can either: Start a shell inside your container and run it from there, which i find quite cumbersome.

alternatively: You can run the commands raw. The commands are specified in `servercmd.go` and their respective endpoint can be found in `conn.go`.
This has the nice effect of letting us craft our own requests.
Here we also see the main security problem, the server is very unprotected, these requests have no kind of confirmation.

One raw curl request might look like this:
```shell
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
        '{
        "To": "1LHroFft5WAxZTqXizQJjBJrPwkVQFAcsa",
        "Amount": 200
        }' \
 'http://localhost:7080/wallet/send'

```

Which read quite easily. Transfer 200 coins to the given public key.

Mining is now a simple get request:
```shell
curl -i -X GET \
 'http://localhost:7080/chain/mining'
 ```



# How to run


## Build

Install deps:


```shell script
go mod tidy
```

build it:

```shell script
go build ./cmd/cli
```

### Create Wallet

```shell script
./cli wallet create -walletname "alice"
./cli wallet create -walletname "bob"
```

### Start two nodes/start server commands

```shell script
./cli server start -nodeaddress 3000 -peernodeaddress 3001 -apiport 8080 -walletname "alice" -ismining=true
./cli server start -nodeaddress 3001 -peernodeaddress 3000 -apiport 8081 -walletname "bob" -ismining=true
```

### Mining empty block to get block reward
```shell script
./cli server miningblock --apiport 8080
```

### Send Transaction to other address
```shell script
./cli server sendtransaction --apiport 8080 --to "172wJyiJZxXWyBW7CYSVddsR5e7ZMxtja9" -amount 100000
```

There are still have other blockchain command, you can find out by type `./cli server`.


Example
------

### Create wallet

```shell script
./cli wallet create -walletname "alice"
```

### Get blocks

 ```shell script
 ./cli server getblocks -apiport 8080
 ```

### Get block hashes

 ```shell script
 ./cli server getblockhashes -apiport 8080
 ```

### Get block height

 ```shell script
 ./cli server getblockheight -apiport 8080
 ```

### Get block utxos

 ```shell script
 ./cli server getutxos -apiport 8080
 ```

### Get wallet address

 ```shell script
 ./cli server getwalletaddress -apiport 8080
 ```

### Get wallet utxos

 ```shell script
 ./cli server getwalletutxos -apiport 8080
 ```

### Get wallet balance

 ```shell script
 ./cli server getwalletbalance -apiport 8080
 ```

### Send transaction

 ```shell script
./cli server sendtransaction --apiport 8080 --to "172wJyiJZxXWyBW7CYSVddsR5e7ZMxtja9" -amount 100000
 ```

### Mining block

 ```shell script
 ./cli server miningblock -apiport 8080
 ```


