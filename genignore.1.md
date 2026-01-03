---
date: January 2026
footer: genignore
header: User Commands
section: 1
title: GENIGNORE
---

# NAME

genignore - A utility for setting up .gitignores

# SYNOPSIS

**genignore** \[*template*\]

# DESCRIPTION

**genignore** generates *.gitignore* files using predefined templates
corresponding to programming languages, build systems, and environments.

If a *template* name is provided, the matching ignore rules are
generated.

Template names are case-sensitive and correspond to the internal
template identifiers shipped with the program.

# ARGUMENTS

*template*

:   The name of the template to generate (for example, **Go**,
    **Python**, **Node**, **Rust**).

# TEMPLATES

**genignore** ships with a large collection of templates covering
many programming languages, build systems, and environments.

The complete list of available template names can be found in the
source file *list.go*.

# EXAMPLES

Generate a Go *.gitignore* file:

    genignore Go

Generate a Python *.gitignore* file:

    genignore Python

# EXIT STATUS

**0**

:   Success.

**\>0**

:   An error occurred (for example, an unknown template).

# SEE ALSO

**gitignore**(5), **git**(1)

# AUTHOR

Written by regarager.
