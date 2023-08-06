# cards_imarshall
Python card deck module with the following functionality:
* Implement a method to shuffle the deck
* Get all cards with the same suit from a shuffled deck
* Order a shuffled deck in the following suit and card number (Diamonds [A-K], Clubs [A-K], Hearts [A-K], Spades [A-K])
* Extra Credit, implement at your own sort method.
* Write unit tests to for each method.

### Installation
```
pip install cards-imarshall
```
or download .tar from 
https://pypi.org/project/cards-imarshall/

### Requirements
Python 2.7.x or Python 3.x

### Usage
```python
>>>from cards_imarshall import cards
>>>d1 = cards.Deck()
>>>d1.deck_as_text()
'AD | 2D | 3D | 4D | 5D | 6D | 7D | 8D | 9D | 10D | JD | QD | KD | AC | 2C | 3C | 4C | 5C | 6C | 7C | 8C | 9C | 10C | JC | QC | KC | AH | 2H | 3H | 4H | 5H | 6H | 7H | 8H | 9H | 10H | JH | QH | KH | AS | 2S | 3S | 4S | 5S | 6S | 7S | 8S | 9S | 10S | JS | QS | KS'
>>>d1.deck_as_text(verbose=True)
'Ace of Diamonds | 2 of Diamonds | 3 of Diamonds | 4 of Diamonds | 5 of Diamonds | 6 of Diamonds | 7 of Diamonds | 8 of Diamonds | 9 of Diamonds | 10 of Diamonds | Jack of Diamonds | Queen of Diamonds | King of Diamonds | Ace of Clubs | 2 of Clubs | 3 of Clubs | 4 of Clubs | 5 of Clubs | 6 of Clubs | 7 of Clubs | 8 of Clubs | 9 of Clubs | 10 of Clubs | Jack of Clubs | Queen of Clubs | King of Clubs | Ace of Hearts | 2 of Hearts | 3 of Hearts | 4 of Hearts | 5 of Hearts | 6 of Hearts | 7 of Hearts | 8 of Hearts | 9 of Hearts | 10 of Hearts | Jack of Hearts | Queen of Hearts | King of Hearts | Ace of Spades | 2 of Spades | 3 of Spades | 4 of Spades | 5 of Spades | 6 of Spades | 7 of Spades | 8 of Spades | 9 of Spades | 10 of Spades | Jack of Spades | Queen of Spades | King of Spades'
>>>d1.shuffle_deck
>>>d1.deck_as_text()
'3D | 2D | AD | JH | 2C | QC | 2S | 10D | 9D | QH | 6D | QS | AH | 8H | 3H | 4S | 4H | 5D | 4D | 7C | JD | 2H | JC | JS | 10H | 8C | 10S | 5H | 3S | AC | 7D | 4C | 6H | 9H | 3C | 9S | KS | 5S | 8S | 7S | 10C | KC | 6S | 5C | KH | KD | 9C | QD | 7H | 8D | 6C | AS'
>>>d1.order_deck()
>>>d1.deck_as_text()
'AD | 2D | 3D | 4D | 5D | 6D | 7D | 8D | 9D | 10D | JD | QD | KD | AC | 2C | 3C | 4C | 5C | 6C | 7C | 8C | 9C | 10C | JC | QC | KC | AH | 2H | 3H | 4H | 5H | 6H | 7H | 8H | 9H | 10H | JH | QH | KH | AS | 2S | 3S | 4S | 5S | 6S | 7S | 8S | 9S | 10S | JS | QS | KS'
>>>d1.shuffle_deck()
>>>d1.deck_as_text()
'7C | 2S | QD | 4S | 10H | 3C | 9D | JC | JD | JH | KD | KC | 6S | AH | 7D | JS | 7H | QC | 6C | QS | 4H | 10C | KH | AS | AD | 10D | 7S | 9C | 8S | 4C | 8D | 10S | QH | 8H | 5D | 6D | 9S | 3H | 3D | 5H | 2D | 5C | 4D | 2C | 3S | 6H | AC | 8C | 5S | KS | 9H | 2H'
>>>d1.shuffle_deck()
'7C | 2S | QD | 4S | 10H | 3C | 9D | JC | JD | JH | KD | KC | 6S | AH | 7D | JS | 7H | QC | 6C | QS | 4H | 10C | KH | AS | AD | 10D | 7S | 9C | 8S | 4C | 8D | 10S | QH | 8H | 5D | 6D | 9S | 3H | 3D | 5H | 2D | 5C | 4D | 2C | 3S | 6H | AC | 8C | 5S | KS | 9H | 2H'
>>>h1 = d1.get_suit("Club")
>>>h1.deck_as_text()
'7C | 3C | JC | KC | QC | 6C | 10C | 9C | 4C | 5C | 2C | AC | 8C'
>>>h1.quick_sort_deck()
'AC | 2C | 3C | 4C | 5C | 6C | 7C | 8C | 9C | 10C | JC | QC | KC'
```

### Testing
```
$cd [package install directory]
$python test_cards.py
.......
----------------------------------------------------------------------
Ran 7 tests in 0.002s

OK

```






