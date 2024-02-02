# Discord Bot


[Русская версия](README_RUS.md)

Bot written in golang using https://github.com/bwmarrin/discordgo


## Content
- [Stack](#stack)
- [User guide](#user-guide)
- [Project structure](#project-structure)
- [Functionality](#functionality)

## Stack
* discordgo - Library for creating discord applications on golang
* cleanenv - Reading the configuration

## User guide
Welcome to my first bot in Discord here is all available commands:
To use the bot please use " /bot " prefix, every command should start with this command
* "help" - returns all available commands with short description
* "hangman start" - start a game hangman where you should guess the hidden word
* "hangman <X>" - command to guess a letter from the word, replace <X> with your guess
* "btc now" - command to get current price of bitcoin in USD
* "btc limit <NUMBER>" - set a limit to BTC price, when limit is meet bot will notify you,replace <NUMBER> with target price

## Project structure
### `cmd/main.go`
Project entry point, applications starts here

### `config`
Reads configuration from .env file, to read the configuration is used [cleanenv](https://github.com/ilyakaznacheev/cleanenv)



### `internal/applicator`
Project set up, initializing session, creating handlers and closing session

### `internal/handler`
Main logic of the project

### `pkg`
Utility functions


## Functionality

### Bitcoin
Fetching and placing a limit to BTC price in US dollars:
 ### `btc now` 
 Will display current bitcoin price, bitcoin price is fetched every 10 seconds in goroutine asynchronously

### `btc limit <NUMBER>`
Sets a limit to BTC price, when limit is reached sends a notifying message to user, replace <NUMBER> with desired price

### Hangman game
Simple game of word guessing has 2 functions:
 ### `hangman start` 
 generate random word and draw a hangman picture hangman 

### `hangman <LETTER>`
User inputs his guess replacing the <LETTER> letter is compared with word and result is printed to user
