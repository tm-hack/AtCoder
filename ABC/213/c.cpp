#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 200'009

bool compare_by_second(pair<int, int> x, pair<int, int> y) {
    if (x.second != y.second) {
        return x.second < y.second;
    } else {
        return x.first < y.first;
    }
}

int main() {
    // 入力
    int H, W, N;
    cin >> H >> W >> N;
    pair<int, int> aa[N + 1];
    pair<int, int> bb[N + 1];

    for (int i = 1; i <= N; i++) {
        int a, b;
        cin >> a >> b;
        aa[i].first = a;
        aa[i].second = i;
        bb[i].first = b;
        bb[i].second = i;
    }

    // 操作
    int cnt_a = 1;
    sort(aa, aa + N);
    for (int i = 1; i <= H; i++) {
        aa[i].first = cnt_a;

        if (i > 1 && aa[i] != aa[i - 1]) {
            cnt_a += 1;
        }
    }

    int cnt_b = 1;
    sort(bb, bb + N);
    for (int i = 1; i <= W; i++) {
        bb[i].first = cnt_b;

        if (i > 1 && bb[i] != bb[i - 1]) {
            cnt_b += 1;
        }
    }

    sort(aa, aa + N, compare_by_second);
    sort(bb, bb + N, compare_by_second);

    // 出力
    for (int i = 1; i <= N; i++) {
        cout << aa[i].first << " " << bb[i].first << endl;
    }

    return 0;
}
