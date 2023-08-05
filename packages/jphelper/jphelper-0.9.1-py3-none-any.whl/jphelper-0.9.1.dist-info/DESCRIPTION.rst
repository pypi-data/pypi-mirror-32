jphelper
========
A python library for processing your japanese learning materials.

Installation
------------

.. code::

    pip install jphelper


Usage
-----

Convert romaji to hiragana or katakana.

.. code:: python

    from jphelper import kanaize
    kanaize('ohayou')  # おはよう
    kanaize('puroguramingu', to_katakana=True)  # プログラミング
    kanaize('jyoukyou')  # じょうきょう
    kanaize('gakkou')  # がっこう


Pair kanji and its furigana.

.. code:: python

    from jphelper import match_reading
    match_reading('学校', 'がっこう')  # 学校[がっこう]
    match_reading('青い空', 'あおいそら')  # 青[あお]い 空[そら]
    match_reading('すばしっこい茶色の狐', 'すばしっこいちゃいろのきつね')  #すばしっこい 茶色[ちゃいろ]の 狐[きつね]


Convert arabic numeral to japanese.

.. code:: python

    from jphelper.number import to_japanese
    to_japanese(3907)  # さんぜんきゅうひゃくなな
    to_japanese(524, use_kanji=True)  # 五百二十四
    to_japanese(-8437)  # マイナスはっせんよんひゃくさんじゅうなな
    to_japanese(0.456, decimal_limit=2)  # ゼロてんよんご
    to_japanese(12500, separator='、')  # いち、まん、に、せん、ご、ひゃく、ゼロ


Playing with numbers.

.. code:: python

    from jphelper.number import shorten, kanji_grouping
    shorten(12345678987)  # 123.4億
    kanji_grouping(-15124)  # -1万5124
    kanji_grouping(-2345678.90123, use_hiragana=True, use_minus_sign=False)  # マイナス234まん5678てん90123


