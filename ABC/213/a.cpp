#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int main() {
    // ε₯ε
    int A, B;
    cin >> A >> B;

    bitset<8> bs_A(A);
    bitset<8> bs_B(B);

    // εΊε
    cout << (bs_A ^ bs_B).to_ullong() << endl;
    return 0;
}
