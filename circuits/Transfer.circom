pragma circom 2.1.6;

include "./circomlib/circuits/poseidon.circom";
template Balance2Hash() {
    signal input balance;
    signal input r;
    signal output hash;

    component p2k;
    p2k = Poseidon(2);
    p2k.inputs[0] <== balance;
    p2k.inputs[1] <== r;

    hash <== p2k.out;

}


template Transfer() {
    signal input balance;
    signal input balanceSecret;

    signal input value;
    signal input valueSecret;

    // add:1 ; sub:0
    signal input add;

    signal newBalance;

    signal output balanceHash;
    signal output valueHash;
    signal output newBalanceHash;

    add * (1 - add) === 0;
    assert(balance >= 0 );
    assert(value >= 0 );

    component bh = Balance2Hash();
    bh.balance <== balance;
    bh.r <== balanceSecret;
    balanceHash <== bh.hash;

    component vh = Balance2Hash();
    vh.balance <== value;
    vh.r <== valueSecret;
    valueHash <== vh.hash;

    newBalance <==  balance - value + add * ( value + value );
    // balance - value *(1-add) + value * add
    // b - v + add * v + add * v
    // b - v + add * (v +v )

    assert(newBalance >=0 );

    component nh =Balance2Hash();
    nh.balance <== newBalance;
    nh.r <== balanceSecret;
    newBalanceHash <== nh.hash;

}

component main {public [add]}= Transfer();

