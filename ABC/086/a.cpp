#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>

using namespace std;

int a, b;

int main() {
    cin >> a >> b;
    if (a * b % 2 == 0) cout << "Even" << endl;
    if (a * b % 2 == 1) cout << "Odd" << endl;
    return 0;
}
