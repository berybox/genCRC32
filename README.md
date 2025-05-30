# genCRC32: Collision-Free CRC32 Hashing for DNA Sequences

## Overview

`genCRC32` is a specialized hashing tool developed in Go for efficiently encoding DNA sequences into 32-bit hash codes without collisions for sequences up to 16 nucleotides (16-mers). It employs a preprocessing step to optimize sequences specifically for CRC32 hashing, making it highly suitable for bioinformatics applications such as genome assembly, sequence alignment, and variant detection.

## Features

* **Collision-Free Hashing:** Ensures unique hash values for DNA sequences (k-mers) up to 16 bases in length.
* **Efficiency:** Utilizes CRC32 for rapid computation, ideal for high-throughput genomic data analysis.
* **Simplicity and Portability:** Implemented purely with Go's standard library for easy integration and cross-platform compatibility.

## Installation

Download and run binary file from Releases

or clone the repository and build using Go:

```bash
git clone https://github.com/berybox/genCRC32.git
cd genCRC32/cmd
go build
```



## Usage

### Hash a DNA sequence

```bash
./genCRC32 hash -s <DNA_sequence> [-p <polynomial>]
```

* `-s, --sequence`: (Required) DNA sequence to hash.
* `-p, --polynomial`: (Optional) CRC32 polynomial (default: first polynomial from the built-in list).

Example:

```bash
./genCRC32 hash -s ATCGATCGATCGATCG
```

### List available CRC32 polynomials

```bash
./genCRC32 list
```

## Algorithm

`genCRC32` combines a preprocessing step (`gen32`) and standard CRC32 hashing:

### Preprocessing (`gen32`)

* DNA sequence is compressed by pairing nucleotides into bytes, reducing sequence complexity.
* Bitwise operations ensure consistency regardless of nucleotide case (case-insensitive).

### CRC32 Hashing

* Processed sequences are hashed with selected CRC32 polynomials proven to avoid collisions.
* Recommended polynomials (collision-free):

  * `0x8741c726`, `0x87496166`, `0x8e2371ef`, `0x8ee5368f`, `0x8efd4bcd`, `0x945d045d`, `0x9d9947fd`, `0xeb31d82e`

## Performance

* Collision-free hashing verified for k-mers from 4 to 16 bases.
* Minimal computational overhead due to highly optimized bitwise operations and CRC32 hashing.


## Reference

This work builds upon established bioinformatics methodologies and CRC32 principles. Detailed explanations and tests can be found in the associated manuscript available upon publication. Currently under review.
