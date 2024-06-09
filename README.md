# copy
Custom implementation of the linux cp command with progress reporting

## Usage

This program essentially follows the exact same usage as the regular `cp` command.

``` console
copy [OPTIONS] <SRC> <DEST>
```

### Options:
- `-R`, `--recursive` - Copy directory recursively (Same as stock `cp`)
- `-p`, `--progress` - Enable the progress bar to be displayed, Useful for copying large directories or files.

## Installation

1. Clone the repo

``` console
$ git clone https://github.com/joshburnsxyz/copy
```

2. Run the `make` command.

``` console
$ cd copy
$ make
```

3. Run the `make install` command

``` console
$ sudo make install
```

4. Verify installation

``` console
copy --help
```

