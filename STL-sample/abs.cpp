#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

// abs(x)：xの絶対値を返す
int main() {
    // 入力
    double a, b;
    cin >> a >> b;

    // 出力
    printf("%.12lf\n", abs(a - b));
    return 0;
}
