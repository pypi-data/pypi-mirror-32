LunarCalendar: A Lunar-Solar Converter
======================================

.. image::
  https://img.shields.io/pypi/v/LunarCalendar.svg
  :target: https://pypi.python.org/pypi/LunarCalendar
  :alt: Last stable version (PyPI)

Overview
--------

If you aren't familiar with English, can read `Chinese(中文) <https://github.com/wolfhong/LunarCalendar/blob/develop/README_zh.rst>`_ here.

LunarCalendar is a Lunar-Solar Converter, containing a number of lunar and solar holidays in China.

Korean and Japanese lunar calendar is the same as Chinese calendar, but has different holidays.
If `LunarCalendar` is extended, Korean and Japanese holidays are easily included, with their languages.

LunarCalendar supports the time range 1900-2100. But if you have a need for the time range, you can use ``generate.html`` to extend it.

LunarCalendar is inspired by `Lunar-Solar-Calendar-Converter <https://github.com/isee15/Lunar-Solar-Calendar-Converter>`_.


Features
--------

* Accurate raw data, synchronize with Microsolf's ``ChineseLunisolarCalendar`` class
* Easy to extend holidays and languages
* included Lunar Festivals, such as: MidAutumn Festival, Chinese New Year Eve, DragonBoat Festivals
* included Solar Festivals without fixed dates, such as: Mother's Day, Easter
* Added legality check of the lunar and solar date


Install
-------

LunarCalendar can be installed from the PyPI with `easy_install`::

   $ easy_install LunarCalendar

Or pip::

   $ pip install LunarCalendar


Console Commands
----------------

A console command called `lunar-find` can be used to find the date of the festival, using it's chinese name.
Default to this year. Supporting alias of the festival.

.. code-block:: console

    $ lunar-find 重阳
    重阳节 on 2018: 2018-10-17

    $ lunar-find 登高节 2019
    重阳节 on 2019: 2019-10-07


Quickstart
----------

Solar to Lunar:

.. code-block:: python

    import datetime
    from lunarcalendar import Converter, Solar, Lunar, DateNotExist

    solar = Solar(2018, 1, 1)
    print(solar)
    lunar = Converter.Solar2Lunar(solar)
    print(lunar)
    solar = Converter.Lunar2Solar(lunar)
    print(solar)
    print(solar.to_date(), type(solar.to_date()))

Lunar to Solar:

.. code-block:: python

    lunar = Lunar(2018, 2, 30, isleap=False)
    print(lunar)
    solar = Converter.Lunar2Solar(lunar)
    print(solar)
    lunar = Converter.Solar2Lunar(solar)
    print(lunar)
    print(lunar.to_date(), type(lunar.to_date()))
    print(Lunar.from_date(datetime.date(2018, 4, 15)))

Legality check for solar and lunar date. 2018-2-15(Leap Month) does not exist, but 2012-4-4(Leap Month) exists:

.. code-block:: python

    Lunar(2012, 4, 4, isleap=True)  # date(2012, 5, 24)
    try:
        lunar = Lunar(2018, 2, 15, isleap=True)
    except DateNotExist:
        print(traceback.format_exc())

Print all the festivals included, with Chinese and English. Other languages are welcome to extend(Fork & Pull Request).

.. code-block:: python

    from lunarcalendar.festival import festivals

    # print festivals, using English or Chinese
    print("----- print all festivals on 2018 in chinese: -----")
    for fest in festivals:
        print(fest.get_lang('zh'), fest(2018))

    print("----- print all festivals on 2017 in english: -----")
    for fest in festivals:
        print(fest.get_lang('en'), fest(2017))

Output:

.. code-block:: shell

    ......
    母亲节 2018-05-13
    父亲节 2018-06-17
    中秋节 2018-09-24
    感恩节 2018-11-22
    重阳节 2018-10-17
    春节 2018-02-16
    中元节 2018-08-25
    七夕节 2018-08-17
    腊八节 2019-01-13
    清明节 2018-04-05
    除夕 2019-02-04
    寒衣节 2018-11-08
    元宵节 2018-03-02
    龙抬头 2018-03-18
    端午节 2018-06-18
    ......


Contribution
------------

Including festival standards:

* Common holidays in the the country, such as: Christmas, Halloween, etc.
* Lunar holidays.
* Solar holidays without fixed dates, such as: Mother's Day, Easter, etc.

Supporting Chinese and English only now. If you want to add Korean or Japanese supports, modify ``lunarcalendar/festival.py`` to add holidays and languages.

Some unusual holidays may not be included, `welcom to extend <https://github.com/wolfhong/LunarCalendar/issues>`_.



About
-----

* `Homepage <http://github.com/wolfhong/LunarCalendar>`_
* `PyPI <https://pypi.python.org/pypi/LunarCalendar>`_
* `Issue tracker <https://github.com/wolfhong/LunarCalendar/issues?status=new&status=open>`_
