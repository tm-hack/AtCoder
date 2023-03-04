#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>

using namespace std;

double N;

bool check(double x, double y) {
    if (pow(x, 3) + x <= y) {
        return true;
    } else {
        return false;
    }
}

int main() {
    // 入力
    cin >> N;

    // 二分探索を行う
    double left = 0;
    double right = N;
    while (right - left > 1e-4) {
        double Mid = (left + right) / 2;
        bool Answer = check(Mid, N);
        if (Answer == true) left = Mid;
        if (Answer == false) right = Mid;
    };

    cout << left << endl;
    return 0;
}
