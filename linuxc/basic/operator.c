#include <stdio.h>

int main() {
    printf("%d\n", 5 / 3);  // 1
    printf("%d\n", 5 / -3); // -1
    printf("%d\n", -5 / 3); // -1

    printf("%d\n", 5 % 3);  // 2
    printf("%d\n", -5 % 3); // -2
    printf("%d\n", 5 % -3); // 2
    
    return 0;
}
