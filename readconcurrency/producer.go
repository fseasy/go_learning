package concurrencyread

import (
	"os"
	"bufio"
)

// NumberedLine ..
//   sentence with number
type NumberedLine struct {
	lineNumber int
	sentence   string
}

// LineNumber ..
//  	get line number
func (nuLine *NumberedLine) LineNumber() int {
	return nuLine.lineNumber
}

// Sentence ..
// 	get sentence
func (nuLine *NumberedLine) Sentence() string {
	return nuLine.sentence
}

// SetLineNumber ..
//	set line number
func (nuLine *NumberedLine) SetLineNumber(nu int) {
	nuLine.lineNumber = nu
}

// SetSentence ..
//	set sentence
func (nuLine *NumberedLine) SetSentence(sentence string) {
	nuLine.sentence = sentence
}

// LineProducer ..
// 	produce line with number.
type LineProducer struct {
	lineQueue chan NumberedLine
	f *os.File
	isDone bool
}

// GetNumberedLine ..
//	get numbered line
func (producer *LineProducer) GetNumberedLine() (lineNumber int, sentence string) {
	nuLine := <- producer.lineQueue
	lineNumber = nuLine.LineNumber()
	sentence = nuLine.Sentence()
	return
} 

// LineQueue ..
// 	get line producer
func (producer *LineProducer) LineQueue() chan NumberedLine {
	return producer.lineQueue
}

func (producer *LineProducer) readFileAndFilQueue() {
	
	scanner := bufio.NewScanner(producer.f)
	scanner.Split(bufio.ScanLines)
	lineNumber := 0
	for scanner.Scan() {
		text := scanner.Text()
		nuLine := NumberedLine{lineNumber: lineNumber, sentence: text}
		producer.lineQueue <- nuLine
		lineNumber++
	}
	producer.isDone = true
}

// ProduceBlocked ..
//	produce synchronously
func (producer *LineProducer) ProduceBlocked() {
	producer.readFileAndFilQueue()
}

// Produce ..
//	produce asynchronous
func (producer *LineProducer) Produce(){
	go producer.readFileAndFilQueue()
}
