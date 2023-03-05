#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int main() {
    // 入力
    long long N;
    cin >> N;

    // 計算
    vector<char> lower_char;
    lower_char.push_back('z');

    for (int i = 0; i < 25; i++) {
        char c = 'a' + i;
        lower_char.push_back(c);
    }

    long long a = N;
    string Answer = "";
    while (a > 0) {
        Answer.push_back(lower_char[a % 26]);
        if (a % 26 == 0) {
            a -= 1;
        }
        a = a / 26;
    };
    reverse(Answer.begin(), Answer.end());

    // 出力
    cout << Answer << endl;
    return 0;
}
