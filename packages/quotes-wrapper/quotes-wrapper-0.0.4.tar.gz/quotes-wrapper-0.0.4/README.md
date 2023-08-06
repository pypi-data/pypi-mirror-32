# Quotes Wrapper

Access quotes from specified URL.

The API should return a valid JSON.

Eg.:

```js
{
    "quotes": [
        "Beautiful is better than ugly.",
        "Explicit is better than implicit.",
        "Simple is better than complex.",
        "Complex is better than complicated.",
        "Flat is better than nested.",
        "Sparse is better than dense.",
        "Readability counts.",
        "Special cases aren't special enough to break the rules.",
        "Although practicality beats purity.",
        "Errors should never pass silently.",
        "Unless explicitly silenced.",
        "In the face of ambiguity, refuse the temptation to guess.",
        "There should be one-- and preferably only one --obvious way to do it.",
        "Although that way may not be obvious at first unless you're Dutch.",
        "Now is better than never.",
        "Although never is often better than *right* now.",
        "If the implementation is hard to explain, it's a bad idea.",
        "If the implementation is easy to explain, it may be a good idea.",
        "Namespaces are one honking great idea -- let's do more of those!"
    ]
}
```

## Functions

`get_quotes`: Returns all quotes from API.

Eg.: 


```js
{
    "quotes": [
        "Beautiful is better than ugly.",
        "Explicit is better than implicit.",
        "Simple is better than complex.",
        "Complex is better than complicated.",
        "Flat is better than nested.",
        "Sparse is better than dense.",
        "Readability counts.",
        "Special cases aren't special enough to break the rules.",
        "Although practicality beats purity.",
        "Errors should never pass silently.",
        "Unless explicitly silenced.",
        "In the face of ambiguity, refuse the temptation to guess.",
        "There should be one-- and preferably only one --obvious way to do it.",
        "Although that way may not be obvious at first unless you're Dutch.",
        "Now is better than never.",
        "Although never is often better than *right* now.",
        "If the implementation is hard to explain, it's a bad idea.",
        "If the implementation is easy to explain, it may be a good idea.",
        "Namespaces are one honking great idea -- let's do more of those!"
    ]
}
```

`get_quote(index)`: Returns a single quote from the API.

Eg.: 

```js
{
    "quote": [
        "There should be one-- and preferably only one --obvious way to do it.",
    ]
}
```

## Configuration

Before using the lib you should configure the API endpoint.

```python
from quotes_wrapper import configure

configure('QUOTES_URL', 'type url here')
```

## Errors

1. **FetchingQuoteError**: When the API response status is not _200 OK_. Raised by `get_quotes` and `get_quote`
1. **QuoteNotFoundError**: When you try to get a quote that doesn't exist. Raised by `get_quote`.
1. **UnknownConfigError**: When trying to configure something that doesn't exist. Raises by `configure`.
