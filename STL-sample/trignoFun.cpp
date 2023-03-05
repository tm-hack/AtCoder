#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int main() {
    // sin(x): sin(x)の値をdouble型で返す。
    // cos(x): cos(x)の値をdouble型で返す。
    // tan(x): cos(x)の値をdouble型で返す。

    // 入力
    double x;
    cin >> x;

    // 出力
    printf("%.12lf\n", sin(x / 180.0 * M_PI));
    printf("%.12lf\n", cos(x / 180.0 * M_PI));
    printf("%.12lf\n", tan(x / 180.0 * M_PI));

    // asin(x): arc-sin(x)の値をdouble型で返す
    // acos(x): arc-cos(x)の値をdouble型で返す
    // atan(x): arc-tan(x)の値をdouble型で返す

    // 入力
    double y;
    cin >> y;

    // 出力
    printf("%.12lf\n", asin(y) * 180.0 / M_PI);
    printf("%.12lf\n", acos(y) * 180.0 / M_PI);
    printf("%.12lf\n", atan(y) * 180.0 / M_PI);

    return 0;
}
