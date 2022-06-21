#include <iostream>

using namespace std;

int main() {
    int treeN, rangeD;
    cin >> treeN >> rangeD;

    int range = rangeD * 2 + 1;

    int i;
    for (i = 1; ; ++i) {
        treeN -= range;
        if (treeN <= 0) break;
    }

    cout << i << endl;
}
