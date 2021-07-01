# Names

If an entity is declared within a function, it is *local* to that function. If declared outside of a function, however, it is visible in all files of the package to which it belongs.

The case of the first letter of a name determines its visibility across package boundaries. If the name begins with an upper-case-letter, it is *exported*, which means that it is visible and accessible outside its own package and may be referred to by other parts of the program.

Package names themselves are always in lower case.

Stylistically, Go programmers use *"camel case"*, that is, interior capital letters are preferred over interior underscores.
