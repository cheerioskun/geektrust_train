# Pre-requisites
* Go 1.15.7/1.16.8/1.17
* go tool

# How to run the code

We have provided scripts to execute the code. 

Use `run.sh` if you are Linux/Unix/macOS Operating systems and `run.bat` if you are on Windows.  Both the files run the commands silently and prints only output from the input file `sample_input/input1.txt`. You are supposed to add the input commands in the file from the appropriate problem statement. 

Internally both the scripts run the following commands 

 * `go build .` - This will build an executable by the name geektrust in the directory $GOPATH/src/geektrust besides the main.go file .
 * Execute the file from the directory $GOPATH/src/geektrust using the command
`./geektrust sample_input/input1.txt`

We expect your program to take the location to the text file as parameter. Input needs to be read from a text file, and output should be printed to the console. The text file will contain only commands in the format prescribed by the respective problem.

This main file, main.go should receive in the command line argument and parse the file passed in. Once the file is parsed and the application processes the commands, it should only print the output.

 # Running the code for multiple test cases

 Please fill `input1.txt` and `input2.txt` with the input commands and use those files in `run.bat` or `run.sh`. Replace `./geektrust sample_input/input1.txt` with `./geektrust sample_input/input2.txt` to run the test case from the second file. 

 # How to execute the unit tests

 The unit tests are ran and the coverage is calculated using the library `gotestsum`. This is independent of your solution and there is no need to add any dependency. However this will work only if you use Go Modules for dependency management.

We execute the unit tests by running the following command from the directory $GOPATH/src/geektrust

`gotestsum --hide-summary=all ./...`
We check for the coverage of unit tests by executing the following command. from the directory $GOPATH/src/geektrust

`gotestsum --hide-summary=all -- -coverprofile=cover.out ./...`

# Help

You can refer our help documents [here](https://help.geektrust.com)
You can read build instructions [here](https://github.com/geektrust/coding-problem-artefacts/tree/master/Go)