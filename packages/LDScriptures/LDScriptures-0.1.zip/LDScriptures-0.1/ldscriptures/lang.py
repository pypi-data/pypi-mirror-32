# coding: latin-1
from .utils import *
from . import exceptions
from . import utils

default = 'eng'

langs = ['eng', 'por', 'spa']

class Eng:  # English

    ot_books = ['Genesis', 'Exodus', 'Leviticus', 'Numbers', 'Deuteronomy', 'Joshua', 'Judges', 'Ruth', '1 Samuel', '2 Samuel', '1 Kings', '2 Kings',
                '1 Chronicles', '2 Chronicles', 'Ezra', 'Nehemiah', 'Esther', 'Job', 'Psalms', 'Proverbs', 'Ecclesiastes', 'Song of Solomon', 'Isaiah',
                'Jeremiah', 'Lamentations', 'Ezekiel', 'Daniel', 'Hosea', 'Joel', 'Amos', 'Obadiah', 'Jonah', 'Micah', 'Nahum', 'Habakkuk', 'Zephaniah',
                'Haggai', 'Zechariah', 'Malachi']

    nt_books = ['Matthew', 'Mark', 'Luke', 'John', 'Acts', 'Romans', '1 Corinthians', '2 Corinthians', 'Galatians', 'Ephesians', 'Philippians',
                'Colossians', '1 Thessalonians', '2 Thessalonians', '1 Timothy', '2 Timothy', 'Titus', 'Philemon', 'Hebrews', 'James', '1 Peter',
                '2 Peter', '1 John', '2 John', '3 John', 'Jude', 'Revelation']

    bofm_books = ['1 Nephi', '2 Nephi', 'Jacob', 'Enos', 'Jarom', 'Omni', 'Words of Mormon', 'Mosiah', 'Alma', 'Helaman', '3 Nephi', '4 Nephi',
                  'Mormon', 'Ether', 'Moroni']

    pgp_books = ['Moses', 'Abraham', 'Joseph Smith-Matthew', 'Joseph Smith-History', 'Articles of Faith']
    
    dc_testament = ['Doctrine and Covenants']

class Spa:  # Spanish
 
    ot_books = ['G�nesis', '�xodo', 'Lev�tico', 'N�meros', 'Deuteronomio', 'Josu�', 'Jueces', 'Rut', '1 Samuel', '2 Samuel', '1 Reyes', '2 Reyes',
                '1 Cr�nicas', '2 Cr�nicas', 'Esdras', 'Nehem�as', 'Ester', 'Job', 'Salmos', 'Proverbios', 'Eclesiast�s', 'Cantares', 'Isa�as',
                'Jerem�as', 'Lamentaciones', 'Ezequiel', 'Daniel', 'Oseas', 'Joel', 'Am�s', 'Abd�as', 'Jon�s', 'Miqueas', 'Nah�m', 'Habacuc',
                'Sofon�as', 'Hageo', 'Zacar�as', 'Malaqu�as']
    
    nt_books = ['Mateo', 'Marcos', 'Lucas', 'Juan', 'Hechos', 'Romanos', '1 Corintios', '2 Corintios', 'G�latas', 'Efesios', 'Filipenses',
                'Colosenses', '1 Tesalonicenses', '2 Tesalonicenses', '1 Timoteo', '2 Timoteo', 'Tito', 'Filem�n', 'Hebreos', 'Santiago', '1 Pedro',
                '2 Pedro', '1 Juan', '2 Juan', '3 Juan', 'Judas', 'Apocalipsis']
    
    bofm_books = ['1 Nefi', '2 Nefi', 'Jacob', 'En�s', 'Jarom', 'Omni', 'Palabras de Morm�n', 'Mos�ah', 'Alma', 'Helam�n', '3 Nefi', '4 Nefi',
                  'Morm�n', '�ter', 'Moroni']
    
    pgp_books = ['Mois�s', 'Abraham', 'Jos� Smith-Mateo', 'Jos� Smith-Historia', 'Art�culos de Fe']
    
    dc_testament = ['Doctrina y Convenios']

