
# Example Walkthrough

I have been given a mysterious SFM file 'dict_data.sfm' to tidy up in
preparation for import to [FLEx][SIL-FLEX]. Let's take a quick look:

### Initial Summary

```
> sfm_sniffer.py --summary dict_data.sfm
\gn : gloss (national)     : occurrences=2480 : type=text                  : exceptions=26
\lx : lexeme               : occurrences=2474 : type=single term           : exceptions=7
\sn : sense number         : occurrences=2456 : type=enumeration           : exceptions=28
\ps : part of speech       : occurrences=2450 : type=enumeration           : exceptions=79
\ge : gloss (english)      : occurrences= 511 : type=optional, word        : exceptions=12
\gr : gloss (regional)     : occurrences= 500 : type=optional, single term : exceptions=11
\glo: ???                  : occurrences= 354 : type=text                  : exceptions=0
\glv: ???                  : occurrences= 346 : type=optional, single term : exceptions=6
\glm: ???                  : occurrences=   5 : type=NULL type             : exceptions=0
\ps_: ???                  : occurrences=   1 : type=word                  : exceptions=0
```

Immediately I see that this SFM file uses the 'standard' MDF backslash
codes with the addition of a few custom tags. I can [modify the
tag definitions](#modifying-tag-definitions) recognised by sfm_sniffer
to generate labels for the unknown tags, although it is not necessary
to do so.

There is a suspicious-looking \ps_ tag that only occurs 1 time. I
treat that as a typo and edit the SFM file to change it to a \ps tag.

Interestingly, the \glm tag only occurs 5 times and has
[NULL type](#null-type), meaning that it has no value. Since there
are 0 exceptions, this tag _never_ has a value and so is redundant.
I choose to edit the SFM file and remove this tag.

### Tag Type Analysis

Looking carefully at the [types](#value-types) I am surprised to
see that \sn ('sense number') has been identified as an _enumeration_
type; I expected it to be a _number_ type. Similarily, \ge
('gloss(english)') has been identified as an _optional, word_; I would
not expect it to be optional, so expected it to be of type _word_.

Let's try stricter type checking:

```
> sfm_sniffer.py --summary --strictest dict_data.sfm
\gn : gloss (national)     : occurrences=2480 : type=text           : exceptions=26
\lx : lexeme               : occurrences=2474 : type=word           : exceptions=149
\sn : sense number         : occurrences=2456 : type=number         : exceptions=41
\ps : part of speech       : occurrences=2451 : type=enumeration    : exceptions=79
\ge : gloss (english)      : occurrences= 511 : type=word           : exceptions=63
\gr : gloss (regional)     : occurrences= 500 : type=optional word  : exceptions=54
\glo: gloss (other)        : occurrences= 354 : type=optional, single term : exceptions=3
\glv: gloss (vernacular)   : occurrences= 346 : type=optional word  : exceptions=23
```

These results reveal:

* \sn ('sense number') is normally a number, but there are 41 instances
when it is not a number. 28 of those instances are not
even recognised values (probably empty values).
* \ge ('gloss (english)') is normally a word. But there are 51 (=63-12)
instances where it has no value. There are 12 instances where it is
a phrase or general text, as opposed to a word.
* \lx ('lexeme') is most commonly a single word. There are 142
(=149-7) instances where it is a phrase as opposed to a word; this
is not a concern. However, the remaining 7 exceptions probably need
fixed.
* \ps ('part of speech') is still considered an enumeration, as
expected, but with 79 unexpected values.

To look in more detail at the exceptions associated with \ps ('part of
speech'), let's run a full report:

```
> sfm_sniffer.py dict_data.sfm
glo: gloss (other)        : occurrences= 354: type=text   : exceptions=0
====================================
\ps : part of speech      : occurrences=2451: type=enumeration
Example values:
adj,adj adv,adj num,adj poss,adj poss.,adj?,adv,adv inter,adv tm,conj,dem,ex,excl,excl.,id,int,int.,inter,interr.,intj,n,num,num.,part.,pers,poss.,pr.,prn,pro,pro pers,pro pers.,pron,pron.,prp,prp.,temp,temp.,temps,v,vi,vt
79 exceptions for \ps of type 'enumeration':
line  855: \ps v. int
line 1875: \ps v. int.
line 1879: \ps <no value>
line 1883: \ps v. int.
line 1897: \ps v. int.
line 1947: \ps <no value>
    >>> snip <<<
line 11187: \ps v. int.
line 11355: \ps v  - v. int. bate
line 11588: \ps <no value>
====================================
\gn : gloss (national)    : occurrences=2480: type=text
26 exceptions for \gn of type 'text':
line  817: \gn <no value>
line  859: \gn <no value>
line  875: \gn <no value>
    >>> snip <<<
line 10491: \gn <no value>
line 10642: \gn <no value>
====================================
\ge : gloss (english)     : occurrences= 511: type=optional, word
12 exceptions for \ge of type 'optional, word':
line  110: \ge and, with
line  530: \ge scorpion, millipede
line  758: \ge cheap, inexpensive
line 1152: \ge pure, holy
    >>> snip <<<
line 11205: \ge lack, absence
line 11333: \ge uncover, reveal
====================================
\gr : gloss (regional)    : occurrences= 500: type=optional, single term
11 exceptions for \gr of type 'optional, single term':
line  759: \gr bootu, koyɗum
line 1227: \gr bappa/(pat.), kawu/(mat.)
line 1481: \gr suŋku, dala
    >>> snip <<<
line 11250: \gr rufugo, yeeraande
line 11431: \gr gulɗum, ŋguleeŋga
====================================
\lx : lexeme              : occurrences=2474: type=single term
7 exceptions for \lx of type 'single term':
line    1: \lx <no value>
line 2335: \lx eptsá - v. int. fatsa
line 2470: \lx ékséɓé, ésséɓá
line 2474: \lx ékslá, alá
line 2712: \lx fá wé...
line 4025: \lx icá  - v.int. ɗatsa
line 11051: \lx ŋá (v.int. ŋɛŋa)
====================================
\glv: gloss (vernacular)  : occurrences= 346: type=optional, single term
6 exceptions for \glv of type 'optional, single term':
line   67: \glv a....de
line   91: \glv atə, al
line  113: \glv atə, al
line  354: \glv ciɓa, barama
line 6622: \glv hərmana (pl. har hay)
line 11208: \glv ɓeca (n.)
```

This shows a number of things...

### Enumerated Type Listings

Since \ps ('part of speech') is an _enumerated_ tag, the list of tag
values is shown:

```
Example values:
adj,adj adv,adj num,adj poss,adj poss.,adj?,adv,adv inter,adv tm,conj,dem,ex,excl,excl.,id,int,int.,inter,interr.,intj,n,num,num.,part.,pers,poss.,pr.,prn,pro,pro pers,pro pers.,pron,pron.,prp,prp.,temp,temp.,temps,v,vi,vt
```

In this case, we can see that there are often several variants of a
value, for example: `ex,excl,excl.` or `int,int.` Using a text editor
with search and replace functionality, the list of values used for the
\ps tag can quickly be rationalised.

Note that the values of `v. int.` are not recognised as valid
because of the repeated period ('.'). These values cannot be
distinguished from text. I could choose to apply search and replace to
substitute an valid enumeration value, e.g. `v-int.`, or I could ignore
these exceptions.

### Modifying Tag Definitions

**TODO**

* describe use of --tag-dict and --dump commands

[SIL-FLEX]: https://software.sil.org/fieldworks/
[SIL-SOLID]: https://software.sil.org/solid/
[MDF_2000]: https://downloads.sil.org/legacy/shoebox/MDF_2000.pdf
