# golang-trains
Trains, takes a network of nodes and connections and can determine shortest path from A-B. Distance for a specific trip.

TO FIX
- [ ] more refactoring's
- [ ] name space
- [ ] error tests  
- [ ] go lint / go vet fixes
- [ ] more idiomatic golang stuff  
- [ ]execute in a goroutine pipeline

TO DO
- [ ] cli application ```go-trains```  
- [ ] cli args ```go-trains --help```  
- [ ] cli args ```go-trains -f inputfile.txt --distance A-B-C```  
- [ ] cli args ```go-trains -i "AB1,BC3.." --distance A-B```  
- [ ] cli args ``` --shortest-trip```  
- [ ] cli args ```--question-file questionsfile.txt```
- [ ] JUNIT xml parsing of test output
- [ ] Docker file to create a build and run test container
- [ ] travis build (or other build server)
- [ ] script file with run commands ```run.sh```

Run tests
go test -v ./trains  
goconvey -host 0.0.0.0  
