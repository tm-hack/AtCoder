#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 69
#define MAX_S 10009

int N, S, A[MAX_N];
bool dp[MAX_N][MAX_S];

int main() {
    // 入力
    int N;
    cin >> N >> S;
    for (int i = 1; i <= N; i++) cin >> A[i];

    // 動的計画法（i = 0）
    dp[0][0] = true;
    for (int i = 1; i <= S; i++) dp[0][i] = false;

    // 動的計画法（i >= 1）
    for (int i = 1; i <= N; i++) {
        for (int j = 0; j <= S; j++) {
            if (j < A[i]) {
                if (dp[i - 1][j] == true)
                    dp[i][j] = true;
                else
                    dp[i][j] = false;
            }

            if (j >= A[i]) {
                if (dp[i - 1][j] == true || dp[i - 1][j - A[i]] == true)
                    dp[i][j] = true;
                else
                    dp[i][j] = false;
            }
        }
    }

    // 出力
    if (dp[N][S] == true)
        cout << "Yes" << endl;
    else
        cout << "No" << endl;
    return 0;
}
