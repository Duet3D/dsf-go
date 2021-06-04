# Duet Software Framework Go Bindings

Find out more about [Duet Software Framework](https://github.com/Duet3D/DuetSoftwareFramework).

Find out more about [Duet Software Framework Go Bindings](https://github.com/Duet3D/dsf-go).

Get in touch with the community at [Duet Software Framework Forum](https://forum.duet3d.com/category/31/dsf-development) for bug reports, discussion and any kind of exchange.

## Differences
* A few functionalities had to be left out since there was no good representation in Go
* Since Go has no implicit type conversion there will be As<Type>() methods provided instead
* Currently there is no notification mechanism for object model updates
* In some cases zero values were chosen instead of nil that would be used by upstream
* Geometry was renamed to Kinematics
