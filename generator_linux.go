//go:generate wget --no-clobber http://mirror.gutscheinrausch.de/gnu/bc/bc-1.07.tar.gz
//go:generate rm -rf bc-1.07/
//go:generate tar xzf bc-1.07.tar.gz
//go:generate sh -c "cp bc-1.07/AUTHORS bc-1.07/COPYING* bc-1.07/NEWS bc-1.07/README ."
//go:generate sh -c "cd bc-1.07/ && ./configure CC=99c"
//go:generate make -C bc-1.07/
//go:generate mkdir assets
//go:generate cp bc-1.07/bc/bc assets
//go:generate assets

package main
