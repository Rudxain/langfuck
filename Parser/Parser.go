package parser

import (
	. "LanguageFuck/Encrypter"
	. "LanguageFuck/Lexer"
	. "LanguageFuck/Types"
)

type Parser struct {
	Tokens *[]*Token
	Swap   map[string]string
	Enc    *Encrypter
}

func ParserInit(tokens *[]*Token, key int) *Parser {
	return &Parser{
		tokens,
		make(map[string]string),
		EncrypterInit(key),
	}
}

func (pr *Parser) Parse(l *Lexer, decrypt bool) {
	var enc string
	var token *Token
	var tokens []*Token = (*pr.Tokens)
	var tokens_len int = len(*pr.Tokens)
	for i := 0; i < tokens_len; i++ {
		token = tokens[i]
		token_content := l.GetTokenContent(token)

		if token.Kind == TOKEN_IMPORTED {
			if decrypt {
				enc = pr.Enc.Decrypt(token_content)
			} else {
				enc = pr.Enc.Encrypt(token_content)
			}
			for i+1 < tokens_len && l.GetTokenContent(tokens[i+1]) == "." {
				i += 2
			}
			pr.Swap[token_content] = enc
		}

		if token.Kind == TOKEN_SYMBOL {
			if decrypt {
				enc = pr.Enc.Decrypt(token_content)
			} else {
				enc = pr.Enc.Encrypt(token_content)
			}
			pr.Swap[token_content] = enc
		}
	}
}