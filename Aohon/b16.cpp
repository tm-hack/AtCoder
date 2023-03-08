#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 200'009

int N;
int h[MAX_N];
int dp[MAX_N];

int main() {
    // 入力
    int N;
    cin >> N;
    for (int i = 1; i <= N; i++) cin >> h[i];

    // 動的計画法
    dp[1] = 0;
    dp[2] = abs(h[2] - h[1]);
    if (N >= 3) {
        for (int i = 3; i <= N; i++) {
            dp[i] = min(dp[i - 1] + abs(h[i] - h[i - 1]),
                        dp[i - 2] + abs(h[i] - h[i - 2]));
        }
    }

    // 出力
    cout << dp[N] << endl;
    return 0;
}
