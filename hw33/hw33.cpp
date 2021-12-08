#include <iostream>

int a = 4;

int triple_num(int a) {
    return a*3;
}
int main() {
    std::cout << triple_num(4) << "\n";
    std::cout << "hello\n";
    int x = 6;
    if (x < 10) {
        std::cout << "true\n";
    }
    for (int i = 0; i < 5; i++) {
        std::cout << i << "\n";
    }
    int scores[4] = {22, 30, 28, 27};
    std::cout << scores[1];
}
