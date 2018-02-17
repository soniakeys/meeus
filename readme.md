# Meeus

[![Build Status](https://travis-ci.org/soniakeys/meeus.png)](https://travis-ci.org/soniakeys/meeus) [![GoDoc](https://godoc.org/github.com/soniakeys/meeus?status.svg)](https://godoc.org/github.com/soniakeys/meeus) [![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/soniakeys/meeus)

Selected algorithms from the book "Astronomical Algorithms"
by Jean Meeus, following the second edition, copyright 1998,
with corrections as of August 10, 2009.

## Package organization

Package meeus is a documentation-only package.  Algorithms are implemented
in subdirectories under meeus, one for each chapter of the book.  In addition
there is a package "base" with additional functions that may not be described
in the book but are useful with multiple other packages.

See meeus package documentation for a chapter title cross-reference.

## Install

### Go get

Technically, `go get github.com/soniakeys/meeus/...` is sufficient.

The tests also require the sexagesimal package, so use the -t option to prompt
`go get` to find it as a test dependency:

    go get -t github.com/soniakeys/meeus/...

### Git clone

Alternatively, you can clone the repository into an appropriate place under
your GOPATH.  To clone into the same place as `go get` for example, assuming
the default GOPATH of ~/go, you would cd to `~/go/src/github.com/soniakeys`
before running the clone command.

    cd <somewhere under GOPATH>
    git clone https://github.com/soniakeys/meeus

The meeus package depends on the external package github.com/soniakeys/unit,
so you may need to `go get` or `git clone` the unit package.  And again, to
run tests, you will need the sexagesimal package.  You can use `go get` or
`git clone` as you prefer.

### Dep

You can also use dep (https://golang.github.io/dep/) to vendor the sexagesimal
and unit packages.  If you do this then these dependencies do not otherwise
need to be installed.  That is, you don't need the `-t` on `go get` and you
don't need to `git clone` sexagesimal or unit.

To use dep, first read about dep on the website linked above and install it.
Then install meeus with either `go get` or `git clone`.  Finally, from the
installed meeus directory, type

    dep ensure

This will "vendor" the sexagesimal and unit packages, installing them under the
`vendor` subdirectory and also installing specific versions of sexagesimal and
unit known to work with the version of meeus that you just installed.

## VSOP87

Routines of the `planetposition` package require "VSOP87" data files.  These
files should be available from public sources, for example
[VisieR](http://cdsarc.u-strasbg.fr/viz-bin/qcat?VI/81/).  The files needed
by planetposition are the VSOP87 "B" files.  It is sufficient to download
the eight files

    VSOP87B.ear  VSOP87B.mar  VSOP87B.nep  VSOP87B.ura
    VSOP87B.jup  VSOP87B.mer  VSOP87B.sat  VSOP87B.ven

There are no requirements on where you place these files in your file system
but you may find it convenient to create a directory for them and set an
environment variable `VSOP87` to this directory.

### Tests

With all eight VSOP87 files dowloaded as described above, and with an
environment variable set to their location, then from the meeus directory

    go test ./...

works as usual to run all tests in all subdirectories of meeus.

To run all tests except for those requiring planet positions computed from
the VSOP87 files, use

    go test -tags nopp ./...

("nopp" for no planet positions)

## Copyright and license

All software in this repository is copyright Sonia Keys and licensed with the
MIT license.

