validatedoc
====

Validates my extremely specific opinions of what python function docstrings
should look like.

I recommend using this with pydocstyle, which velidates pep257.  This is more
strict and more opinionated than that.

Installing
---
```
pip install validatedoc
```

Usage
---
```
validatedoc file1.py dir/file2.py file2.py
```

More info
---
Example of how it expects docs to be formated:

```
def foo(bar, baz, qux) -> zam:
    """Do foo thing to bar baz and qux and produce zam.

    :param bar: a thing
    :param baz: another thing
    :param qux: more
    :raises FooError: when foo goes wrong
    :raises ValueError: if a value is wrong
    :return: the value of Zam
    """
    ...
```

Specifically, it ensures that:
- There is an empty newline between the description and parameters list
- Parameters are in the format `:param name: description`
- Exception docstrings are in the format `:raises exception: description`
- Return docstrings look like `:return: description`, and come after all raises
    and param doc strings
- There are no empty newlines in the param/raises/return list
- Lines in the param/raises/return list either start an item or have a four
    space indent (to continue the previous line)
- The params match exactly the arguments
- Return docstrings are last
- If the function is annotated to return None, no return should be present
- If the function specifies a return other than none, return is present

In the future, I also want to:
- [ ] Verify that any exceptions raised in the function have a raises docstring
- [ ] Better support \*args and \*\*kwargs (right now they're considered
    normal args that require a single param each if they're in the function
    definition)
- [ ] Validate docstrings on things other than functions
- [ ] Validate the short and long descriptions in some way
- [ ] Something with types
- [ ] Require param/raises/return lines to start with a lowercase or an
    uppercase (and decide which to use)
