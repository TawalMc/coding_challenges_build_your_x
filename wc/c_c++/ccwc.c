#include "ccwc.h"
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void check_args(int argc, char const *argv[], struct ccwc_options *options, int *f_path_pos)
{
    if (argc < 2)
    {
        fprintf(stderr, "file arg missed\n");
        exit(EXIT_FAILURE);
    }

    for (size_t i = 1; i < argc; i++)
    {
        if (argv[i][0] == '-')
        {
            if (strcmp(argv[i], "-l") == 0)
            {
                options->l = true;
            }
            else if (strcmp(argv[i], "-w") == 0)
            {
                options->w = true;
            }
            else if (strcmp(argv[i], "-m") == 0)
            {
                options->m = true;
            }
            else if (strcmp(argv[i], "-c") == 0)
            {
                options->c = true;
            }
            else
            {
                fprintf(stderr, "invalid option %s\n", argv[i]);
                exit(EXIT_FAILURE);
            }
        }
        else
        {
            *f_path_pos = i;
        }
    }
    if (!options->l && !options->w && !options->m && !options->c)
    {
        *options = (struct ccwc_options){
            .l = true,
            .w = true,
            .m = true,
            .c = true};
    }
}

void count_l_w(int c, struct ccwc_options const *options, struct ccwc_counts *counts, struct ccwc_carac_type *types)
{
    if (options->l && c == '\n')
        counts->l++;

    if (options->w && c)
    {
        types->was_c = types->is_c;
        if (isspace(c))
            types->is_c = false;
        else
            types->is_c = true;
        if ((types->was_c != types->is_c) && types->was_c)
            (counts->w)++;
    }
}