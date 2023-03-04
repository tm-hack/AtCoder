#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>

using namespace std;

int N, A, B;

int main() {
    // 入力
    cin >> N >> A >> B;

    // 全探索
    int Answer = 0;
    for (int i = 1; i <= N; i++) {
        int digit_sum = 0;
        string str = to_string(i);
        for (int j = 0; j < str.size(); j++) {
            digit_sum += str.at(j) - '0';
        }

        if (digit_sum >= A && digit_sum <= B) {
            Answer += i;
        }
    }

    // 出力
    cout << Answer << endl;
    return 0;
}
