# genignore
## Utility for setting up .gitignore files

### Usage

Usage: `genignore [template]` \
For a full list of templates, see `list.go`.

### Development

To setup the project for development, simply run `git clone https://github.com/regarager/genignore`.

To build, run `make`. If you want to use more cores, just run it with the `-j<numberOfCores>` flag.

### Installation

To install, just run `make install`, using `sudo` where necessary. It will automatically install the binary and manpages.

### Updating The Template Store

To update the cache of templates the program downloads on setup, run the program with the `update` subcommand (i,e `genignore update`). The program will automatically clear the template store, restart itself, and download the new versions.

### License
This project is licensed under the GNU General Public License v3.0 (GPLv3). All past and future versions of `genignore` are covered by this license. See the LICENSE file for full details.
