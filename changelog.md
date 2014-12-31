## v1 31 Dec 2014

Status:  No changes in the last year!  89 GitHub stars currently, otherwise
little feedback.  Some feedback suggested changes to improve algorithms or
wrap existing functions in a more convenient API.  I declined these in the
interest of keeping this library as close as practical to the algorithms in
the book.  Otherwise, bug reports or other issues are always welcome.

There were a few minor changes just after the v0.2 tag.  Also I do have a
rewrite of of the angle formatting stuff ready to merge now.  This changes
the angle formatting API so it's time for a new tag.  The v1 tag will reference
this code that has been stable for the last year, including the minor changes
just after v0.2, but before API breaking changes to the angle formatting.

Happy New Year!

## v0.2 2013-10-11

* All chapters now implemented.

Upcoming work will change the API a bit.

## v0.1 2013-09-29

Tag added with the library starting to get some attention.  Existing code is
pretty good shape but there is some work remaining.  Here's what I can think
of off the top of my head:

* Implement remaining chapters.  There are just a few left.

* Change the API in a few places.  As packages from later chapters have
imported earlier ones, I've seen a few little things I'd like to go back and
change.  The biggest change will be for angle formatting.  I'll keep the
basic idea but allow more flexibility, probably at the expense of adding
some API methods.

* Add correctness tests in addition to the examples from the book.  This is
possible in many cases by drawing on sources outside the book.

* Review everything for consistency in style and naming.  I learned a few
things as I went along and some of the earlier packages need updates.
