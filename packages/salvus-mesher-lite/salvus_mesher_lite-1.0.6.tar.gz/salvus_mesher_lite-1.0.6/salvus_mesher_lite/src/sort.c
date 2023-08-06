// This file is part of the lite version of the SalvusMesher package intended
// to produce meshes for AxiSEM3D. If you are looking for the full version
// head over to http://mondaic.com.
//
// :copyright:
//     Copyright (C) 2016-2018 Salvus Development Team <www.mondaic.com>,
//                             ETH Zurich
// :license:
//     GNU General Public License, Version 3 [academic use only]
//     (http://www.gnu.org/copyleft/gpl.html)
#include <stdlib.h>


// Custom struct to be able to sort array but retain indices at the same time.
typedef struct {
    double value;
    size_t index;
} value_index;


int compare_value_index(const void *vi1, const void *vi2)
{
    const value_index *_vi1 = vi1;
    const value_index *_vi2 = vi2;
    if (_vi1->value < _vi2->value)
        return -1;
    else if (_vi1->value > _vi2->value)
        return +1;
    else
        return 0;
}


void lexsort_internal_loop(
        long long int dim_counts,
        long long int dim,
        long long int points_count,
        double* points,
        long long int* loc,
        long long int* segstart,
        long long int* segend,
        long long int segcount) {

    int cur_temp_count;
    long long int i, j, k, s, e, len;
    value_index* temp_array;
    double* temp_double;
    long long int* temp_long_long;

    // Allocate temp arrays. Will be reallocated if deemed too small down the
    // line.
    cur_temp_count = 2048;
    temp_array = (value_index *) malloc(cur_temp_count * sizeof(value_index));
    temp_double = (double *) malloc(cur_temp_count * sizeof(double));
    temp_long_long = (long long int *) malloc(cur_temp_count * sizeof(long long int));

    for (i=0; i<segcount; i++) {
        s = segstart[i];
        e = segend[i];
        len = e - s;

        // Reallocate temp arrays if necessary.
        if (len > cur_temp_count) {
            cur_temp_count = len + 1024;
            temp_array = (value_index *) realloc(temp_array, cur_temp_count * sizeof(value_index));
            temp_double = (double *) realloc(temp_double, cur_temp_count * sizeof(double));
            temp_long_long = (long long int *) realloc(temp_long_long, cur_temp_count * sizeof(long long int));
        }

        // Get all points that matter.
        for (j=0; j<len; j++) {
            temp_array[j].value = points[dim * points_count + s + j];
            temp_array[j].index = j;
        }

        // Sort.
        qsort(temp_array, len, sizeof(value_index), compare_value_index);

        for (k=0; k<dim_counts; k++) {
            for (j=0; j<len; j++) {
                temp_double[j] = points[k * points_count + s + temp_array[j].index];
            }
            for (j=0; j<len; j++) {
                points[k * points_count + s + j] = temp_double[j];
            }
        }

        for (j=0; j<len; j++) {
            temp_long_long[j] = loc[s + temp_array[j].index];
        }
        for (j=0; j<len; j++) {
            loc[s + j] = temp_long_long[j];
        }
    }
    free(temp_array);
    free(temp_double);
    free(temp_long_long);
}