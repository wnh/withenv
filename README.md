
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
 - The `.env` file  contains `NAME=value` pairs each on its own line.
 - The name is specified as the text before the *first* equals. Thus
   `foo=bar=baz` will result in a variable named `foo` set to the value
   `bar=baz`.  
 - Blank lines are ignored
 - All leading trailing white space is removed
 - Comments are line based. Comment lines use the first non-whitespace character '#'


## TODO

- Make shell scripts work without addind `sh` infront of them
- RTests maybe? I dunno, could be useful.
- Proper handling of quoted values for compatibility with some other environment tools

## Version 
* Version 0.1

## Contact
* Homepage: https://github.com/wnh
* e-mail: harding.will@gmail.com
