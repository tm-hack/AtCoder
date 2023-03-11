#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

#define MAX_N 200'009

int N;
int A[MAX_N];
bool B[MAX_N];

int main() {
    // 入力
    cin >> N;
    for (int i = 1; i <= N; i++) cin >> A[i];

    // 初期化
    for (int i = 1; i <= N; i++) B[i] = false;

    // 計算
    for (int i = 1; i <= N; i++)
        if (!B[i]) B[A[i]] = true;

    int Answer = 0;
    for (int i = 1; i <= N; i++)
        if (!B[i]) Answer += 1;

    // 出力
    cout << Answer << endl;
    for (int i = 1; i <= N; i++) {
        if (!B[i]) cout << i << " ";
    }
    cout << endl;
    return 0;
}
