
NAME
    withenv -- run command with modified environment

SYNOPSIS
    withenv [--<file>] command...

DESRIPTION
    *withenv* runs reads in a environment file from your current working directory,
    loads the contained variables into the environment and runs the provided
    command with the new environment.
    
    Default environment file name is '.env' 

    --<file>
        Loads '.<file>.env' instead of default environment file

        Example: 
          `withenv --local ./server`
        will load variable fom the file '.local.env'

FILE FORMAT
    The `.env` file  contains `NAME=value` pairs each on its own line.
    
    The name is specified as the text before the *first* equals. Thus
    `foo=bar=baz` will result in a variable named `foo` set to the value
    `bar=baz`.  
    
    Blank lines are ignored
    
    All leading trailing white space is removed
    
    Comments are line based. Comment lines use the first non-whitespace character '#'

TODO
    Make shell scripts work without addind `sh` infront of them
    Proper handling of quoted values for compatibility with some other environment tools

CONTACT
    Homepage: https://github.com/wnh
    e-mail: harding.will@gmail.com
