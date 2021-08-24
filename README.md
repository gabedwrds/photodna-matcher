photodna-matcher finds matches in a set of PhotoDNA hashes and is presented for research purposes only.

## input

Provide a file list `hashes.csv` with one line per file. Each line includes a filename and a PhotoDNA hash separated by a comma. PhotoDNA hashes must be 192 characters long (144 bytes encoded with base64).

To generate hashes themselves, check out [jPhotoDNA](https://github.com/jankais3r/jPhotoDNA) or any forensics package implementing PhotoDNA.

## output

For each file in the set, the output will list its ID (the line number in the input CSV), filename, PhotoDNA distance from its nearest match, and the ID and filename of that nearest match.

You might want to pipe this through `sort -k3 -n` to get the closest matches first.

## example

input CSV:
```
ILSVRC2012_test_00000001.JPEG,BgwBORAEAFEIAwBKABIAZAV4DbBC0I3PDRo3DxAsRyEZNElfGDtlbE01sExpFcYNDAkWqgMTLc0GFWLPEAdszxMKRNEEDQnCCR4nxwcMF9oCBAjWAAUB2wAkAPwMahT/ECctLRQvJS0RLBk2Bz4WUCQ/RGNQJYciBiUSHRgjGRoYLTQcExxUDicGYgMkADUA
ILSVRC2012_test_00000002.JPEG,TCRUFWgzxAojXKUVLjlRW0oPFVEiFA4nbAc5HIcLNmMkZHoSKpeLPy0qeChDAzgQLwkXL2onIlI3iyP/tq5V/3M8T0YjCh8KRQQ2IhoyHCwuXbBwnST8RXVjgS4XLxoqMhM5IglPJj5JMRWMch8eZYlAMKMbgFIlNA0QJhOeHBpwTQs8TCsZMI4pQVE0QRpQ
...
```

results:
```
0 ILSVRC2012_test_00000001.JPEG 232274 from 60885 ILSVRC2012_test_00060886.JPEG
1 ILSVRC2012_test_00000002.JPEG 281942 from 86213 ILSVRC2012_test_00086214.JPEG
...
```