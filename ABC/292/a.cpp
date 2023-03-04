#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <cstdlib>
#include <iostream>
#include <string>
#include <vector>

using namespace std;

string S;

int main() {
    // 入力
    cin >> S;

    // 操作
    transform(S.begin(), S.end(), S.begin(), ::toupper);

    // 出力
    cout << S << endl;
    return 0;
}
