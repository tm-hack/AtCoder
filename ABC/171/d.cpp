#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 100'009
long long A[MAX_N];

int main() {
    // 入力
    long N;
    cin >> N;
    for (long i = 0; i < N; i++) {
        long a;
        cin >> a;
        A[a]++;
    }

    long long sum_score = 0;
    for (long i = 0; i < MAX_N; i++) {
        sum_score += i * A[i];
    }

    // オペレーション
    long Q;
    cin >> Q;
    for (long i = 0; i < Q; i++) {
        long B, C;
        cin >> B >> C;
        sum_score = sum_score - (B - C) * A[B];

        // 出力
        cout << sum_score << endl;

        // 後処理
        A[C] += A[B];
        A[B] = 0;
    }

    return 0;
}
