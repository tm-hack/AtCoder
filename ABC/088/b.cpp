#include <math.h>

#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

int N;
vector<int> A;

int main() {
    // 入力
    cin >> N;
    int a;
    for (int i = 0; i < N; i++) {
        cin >> a;
        A.push_back(a);
    }

    // ゲームを実施
    sort(A.begin(), A.end());

    int cnt = 1;
    int score_A = 0;
    int score_B = 0;
    while (!A.empty()) {
        if (cnt % 2 == 1) score_A += A.back();
        if (cnt % 2 == 0) score_B += A.back();
        A.pop_back();
        cnt++;
    }

    // 出力
    cout << (score_A - score_B) << endl;
    return 0;
}
