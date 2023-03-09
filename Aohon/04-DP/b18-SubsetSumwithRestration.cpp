#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 69
#define MAX_S 10009

int N, S, A[MAX_N];
bool dp[MAX_N][MAX_S];
vector<int> Answer;

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

    // 答えの復元
    int score = S;
    for (int i = N; i >= 1; i--) {
        if (score == 0) break;

        if (dp[i][score] == true && score - A[i] >= 0 &&
            dp[i - 1][score - A[i]] == true) {
            Answer.push_back(i);
            score = score - A[i];
        }
    }

    // 出力
    if (Answer.size() == 0) {
        cout << "-1" << endl;
    } else {
        cout << Answer.size() << endl;
        for (int i = 0; i < Answer.size(); i++) {
            if (i >= 1) {
                cout << " ";
            }
            cout << Answer[i];
        }
        cout << endl;
    }

    return 0;
}
