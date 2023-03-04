#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>
#include <vector>

using namespace std;

int N;
vector<int> D;

int main() {
    // 入力
    cin >> N;
    int d;
    for (int i = 0; i < N; i++) {
        cin >> d;
        D.push_back(d);
    };

    // 探索
    sort(D.begin(), D.end());
    int Answer = 1;
    for (int i = 1; i < N; i++) {
        if (D[i - 1] < D[i]) {
            Answer += 1;
        }
    }

    // 出力
    cout << Answer << endl;
    return 0;
}
