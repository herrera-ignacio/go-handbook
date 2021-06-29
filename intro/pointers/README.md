# Pointers

Go provides pointers, that is, **values that contain the address of a variable**. In some languages, notably C, pointers are relatively unconstrained. In other languages, pointers are disguised as "references", and there's not much that can be done with them except pass them around. Go takes a position somewhere in the middle. **Pointers are explicitly visible**. The `&` operator yields the address of a variable, and the `*` operator retrieves the variable that the pointer refers to, but **there is no pointer arithmetic**.
