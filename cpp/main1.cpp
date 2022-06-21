#include <iostream>

using namespace std;

int main() {
    int l, it = 0, a[200100];
    while(cin >> l){
        a[it] = l;
        if (l == NULL) break;
    }

    int n = a[0] - 1;

    for (int i = 1; i < 200100; i++) {
        int count = 1;
        int tmp = 0;
        for (int j = 1; j < 200100; j++) {
            if (i == j) continue;
            if (a[j] > tmp) tmp = a[j];
            if (count++ > n) break;
        }
        cout << tmp;
    }
}
