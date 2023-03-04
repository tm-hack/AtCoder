#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>

using namespace std;

int N;
int A[209];

int main() {
    // 入力
    cin >> N;
    for (int i = 1; i <= N; i++) cin >> A[i];

    int Answer = 0;
    while (true) {
        bool flag = true;
        for (int i = 1; i <= N; i++) {
            if (A[i] % 2 == 1) flag = false;
            if (A[i] % 2 == 0) A[i] = A[i] / 2;
        }
        if (!flag) break;
        Answer += 1;
    }

    // 出力
    cout << Answer << endl;
}
