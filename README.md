# DsDoc

API Document generator for DSLinks. This tool will scan the files in the current
directory, and any sub-directories and parse source files for any DsDoc comments.

*The tool will ignore any hidden directories, and will not traverse symlinks.*

The DsDoc tool currently supports Dart, Java, C, and C++ source files for parsing.

Once the source files have been parsed it will build a document tree of the
various nodes and actions which comprise the link and output the documentation
in markdown or plain text. Included in the output is a hierarchy tree of link.

## Installation

Install the tool with `go install`.

```
go install github.com/iot-dsa/dsdoc
```

## Run the tool

To run the tool, ensure your [gopath is in your PATH](https://golang.org/doc/code.html#GOPATH)
then, from the root of your link source files, run `dsdoc`.

By default the tool will create your DSLink API documentation in a file called `api.md`

# Writing DsDocs

DsDocs use a special comment form with Annotations to delimit the documentation.
Since the tool will parse all files prior to setting up the hierarchy, it is
recommend that the DsDoc comments appear in your source files near where the nodes
are declared.

## DsDoc Comments

DsDoc comments must start with `//*`

DsDoc comments must be specified in a single comment block rather than broken up
into multiple areas. The DsDoc tool ignores any code in the files and does **not**
use the code in any way to influence the documentation.

```
//* line one
//* line two
```

DsDoc comments may be at the end of a series of code lines, or on lines on their
own.
```
some code;  //* line one
other code; //* line two
```

At least one non-dsdoc comment line must separate blocks of comments. This
line may or may not contain code.
```
//* block one
//* block one still
  
//* block two
```

## DsDoc Format

DsDocs currently have two specific formats. One for documenting Nodes and one
for documenting Actions. These also share some common attributes which must be
declared. Other annotations are optional as noted.

A DsDoc must start with either a `@Node` or an `@Action` annotation. They optionally
may be followed by a path name when the node or action has a fixed path.

Following the `@Node` or `@Action` annotation, all other annotations may be in
any order you choose, however they must all start on their own line.

### `@Node [pathName]`

The `pathName` is optional and may be omitted. If a node has a fixed name
it should be included. A node cannot be invoked by a user of the link. 
It may or may not contain a value which in turn may or may not be writable. 
They may also make up part of the hierarchy of a link.

### `@Action [pathName]`

The `pathName` is optional and may be omitted. However it is very rare that
an Action has a dynamic name, and usually have a fixed name which should be
included. An action is a node which may be invoked by a user of the link, often
times with various parameters, and may or may not have a return value of varying
types.


### `@MetaType [type]`

If a Node or Action does not have a fixed name, such as may be the case if a
node receives a name from either user input or via an external API call, then
the DsDoc must include an `@MetaType` annotation which represents an internal
type for the link. The `type` argument is required in this case.  
It is an error to define both a name and MetaType.

### `@Is [isType]`

This attribute is optional but recommended. If a Node or Action implements a 
specific `$is` attribute, as are used with Profiles, specify the case-sensitive 
`$is` name here.
