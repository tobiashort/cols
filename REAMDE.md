# cols

## Usage

```
Usage: cols [-n COLS] [FILE]

  ATTENTION
	Be aware that this program loads the whole content
	of the file into memory.

  FILE
	Optional file otherwise reads from stdin.

  -n int
    	Number of columns (default 4)
```

## Example

```shell
% cat names.txt | head -n 100 | cols -n 5
Aaren      Addia       Adore       Agathe     Aili   
Aarika     Addie       Adoree      Aggi       Ailina 
Abagael    Addy        Adorne      Aggie      Ailis  
Abagail    Adel        Adrea       Aggy       Ailsun 
Abbe       Adela       Adria       Agna       Ailyn  
Abbey      Adelaida    Adriaens    Agnella    Aime   
Abbi       Adelaide    Adrian      Agnes      Aimee  
Abbie      Adele       Adriana     Agnese     Aimil  
Abby       Adelheid    Adriane     Agnesse    Aindrea
Abbye      Adelice     Adrianna    Agneta     Ainslee
Abigael    Adelina     Adrianne    Agnola     Ainsley
Abigail    Adelind     Adriena     Agretha    Ainslie
Abigale    Adeline     Adrienne    Aida       Ajay   
Abra       Adella      Aeriel      Aidan      Alaine 
Ada        Adelle      Aeriela     Aigneis    Alameda
Adah       Adena       Aeriell     Aila       Alana  
Adaline    Adey        Afton       Aile              
Adan       Adi         Ag          Ailee             
Adara      Adiana      Agace       Aileen            
Adda       Adina       Agata       Ailene            
Addi       Adora       Agatha      Ailey
```
