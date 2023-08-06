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
// build facet array for the find_surface function
void facets(
        long long int ndim,
        long long int nelem,
        long long int nfacetsperelem,
        long long int nnodesperfacet,
        long long int npointsperelem,
        long long int *connectivity,
        long long int *facets,
        long long int *faces){

    long long int i, j, k, l;

     for (i=0; i<nelem; i++) {
         for (j=0; j<nfacetsperelem; j++) {
             l = i * nfacetsperelem + j;
             for (k=0; k<nnodesperfacet; k++) {
                 faces[l * nnodesperfacet + k] =
                     connectivity[i * npointsperelem +
                                  facets[j * nnodesperfacet + k]];
             }
         }
     }
}