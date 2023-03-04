#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>
#include <map>
#include <vector>

using namespace std;

int N;
int t[100009];
int x[100009], y[100009];

int main() {
    // 入力
    cin >> N;
    t[0] = 0;
    x[0] = 0;
    y[0] = 0;
    for (int i = 1; i <= N; i++) {
        cin >> t[i] >> x[i] >> y[i];
    }

    // 探索
    string Answer = "Yes";
    for (int i = 1; i <= N; i++) {
        int pt = t[i] - t[i - 1];
        int px = abs(x[i] - x[i - 1]);
        int py = abs(y[i] - y[i - 1]);

        if (pt < px + py) {
            Answer = "No";
        }

        if ((pt - px - py) % 2 == 1) {
            Answer = "No";
        }
    }

    // 出力
    cout << Answer << endl;
    return 0;
};
