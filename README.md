# IMT2681-3-parser

Repo for creating the user input parser for the bot

## Storing the values

We created a JSON-file containing the supported currencies. The format is as following

```JSON

	"nok": "norwegian kronie",
	"sek": "swedish krona"

```

## The parsing algorithm

The parsing of the text happens in three steps. Since the specification all had input statements with the word *and*, we split the text into two seperate strings. The base string will be read out of the first string, and the target from the second. There is two methods for reading the base and target currency. First will look for the abbreviation, the second will look for the full name of the currency.

### Method two - look for the full text

A loop will run the string.contains() on the two pieces on text, split on *and*. When/if there is a match, the algorithm will not look for Abreviations. This is the method that is run first. 

### Method - look for abreviations

Here, we split the string on space using string.split(" ") on each of the two pieces of string. Then we look for instances of abbreviation in the two pieces of string. To do this we use a string search algorithm using hash maps. This is very effective. 

# Weaknesses

One of the most obvious weaknesses, is when a string asking for currency data contains more than one instance of *and*. 
