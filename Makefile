PREFIX ?= /usr/local

genignore: main.go
	go build -o genignore

genignore.1: genignore.1.md
	rm -f genignore.1
	pandoc genignore.1.md -s -t man -o genignore.1

MAN = genignore.1

OBJS = genignore

install-bin: $(OBJS)
	cp genignore $(PREFIX)/bin/

install-man: $(MAN)
	cp genignore.1 $(PREFIX)/share/man/man1/

create-prefix:
	mkdir -p $(PREFIX)
	mkdir -p $(PREFIX)/bin
	mkdir -p $(PREFIX)/share/man/man1

install: create-prefix install-bin install-man

clean:
	rm -rf $(OBJS)

all: $(OBJS) $(MAN)

.PHONY: all clean install-bin install-man install