class Por:  # Portuguese

    ot_books = ['G�nesis', '�xodo', 'Lev�tico', 'N�meros', 'Deuteron�mio', 'Josu�', 'Ju�zes', 'Rute', '1 Samuel', '2 Samuel', '1 Reis', '2 Reis',
                '1 Cr�nicas', '2 Cr�nicas', 'Esdras', 'Neemias', 'Ester', 'J�', 'Salmos', 'Prov�rbios', 'Eclesiastes', 'Cantares de Salom�o',
                'Isa�as', 'Jeremias', 'Lamenta��es', 'Ezequiel', 'Daniel', 'Oseias', 'Joel', 'Am�s', 'Obadias', 'Jonas', 'Miqueias', 'Naum',
                'Habacuque', 'Sofonias', 'Ageu', 'Zacarias', 'Malaquias']
    
    nt_books = ['Mateus', 'Marcos', 'Lucas', 'Jo�o', 'Atos', 'Romanos', '1 Cor�ntios', '2 Cor�ntios', 'G�latas', 'Ef�sios', 'Filipenses',
                'Colossenses', '1 Tessalonicenses', '2 Tessalonicenses', '1 Tim�teo', '2 Tim�teo', 'Tito', 'Filemom', 'Hebreus', 'Tiago', '1 Pedro',
                '2 Pedro', '1 Jo�o', '2 Jo�o', '3 Jo�o', 'Judas', 'Apocalipse']
    
    bofm_books = ['1 N�fi', '2 N�fi', 'Jac�', 'Enos', 'Jarom', '�mni', 'Palavras de M�rmon', 'Mosias', 'Alma', 'Helam�', '3 N�fi', '4 N�fi',
                  'M�rmon', '�ter', 'Mor�ni']
    
    pgp_books = ['Mois�s', 'Abra�o', 'Joseph Smith-Mateus', 'Joseph Smith-Hist�ria', 'Regras de F�']
    
    dc_testament = ['Doutrina e Conv�nios']


eng = Eng
spa = Spa
por = Por


def lang_verify(language):
    if language in langs:
        return eval(language)
    else:
        raise exceptions.InvalidLang('The language "{}" is not a valid language. Try one: {}.'.format(str(language), str(langs)))


def match_scripture(book_name, language):
    language = lang_verify(language)
    book_name = book_name.lower()
    scripture = ''
    
    if book_name in utils.lower_list(language.ot_books):
        scripture = 'ot'
    elif book_name in utils.lower_list(language.nt_books):
        scripture = 'nt'
    elif book_name in utils.lower_list(language.bofm_books):
        scripture = 'bofm'
    elif book_name in utils.lower_list(language.pgp_books):
        scripture = 'pgp'
    elif book_name in utils.lower_list(language.dc_testament):
        scripture = 'dc-testament'
    else:
        raise exceptions.InvalidBook('The book \'{}\' does not exist.'.format(str(book_name)))
    
    return scripture


def get_book_code(book, language):
    language = lang_verify(language)
    book = book.lower()
    
    codes = None
    
    if book in lower_list(language.ot_books):
        scripture = language.ot_books
        codes = utils.ot_data['codes']
    elif book in lower_list(language.nt_books):
        scripture = language.nt_books
        codes = utils.nt_data['codes']
    elif book in lower_list(language.bofm_books):
        scripture = language.bofm_books
        codes = utils.bofm_data['codes']
    elif book in lower_list(language.pgp_books):
        codes = utils.pgp_data['codes']
        scripture = language.pgp_books
    elif book in lower_list(language.dc_testament):
        codes = ['dc']
        scripture = language.dc_testament
    else:
        raise exceptions.InvalidBook('The book \'{}\' does not exist.'.format(str(book)))
    
    return codes[item_position(book, scripture)]


def translate_book_name(from_lang, to_lang, book_name):
    from_lang = lang_verify(from_lang)
    to_lang = lang_verify(to_lang)
    
    scripture = match_scripture(book_name, from_lang)
    
    if scripture == 'ot':
        position = item_position(book_name, from_lang.ot_books)
        if position != -1:
            return to_lang.ot_books[position]
            
    if scripture == 'nt':
        position = item_position(book_name, from_lang.nt_books)
        if position != -1:
            return to_lang.nt_books[position]
    
    if scripture == 'bofm':
        position = item_position(book_name, from_lang.bofm_books)
        if position != -1:
            return to_lang.bofm_books[position]
    
    if scripture == 'pgp':
        position = item_position(book_name, from_lang.pgp_books)
        if position != -1:
            return to_lang.pgp_books[position]
