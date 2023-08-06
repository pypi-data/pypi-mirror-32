#!/usr/bin/python
# -*- coding: utf-8 -*-
from __future__ import absolute_import, division, print_function, unicode_literals
import csv
import os
import logging
try:
    from StringIO import StringIO
except ImportError:
    from io import StringIO
import requests

from anycsv import dialect
from anycsv import encoding
from anycsv.csv_model import Table
from anycsv import io_tools
from anycsv import exceptions
import gzip
import io

DEFAULT_ENCODING='utf-8'
ENC_PRIORITY=['magic', 'lib_chardet', 'header', 'default']

def reader(filename=None, url=None, content=None, skip_guess_encoding=False, delimiter=None, sniff_lines=100, max_file_size=-1):
    """

    :param filename:
    :param url:
    :param content: The content of a CSV file as a string
    :param skip_guess_encoding: If true, the parser uses utf-8
    :param delimiter:
    :param sniff_lines:
    :return:
    """
    logger = logging.getLogger(__name__)

    if not filename and not url and not content:
        raise exceptions.AnyCSVException('No CSV input specified')

    meta = sniff_metadata(filename, url, content, skip_guess_encoding=skip_guess_encoding, sniffLines=sniff_lines)
    table = Table(url=url, filename=filename)

    table.dialect = meta['dialect']
    if delimiter:
        table.delimiter = delimiter
        if 'delimiter' in table.dialect and table.dialect['delimiter'] != delimiter:
            logger.warning('The given delimiter differs from the guessed delimiter: ' + dialect['delimiter'])
    elif 'delimiter' in table.dialect:
        table.delimiter = table.dialect['delimiter']
    else:
        raise exceptions.NoDelimiterException('No delimiter detected')

    if 'quotechar' in table.dialect:
        table.quotechar = table.dialect['quotechar']

    table.encoding = meta['used_enc']

    if content:
        if max_file_size!=-1  and len(content)> max_file_size:
            raise exceptions.FileSizeException("Maximum file size exceeded {} > {} ".format(len(content), max_file_size))
        input = io.TextIOBase(content, encoding=table.encoding)
    elif filename and os.path.exists(filename):
        if filename[-3:] == '.gz':
            if max_file_size != -1 and os.stat(filename).st_size/0.4 > max_file_size: #assuming 40% compression / com/orig=0.4 -> orig = com/0.4
                raise exceptions.FileSizeException(
                    "Maximum file size exceeded {} > {} ".format(os.stat(filename).st_size, max_file_size))

            input = io.TextIOWrapper(gzip.open(filename, 'r'), encoding=table.encoding)
        else:
            if max_file_size != -1 and os.stat(filename).st_size  > max_file_size:
                raise exceptions.FileSizeException(
                    "Maximum file size exceeded {} > {} ".format(os.stat(filename).st_size, max_file_size))
            input = io.open(filename,encoding=table.encoding)
    elif url:
        input = URLHandle(url,max_file_size,encoding=table.encoding)
    else:
        raise exceptions.AnyCSVException('No CSV input specified')






    if table.encoding and ('utf-8' in table.encoding or 'utf8' in table.encoding):
        table.csvReader = UnicodeReader(input,
                                        delimiter=table.delimiter,
                                        quotechar=table.quotechar)
    else:
        table.csvReader = EncodedCsvReader(input,
                                     encoding=table.encoding,
                                     delimiter=table.delimiter,
                                     quotechar=table.quotechar)

    return table


def sniff_metadata(fName= None, url=None, content=None, header=None, sniffLines=100, skip_guess_encoding=False):
    logger = logging.getLogger(__name__)
    id = url if url is not None else fName

    if not any([fName,url]) and not any([content, header]):
        #we need at least one of the three, so return empty results
        return {}

    if not any([content, header]) and any([fName, url]):
        res = io_tools.getContentAndHeader(fName=fName, url=url, download_dir="/tmp/", max_lines=sniffLines)
        content, header = res['content'], res['header']

    
    logger.debug('(%s) Extracting CSV meta data ', id)
    meta = extract_csv_meta(header=header, content=content, skip_guess_encoding=skip_guess_encoding)
    logger.debug("(%s) Meta %s ", id, meta)

    return meta


