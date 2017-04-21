---
### golang deep copy

Recursively parse and duplicate structures, maps, slices and
combinations of those for objects composed of the default types.

ignores channels

- is there a sane meaning to duplicating an object with channels?
- possible extension to include creating channels with duplicate depth
  and direction is being considered

---

*Bugs*

Lots still working on simplifying recursion and pointer vs interface
of pointer selection methods

Unset pointers within objects copy don't copy correctly
 
