#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>

using namespace std;

string S;

int main() {
    cin >> S;

    int Answer = 0;
    for (int i = 0; i < S.size(); i++) {
        if (S.at(i) == '1') Answer += 1;
    }

    cout << Answer << endl;
    return 0;
}
