Meeus
=====

[GoDoc](http://godoc.org/github.com/soniakeys/meeus)

Selected algorithms from the book "Astronomical Algorithms"
by Jean Meeus, following the second edition, copyright 1998,
with corrections as of August 10, 2009.

Generally each subpackage implements algorithms of one chapter of the book.
In addition there is a package "base" with additional functions that
may not be described in the book but are useful with multiple other packages.

Cross reference from chapters to package names
----------------------------------------------
<table>
	<tr><th>Chapter</th><th>Package name</th><th>Status</th></tr>
	<tr><td>1.  Hints and Tips</td><td>hints</td><td>complete</td></tr>
	<tr><td>2.  About Accuracy</td><td>accuracy</td><td>complete</td></tr>
    <tr><td>3.  Interpolation</td><td>interp</td><td>under renovation</td></tr>
    <tr><td>4.  Curve Fitting</td><td>fit</td><td>complete</td></tr>
    <tr><td>5.  Iteration</td><td>iterate</td><td>complete</td></tr>
    <tr><td>6.  Sorting Numbers</td><td>sort</td><td>complete</td></tr>
    <tr><td>7.  Julian Day</td><td>julian</td><td>complete</td></tr>
    <tr><td>8.  Date of Easter</td><td>easter</td><td>complete</td></tr>
    <tr><td>9.  Jewish and Moslem Calendars</td><td>jm</td><td>complete</td></tr>
    <tr><td>10. Dynamical Time and Universal Time</td><td>deltat</td><td>complete</td></tr>
    <tr><td>11. The Earth’s Globe</td><td>globe</td><td>complete</td></tr>
    <tr><td>12. Sidereal Time at Greenwich</td><td>sidereal</td><td>complete</td></tr>
    <tr><td>13. Transformation of Coordinates</td><td>coord</td><td>complete</td></tr>
    <tr><td>14. The Parallactic Angle, and three other Topics</td><td>parallactic</td><td>complete</td></tr>
    <tr><td>15. Rising, Transit, and Setting</td><td>rise</td><td>complete</td></tr>
    <tr><td>16. Atmospheric Refraction</td><td>refraction</td><td>complete</td></tr>
    <tr><td>17. Angular Separation</td><td>angle</td><td>complete</td></tr>
    <tr><td>18. Planetary Conjunctions</td><td>conjunction</td><td>complete</td></tr>
    <tr><td>19. Bodies in Straight Line</td><td>line</td><td>complete</td></tr>
    <tr><td>20. Smallest Circle containing three Celestial Bodies</td><td>circle</td><td>complete</td></tr>
    <tr><td>21. Precession</td><td>precess</td><td>partial</td></tr>
    <tr><td>22. Nutation and the Obliquity of the Ecliptic</td><td>nutation</td><td>complete</td></tr>
    <tr><td>23. Apparent Place of a Star</td><td>apparent</td><td>complete</td></tr>
    <tr><td>24. Reduction of Ecliptical Elements from one Equinox to another one</td><td>elementequinox</td><td>complete</td></tr>
    <tr><td>25. Solar Coordinates</td><td>solar</td><td>partial</td></tr>
    <tr><td>26. Rectangular Coordinates of the Sun</td><td>solarxyz</td><td>complete</td></tr>
    <tr><td>27. Equinoxes and Solstices</td><td>solstice</td><td>complete</td></tr>
    <tr><td>28. Equation of Time</td><td>eqtime</td><td>complete</td></tr>
    <tr><td>29. Ephemeris for Physical Observations of the Sun</td><td>solardisk</td><td>complete</td></tr>
    <tr><td>30. Equation of Kepler</td><td>kepler</td><td>complete</td></tr>
    <tr><td>31. Elements of Planetary Orbits</td><td>planetelements</td><td>partial</td></tr>
    <tr><td>32. Positions of the Planets</td><td>planetposition</td><td>partial</td></tr>
    <tr><td>33. Elliptic Motion</td><td>elliptic</td><td>partial</td></tr>
    <tr><td>34. Parabolic Motion</td><td>parabolic</td><td>complete</td></tr>
    <tr><td>35. Near-parabolic Motion</td><td>nearparabolic</td><td>complete</td></tr>
    <tr><td>36. The Calculation of some Planetary Phenomena</td><td>planetary</td><td>partial</td></tr>
    <tr><td>37. Pluto</td><td>pluto</td><td>complete</td></tr>
    <tr><td>38. Planets in Perihelion and in Aphelion</td><td>perihelion</td><td>complete</td></tr>
    <tr><td>39. Passages through the Nodes</td><td>node</td><td>complete</td></tr>
    <tr><td>40. Correction for Parallax</td><td>parallax</td><td>complete</td></tr>
    <tr><td>41. Illuminated Fraction of the Disk and Magnitude of a Planet</td><td>illum</td><td>complete</td></tr>
    <tr><td>42. Ephemeris for Physical Observations of Mars</td><td>mars</td><td>complete</td></tr>
    <tr><td>43. Ephemeris for Physical Observations of Jupiter</td><td>jupiter</td><td>complete</td></tr>
    <tr><td>44. Positions of the Satellites of Jupiter</td><td>jupitermoons</td><td>complete</td></tr>
    <tr><td>45. The Ring of Saturn</td><td>saturnring</td><td>complete</td></tr>
    <tr><td>46. Positions of the Satellites of Saturn</td><td>saturnmoons</td><td>complete</td></tr>
    <tr><td>47. Position of the Moon</td><td>moon</td><td>complete</td></tr>
    <tr><td>48. Illuminated Fraction of the Moon&#39s Disk</td><td>moonillum</td><td>complete</td></tr>
    <tr><td>49. Phases of the Moon</td><td>moonphase</td><td>complete</td></tr>
    <tr><td>50. Perigee and apogee of the Moon</td><td>apsis</td><td>partial</td></tr>
    <tr><td>51. Passages of the Moon through the Nodes</td><td>moonnode</td><td>complete</td></tr>
    <tr><td>52. Maximum declinations of the Moon</td><td>moonmaxdec</td><td>complete</td></tr>
    <tr><td>Non-Meeus useful functions</td><td>base</td><td></td></tr>
</table>
