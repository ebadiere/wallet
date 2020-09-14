# wallet


solc commands:

solc -o ./contracts/abi --abi --base-path /Users/ericbadiere/projects/wallet  ./contracts/token/ERC20/ERC20.sol

solc -o ./contracts/bin --base-path /Users/ericbadiere/projects/wallet  ./contracts/token/ERC20/ERC20.sol

abigen --bin=./contracts/bin/ERC20.bin --abi=./contracts/abi/ERC20.abi -pkg=ERC20 --out=./contracts/token/ERC20/ERC20.go

ToDo
Deploy an WEth version of the Uni1KyberArbitraguer, then complete the arbitrage finder

implement the ethToToken uniswap swap function in an isolated function easy to test


EthLendOracle:

Gas Price:  90000000000

Gas Price:  99000000000
0x517B5964dB0f8D8Beb40cd6894e0fA9eBEbeeAf0
0x3cc5c913f0962ba6f46e5b57ac98bfdd11a32f783e51f5b25147da070911f2c5


