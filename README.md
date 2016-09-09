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
`pathName` may not contain spaces.

### `@Action [pathName]`

The `pathName` is optional and may be omitted. However it is very rare that
an Action has a dynamic name, and usually have a fixed name which should be
included. An action is a node which may be invoked by a user of the link, often
times with various parameters, and may or may not have a return value of varying
types.  
`pathName` may not contain spaces.

### `@MetaType [type]`

If a Node or Action does not have a fixed name, such as may be the case if a
node receives a name from either user input or via an external API call, then
the DsDoc must include an `@MetaType` annotation which represents an internal
type for the link. The `type` argument is required in this case.  
It is an error to define both a name and MetaType.  
`type` cannot contain spaces.

### `@Is [isType]`

This attribute is optional but recommended. If a Node or Action implements a
specific `$is` attribute, as are used with Profiles, specify the case-sensitive
`$is` name here.  
`isType` cannot contain spaces.

### `@Parent [parentName]`

This attribute is required by all DsDocs (both Node and Action types). The value
should be the name or `MetaType` of the parent of this node or action. If the
node or action is on the root of the link, then you should use the special
value `root`.

### `Short Description`

A Short description is required for all nodes and actions. It does not start with
an annotation, rather it is denoted by being preceded and followed by an empty
line.

### `Long Description`

A long description is optional, but follows the same conventions as
Short Description by being preceded and followed by an empty line. This description
may span multiple lines, but not multiple paragraphs. It should provide detailed
information about the purpose of node or action.

### `@Param [name] [type] [Description]`

The `@Param` annotation is optional for Action DsDocs. It is not valid for Node
DsDocs. They represent the parameters passed to the action. If provided, the
`name`, `type` and `Description` are required. Multiple `@Param` annotations may
be specified.  
`name` cannot contain spaces.  
`type` cannot contain spaces, and should be the DSA type of parameter, including
bool, num, string, and enum.  
`Description` is a long description of what the value represents, and may span
multiple lines.

### `@Return [type]`

The `@Return` annotation is optional for Action DsDocs. It is not valid for Node
type DsDocs. This represents the return type of the Action. If provided, the
`type` must be specified.  
`type` Should be the return type specified by the Action, and may include, value,
table, and stream.

### `@Column [name] [type] [Description]`

The `@Column` annotation is optional for Action DsDocs and should be provided if
a `@Return` annotation is present. It is not valid for Node DsDoc types.
This represents the columns returned by the Action. If provided, the `name`,
`type`, and `Description` are required. Multiple `@Column` annotations may be
specified.  
`name` may not contain spaces.  
`type` may not contain spaces. It should be a DSA type such as bool, num, and
string.  
`Description` is a long description of what the column represents, and may span
multiple lines.

### `@Value [type]`

The `@Value` annotation is optional for Node DsDocs. It is not valid for Action
DsDocs. A value should be specified if this node has a readable or writable
value.  
`type` must be specified if the annotation is provided. It should represent the
DSA type of the value including num, bool, and string.

## Examples

The following are several examples illustrating a fictional link.

```
//* @Action Add_Device
//* @Is addDeviceCmd
//* @Parent root
//*
//* Adds a device to the link.
//*
//* Add Device accepts the URL and name of the device to add to the link. It
//* will verify the URL is accessible and return an error message if it fails.
//*
//* @Param url string The URL of the device.
//* @Param name string The name for the device, it will be added to the link
//* under this name.
//*
//* @Return value
//* @Column success bool A boolean which represents if the action succeeded or
//* not. Returns false on failure and true on success.
//* @Column message string If the action succeeds, this will be "Success!", on
//* failure, it will return the error message.
```

```
//* @Node
//* @MetaType DeviceNode
//* @Is deviceNode
//* @Parent root
//*
//* A device which has been added to the link.
//*
//* When added to the link, a device will appear as the name provided. This
//* node maintains the connection with the remote host.
```

```
//* @Action Remove_Device
//* @Is removeDeviceCmd
//* @Parent DeviceNode
//*
//* Removes a device from the link.
//*
//* @Return value
//* @Column success bool A boolean which represents if the action succeeded or
//* not. Returns false on failure and true on success.
//* @Column message string If the action succeeds, this will be "Success!", on
//* failure, it will return the error message.
```

```
//* @Node version
//* @Parent DeviceNode
//*
//* A hierarchy node which holds version value nodes.
```

```
//* @Node versionNumber
//* @Parent version
//*
//* String which holds the full version number.
//*
//* @Value string
```

```
//* @Node releaseDate
//* @Parent version
//*
//* String which holds the release date of the current version.
//*
//* @Value string
```

## Output

The above examples will output the following in api.md:

### root  

Type: Node   

Short: Root node of the DsLink  


---

### Add_Device  

Type: Action   
Is: addDeviceCmd   
Parent: [root](#root)  

Short: Adds a device to the link.  

Long: Add Device accepts the URL and name of the device to add to the link. It will verify the URL is accessible and return an error message if it fails.  

Params:  

Name | Type | Description
--- | --- | ---
url | string | The URL of the device.
name | string | The name for the device, it will be added to the link under this name.

Return type: value   
Columns:  

Name | Type | Description
--- | --- | ---
success | bool | A boolean which represents if the action succeeded or not. Returns false on failure and true on success.
message | string | If the action succeeds, this will be "Success!", on failure, it will return the error message.

---

### DeviceNode  

Type: Node   
Is: deviceNode   
Parent: [root](#root)  

Short: A device which has been added to the link.  

Long: When added to the link, a device will appear as the name provided. This node maintains the connection with the remote host.  


---

### Remove_Device  

Type: Action   
Is: removeDeviceCmd   
Parent: [DeviceNode](#DeviceNode)  

Short: Removes a device from the link.  

Return type: value   
Columns:  

Name | Type | Description
--- | --- | ---
success | bool | A boolean which represents if the action succeeded or not. Returns false on failure and true on success.
message | string | If the action succeeds, this will be "Success!", on failure, it will return the error message.

---

### version  

Type: Node   
Parent: [DeviceNode](#DeviceNode)  

Short: A hierarchy node which holds version value nodes.  


---

### versionNumber  

Type: Node   
Parent: [version](#version)  

Short: String which holds the full version number.  

Value Type: string

---

### releaseDate  

Type: Node   
Parent: [version](#version)  

Short: String which holds the release date of the current version.  

Value Type: string

---

```
- root
 |- Add_Device
 |- DeviceNode
 | |- Remove_Device
 | |- version
 | | |- versionNumber
 | | |- releaseDate

```
