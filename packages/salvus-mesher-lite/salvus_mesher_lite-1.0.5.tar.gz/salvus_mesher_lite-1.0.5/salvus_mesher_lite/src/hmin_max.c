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
#include <math.h>

// find elementwise the shortest edge
void hmin(
        long long int ndim,
        long long int nelem,
        long long int npoints,
        long long int nedges,
        long long int npointsperelem,
        long long int *connectivity,
        long long int *edges,
        double* points,
        double* hmin){

    long long int i, j, k, idx1, idx2;
    double dist;

    #pragma omp parallel for private(i, j, k, idx1, idx2, dist)
    for (i=0; i<nelem; i++) {
        for (j=0; j<nedges; j++) {
            dist = 0.;
            idx1 = connectivity[i * npointsperelem + edges[j * 2 + 0]];
            idx2 = connectivity[i * npointsperelem + edges[j * 2 + 1]];
            for (k=0; k<ndim; k++) {
                dist = dist + pow(points[idx1 * ndim + k] - points[idx2 * ndim + k], 2);
            }

            dist = sqrt(dist);
            hmin[i] = fmin(hmin[i], dist);
        }
    }
}

// find elementwise the longest edge
void hmax(
        long long int ndim,
        long long int nelem,
        long long int npoints,
        long long int nedges,
        long long int npointsperelem,
        long long int *connectivity,
        long long int *edges,
        double* points,
        double* hmax){

    long long int i, j, k, idx1, idx2;
    double dist;

    #pragma omp parallel for private(i, j, k, idx1, idx2, dist)
    for (i=0; i<nelem; i++) {
        for (j=0; j<nedges; j++) {
            dist = 0.;
            idx1 = connectivity[i * npointsperelem + edges[j * 2 + 0]];
            idx2 = connectivity[i * npointsperelem + edges[j * 2 + 1]];
            for (k=0; k<ndim; k++) {
                dist = dist + pow(points[idx1 * ndim + k] - points[idx2 * ndim + k], 2);
            }

            dist = sqrt(dist);
            hmax[i] = fmax(hmax[i], dist);
        }
    }
}