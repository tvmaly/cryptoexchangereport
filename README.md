# cryptoexchangereport
crypto exchange metrics important to crypto traders

## Interfaces used in the system

### Quotes Interfaces

	type ExchangeQuote interface {
		GenericQuote() *GenericQuote
	}

	type Quote interface {
		MidPoint() float64
		SpreadPercentage() float64
		Weighted() bool
	}

## Layout of Code
* docs - documentation
* main - code to generate the binaries
* vendor - 3rd party code
* lib - library code shared across systems
** constants - constants used by all code
** client - client based code specific only to clients
** server - server based code specific only to servers
** parsing
*** book - models to handle depth of book data
*** quotes - models to handle ticker data
*** trades - models to handle trade data
*** utility - utility code common to parsing


