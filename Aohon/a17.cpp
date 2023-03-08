#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 100'009

int N, A[MAX_N], B[MAX_N];
int dp[MAX_N];
vector<int> Answer;

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

    // 答えの復元
    int Place = N;
    while (true) {
        Answer.push_back(Place);
        if (Place == 1) break;

        if (dp[Place - 1] + A[Place] == dp[Place]) {
            Place = Place - 1;
        } else {
            Place = Place - 2;
        }
    }

    reverse(Answer.begin(), Answer.end());

    // 出力
    cout << Answer.size() << endl;
    for (int i = 0; i < Answer.size(); i++) {
        if (i >= 1) cout << " ";
        cout << Answer[i];
    }
    cout << endl;
    return 0;
}
