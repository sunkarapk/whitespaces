# whitespaces

__Fix all the whitespaces!!__

## Installation

You can install it using golang

```
go get github.com/pksunkara/whitespaces
```

Or select one of the binaries below

 * Architecture i386 [ [linux](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_linux_386.tar.gz?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_darwin_386.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_freebsd_386.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_openbsd_386.zip?direct) ]
 * Architecture amd64 [ [linux](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_linux_amd64.tar.gz?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_darwin_amd64.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_freebsd_amd64.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_openbsd_amd64.zip?direct) ]
 * Architecture arm [ [linux](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_linux_arm.tar.gz?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_freebsd_arm.zip?direct) ]

You can install using deb files too.

 * [i386.deb](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_i386.deb?direct)
 * [amd64.deb](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_amd64.deb?direct)
 * [armhf.deb](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.0_armhf.deb?direct)

## Usage

```
whitespaces [options] FILES
```

#### Fix all whitespaces

```
whitespaces FILES
```

* Fixes trailing whitespaces
* Fixes whitespaces in blank lines
* Adds newline character at the end of file

#### Check for whitespaces

```
whitespaces -c FILES
```

Checks for whitespaces and exits if found.

```
whitespaces -c -f FILES
```

Continue checking for whitespaces in all files and then exit

_If you like this project, please watch this and follow me_

## Contributors
Here is a list of [Contributors](http://github.com/pksunkara/whitespaces/contributors)

### TODO

- Some more cases

__I accept pull requests and guarantee a reply back within a day__

## License
MIT/X11

## Bug Reports
Report [here](http://github.com/pksunkara/whitespaces/issues). __Guaranteed reply within a day__.

## Contact
Pavan Kumar Sunkara (pavan.sss1991@gmail.com)

Follow me on [github](https://github.com/users/follow?target=pksunkara), [twitter](http://twitter.com/pksunkara)
