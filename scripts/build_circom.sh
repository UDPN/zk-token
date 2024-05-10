#!/usr/bin bash

mkdir -p build/circuits

# build BalanceHash.circom
circom circuits/BalanceHash.circom --r1cs --wasm -o ./build/circuits
snarkjs groth16 setup ./build/circuits/BalanceHash.r1cs ./circuits/ptau/powersOfTau28_hez_final_12.ptau ./build/circuits/BalanceHash.zkey
snarkjs zkey export verificationkey ./build/circuits/BalanceHash.zkey ./build/circuits/BalanceHash_verification_key.json
snarkjs zkey export solidityverifier ./build/circuits/BalanceHash.zkey ./contracts/BalanceVerifier.sol
solcjs ./contracts/BalanceVerifier.sol --bin  --abi --optimize -o ./build/contracts/bin
./abigen/abigen --bin ./build/contracts/bin/contracts_BalanceVerifier_sol_Groth16Verifier.bin --abi ./build/contracts/bin/contracts_BalanceVerifier_sol_Groth16Verifier.abi --pkg contracts --type BalanceVerifier --out ./core/contracts/BalanceVerifier.go


#build Transfer.circom
circom circuits/Transfer.circom --r1cs --wasm -o ./build/circuits
snarkjs groth16 setup ./build/circuits/Transfer.r1cs ./circuits/ptau/powersOfTau28_hez_final_12.ptau ./build/circuits/Transfer.zkey
snarkjs zkey export verificationkey ./build/circuits/Transfer.zkey ./build/circuits/Transfer_verification_key.json
snarkjs zkey export solidityverifier ./build/circuits/Transfer.zkey ./contracts/TransferVerifier.sol
solcjs ./contracts/TransferVerifier.sol --bin  --abi --optimize -o ./build/contracts/bin
./abigen/abigen --bin ./build/contracts/bin/contracts_TransferVerifier_sol_Groth16Verifier.bin --abi ./build/contracts/bin/contracts_TransferVerifier_sol_Groth16Verifier.abi --pkg contracts --type TransferVerifier --out ./core/contracts/TransferVerifier.go

solcjs ./contracts/HideToken.sol --bin  --abi --optimize -o ./build/contracts/bin
./abigen/abigen --bin ./build/contracts/bin/contracts_HideToken_sol_HideToken.bin --abi ./build/contracts/bin/contracts_HideToken_sol_HideToken.abi --pkg contracts --type HideToken --out ./core/contracts/HideToken.go
