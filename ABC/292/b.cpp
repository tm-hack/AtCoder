#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>
#include <vector>

using namespace std;

int N, Q;
int player[109];

int main() {
    // 入力
    cin >> N >> Q;

    for (int i = 1; i <= N; i++) player[i] = 0;

    // イベント解析
    for (int i = 0; i < Q; i++) {
        int event, no;
        cin >> event >> no;

        if (event == 1) {
            player[no] += 1;
        }

        if (event == 2) {
            player[no] += 2;
        }

        if (event == 3) {
            if (player[no] >= 2) {
                cout << "Yes" << endl;
            } else {
                cout << "No" << endl;
            }
        }
    }

    return 0;
}
