#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 200'009

int N, M;

int main() {
    // 入力
    cin >> N >> M;
    vector<int> tye[MAX_N];

    for (int i = 1; i <= M; i++) {
        cin >> score[i].first;
        score[i].second = i;
    }
    sort(score, score + MAX_N, greater<pair<int, int>>());

    // 出力
    cout << "Good Luck!" << endl;
    return 0;
}
