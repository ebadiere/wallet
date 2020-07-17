# wallet


solc commands:

solc -o ./contracts/abi --abi --base-path /Users/ericbadiere/projects/wallet  ./contracts/token/ERC20/ERC20.sol

solc -o ./contracts/bin --base-path /Users/ericbadiere/projects/wallet  ./contracts/token/ERC20/ERC20.sol

abigen --bin=./contracts/bin/ERC20.bin --abi=./contracts/abi/ERC20.abi -pkg=ERC20 --out=./contracts/token/ERC20/ERC20.go

ToDo
Deploy an WEth version of the Uni1KyberArbitraguer, then complete the arbitrage finder

implement the ethToToken uniswap swap function in an isolated function easy to test

