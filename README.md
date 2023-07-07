findfile is a simple tool that finds matching files using positional arguments as filters.

a common prefix is found and stripped during each positional argument processing.
This style of matching allows filters to be iterively extended.

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [Examples](#examples)
- [Usage](#usage)

<!-- markdown-toc end -->


# Examples

     # Find files and directories matching 'foo' and 'bar'
     findfile foo bar

     # Find files and directories matching 'foo' and 'bar' (case sensitive)
     findfile -C foo bar
     
     # Find just directories matching fizz
     findfile -D fizz
     


# Dev


    go install github.com/goreleaser/goreleaser@latest
