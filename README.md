
withenv
======
**withenv** runs reads in a `.env` file from your current working directory,
loads the contained variables into the environment and runs the provided
command with the new environment.

## Installing
**withenv** is written in go so you're gonna need that for sure. 

To install:

``
go install github.com/wnh/withenv
``

## Running
``
$ withenv command arg1 .. argN
``

## File format
The `.env` file  contains `NAME=value` pairs each on its own line.  All leading
trailing white space is preserved.  The name is specified as the text before
the *first* equals. Thus `foo=bar=baz` will result in a variable named `foo`
set to the value `bar=baz`.  Blank lines are not allowed at this time, nor are
comments.


## TODO

- Make blank lines need to work. Seriously.
- Comments would be nice.
- Proper handling of quoted values for compatibility with some other environment tools

## Version 
* Version 0.1

## Contact
#### Developer/Company
* Homepage: https://github.com/wnh
* e-mail: harding.will@gmail.com
