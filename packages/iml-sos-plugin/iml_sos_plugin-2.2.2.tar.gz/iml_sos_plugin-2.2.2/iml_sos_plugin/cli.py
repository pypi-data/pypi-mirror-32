# Copyright (c) 2018 Intel Corporation. All rights reserved.
# Use of this source code is governed by a MIT-style
# license that can be found in the LICENSE file.

import sys
from subprocess import call
from itertools import chain, imap


def flat_map(fn, xs):
    return chain.from_iterable(imap(fn, xs))


def main():
    args = sys.argv[1:]

    plugins = [
        "iml",
        "corosync",
        "pacemaker",
        "kernel",
        "pci",
        "logs",
        "processor",
        "memory",
        "multipath",
        "filesys",
        "block",
        "yum",
        "systemd",
    ]

    code = call(["sosreport"] + args + ["--batch", "--log-size=0"] +
                list(flat_map(lambda x: ["--only", x], plugins)))

    sys.exit(code)

def chroma_diagnostics():
    print "chroma-diagnostics no longer exists. Please use 'iml-diagnostics' instead."
