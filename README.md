# remember-me-not

## How to use?

Add your project information in run_information.json in config folder.
{
    "Projects":[
       {
          "id": "1",
          "Name":"Test Script",
          "Directory":"./scripts/",
          "RunInstruction":"./test.sh",
          "StopInstruction": "./test.sh"
       }
    ]
 }
 
 ### Configuration Definition
 
 #### id 
 
 Should be unique
 
 #### Name
 
 Name of project. Can be anything.
 
 #### Directory
 
 Directory where project's run command should be run. The code changes directory to that directory and runs the run instruction.
 
 #### RunInstruction
 
 Command which runs the project
 
  #### StopInstruction
 
 Command which stops the project

## How to run remember-me-not?

go run main.go 

You should see following UI at http://localhost:1010

<img width="1774" alt="Screenshot 2022-01-27 at 12 23 01 AM" src="https://user-images.githubusercontent.com/33624864/151227860-d19a4a4a-a708-432e-91d6-18bc6b8f380e.png">

## Run your project

Press run button.

## Stop your project

Press stop button.
