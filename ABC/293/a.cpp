#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

string S;

int main() {
    // 入力
    cin >> S;

    for (int i = 0; i < S.size() / 2; i++) {
        swap(S[2 * i], S[2 * i + 1]);
    }

    // 出力
    cout << S << endl;
    return 0;
}
