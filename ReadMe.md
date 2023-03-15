
# test_with_golang
A test application, that will read a .csv file, and do some operations with it.

In order to execute it do the following lines:

    go build

and then

    go run .
This will run a small web server in localhost:8080
In order to test the operations run the following lines into the command window:

    curl -F file=@"D:/Arquivos/Projects/eCore/matrix.csv" localhost:8080/echo

*This one will return the file as a matrix format*

    curl -F file=@"D:/Arquivos/Projects/eCore/matrix.csv" localhost:8080/flatten

*This one will return the values of the file as a single line*

    curl -F file=@"D:/Arquivos/Projects/eCore/matrix.csv" localhost:8080/sum

*This one will sum all the values in the .csv file, and then return it*

    curl -F file=@"D:/Arquivos/Projects/eCore/matrix.csv" localhost:8080/multiply

*This one will multiply all the values in the .csv file, and then return it*

    curl -F file=@"D:/Arquivos/Projects/eCore/matrix.csv" localhost:8080/invert

*This one will return the file as a matrix inverted format*

There's tests for the operations called by the curl as well
To run the tests, do the following

    go test ./tests
