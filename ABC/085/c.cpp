#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>
#include <vector>

using namespace std;

int N;
long Y;
vector<int> A;

int main() {
    // 入力
    cin >> N >> Y;

    // 探索
    long x = -1;
    long y = -1;
    long z = -1;
    for (int i = 0; i <= N; i++) {
        for (int j = 0; j <= N - i; j++) {
            if (10000 * i + 5000 * j + 1000 * (N - i - j) == Y) {
                x = i;
                y = j;
                z = N - i - j;
            }
        }
    }

    // 出力
    cout << x << " " << y << " " << z << endl;
    return 0;
}
