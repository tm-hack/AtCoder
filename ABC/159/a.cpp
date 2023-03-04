#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>
#include <vector>

using namespace std;

int N, M;

int main() {
    // 入力
    cin >> N >> M;

    // 計算
    int Answer_N = (N >= 2) ? N * (N - 1) / 2 : 0;
    int Answer_M = (M >= 2) ? M * (M - 1) / 2 : 0;

    // 出力
    cout << Answer_N + Answer_M << endl;
    return 0;
}
