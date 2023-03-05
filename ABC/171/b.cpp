#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int main() {
    // 入力
    int N, K;
    cin >> N >> K;
    vector<int> p(N);
    for (int i = 0; i < N; i++) {
        cin >> p[i];
    }

    // 計算
    sort(p.begin(), p.end());
    int Answer = 0;
    for (int i = 0; i < K; i++) {
        Answer += p[i];
    }

    // 出力
    cout << Answer << endl;
    return 0;
}
