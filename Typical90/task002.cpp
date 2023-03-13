#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

double N;

// 正しいカッコが作れるかどうかを判定する
bool is_kakko(bitset<21> bs1) {
    int kakko[2] = {0};
    for (int i = 0; i < N; i++) {
        kakko[bs1[i]]++;
        if (kakko[0] < kakko[1]) {
            return false;
        }
    }

    if (kakko[0] != kakko[1]) {
        return false;
    } else {
        return true;
    }
}

int main() {
    // 入力
    cin >> N;

    // 探索
    for (int i = pow(2, N) - 1; i >= 0; i--) {
        bitset<21> bs1(i);
        if (is_kakko(bs1)) {
            // 出力
            for (int i = N - 1; i >= 0; i--) {
                if (bs1[i]) {
                    cout << "(";
                } else {
                    cout << ")";
                }
            }
            cout << endl;
        }
    }
    return 0;
}
