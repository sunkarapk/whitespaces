# whitespaces

__Fix all the whitespaces!!__

## Installation

You can download the binaries

 * Architecture i386 [ [linux](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_linux_386.tar.gz?direct) / [windows](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_windows_386.zip?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_darwin_386.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_freebsd_386.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_openbsd_386.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_netbsd_386.zip?direct) / [dragonfly](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_dragonfly_386.zip?direct) / [plan9](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_plan9_386.zip?direct) ]
 * Architecture amd64 [ [linux](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_linux_amd64.tar.gz?direct) / [windows](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_windows_amd64.zip?direct) / [darwin](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_darwin_amd64.zip?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_freebsd_amd64.zip?direct) / [openbsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_openbsd_amd64.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_netbsd_amd64.zip?direct) / [dragonfly](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_dragonfly_amd64.zip?direct) / [solaris](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_solaris_amd64.zip?direct) ]
 * Architecture arm [ [linux](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_linux_arm.tar.gz?direct) / [freebsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_freebsd_arm.zip?direct) / [netbsd](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_netbsd_arm.zip?direct) ]

Or by using deb packages

[ [i386.deb](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_i386.deb?direct) / [amd64.deb](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_amd64.deb?direct) / [armhf.deb](https://dl.bintray.com//content/pksunkara/utils/whitespaces_1.1.1_armhf.deb?direct) ]

Or by using golang

```
go get github.com/pksunkara/whitespaces
```

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
