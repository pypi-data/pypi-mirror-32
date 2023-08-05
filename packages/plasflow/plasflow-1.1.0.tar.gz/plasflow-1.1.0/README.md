[![Anaconda-Server Badge](https://anaconda.org/smaegol/plasflow/badges/installer/conda.svg)](https://anaconda.org/smaegol/plasflow) [![PyPI version](https://badge.fury.io/py/plasflow.svg)](https://badge.fury.io/py/plasflow)

# PlasFlow 1.1

PlasFlow is a set of scripts used for prediction of plasmid sequences in metagenomic contigs. It relies on the neural network models trained on full genome and plasmid sequences and is able to differentiate between plasmids and chromosomes with accuracy reaching 96%. It outperforms other available solutions for plasmids recovery from metagenomes and incorporates the thresholding which allows for exclusion of incertain predictions. PlasFlow has been published in _Nucleic Acids Research_ (https://doi.org/10.1093/nar/gkx1321).

# Table of contents

- [News](#news)
- [Requirements](#requirements)
- [Installation](#installation)

  - [Conda-based](#conda-based---recommended)
  - [Pip installer](#pip-installer)
  - [Manual installation](#manual-installation)

- [Getting started](#getting-started)

- [Output](#output)

- [Test dataset](#test-dataset)
- [Detailed information](#detailed-information)
- [Citation](#citation)
- [TBD](#tbd)
- [Support](#support)

## News

#### 2018-05-25 Version 1.1 released

New version (1.1) released, which is better suited for large datasets. It can be downloaded from conda and pypi, but the simplest way to upgrade is to replace PlasFlow.py file in you previous installation with the current one.
If you still encounter problems with the new version, try to use smaller numbers for the `--batch_size` option.


## Requirements:

- Python 3.5
- Python packages:

  - Scikit-learn 0.18.1
  - Numpy
  - Pandas
  - [TensorFlow 0.10.0](https://storage.googleapis.com/tensorflow/linux/cpu/tensorflow-0.10.0rc0-cp35-cp35m-linux_x86_64.whl)
  - rpy2 >= 2.7.8
  - scipy
  - biopython

- R 3.25

- R packages:

  - [Biostrings](https://bioconductor.org/packages/release/bioc/html/Biostrings.html)

## Installation

### Conda-based - recommended

Conda is recommended option for installation as it properly resolve all dependencies (including R and Biostrings) and allows for installation without messing with other packages installed. Conda can be used both as the [Anaconda](https://www.anaconda.com/download/), and [Miniconda](https://conda.io/miniconda.html) (which is easier to install and maintain).

After the installation it is required to add [bioconda](https://bioconda.github.io/) channel, required for [Biostrings](https://bioconductor.org/packages/release/bioc/html/Biostrings.html) package installation:

```
conda config --add channels bioconda
```

To exclude the possibility of dependencies conflicts its encouraged to create spearate conda environment for Plasflow using command:

```
conda create --name plasflow python=3.5
```

Python 3.5 is required becuase of TensorFlow requirements.

to activate created environment type:

```
source activate plasflow
```

PlasFlow can be easily installed as an Anaconda package from my Anaconda channel using:

```
conda install plasflow -c smaegol
```

With this command all required dependencies are installed into created conda environment. When installation is finished PlasFlow can be invoked as described in the [Getting started](#getting-started) section.

When you decide to finish your work with PlasFlow, you can simply deactivate current anaconda environment with command:

```
source deactivate
```

### Pip installer

There is a possibility of pip based installation. However, some requirements have to be met:

1. Python 3.5 is required (due to TensorFlow requirements)
2. TensorFlow has to be installed manually:

```
pip install https://storage.googleapis.com/tensorflow/linux/cpu/tensorflow-0.10.0rc0-cp35-cp35m-linux_x86_64.whl
```

then install PlasFlow with

```
pip install plasflow
```

However, models used for prediction have to be downloaded separately (for example using `git clone https://github.com/smaegol/PlasFlow`).

### Manual installation

Of course, PlasFlow repo can be cloned using

```
git clone https://github.com/smaegol/PlasFlow
```

but in that case all dependencies have to be installed manually. TensorFlow can be installed as specified above:

```
pip install https://storage.googleapis.com/tensorflow/linux/cpu/tensorflow-0.10.0rc0-cp35-cp35m-linux_x86_64.whl
```

python dependencies can be installed using pip:

```
pip install numpy pandas scipy rpy2 scikit-learn biopython
```

to install R Biostrings go to <https://bioconductor.org/packages/release/bioc/html/Biostrings.html> and follow instructions therein.

## Getting started

PlasFlow is designed to take a metagenomic assembly and identify contigs which may come from plasmids. It outputs several files, from which the most important is a tabular file containing all predictions (specified with `--output` option).

Prior to the PlasFlow invocation it is highly recommended to filter sequences by length, leaving only those longer than 1000 bp. PlasFlow, similarly to other kmer-based methods, does not perform well on short sequences, as it is hard to get proper kmer coverage from them. Hence, results for short sequences are unreliable. As metagenomic assemblies usually contain large number of short contigs additional filtering test can improve results and speed up the PlasFlow. It can also prevent too high RAM usage.

To filter sequences using provided Perl script type:

```
filter_sequences_by_length.pl -input input_dataset.fasta -output filtered_output.fasta -thresh sequence_length_threshold
```

where sequence length threshold have to be provided in base pairs. Filtered fasta file can be then used directly for PlasFlow prediction.

Due to a high RAM usage during the kmer frequencines calculation it is suggested to run calculations on data sets not larger than 50 000 sequences. We are working to improve this in the next release.

Options available in PlasFlow include:

- `--input` - specifies input fasta file with assembly contigs to classify [required]
- `--output` - a name of the tsv file with the tabular output of classification [required]
- `--threshold` - manually specified threshold for probability filtering (default = 0.7)
- `--labels` - manually specified custom location of labels file (used for translation from numeric output to actual class names)
- `--models` - custom location of models used for prediction (have to be specified if PlasFlow was installed using pip)
- `--batch_size` - how many sequences can be used in the single batch of kmers frequency calculation


## Output

The most important output of PlasFlow is a tabular file containing all predictions (specified with `--output` option), consiting of several columns including:

contig_id | contig_name | contig_length | id | label | ...
--------- | ----------- | ------------- | -- | ----- | ---


where:

- `contig_id`is an internal id of sequence used for the classification
- `contig_name` is a name of contig used in the classification
- `contig_length` shows the length of a classified sequence
- `id` is an internal id of a produced label (classification)
- `label` is the actual classification
- `...` represents additional columns showing probabilities of assignment to each possible class

Sequences can be classified to 26 classes including: chromosome.Acidobacteria, chromosome.Actinobacteria, chromosome.Bacteroidetes, chromosome.Chlamydiae, chromosome.Chlorobi, chromosome.Chloroflexi, chromosome.Cyanobacteria, chromosome.DeinococcusThermus, chromosome.Firmicutes, chromosome.Fusobacteria, chromosome.Nitrospirae, chromosome.other, chromosome.Planctomycetes, chromosome.Proteobacteria, chromosome.Spirochaetes, chromosome.Tenericutes, chromosome.Thermotogae, chromosome.Verrucomicrobia, plasmid.Actinobacteria, plasmid.Bacteroidetes, plasmid.Chlamydiae, plasmid.Cyanobacteria, plasmid.DeinococcusThermus, plasmid.Firmicutes, plasmid.Fusobacteria, plasmid.other, plasmid.Proteobacteria, plasmid.Spirochaetes.

If the probability of assignment to given class is lower than threshold (default = 0.7) then the sequence is treated as unclassified.

Additionaly, PlasFlow produces fasta files containing input sequences binned to plasmids, chromosomes and unclassified.

## Test dataset

Test dataset is located in the `test` folder (file `Citrobacter_freundii_strain_CAV1321_scaffolds.fasta`). It is the SPAdes 3.9.1 assembly of Citrobacter freundii strain CAV1321 genome (NCBI assembly ID: GCA_001022155.1), which contains 1 chromosome and 9 plasmids. In the same folder the results of classification can be found in the form of tsv file (`Citrobacter_freundii_strain_CAV1321_scaffolds.fasta.PlasFlow.tsv`) and fasta files containing identified bins (`Citrobacter_freundii_strain_CAV1321_scaffolds.fasta.PlasFlow.tsv_chromosomes.fasta`, `Citrobacter_freundii_strain_CAV1321_scaffolds.fasta.PlasFlow.tsv_plasmids.fasta` and `Citrobacter_freundii_strain_CAV1321_scaffolds.fasta.PlasFlow.tsv_unclassified.fasta`).

To invoke PlasFlow on the test dataset please copy the `test/Citrobacter_freundii_strain_CAV1321_scaffolds.fasta` file to you current working directory and type:

```
PlasFlow.py --input Citrobacter_freundii_strain_CAV1321_scaffolds.fasta --output test.plasflow_predictions.tsv --threshold 0.7
```
The predictions will be located in the `test.plasflow_predictions.tsv` file and can be compared to results available in the `test/Citrobacter_freundii_strain_CAV1321_scaffolds.fasta.PlasFlow.tsv`.


## Detailed information

Detailed information concerning the alogrithm and assumptions on which the PlasFlow is based can be found in the publication "_PlasFlow - Predicting Plasmid Sequences in Metagenomic Data Using Genome Signatures_" (_Nucleic Acids Research_, submitted). The flowchart illustrating major steps of training and prediction is shown below

![PlasFlow Flowchart](https://github.com/smaegol/PlasFlow/blob/master/flowchart.png)

All models tested and described in the manuscript can be found in the seperate repository: <https://github.com/smaegol/PlasFlow_models>

Scripts used for the preparation of training dataset and for neural network training are available in the `scripts` subfolder as well in the separate repository: <https://github.com/smaegol/PlasFlow_processing>

## Citation

Please cite the following paper when using PlasFlow for your own research.

> Krawczyk PS, Lipinski L, Dziembowski A.
> Nucleic Acids Res. 2018 Apr 6;46(6):e35. doi: 10.1093/nar/gkx1321.

## TBD

In next releases we plan to retrain models using the most recent TensorFlow release. During the development of PlasFlow there was a lot of changes in the TensorFlow library and the newest version is not compatible with models trained for TensorFlow. However, retraining requires signficant computational effort and recoding. As we want to include _Archaea_ sequences (which are missed now) in the models, we plan to train new models with the latest TensorFlow version and release new version of PlasFlow in the second part of 2018.

## Support

Any issues connected with the PlasFlow should be addressed to Pawel Krawczyk (p.krawczyk (at) ibb.waw.pl).
