# golang-trains  
[![Build Status](https://travis-ci.org/joesustaric/golang-trains.svg?branch=master)](https://travis-ci.org/joesustaric/golang-trains)  
Trains, takes a network of nodes and connections and can determine shortest path from A-B. Distance for a specific trip.

TO DO
- [ ] Finish the problem
- [ ] Refactor
- [ ] better ReadMe description
- [ ] Docker file to create a build env
- [ ] cli application ```go-trains```  
- [ ] cli args ```go-trains --help```  
- [ ] cli args ```go-trains --file-input inputfile.txt --questions questions-file.txt```  
- [ ] cli args ```go-trains --input "AB1,BC3.." --distance A-B```  
- [ ] cli args ```go-trains --input "AB1,BC3.." --shortest-trip A-B``` etc..  
- [ ] JUNIT xml parsing of test output
- [x] travis build (or other build server)
- [ ] script file with run commands ```run.sh```  

Refactoring Fixes
- [ ] more refactoring's
- [ ] namespace
- [ ] error tests  
- [ ] go lint / go vet fixes
- [ ] more idiomatic golang stuff  
- [ ] execute in a goroutine pipeline?

Extra TODO's  
- [ ] trains as a web service?

#### Run tests
go test -v ./trains  
goconvey -host 0.0.0.0  
