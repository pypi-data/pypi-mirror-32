#!/usr/bin/env python
#
# -----------------------------------------------------------------------------
# Copyright (c) 2017 The Regents of the University of California
#
# This file is part of kevlar (http://github.com/dib-lab/kevlar) and is
# licensed under the MIT license: see LICENSE.
# -----------------------------------------------------------------------------

import glob
import pytest
import re
from tempfile import NamedTemporaryFile
import screed
import kevlar
from kevlar.tests import data_file, data_glob


@pytest.fixture
def triomask():
    mask = khmer.Counttable(19, 1e4, 4)
    mask.consume('TGAGGGGACTAGGTGATCAGGTGAGGGTTTCCCAGTTCCCGAAGATGACT')
    mask.consume('GATCTTTCGCTCCCTGTCATCAAGGAGTGATACGCGAAGTGCGTCCCCTT')
    mask.consume('GAAGTTTTGACAATTTACGTGAGCCCTACCTAGCGAAACAACAGAGAACC')
    return mask


@pytest.mark.parametrize('mask,numbands,band', [
    (None, None, None),
    (None, 9, 2),
    (triomask, None, None),
    (triomask, 23, 19),
])
def test_load_threading(mask, numbands, band):
    # Smoke test: make sure things don't explode when run in "threaded" mode.
    infiles = data_glob('trio1/case1.fq')
    sketch = kevlar.count.load_sample_seqfile(
        infiles, 19, 1e7, mask=mask, numbands=numbands, band=band, numthreads=2
    )


@pytest.mark.parametrize('infile,testout,numbands,band,kmers_stored', [
    ('case', 'case', 0, 0, 973),
    ('ctrl1', 'ctrl1', 0, 0, 973),
    ('ctrl2', 'ctrl2', 0, 0, 966),
    ('case', 'case-band-2-1', 2, 1, 501),
    ('case', 'case-band-16-7', 16, 7, 68),
])
def test_count_simple(infile, testout, numbands, band, kmers_stored, capsys):
    infile = data_file('simple-genome-{}-reads.fa.gz'.format(infile))
    testout = data_file('simple-genome-{}.ct'.format(testout))
    with NamedTemporaryFile() as outfile:
        arglist = ['count', '--ksize', '25', '--memory', '10K',
                   '--num-bands', str(numbands), '--band', str(band),
                   outfile.name, infile]
        args = kevlar.cli.parser().parse_args(arglist)
        kevlar.count.main(args)
        out, err = capsys.readouterr()

        assert '600 reads processed' in str(err)
        assert '{:d} distinct k-mers stored'.format(kmers_stored) in str(err)

        outputfilename = outfile.name + '.counttable'
        with open(outputfilename, 'rb') as f1, open(testout, 'rb') as f2:
            assert f1.read() == f2.read()


def test_count_threading():
    with NamedTemporaryFile(suffix='.counttable') as outfile:
        infile = data_file('trio1/case1.fq')
        arglist = ['count', '--ksize', '19', '--memory', '500K',
                   '--threads', '2', outfile.name, infile]
        args = kevlar.cli.parser().parse_args(arglist)
        kevlar.count.main(args)

    # No checks, just doing a "smoke test" to make sure things don't explode
    # when counting is done in "threaded" mode.


def test_count_problematic():
    arglist = [
        'count', '--ksize', '21', '--memory', '200K', '--band', '2',
        'bogusoutput', data_file('trio1/ctrl1.fq')
    ]
    args = kevlar.cli.parser().parse_args(arglist)
    with pytest.raises(ValueError) as ve:
        kevlar.count.main(args)
    assert 'Must specify --num-bands and --band together' in str(ve)

    arglist = [
        'count', '--ksize', '21', '--memory', '97',
        'bogusoutput', data_file('trio1/ctrl1.fq')
    ]
    args = kevlar.cli.parser().parse_args(arglist)
    with pytest.raises(kevlar.sketch.KevlarUnsuitableFPRError):
        kevlar.count.main(args)


def test_effcount_smoketest():
    with NamedTemporaryFile(suffix='.ct') as o1, \
            NamedTemporaryFile(suffix='.ct') as o2, \
            NamedTemporaryFile(suffix='.ct') as o3:

        arglist = [
            'effcount', '--sample', data_file('trio1/ctrl1.fq'),
            '--sample', data_file('trio1/ctrl2.fq'),
            '--sample', data_file('trio1/case2.fq'),
            '--ksize', '21', '--memory', '200K', '--memfrac', '0.005',
            '--max-fpr', '0.1', '--threads', '2', o1.name, o2.name, o3.name
        ]
        args = kevlar.cli.parser().parse_args(arglist)
        kevlar.effcount.main(args)


def test_effcount_problematic():
    arglist = [
        'effcount', '--sample', data_file('trio1/ctrl1.fq'), '--ksize', '21',
        '--memory', '200K', 'bogusoutput'
    ]
    args = kevlar.cli.parser().parse_args(arglist)
    with pytest.raises(AssertionError):
        kevlar.effcount.main(args)

    arglist = [
        'effcount', '--sample', data_file('trio1/ctrl1.fq'),
        '--sample', data_file('trio1/ctrl2.fq'), '--ksize', '21',
        '--memory', '200K', 'bogusoutput'
    ]
    args = kevlar.cli.parser().parse_args(arglist)
    with pytest.raises(ValueError) as ve:
        kevlar.effcount.main(args)
    assert 'number of outfiles must match number of declared' in str(ve)

    arglist = [
        'effcount', '--sample', data_file('trio1/ctrl1.fq'),
        '--sample', data_file('trio1/ctrl2.fq'), '--ksize', '21',
        '--memory', '200K', '--band', '2', 'bogusoutput1', 'bogusoutput2'
    ]
    args = kevlar.cli.parser().parse_args(arglist)
    with pytest.raises(ValueError) as ve:
        kevlar.effcount.main(args)
    assert 'Must specify --num-bands and --band together' in str(ve)
