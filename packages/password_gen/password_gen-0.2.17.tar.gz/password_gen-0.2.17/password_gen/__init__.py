# During my recent encounter on an article about the vulnerability of weak passwords, I was enlightened on the dangers of weak passwords, which includes single-handedly serving as the "weakest link" to a web server. In details, if my weak password is guessed by an hacker, the whole web server of the site is vulnerable.

# Olu Gbadebo
# Aug 21st 2017
# Password Generator: generates specifically rendom passwords.

# Each password generated as a format (this helps me to remember the password):
# 1: a digit, a randomized-case word, a special character, another word, another special character, a third word, a third special character, one last word, second digit
# 2: a special character, a randomized-case word, a digit, another word, another digit, third word, third digit, last word, special character
# 3: a randomized-case word, a digit, another word, another digit, third word, third digit, last word
# 4: a randomized-case word, a special character, another word, another special character, third word, third special character, last word

# imports
from pkg_resources import resource_filename, Requirement
import os
import re
import sys
import random
import warnings
import linecache

# range of digits
DIGITS = [0, 9]
# range of special characters (ascii)
CHARACTERS = [33, 47]

# list of words
DICTIONARY = []

# pattern of passwords
PATTERN = []

# randomizers
# returns a digit
def randomize_digit(dig_range):
    return random.randint(dig_range[0], dig_range[1])

# retuns a word with varied cases
def randomize_case(size):
    random_word = get_word(size)
    result_word = ''
    for every_character in random_word:
        # frequency of capitalizing a every_character
        if (random.randint(0,4) == 1):
            result_word += str(every_character).upper()
        else:
            result_word += every_character
    return result_word

# returns a special character
def randomize_character(char_range):
    if (char_range[0] >= 33 and char_range[1] <= 47):
        return chr(random.randint(char_range[0], char_range[1]))
    else:
        raise ValueError('Invalid special character range')

# get a random word from list of most common words
def get_word(size):
    if len(DICTIONARY) is not 0 and size is 0:
        return DICTIONARY[random.randint(0, len(DICTIONARY) - 1)]
    elif len(DICTIONARY) is not 0 and size is not 0:
        c_DICT = DICTIONARY
        random_index = random.randint(0, len(c_DICT) - 1)
        while len(c_DICT) != 0 and len(c_DICT[random_index]) != size:
            if len(c_DICT) == 0:
                print("** Specify a size that matches the lenght of at least one word in your dictionary **\n")
            c_DICT.pop(random_index)
            random_index = random.randint(0, len(c_DICT) - 1)
        return c_DICT[random_index]
    # get a word from a random line in dictionary.txt and remove the trailing white space
    try:
        filename = resource_filename("password_gen", "data/dictionary.txt")
        filename_e = resource_filename("password_gen", "data/dictionary-extra.txt")
        d = open(filename, 'r')
        d_e = open(filename_e, 'r')
    except Exception as e:
        raise Exception("Dictionary file open failed")
    else:
        if size == 0:
            return linecache.getline(filename, random.randint(1, 8829)).split()[0]
        else:
            random_word = linecache.getline(filename_e, random.randint(1, 8829)).split()[0]
            while len(random_word) != size:
                random_word = linecache.getline(filename_e, random.randint(1, 8829)).split()[0]
            return random_word

# populate with default ranges
def populate_digit(dig_range):
    if len(dig_range) is not 2:
        raise ValueError('Insufficient or excessive digit range array provided')
    if dig_range[0] >= 0 and dig_range[1] <= 9:
        DIGITS[0] = int(dig_range[0])
        DIGITS[1] = int(dig_range[1])
    else:
        raise ValueError('Invalid range for digits')

def populate_char(char_range):
    if len(char_range) is not 2:
        raise ValueError('Insufficient or excessive special characters range array provided')
    if char_range[0] >= 33 and char_range[1] <= 47:
        CHARACTERS[0] = int(char_range[0])
        CHARACTERS[1] = int(char_range[1])
    else:
        raise ValueError('Invalid ascii range for special characters')

