#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>
#include <string>
#include <string_view>
#include <vector>

using namespace std;

string str;

int main() {
    // 入力
    cin >> str;

    reverse(str.begin(), str.end());

    string divides[4] = {"dream", "dreamer", "erase", "eraser"};
    for (string &d : divides) {
        reverse(d.begin(), d.end());
    }

    // 探索
    while (!str.empty()) {
        bool found = false;
        for (string &d : divides) {
            if (str.substr(0, d.size()) == d) {
                str = str.substr(d.size(), str.size() - d.size());
                found = true;
                break;
            }
        }
        if (!found) {
            cout << "NO" << endl;
            return 0;
        }
    }

    // 出力
    cout << "YES" << endl;
    return 0;
}
