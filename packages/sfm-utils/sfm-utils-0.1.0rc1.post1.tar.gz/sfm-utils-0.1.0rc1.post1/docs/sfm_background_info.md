# Background

Standard Format Marker files contain lexicographical data structured
using tags (backslash codes). For example:

```
\lx dapan
\ps m
\ge spear
\de three-pronged spear with barbs, used for eels
\ee This is similar to the unbarbed fv:nasel used for crayfish.
\mr dapa-n
\dt 14/Apr/93
\lx gebhaa
\ps n
\sn 1
\ge husband
\lt big person.
\lf SynD = namorit
\le Rana dialect
\sn 2
\ge clan_head
\lf SynD = tean elen
\le Rana dialect
\mr geba-haa
\cf haa
\ce big, important, loud
```

A significant difficulty inherent in such SFM data is the lack of a
defined schema. Each tag has a meaning. Some tags should only be used
with a certain values. The tags should form a consistent hierarchical
tree structure. However, the lack of defined schema - and the fact that
SFM files are typically edited by multiple people - means that SFM data
frequently contains inconsistencies and errors.

There is an inherent structure in such SFM data (although it may be
hard to see). Tools such as [FLEx][FLEX]
