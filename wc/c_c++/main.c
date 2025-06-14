#include <stdlib.h>
#include <stdio.h>
#include <wchar.h>
#include <locale.h>

#include "ccwc.h"

void print_args(int argc, char const *argv[])
{
    printf("main %d\n", argc);
    for (size_t i = 0; i < argc; i++)
    {
        printf("arg[%ld]: %s, ", i, argv[i]);
    }
    printf("\n");
}

int main(int argc, char const *argv[])
{
    struct ccwc_options options = {
        .l = false,
        .w = false,
        .m = false,
        .c = false};

    int f_path_pos = 1;
    check_args(argc, argv, &options, &f_path_pos);

    if (setlocale(LC_ALL, "") == NULL)
    {
        fprintf(stderr, "Error setting locale.\n");
        exit(EXIT_FAILURE);
    }

    struct ccwc_counts counts = {
        .l = 0,
        .w = 0,
        .m = 0,
        .c = 0,
    };

    struct ccwc_carac_type types = {
        .was_c = false,
        .is_c = false,
    };

    bool already_read = false;
    if (options.c)
    {
        FILE *file = fopen(argv[f_path_pos], "rb");
        if (!file)
        {
            perror("fopen() failed");
            return EXIT_FAILURE;
        }

        int c;
        while ((c = fgetc(file)) != EOF)
        {
            counts.c++;
            count_l_w(c, &options, &counts, &types);
            already_read = true;
        }
        fclose(file);
    }
    if (options.m)
    {
        FILE *file = fopen(argv[f_path_pos], "r");
        if (!file)
        {
            perror("fopen() failed");
            return EXIT_FAILURE;
        }

        wchar_t c;
        while ((c = fgetwc(file)) != WEOF)
        {
            counts.m++;
            if (!already_read)
            {
                count_l_w(c, &options, &counts, &types);
            }
                }
        fclose(file);
    }

    printf("(stats) l: %ld, w: %ld, m: %ld, c: %ld\n",
           counts.l, counts.w, counts.m, counts.c);

    return 0;
}
