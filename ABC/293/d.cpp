#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int N, M;

int main() {
    // 入力
    cin >> N >> M;

    vector<int> u(M), v(M);
    for (int i = 0; i < M; i++) {
        string a, b;
        cin >> u[i] >> a >> v[i] >> b;
        u[i]--;
        v[i]--;
    }

    // グラフ生成
    dsu ufd(N);

    // 計算
    int Answer_1 = 0;
    for (int i = 0; i < M; i++) {
        if (ufd.leader(u[i]) == ufd.leader(v[i])) Answer_1++;
        ufd.merge(u[i], v[i]);
    }
    int Answer_2 = ufd.groups().size() - Answer_1;

    // 出力
    cout << Answer_1 << " " << Answer_2 << endl;
    return 0;
}
