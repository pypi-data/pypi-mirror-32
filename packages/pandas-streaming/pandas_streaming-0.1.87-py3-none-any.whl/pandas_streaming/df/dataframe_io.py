# -*- coding: utf-8 -*-
"""
@file
@brief Saves a dataframe into a zip files.
"""
import io
import os
import zipfile
import pandas
import numpy


def to_zip(df, zipfilename, zname="df.csv", **kwargs):
    """
    Saves a dataframe into a zip file.
    It can be read by @see fn to_zip.

    @param      df          dataframe or numpy array
    @param      zipfilename a :epkg:`*py:zipfile:ZipFile` or a filename
    @param      zname       a filename in th zipfile
    @param      kwargs      parameters for :epkg:`pandas:to_csv` or
                            :epkg:`numpy:save`
    @return                 zipfilename

    .. exref::
        :title: Saves and reads a dataframe in a zip file
        :tag: dataframe

        This shows an example on how to save and read a
        :epkg:`pandas:dataframe` directly into a zip file.

        .. runpython::
            :showcode:

            import pandas
            from pandas_streaming.df import to_zip, read_zip

            df = pandas.DataFrame([dict(a=1, b="e"),
                                   dict(b="f", a=5.7)])

            name = "dfs.zip"
            to_zip(df, name, encoding="utf-8", index=False)
            df2 = read_zip(name, encoding="utf-8")
            print(df2)

    .. exref::
        :title: Saves and reads a numpy array in a zip file
        :tag: array

        This shows an example on how to save and read a
        :epkg:`numpy:ndarray` directly into a zip file.

        .. runpython::
            :showcode:

            import numpy
            from pandas_streaming.df import to_zip, read_zip

            arr = numpy.array([[0.5, 1.5], [0.4, 1.6]])

            name = "dfsa.zip"
            to_zip(arr, name, 'arr.npy')
            arr2 = read_zip(name, 'arr.npy')
            print(arr2)
    """
    if isinstance(df, pandas.DataFrame):
        stb = io.StringIO()
        ext = os.path.splitext(zname)[-1]
        if ext == '.npy':
            raise ValueError(
                "Extension '.npy' cannot be used to save a dataframe.")
        df.to_csv(stb, **kwargs)
    elif isinstance(df, numpy.ndarray):
        stb = io.BytesIO()
        ext = os.path.splitext(zname)[-1]
        if ext != '.npy':
            raise ValueError(
                "Extension '.npy' is required when saving a numpy array.")
        numpy.save(stb, df, **kwargs)
    else:
        raise TypeError("Type not handled {0}".format(type(df)))
    text = stb.getvalue()

    if isinstance(zipfilename, str):
        ext = os.path.splitext(zipfilename)[-1]
        if ext != '.zip':
            raise NotImplementedError(
                "Only zip file are implemented not '{0}'.".format(ext))
        zf = zipfile.ZipFile(zipfilename, 'w')
        close = True
    elif isinstance(zipfilename, zipfile.ZipFile):
        zf = zipfilename
        close = False
    else:
        raise TypeError(
            "No implementation for type '{0}'".format(type(zipfilename)))

    zf.writestr(zname, text)
    if close:
        zf.close()


def read_zip(zipfilename, zname="df.csv", **kwargs):
    """
    Reads a dataframe from a zip file.
    It can be saved by @see fn read_zip.

    @param      zipfilename a :epkg:`*py:zipfile:ZipFile` or a filename
    @param      zname       a filename in th zipfile
    @param      kwargs      parameters for :epkg:`pandas:read_csv`
    @return                 :epkg:`pandas:dataframe` or :epkg:`numpy:array`
    """
    if isinstance(zipfilename, str):
        ext = os.path.splitext(zipfilename)[-1]
        if ext != '.zip':
            raise NotImplementedError(
                "Only zip file are implemented not '{0}'.".format(ext))
        zf = zipfile.ZipFile(zipfilename, 'r')
        close = True
    elif isinstance(zipfilename, zipfile.ZipFile):
        zf = zipfilename
        close = False
    else:
        raise TypeError(
            "No implementation for type '{0}'".format(type(zipfilename)))

    content = zf.read(zname)
    stb = io.BytesIO(content)
    ext = os.path.splitext(zname)[-1]
    if ext == '.npy':
        df = numpy.load(stb, **kwargs)
    else:
        df = pandas.read_csv(stb, **kwargs)

    if close:
        zf.close()

    return df
