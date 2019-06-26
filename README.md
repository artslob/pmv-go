# pmv-go

antlr jar version used to generate parser: `4.7.2`

Command setup:
```bash
$ md5sum antlr-4.7.2-complete.jar                                              
58c9cdda732eabd9ea3e197fa7d8f2d6  antlr-4.7.2-complete.jar
$ alias antlr='java -jar "${HOME}/bin/antlr-4.7.2-complete.jar"'
```
Command to generate parser:
```bash
# in project root
cd lang
antlr -Dlanguage=Go -o ../parser Lang.g4
```
