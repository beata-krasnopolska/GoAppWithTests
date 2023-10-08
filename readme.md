The project made in  accordance with the official tutorial of Go in "Learn Go with tests"


Go version: 1.21.1


in order to build project on the terminal enter: go build
the httpserver.exe file appears. In order to run it in the terminal enter: ./httpserver

There is possibility ot test the funcionality with use of curl:
1. in terminal: ./httpserver
2. in additional terminal: curl -X POST http://localhost:5000/players/Pepper (few times, we can tets it wit different name)
3. curl http://localhost:5000/players/Pepper (the number of previously added plyers with name Pepper should be displayed)