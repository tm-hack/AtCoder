#include <math.h>

#include <algorithm>
#include <atcoder/all>
#include <iostream>

using namespace std;

int A, B, C, X;

int main() {
    // ε₯ε
    cin >> A >> B >> C >> X;

    int Answer = 0;
    for (int i = 0; i <= A; i++) {
        for (int j = 0; j <= B; j++) {
            for (int k = 0; k <= C; k++) {
                if (500 * i + 100 * j + 50 * k == X) Answer += 1;
            }
        }
    }

    // εΊε
    cout << Answer << endl;
    return 0;
}