def populate_dict(dictionary):
    # empty current dictionary
    del DICTIONARY[:]
    # populate dictionarywith new words
    for every_word in dictionary:
        DICTIONARY.append(every_word)

def pattern_occurence(pointer, pattern):
    counter = 1
    if pattern[pointer] == 'd':
        while pointer + 1 < len(pattern) and  pattern[pointer] == pattern[pointer + 1]:
            counter += 1
            pointer += 1
        return [[counter, 'd'], pointer + 1]
    elif pattern[pointer] == 'w':
        try:
            size = int(pattern[pointer + 1])
        except Exception as e:
            raise ValueError("A number must follow 'w' to indicate the length of the word.")
        else:
            return [[size, 'w'], pointer + 2]
    else:
        return [[1, pattern[pointer]], pointer + 1]

def parse_pattern(pattern):
    pattern = pattern.lower()

    if len(re.split('[^dw0-9\W]', pattern)) != 1:
        raise ValueError("Invalid pattern")

    pointer = 0
    while pointer < len(pattern):
        temp_pattern, pointer = pattern_occurence(pointer, pattern)
        PATTERN.append(temp_pattern)

def setAllDefault():
    del PATTERN[:]
    del DICTIONARY[:]
    DIGITS = [0, 9]
    CHARACTERS = [33, 47]

def create_password():
    if len(PATTERN) != 0:
        password = pattern_password()
    # randomly select a format
    pw_type = random.randint(1, 4)
    if pw_type == 1:
        password = password_style_1()
    elif pw_type == 2:
        password = password_style_2()
    elif pw_type == 3:
        password = password_style_3()
    else:
        password = password_style_4()

    setAllDefault()

    return password

# create password
def password_style_1():
    # 1: a digit, a randomized-case word, a special character, another word, another special character, a third word, a third special character, one last word, second digit
    return str(randomize_digit(DIGITS)) + randomize_case(0) + randomize_character(CHARACTERS) + randomize_case(0) + randomize_character(CHARACTERS) + randomize_case(0) + randomize_character(CHARACTERS) + randomize_case(0) + str(randomize_digit(DIGITS))

def password_style_2():
    # 2: a special character, a randomized-case word, a digit, another word, another digit, third word, third digit, last word, special character

    return randomize_character(CHARACTERS) + randomize_case(0) + str(randomize_digit(DIGITS)) + randomize_case(0) + str(randomize_digit(DIGITS)) + randomize_case(0) + str(randomize_digit(DIGITS)) + randomize_case(0) + randomize_character(CHARACTERS)

def password_style_3():
    # 3: a randomized-case word, a digit, another word, another digit, third word, third digit, last word

    return randomize_case(0) + str(randomize_digit(DIGITS)) + randomize_case(0) + str(randomize_digit(DIGITS)) + randomize_case(0) + str(randomize_digit(DIGITS)) + randomize_case(0)

def password_style_4():
    # 4: a randomized-case word, a special character, another word, another special character, third word, third special character, last word

    return randomize_case(0) + randomize_character(CHARACTERS) + randomize_case(0) + randomize_character(CHARACTERS) + randomize_case(0) + randomize_character(CHARACTERS) + randomize_case(0)

def pattern_password():
    password = ""
    for el in PATTERN:
        if el[1] == 'd':
            for i in range(el[0]):
                password += str(randomize_digit(DIGITS))
        elif el[1] == 'w':
            password += randomize_case(el[0])
        else:
            password += el[1]
    return password

def generate(pattern=None, dictionary=None, digits=None, characters=None):
    try:
        warnings.warn('Parameters not provided will be replaced with default values.')
        if pattern:
            parse_pattern(pattern)
        if dictionary:
            populate_dict(dictionary)
        if digits:
            populate_digit(digits)
        if characters:
            populate_char(characters)
    except Exception as error:
        raise error
    else:
        return create_password()

if __name__ == '__main__':
    print( generate() )
