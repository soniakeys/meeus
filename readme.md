# Meeus

[![Build Status](https://travis-ci.org/soniakeys/meeus.png)](https://travis-ci.org/soniakeys/meeus) [![GoDoc](https://godoc.org/github.com/soniakeys/meeus?status.svg)](https://godoc.org/github.com/soniakeys/meeus) [![Go Walker](http://gowalker.org/api/v1/badge)](http://gowalker.org/github.com/soniakeys/meeus) [![status](https://sourcegraph.com/api/repos/github.com/soniakeys/meeus/.badges/status.png)](https://sourcegraph.com/github.com/soniakeys/meeus)

Selected algorithms from the book "Astronomical Algorithms"
by Jean Meeus, following the second edition, copyright 1998,
with corrections as of August 10, 2009.

Package meeus is a documentation-only package.  Algorithms are implemented
in subdirectories under meeus, one for each chapter of the book.  In addition
there iss a package "base" with additional functions that may not be described
in the book but are useful with multiple other packages.

See meeus package documentation for a chapter title cross-reference.

## Breaking changes!

7 Jan 2015 the custom formatters are replaced.  Use the v1 tag if you need
the old formatters.
