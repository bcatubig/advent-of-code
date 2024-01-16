#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
int main(int argc, char** argv) {
    if (argc != 2) {
        fprintf(stderr, "Usage: %s <file>\n", argv[0]);
        return EXIT_FAILURE;
    };

    printf("loading: %s\n", argv[1]);
    FILE* in_file = fopen(argv[1], "r");
    if (in_file == NULL) {
        perror("fopen");
        return EXIT_FAILURE;
    };

    char* buf = NULL;
    size_t buf_size = 0;
    ssize_t nread;

    while ((nread = getline(&buf, &buf_size, in_file)) != -1) {
        printf("%s", buf);
        for (int i = 0; i < nread - 1; i++) {
            printf("%c - isNum: %d\n", buf[i], isdigit(buf[i]) != 0);
        }
    }

    fclose(in_file);
    return EXIT_SUCCESS;
}
