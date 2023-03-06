#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int main() {
    // 入力
    int A, B;
    cin >> A >> B;

    bitset<8> bs_A(A);
    bitset<8> bs_B(B);

    // 出力
    cout << (bs_A ^ bs_B).to_ullong() << endl;
    return 0;
}
