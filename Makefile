genignore.1: genignore.1.md
	rm -f genignore.1
	pandoc genignore.1.md -s -t man -o genignore.1

man: genignore.1

all: man

.PHONY: all man
