#include <algorithm>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

int N, X, A[100009];

// 二分探索を行う
// A[i] <= xを満たす最大のiを返す
int search(int x) {
    int L = 1, R = N + 1;
    while (R - L > 1) {
        int M = (L + R) / 2;
        if (A[M] > x) R = M;
        if (A[M] <= x) L = M;
    }
    return L;
}

int main() {
    // 入力
    cin >> N >> X;
    for (int i = 1; i <= N; i++) cin >> A[i];
    sort(A + 1, A + N + 1);

    // 二分探索を行う
    int Answer = search(X);
    cout << Answer << endl;
    return 0;
}
