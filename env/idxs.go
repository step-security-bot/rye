package env

import (
	"fmt"
	"strconv"
)

type Idxs struct {
	words1 [1000]string
	words2 map[string]int
	wordsn int
}

func (e *Idxs) IndexWord(w string) int {
	idx, ok := e.words2[w]
	if ok {
		return idx
	} else {
		e.words1[e.wordsn] = w
		e.words2[w] = e.wordsn
		e.wordsn += 1
		return e.wordsn - 1
	}
}

func (e *Idxs) GetIndex(w string) (int, bool) {
	idx, ok := e.words2[w]
	if ok {
		return idx, true
	}
	return 0, false
}

func (e Idxs) GetWord(i int) string {
	return e.words1[i]
}

func (e Idxs) Probe() {
	fmt.Print("<IDXS: ")
	for i := 0; i < e.wordsn; i++ {
		fmt.Print(strconv.FormatInt(int64(i), 10) + ": " + e.words1[i] + " ")
	}
	fmt.Println(">")
}

func (e Idxs) GetWordCount() int {
	return e.wordsn
}

func NewIdxs() *Idxs {
	var e Idxs
	e.words2 = make(map[string]int)
	e.wordsn = 1

	/*
		BlockType    Type = 1
		IntegerType  Type = 2
		WordType     Type = 3
		SetwordType  Type = 4
		OpwordType   Type = 5
		PipewordType Type = 6
		BuiltinType  Type = 7
		FunctionType Type = 8
		ErrorType    Type = 9
		CommaType    Type = 10
		VoidType     Type = 11
		StringType   Type = 12
		TagwordType  Type = 13
	*/

	// register words for builtin kinds, which the value objects should return on GetKind()

	e.IndexWord("block")
	e.IndexWord("integer")
	e.IndexWord("word")
	e.IndexWord("setword")
	e.IndexWord("opword")
	e.IndexWord("pipeword")
	e.IndexWord("builtin")
	e.IndexWord("function")
	e.IndexWord("error")
	e.IndexWord("comma")
	e.IndexWord("void")
	e.IndexWord("string")
	e.IndexWord("tagword")

	return &e
}
