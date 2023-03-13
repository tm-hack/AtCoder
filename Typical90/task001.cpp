#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int N, L, K;
int A[100'009];

// ansで分割できるかどうかを返す
bool is_split(int ans) {
    int idx_0 = 0;
    int idx_1 = 0;
    int cnt = 0;

    while (cnt < K && idx_1 < N) {
        idx_1++;
        if (A[idx_1] - A[idx_0] >= ans) {
            idx_0 = idx_1;
            cnt++;
        }
    }

    if (cnt == K && L - A[idx_0] >= ans) {
        return true;
    } else {
        return false;
    }
}

int main() {
    // 入力
    cin >> N >> L >> K;
    for (int i = 1; i <= N; i++) cin >> A[i];

    // 答えで二分探索
    int left = 0;
    int right = L;
    while (right - left > 1) {
        if (is_split((left + right) / 2)) {
            left = (left + right) / 2;
        } else {
            right = (left + right) / 2;
        }
    }

    // 出力
    cout << left << endl;
    return 0;
}
