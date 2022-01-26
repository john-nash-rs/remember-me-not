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

<img width="1788" alt="Screenshot 2022-01-27 at 12 16 08 AM" src="https://user-images.githubusercontent.com/33624864/151227035-354b0358-9b3e-4e62-8a8a-7f72d92508ec.png">

## Run your project

Press run button.

## Stop your project

Press stop button.
