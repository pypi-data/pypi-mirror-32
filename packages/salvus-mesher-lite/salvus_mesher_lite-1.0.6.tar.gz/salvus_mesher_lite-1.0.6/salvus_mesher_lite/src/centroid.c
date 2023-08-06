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
// compute the element centroid
void centroid(
        long long int ndim,
        long long int nelem,
        long long int npointsperelem,
        long long int *connectivity,
        double* points,
        double* centroid){

    long long int i, j, k, idx;
    double c;

    #pragma omp parallel for private(i, j, k, idx, c)
    for (i=0; i<nelem; i++) {
        for (j=0; j<ndim; j++) {
            c = 0.;
            for (k=0; k<npointsperelem; k++) {
                idx = connectivity[i * npointsperelem + k];
                c = c + points[idx * ndim + j];
            }
            centroid[i * ndim + j] = c / npointsperelem;
        }
    }
}