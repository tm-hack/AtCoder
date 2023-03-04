#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>

using namespace std;

int N, K;
int A[100009];
int R[100009];

bool check(int i, int j, int K) {
    if (A[j] - A[i] <= K) {
        return true;
    } else {
        return false;
    }
}

int main() {
    // 入力
    cin >> N >> K;
    for (int i = 1; i <= N; i++) cin >> A[i];

    // 各A[i]に対して二分探索を行う
    for (int i = 1; i <= N - 1; i++) {
        int left = 1;
        int right = N + 1;
        while (right - left > 1) {
            int Mid = (left + right) / 2;
            bool Answer = check(i, Mid, K);
            if (Answer == true) left = Mid;
            if (Answer == false) right = Mid;
        }

        R[i] = left;
    }

    // 出力
    long long Answer = 0;
    for (int i = 1; i <= N - 1; i++) Answer += (R[i] - i);
    cout << Answer << endl;
    return 0;
}
