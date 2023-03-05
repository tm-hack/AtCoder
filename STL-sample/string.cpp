#include <bits/stdc++.h>

#include <atcoder/all>

using namespace std;
using namespace atcoder;

int main() {
    // 入力した2文字を連結し、最初の10文字を出力する
    string a, b;
    cin >> a >> b;
    string c = a + b;
    if (c.size() <= 10)
        cout << c << endl;
    else
        cout << c.substr(0, 10) << endl;

    // 入力した文字列sの偶数文字目を'z'に変えて出力する。
    string s;
    cin >> s;
    for (int i = 1; i < s.size(); i += 2) s[i] = 'z';
    cout << s << endl;

    return 0;
}