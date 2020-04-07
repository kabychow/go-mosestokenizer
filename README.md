# Moses Tokenizer for GoLang

GoLang implementation of [Tokenizer & Normalizer from Moses Decoder](https://github.com/moses-smt/mosesdecoder)

## Installation

```bash
go get github.com/khaibin/go-mosestokenizer
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/khaibin/go-mosestokenizer"
    "github.com/khaibin/go-mosestokenizer/nonbreaking_prefix"
)

func main() {
    text := "This is a string"
    lang := "en"

    // Tokenize and get the result as []string
    mosestokenizer.Tokenize(text, lang)

    // Tokenize and get the result as string
    mosestokenizer.TokenizeAsString(text, lang)

    // Normalization
    mosestokenizer.Normalize(text, lang)
    
    prefix := "mr"
    prefix_lang := "en"

    // Returns true if string is non-breaking prefix
    nonbreaking_prefix.Find(prefix, prefix_lang)

    // Returns true if string is non-breaking numeric only prefix
    nonbreaking_prefix.FindNumeric(prefix, prefix_lang)

    // Constants
    //   perluniprops.ALPHA
    //   perluniprops.NUM
    //   perluniprops.ALNUM
}
```

## Publications
The segmentation methods are described in:

[Rico Sennrich, Barry Haddow and Alexandra Birch (2016): Neural Machine Translation of Rare Words with Subword Units Proceedings of the 54th Annual Meeting of the Association for Computational Linguistics (ACL 2016). Berlin, Germany.](https://arxiv.org/abs/1508.07909)


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)