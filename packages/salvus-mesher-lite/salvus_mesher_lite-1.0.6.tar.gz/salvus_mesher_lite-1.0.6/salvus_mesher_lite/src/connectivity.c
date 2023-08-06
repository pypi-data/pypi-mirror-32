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
void connectivity_2D(int nelem_x, int nelem_y, long long int connectivity[][4]) {
    long long int i, j, l, a;

    #pragma omp parallel for private(i, j, l, a)
    for (i = 0; i < nelem_y; i++) {
        for (j = 0; j < nelem_x; j++) {
            l = i * nelem_x + j;
            a = j * (nelem_y + 1);

            connectivity[l][0] = a + i;
            connectivity[l][1] = a + nelem_y + i + 1;
            connectivity[l][2] = a + nelem_y + i + 2;
            connectivity[l][3] = a + i + 1;
        }
    }
}