pragma circom 2.1.6;

include "./circomlib/circuits/poseidon.circom";

template BalanceHash() {
   signal input balance;
   signal input balanceSecret;
   signal input address;
   signal output hash;

   component p2k;
   p2k = Poseidon(2);
   p2k.inputs[0] <== balance;
   p2k.inputs[1] <== balanceSecret;

   hash <== p2k.out;
}


component main {public [ balance,address ]}= BalanceHash();