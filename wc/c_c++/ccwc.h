#pragma once
#include <stdbool.h>
#include <ctype.h>

struct ccwc_options
{
    bool l;
    bool w;
    bool m;
    bool c;
};

struct ccwc_counts
{
    long l;
    long w;
    long m;
    long c;
};

struct ccwc_carac_type
{
    bool was_c;
    bool is_c;
};

void check_args(int argc, char const *argv[], struct ccwc_options *args, int *file_path_pos);
void count_l_w(int c, struct ccwc_options const *options, struct ccwc_counts *counts, struct ccwc_carac_type *types);