def extract_csv_meta(header, content=None, id='', skip_guess_encoding=False):
    logger = logging.getLogger(__name__)
    results = {'used_enc': None,
               'dialect': {}}

    # check if guess encoding is possible
    if not skip_guess_encoding:
        try:
            import anycsv.encoding
        except Exception as e:

            print('Could not import "magic" library. To support encoding detection please install python-magic.')
            skip_guess_encoding = True

    # get encoding
    if skip_guess_encoding:
        results['used_enc'] = DEFAULT_ENCODING
        content_encoded = content#.decode(encoding=results['used_enc'])
        status="META encoding"
    else:
        results['enc'] = encoding.guessEncoding(content, header)

        content_encoded = None
        status="META "
        c_enc = None
        for k in ENC_PRIORITY:
            #we try to use the different encodings
            try:
                if k in results['enc'] and results['enc'][k]['encoding'] is not None:
                    content_encoded = content.decode(encoding=results['enc'][k]['encoding'])
                    c_enc = results['enc'][k]['encoding']
                    status+=" encoding"
                    break
            except Exception as e:
                logger.debug('(%s) ERROR Tried %s encoding: %s', results['enc'][k]['encoding'],id, e)
        if content_encoded:
            results['used_enc'] = c_enc

    # get dialect
    try:
        results['dialect'] = dialect.guessDialect(content_encoded)
        status+=" dialect"
    except Exception as e:
        logger.warning('(%s)  %s',id, str(e))
        results['dialect']={}

    #if fName:
    #    results['charset'] = encoding.get_charset(fName)

    logger.debug("(%s) %s", id, status)
    return results


class URLHandle:
    def __init__(self, url, max_file_size, encoding=None):
        self.url = url
        self.encoding = encoding
        self._init()
        self.max_file_size=max_file_size

    def my_iter_lines(self,result):
        buf = b''
        for chunk in result.iter_content(chunk_size=64 * 1024):
            buf += chunk
            pos = 0
            while True:
                eol = buf.find(b'\n', pos)
                if eol != -1:
                    yield buf[pos:eol]
                    pos = eol + 1
                else:
                    buf = buf[pos:]
                    break
        if buf:
            yield buf

    def _init(self):
        self._count = 0
        req = requests.get(self.url)
        #if self.encoding:
        #self.input = req.iter_lines(chunk_size=1024, encoding=self.encoding   )
        #else:
        self.input = self.my_iter_lines(req)

    def seek(self, offset):
        if offset < self._count:
            self._init()
        while offset < self._count:
            self.next()

    def __iter__(self):
        return self

    def __next__(self):
        _next = next(self.input)#.next()
        self._count += len(_next)
        if self.max_file_size != -1 and self._count > self.max_file_size:
            raise exceptions.FileSizeException(
                "Maximum file size exceeded {} > {} ".format(self._count, self.max_file_size))

        if self.encoding:
            return _next.decode(self.encoding)
        else:
            return str(_next)

    def next(self):
        return self.__next__()



class CsvReader:

    def __init__(self, f, reader, encoding):
        self.f = f
        self.reader = reader
        self.encoding = encoding
        self._start_line = 0
        self.line_num = 0

    def __iter__(self):
        return self

    def seek_line(self, line_number):
        if line_number < self.line_num:
            self.f.seek(0)
            self.line_num = 0
        self._start_line = line_number

    def _next(self):
        while self._start_line > self.line_num:
            next(self.reader)#.next()
            self.line_num += 1
        row = self.reader.__next__()
        self.line_num += 1
        return row


class EncodedCsvReader(CsvReader):
    def __init__(self, f, encoding="utf-8-sig", delimiter="\t", quotechar="'", **kwds):
        if not quotechar:
            quotechar = "'"
        if not encoding:
            encoding = 'utf-8-sig'
        if not delimiter:
            reader = csv.reader(f, quotechar=quotechar.encode('ascii'), **kwds)
        else:
            reader = csv.reader(f, delimiter=delimiter, quotechar=quotechar,
                                     **kwds)
        CsvReader.__init__(self, f, reader, encoding)

    def __next__(self):
        row = self._next()
        result = [str(s) for s in row]
        return result

    def next(self):
        return self.__next__()


class UnicodeReader(CsvReader):
    def __init__(self, f, delimiter="\t", quotechar="'", encoding='utf-8', errors='strict', **kwds):
        if not quotechar:
            quotechar = "'"
        if not encoding:
            encoding = 'utf-8'
        if not delimiter:
            reader = csv.reader(f, quotechar=quotechar.encode('ascii'), **kwds)
        else:
            reader = csv.reader(f, delimiter=delimiter.encode('ascii'), quotechar=quotechar.encode('ascii'),
                                     **kwds)
        self.encoding_errors = errors
        CsvReader.__init__(self, f, reader, encoding)

    def next(self):
        row = self._next()
        encoding = self.encoding
        encoding_errors = self.encoding_errors
        float_ = float
        unicode_ = str
        return [(value if isinstance(value, float_) else
                 unicode_(value, encoding, encoding_errors)) for value in row]


