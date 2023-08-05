# Terms Tree Library

Library to build *terms tree* from indented text file and to search terms in that tree.

May be used for text labeling/classification tasks.

See also: [termstree library](https://pypi.org/project/termsquery/)

## Example

Demo Script: 

```python
import termstree

TERMS_TREE_SRC = """
# comment

Asia
    Japan
        Tokyo [url="https://en.wikipedia.org/wiki/Tokyo"]
        Osaka
    China
        Beijing
        Shanghai
Europe
    England
        London

    Germany [url="https://en.wikipedia.org/wiki/Germany"]
        Berlin
        Munich
"""

terms_tree = termstree.build(TERMS_TREE_SRC, terms_normalizer=None)

text = 'During the 16th century, Munich was a centre of the German counter reformation. Europe ...'

for hit in terms_tree.search_in(text):
    print(hit)
```

Result (list of 'hits' - terms found in the text):
```
Hit(node=Node('Munich'), dhits=1, ihits=0)
Hit(node=Node('Europe'), dhits=1, ihits=1)
Hit(node=Node('Germany', {'url': 'https://en.wikipedia.org/wiki/Germany'}), dhits=0, ihits=1)
```

Every hit corresponds to a term from terms tree and has next attributes:

- node - found term
- dhits (direct hits) - number of direct term occurrences in the text
- ihits (indirect hits) - number of term's children occurrences

