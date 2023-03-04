#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>
#include <string>

using namespace std;

string S;

int main() {
    // 入力
    cin >> S;

    // 判定
    int N = S.length();
    string Answer = "Yes";
    if (S != string(S.rbegin(), S.rend())) Answer = "No";

    string substr1 = S.substr(0, (N - 1) / 2);
    if (substr1 != string(substr1.rbegin(), substr1.rend())) Answer = "No";

    string substr2 = S.substr((N + 3) / 2 - 1);
    if (substr2 != string(substr2.rbegin(), substr2.rend())) Answer = "No";

    // 出力
    cout << Answer << endl;
    return 0;
}
