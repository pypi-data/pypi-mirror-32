# SalvusMesher Lite

This package contains the GPLv3 licensed lite version of the SalvusMesher
package. It is capable of generating meshes with all the bells and whistles
for usage in AxiSEM3D. If you are looking for the full version of the mesher
head over to http://mondaic.com.

## Installation

Trivial with `conda`:

```bash
$ conda install -c conda-forge salvus_mesher_lite
```

## Usage

Two steps:

```bash
$ python -m salvus_mesher_lite.interface AxiSEM --save_yaml=input.yaml
```

This will create a self-explanatory `input.yaml` YAML file. Edit it to your
satisfaction and run

```bash
$ python -m salvus_mesher_lite.interface --input_file=input.yaml
```