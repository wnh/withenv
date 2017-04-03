withenv(1)							    withenv(1)



NAME
       withenv ‐‐ run command with modified environment


SYNOPSIS
       withenv [‐‐<file>] command...


DESRIPTION
       withenv	runs  reads  in  a  environment file from your current working
       directory, loads the contained variables into the environment and  runs
       the provided command with the new environment.

       Default environment file name is ’.env’


OPTIONS
       ‐‐<file>
	      Loads ’.<file>.env’ instead of default environment file



FILES
       The ‘.env‘ file	contains ‘NAME=value‘ pairs each on its own line.

       The  name  is  specified  as  the  text before the *first* equals. Thus
       ‘foo=bar=baz‘ will result in a variable named ‘foo‘ set	to  the  value
       ‘bar=baz‘.

       Blank lines are ignored

       All leading trailing white space is removed

       Comments  are  line  based.  Comment lines use the first non‐whitespace
       character ’#’


EXAMPLES
       ‘withenv ./server‘
	      will run server with the default ’./.env’

       ‘withenv ‐‐local ./server‘
	      will load variable fom the file ’./.local.env’


AUTHORS
       Homepage: https://github.com/wnh
       Email: harding.will@gmail.com


TODO
       Make shell scripts work without addind ‘sh‘ infront of them Proper han‐
       dling  of  quoted  values for compatibility with some other environment
       tools




								    withenv(1)
