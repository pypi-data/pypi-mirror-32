``pseudol10nutil``
==================

Python module for performing pseudo-localization on strings.  Tested against Python 2, Python3, PyPy and PyPy3.


Dependencies
------------

This package has the following external dependencies:

* `six <https://pythonhosted.org/six/>`_ - for Python 2 to 3 compatibility


``PseudoL10nUtil`` class
------------------------

Class for pseudo-localizing strings.  The class currently has the following members:

- ``transforms`` - field that contains the list of transforms to apply to the string.  The transforms will be applied in order.  Default is ``[transliterate_diacritic, pad_length, square_brackets]``
- ``pseudolocalize(s)`` - method that returns a new string where the transforms to the input string ``s`` have been applied.


``pseudol10nutil.transforms`` module
------------------------------------

The following transforms are currently available:

- ``transliterate_diacritic`` - Takes the input string and returns a copy with diacritics added e.g. ``Hello`` -> ``Ȟêĺĺø``.
- ``transliterate_circled`` - Takes the input string and returns a copy with circled versions of the letters e.g. ``Hello`` -> ``Ⓗⓔⓛⓛⓞ``
- ``transliterate_fullwidth`` - Takes the input string and returns a copy with the letters converted to their fullwidth counterparts e.g. ``Hello`` -> ``Ｈｅｌｌｏ``
- ``pad_length`` - Appends a series of characters to the end of the input string to increase the string length per `IBM Globalization Design Guideline A3: UI Expansion <https://www-01.ibm.com/software/globalization/guidelines/a3.html>`_.
- ``angle_brackets`` - Surrounds the input string with '《' and '》' characters.
- ``curly_brackets`` - Surrounds the input string with '❴' and '❵' characters.
- ``square_brackets`` - Surrounds the input string with '⟦' and '⟧' characters.


Format string support
---------------------

When performing pseudo-localization on a string, the process will skip performing pseudo-localization on format strings.  Python style format strings (e.g. ``{foo}``) and printf style format strings (e.g. ``%s``) are supported.  For example::

   Input [1]: Source {source1} returned 0 rows.
   Output [1]: '⟦Șøüȓċê {source1} ȓêťüȓñêđ 0 ȓøẁš.﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ𝟘🚦﹎ЍאǆᾏⅧ㈴㋹⟧

   Input [2]: Source %(source2)s returned 1 row.
   Output [2]: ⟦Șøüȓċê %(source2)s ȓêťüȓñêđ 1 ȓøẁ.﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ𝟘🚦﹎ЍאǆᾏⅧ㈴㋹퓛⟧

   Input [3]: Source %s returned %d rows.
   Output [3]: ⟦Șøüȓċê %s ȓêťüȓñêđ %d ȓøẁš.﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ𝟘🚦﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ⟧


Example usage
-------------

Python 3 example::


   >>> from pseudol10nutil import PseudoL10nUtil
   >>> util = PseudoL10nUtil()
   >>> s = u"The quick brown fox jumps over the lazy dog."
   >>> util.pseudolocalize(s)
   '⟦Ťȟê ʠüıċǩ ƀȓøẁñ ƒøẋ ǰüɱƥš øṽêȓ ťȟê ĺàźÿ đøğ.﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ𝟘🚦﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ𝟘🚦﹎Ѝא⟧'
   >>> import pseudolocalize.transforms
   >>> util.transforms = [pseudol10nutil.transforms.transliterate_fullwidth, pseudol10nutil.transforms.curly_brackets]
   >>> util.pseudolocalize(s)
   '❴Ｔｈｅ ｑｕｉｃｋ ｂｒｏｗｎ ｆｏｘ ｊｕｍｐｓ ｏｖｅｒ ｔｈｅ ｌａｚｙ ｄｏｇ.❵'
   >>> util.transforms = [pseudol10nutil.transforms.transliterate_circled, pseudol10nutil.transforms.pad_length, pseudol10nutil.transforms.angle_brackets]
   >>> util.pseudolocalize(s)
   '《Ⓣⓗⓔ ⓠⓤⓘⓒⓚ ⓑⓡⓞⓦⓝ ⓕⓞⓧ ⓙⓤⓜⓟⓢ ⓞⓥⓔⓡ ⓣⓗⓔ ⓛⓐⓩⓨ ⓓⓞⓖ.﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ𝟘🚦﹎ЍאǆᾏⅧ㈴㋹퓛ﺏ𝟘🚦﹎Ѝא》'

License
-------

This is released under an MIT license.  See the ``LICENSE`` file in this repository for more information.


