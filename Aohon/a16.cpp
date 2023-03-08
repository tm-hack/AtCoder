#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 100'009

int N, A[MAX_N], B[MAX_N];
int dp[MAX_N];

int main() {
    // 入力
    cin >> N;
    for (int i = 2; i <= N; i++) cin >> A[i];
    for (int i = 3; i <= N; i++) cin >> B[i];

    // 動的計画法
    dp[1] = 0;
    dp[2] = A[2];
    for (int i = 3; i <= N; i++) {
        dp[i] = min(dp[i - 1] + A[i], dp[i - 2] + B[i]);
    }

    // 出力
    cout << dp[N] << endl;
    return 0;
}
