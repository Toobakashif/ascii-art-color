#!/bin/bash



go run . "hello world" standard --COLOR=red
echo 'Does it display hello world in red"'
go run . "1 + 1 = 2" shadow --COLOR=green
echo 'Does it display "1 + 1 = 2" in green'
go run . "(%&) ??" standard --COLOR=yellow
echo 'Does it display (%&) ?? in yellow and thinkertoy'
go run . "Better than others?" standard --COLOR=blue --RANDOM
echo 'Does it display text randomly in blue and thinkertoy'
go run . "Better than others?" standard --COLOR=purple --POSITION=2:2
echo 'Does it display second letter in standard font and purple'
go run . "Dog\nCat?" shadow --COLOR=green --POSITION=6:7
 echo 'Does it display two letters in shadow font and green'
go run . "HeY GuYs" shadow --COLOR=cyan --WORD=GuYs
echo 'Does it display GuYs in shadow font and cyan'
go run . "RGB()" standard --COLOR=blue --LETTER=B
echo 'Does it display B in Standard font and blue'