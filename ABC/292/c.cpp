#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

long N;
long memo[2'00'009];

int main() {
    // 入力
    cin >> N;

    // 計算
    // 掛けてxになる整数の組み合わせの個数をメモする
    for (int i = 1; i <= N; i++) {
        for (int j = 1; j * j <= i; j++) {
            if (i % j == 0) {
                memo[i] += 1;
                if (i / j != j) memo[i] += 1;
            }
        }
    }

    // AB+CD=Nとなる組み合わせの個数を計算する
    long Answer = 0;
    for (int i = 1; i <= N; i++) {
        int AB = i;
        int CD = N - i;
        Answer += memo[AB] * memo[CD];
    }

    // 出力
    cout << Answer << endl;
    return 0;
}
