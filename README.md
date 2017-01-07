# tar-append

A native Go CLI tool for appending files to the end of a tar received via stdin and sent to stdout.

## Notes

1. tar-append only appends files to the end of the input archive which is marked by a closed/detached STDIN
2. tar-append will only append files at or below the current working directory.
3. tar-append will only append a single file (the first argument to the program).
4. tar-append will append files with names that already exist in the archive - this is purposeful.
5. tar-append only writes the tarball to STDOUT. Error messages will be sent on STDERR.

## Examples

    mkdir logs
    tar cf - 1.txt 2.txt | tar-append 3.txt | tar-append 4.txt | tar xf - -C ./logs
    ls ./logs
    # 1.txt   2.txt   3.txt   4.txt 

### Authors and Copyright Holders

1. Jeff Nickoloff "jeff@allingeek.com"

### License

This project is released under the MIT license.
