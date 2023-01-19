# BPMN-Struct

An open Golang implementation of the BPMN spec, [as defined here](https://www.omg.org/spec/BPMN/2.0/PDF).

The goal is to implement the entire spec as closely as possible for use in other projects.

### Implementation Notes
Instead of using pointers to specify optional values, I've instead used arrays with 0 or 1 elements. For one, this is closer in form to the actual spec. Two, this is designed to prevent shared memory and keep everything passed by value. I'll work out how to link in the XML and JSON later.