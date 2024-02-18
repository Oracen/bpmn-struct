# BPMN-Struct

An open Golang implementation of the BPMN spec, [as defined here](https://www.omg.org/spec/BPMN/2.0/PDF).

The goal is to implement the entire spec as closely as possible for use in other projects.

## Implementation Notes

### Folder structure

The folder structure is designed to be as close as possible to the way that the BPMN spec itself is documented. This means mirroring the section/subsection structure. While this isn't very idomatic, and makes imports longer than they need to be, the idea is to make the code map onto the section of the spec that describes its use. I stopped short of using the section numbers because that was a bridge too far.

Besides, if you're using BPMN then I'm going to assume a bit of old-fashioned code structure + design doco isn't going to scare you!

### Specifying optional values

Instead of using pointers to specify optional values, I've instead used arrays with 0 or 1 elements. For one, this is closer in form to the actual spec. Two, this is designed to prevent shared memory and keep everything passed by value. I'll work out how to link in the XML and JSON later.
