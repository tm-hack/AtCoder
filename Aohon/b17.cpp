#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 100'009

int N;
int h[MAX_N];
int dp[MAX_N];
vector<int> Answer;

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

    // 答えの復元
    int Place = N;
    while (true) {
        Answer.push_back(Place);
        if (Place == 1) break;

        if (dp[Place] == dp[Place - 1] + abs(h[Place - 1] - h[Place])) {
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
