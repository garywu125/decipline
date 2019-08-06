module Demo : application source code in <remote repository>


.go code file import path
- import <repository> feature packages using <repository path> 
- import remote third party packages  usning <repository path>
- multiple command package locate in commands/ 


起手式:
  - go mod init <repository>: go to project directory, add command  go mod init <repository>
  - cd <project-directory>/commands/; go build