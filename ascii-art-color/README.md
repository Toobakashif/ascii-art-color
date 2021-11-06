Dear Auditor!

We made for you a tester file. You need to give this file permission to it with the command: chmod 777 tester.sh
We hope you find reading this same enjoyable like it was for us.


RGB : Red, Green and Blue 

Color choice:
Color		    hex code   	    RGB code
black		    #000000	 	    rgb(0,0,0)
red		        #FF0000		    rgb(255,0,0)
green		    #00FF00 	    rgb(0,255,0)
yellow		    #FFFF00		    rgb(255,255,0)
blue		    #0000FF 	    rgb(0,0,255)
purple  	    #800080		    rgb(128,0,128)
cyan		    #00FFFF 	    rgb(0,255,255)
white  		    #FFFFFF		    rgb(255,255,255)
lightgray 	    #D3D3D3 	    rgb(211,211,211)
darkgray 	    #A9A9A9		    rgb(169,169,169)
lightred	    #FF7F7F		    rgb(255,127,127)
lightgreen	    #90EE90		    rgb(144,238,144)
lightyellow	    #FFFFE0		    rgb(255,255,224)
lightblue	    #ADD8E6		    rgb(173,216,230)
lightmagenta	#E78BE7		    rgb(231,139,231)
lightcyan 	    #E0FFFF		    rgb(224,255,255)

Usage 1: go run . [STRING] [BANNER] [OPTION]  
EX1. go run . "fish" standard --COLOR=#FF7F7F  //"fish" is string  ,Banner is standard and option is color 

Usage 2: go run . [STRING] [BANNER] [OPTION][OPTION2]

OPTION2:  --POSITION=<begin:end> or --WORD=<word> or --LETTER=<letter> or --RANDOM

EX1. go run . "fish" standard --COLOR=#FF7F7F --LETTER=s
The letter s in standard font and #FF7F7F
EX2. go run . "fish" standard --COLOR=#FF7F7F --RANDOM
A random character in standard font and #FF7F7F
EX3. go run . "Better than others?" standard --COLOR=purple --POSITION=2:2
Here you just have to  write the characters position that you want to color 
EX4. go run . "Better than others?" standard --COLOR=purple --POSITION=2:4
It will color characters starting from the second to the fourth character.
EX5. go run . "Better than others?" standard --COLOR=red  --POSITION=2
Here you just have to  write the characters position that you want to color 
RGB Example:
go run . "fish" standard --COLOR='rgb(224,255,255)' --LETTER=s

In rgb code case you need to use upper commas ' 

TEST:
Try passing as arguments "hello world" standard --color=red.
Does it display the expected result?
Try passing as arguments "1 + 1 = 2" standard --color=green.
Does it display the expected result?

Enjoy life,

Team Markus:
Tooba
Janek Luts