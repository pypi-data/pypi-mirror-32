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
#include <math.h>

// compute the angle in an element corner
// using this stable formula:
//     angle = atan2(norm(cross(a,b)), dot(a,b))
void angle(
        long long int ndim,
        long long int nelem,
        long long int npointsperelem,
        long long int node1,
        long long int node2,
        long long int node3,
        long long int *connectivity,
        double* points,
        double* angle){

    long long int i, idx1, idx2, idx3;
    double vecax, vecay, vecaz;
    double vecbx, vecby, vecbz;
    double vx, vy, vz, a, b;

    if (ndim == 3) {
    }
    else if (ndim == 2) {
        #pragma omp parallel for private(idx1, idx2, idx3, vecax, vecay, \
                                         vecaz, vecbx, vecby, vecbz, vx, vy, \
                                         vz, a, b)
        for (i=0; i<nelem; i++) {
            idx1 = connectivity[i * npointsperelem + node1];
            idx2 = connectivity[i * npointsperelem + node2];
            idx3 = connectivity[i * npointsperelem + node3];

            vecax = points[idx1 * ndim + 0] - points[idx2 * ndim + 0];
            vecay = points[idx1 * ndim + 1] - points[idx2 * ndim + 1];

            vecbx = points[idx3 * ndim + 0] - points[idx2 * ndim + 0];
            vecby = points[idx3 * ndim + 1] - points[idx2 * ndim + 1];

            // cross product
            vz = vecax * vecby - vecay * vecbx;

            // norm
            a = fabs(vz);

            // dot product
            b = vecax * vecbx + vecay * vecby;

            // atan2
            angle[i] = atan2(a, b);
        }
    }
    else exit(1);
}