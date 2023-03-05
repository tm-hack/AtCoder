#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int main() {
    // 入力
    long N, M;
    cin >> N >> M;

    vector<int> u(M), v(M);
    for (int i = 0; i < M; i++) {
        cin >> u[i] >> v[i];
        u[i]--;
        v[i]--;
    }

    // グラフ生成
    dsu ufd(N);
    for (int i = 0; i < N; i++) {
        ufd.merge(u[i], v[i]);
    }

    // 探索
    // 代表点Lの属するグラフの頂点Vの個数と辺Eの個数を記録する
    vector<int> vs(N), es(N);
    for (int i = 0; i < N; i++) {
        vs[ufd.leader(i)]++;
    }

    for (int i = 0; i < M; i++) {
        es[ufd.leader(u[i])]++;
    }

    // 判定及び出力
    if (vs == es) {
        cout << "Yes" << endl;
    } else {
        cout << "No" << endl;
    }

    return 0;
}
