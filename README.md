Paxos
==================
[Paxos](https://en.wikipedia.org/wiki/Paxos_(computer_science)) Paxos algorithm implemented in Golang.

### Requirements:
1. Go version 1.14
2. GNU Makefile

### Project Structure 
    .
    ├── src                    # source code files
    │   ├── bin                # host execuatable file
    │   ├── main.go            # Entry point and init the server
    │   ├── util.go            # helper code
    │   ├── Makefile           # makefile to automate build and testing
    │   ├── proposer.go        # implementations of proposer  
    │   ├── acceptor.go        # implementations of acceptor  
    │   ├── learner.go         # implementations of learner
    │   ├── mid.go             # implementations of middleware to help mitigate the malicious node
    │   └── malicious.go       # implementations of a malicious node  
    ├── go.mod                  # Go Module and dependancies
    ├── go.sum                  # Dependancies versions 
    └── README.md               # README file 
#### to Run General Algorithm: 
- run make 
which will run there servers listening on port 3000, 5000, 8000 respectivly 

#### to Run Malicious version:
- run make testlive for a demo of a system livelock by a malicious node (node with port 5000)
- run make testc for a demo of a system compromised by a malicious node (node with port 5000)


#### Notes:
Running make will generaete executable in bin directory
to remove the executable:
- run make clean 