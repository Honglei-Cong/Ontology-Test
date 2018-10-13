
#include <stdlib.h>
#include <stdio.h>

extern void rust_capitalize(char *);

int main() {
        char str[] = "test hello";
        rust_capitalize(str);
        printf("%s\n", str);
        return 0;
}

