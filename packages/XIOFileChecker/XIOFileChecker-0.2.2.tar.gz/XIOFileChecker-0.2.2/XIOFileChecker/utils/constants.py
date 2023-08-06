#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""
    :platform: Unix
    :synopsis: Constants used in this package.

"""

from datetime import datetime

# Program version
VERSION = '0.2.2'

# Date
VERSION_DATE = datetime(year=2018, month=6, day=1).strftime("%Y-%d-%m")

# CMIP6 filename format
CMIP6_FILENAME_PATTERN = '^(?P<variable_id>[\w.-]+)_' \
                         '(?P<table_id>[\w.-]+)_' \
                         '(?P<source_id>[\w.-]+)_' \
                         '(?P<experiment_id>[\w.-]+)_' \
                         '(?P<variant_label>[\w.-]+)_' \
                         '(?P<grid_label>[^-_]+)' \
                         '(_(?P<period_start>[\w.-]+)-(?P<period_end>[\w.-]+))?' \
                         '\\.nc$'

# Table labels
XIOS_LABEL = 'XIOS'
DR2XML_LABEL = 'DR2XML'

# Table widths
XIOS_WIDTH = len(DR2XML_LABEL) + 2
DR2XML_WIDTH = len(DR2XML_LABEL) + 2

# Facet to ignore
IGNORED_FACETS = ['period_start', 'period_end']

# Help
PROGRAM_DESC = \
    """
    .___ _ ___ ___ _ _ _______ _ _________ _ _________.|n
    ._ _|_|___| __|_| |___ ___| |_ ___ ___| |_ ___ ___.|n
    |_'_| | . | __| | | -_| __| . | -_| __| '_| -_| __||n
    |_,_|_|___|_| |_|_|___|___|_|_|___|___|_,_|___|_| |n|n
                                                       
    XIOFileChecker allows you to easily check the consistency of XIOS inputs|n
    and outputs. This means XIOFileChecker compares what should be written by|n
    XIOS according to the DR2XML files definition and what was really written|n
    on disk.
    
    """

EPILOG = \
    """
    Developed by:|n
    Levavasseur, G. (UPMC/IPSL - glipsl@ipsl.fr)|n
    Moine, M.-P. (CNRM/CERFACS - marie-pierre.moine@cerfacs.fr)

    """

OPTIONAL = \
    """Optional arguments"""

POSITIONAL = \
    """Positional arguments"""

HELP = \
    """
    Show this help message and exit.

    """

LOG_HELP = \
    """
    Logfile directory.|n
    Default is the working directory.|n
    If not, standard output is used.

    """

DEBUG_HELP = \
    """
    Debug mode.

    """

ALL_HELP = \
    """
    Show all results/entries.|n
    Default only print differences.

    """

FULL_TABLE_HELP = \
    """
    Force all columns display.|n
    Default hide facets with unchanged values.

    """

VERSION_HELP = \
    """
    Program version.

    """

DIRECTORY_HELP = \
    """
    Input directory including XIOS filedef and netCDF output files.

    """

XML_HELP = \
    """
    Input directory including XIOS filedef, if different|n
    from simulation root directory.

    """

SET_FILTER_HELP = \
    """
    Filter facet values matching the regular expression.|n
    Duplicate the flag to set several filters.|n
    Default includes all values (i.e., no filters).
    
    """

MAX_PROCESSES_HELP = \
    """
    Number of maximal processes to simultaneously treat|n
    several files. Max is the CPU count.|n
    Set to 1 seems sequential processing.|n
    Set to -1 uses the max CPU count.|n
    Default is set to 4 processes.

    """
