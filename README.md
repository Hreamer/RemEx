# What is RemEx? 
RemEx is a Remote Code Compilation and Execution Tool  
This tool will compile code you send it and than return the output  
To ensure safety of the machine that is running the server side of the tool, as remote code execution is one of the holy grails of security flaws, the Server.go file runs inside a docker container  
The Client.go file just sends the specified file over the wire. 

## Applications 
This project has specific applications in the embedded sphere and I can see a world where this is something similar to what runs on LeetCode when you click submit. 

## What Inspired Me?  
Just wanted to make a personal project that involved Golang as I wanted to learn the language and I had the idea for this project ever since I heard about the Log4J Exploit. 
