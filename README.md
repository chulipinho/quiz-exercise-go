
# Quiz Game

A simple quiz game made in [Go](golang.org). Exercise provided by [Gophercises](https://gophercises.com/)

The user will be asked questions and must provide the answers. A timer (30s default) will be running in the background and as soon as it goes off the application will finish displaying the users result. In case the user manages to finish all questions in time, the application will also finish showing the results.

## How to use

Clone the repository:

```
git clone https://github.com/chulipinho/quiz-exercise-go.git
```

Open the folder in terminal:
```
cd quiz-exercise-go
```
Start the application:
```
go run cmd/quiz/main.go
```
Optionaly you can use the flags ```-t``` to provide the time limit and/or ```-p``` to provide a .csv file with problems. 

The .csv must be formatted "question,answer"(without quotes). You can look at [problems.csv](data/problems.csv) for an example.

## Solutions Applied

The application itself is pretty simple. The tricky part comes when implementing the timer. To make it work properly it war necessary to take advantage of Go Routines and Go Channels.

## Author:

Fellipe Luz Souza Machado

fellipe.luz.machado@gmail.com

[LinkedIn](https://www.linkedin.com/in/fellipe-luz/)

[GitHub](github.com/chulipinho)
