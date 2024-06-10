# Ascii-art

### Objectives

_Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big._

What we mean by a graphic representation using ASCII, is to write the string received using ASCII characters, as you can see in the example below:

- This project should handle an input with numbers, letters, spaces, special characters and \n.
- Take a look at the ASCII manual.

Explain :

```go
for index, line := range asciiLines {
		p := index / 9
		modulo := index % 9
		if modulo > 0 {
			if asciiFileLines[p] == nil {
				asciiFileLines[p] = make([]string, 8)
			}
			asciiFileLines[p][modulo-1] = line
		}
	}
```

To find the line :

- 9 / 9 = 1 modulo 9 % 9 = 0
- 10 / 9 = 1.1111 modulo 10 % 9 = 1
- 11 / 9 = 1.2222 modulo 11 % 9 = 2
- 12 / 9 = 1.3333 modulo 12 % 9 = 3
- ect...

```
/ n° ligne fichier : 0 / n° ascii :0 / n° ligne :0
      / n° ligne fichier : 1 / n° ascii :0 / n° ligne :1
      / n° ligne fichier : 2 / n° ascii :0 / n° ligne :2
      / n° ligne fichier : 3 / n° ascii :0 / n° ligne :3
      / n° ligne fichier : 4 / n° ascii :0 / n° ligne :4
      / n° ligne fichier : 5 / n° ascii :0 / n° ligne :5
      / n° ligne fichier : 6 / n° ascii :0 / n° ligne :6
      / n° ligne fichier : 7 / n° ascii :0 / n° ligne :7
      / n° ligne fichier : 8 / n° ascii :0 / n° ligne :8
/ n° ligne fichier : 9 / n° ascii :1 / n° ligne :0
 _  / n° ligne fichier : 10 / n° ascii :1 / n° ligne :1
| | / n° ligne fichier : 11 / n° ascii :1 / n° ligne :2
| | / n° ligne fichier : 12 / n° ascii :1 / n° ligne :3
| | / n° ligne fichier : 13 / n° ascii :1 / n° ligne :4
|_| / n° ligne fichier : 14 / n° ascii :1 / n° ligne :5
(_) / n° ligne fichier : 15 / n° ascii :1 / n° ligne :6
    / n° ligne fichier : 16 / n° ascii :1 / n° ligne :7
    / n° ligne fichier : 17 / n° ascii :1 / n° ligne :8
```
