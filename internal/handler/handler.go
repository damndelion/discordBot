package handler

import (
	"discord-go-bot/config"
	"discord-go-bot/internal/handler/dto"
	"discord-go-bot/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// handler holds session details
type handler struct {
	session *discordgo.Session
	cfg     *config.Config
}

// NewHandler creates new session handler
func NewHandler(session *discordgo.Session, cfg *config.Config) {
	r := &handler{session, cfg}
	session.AddHandler(r.messageHandler)
	fetchBTCPrice()
}

// messageHandler manages bot commands
func (h *handler) messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")
	if args[0] != h.cfg.Prefix {
		return
	}

	if args[1] == "hangman" {
		h.handleHangmanGame(s, m, args)
	} else if args[1] == "btc" {
		h.handleBtcPrice(s, m, args)
	} else if args[1] == "help" {
		h.handleHelp(s, m, args)
	}
}

// handleBtcPrice function handles btc commands
// now command returns current price of BTC
// limit command sets stop limit for BTC price if limit is triggered it will notify the user
func (h *handler) handleBtcPrice(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	if args[2] == "now" {
		embed := discordgo.MessageEmbed{
			Title: "Current Bitcoin(BTC) price: " + fmt.Sprintf("%.2f", btcPrice) + "$",
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		if err != nil {
			log.Fatal(err)
		}
	} else if args[2] == "limit" {
		if notificationChan == nil {
			notificationChan = make(chan string)
			go func(session *discordgo.Session, channelID string) {
				msg := <-notificationChan
				embed := discordgo.MessageEmbed{
					Title: msg,
				}
				if _, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed); err != nil {
					log.Println("Error sending notification:", err)
				}
			}(s, m.ChannelID)
		}

		embed := discordgo.MessageEmbed{
			Title: "Your limit was set to " + args[3] + "$",
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		if err != nil {
			log.Fatal(err)
		}
	}
}

var btcPrice float64
var btcLimit float64
var notificationChan chan string

// fetchBTCPrice get price of BTC in USD every ten seconds
// check if user limit was hit and send message to channel
func fetchBTCPrice() {
	go func() {
		for {
			if btcLimit > 0 && btcPrice <= btcLimit {
				notificationChan <- "BTC price has reached or fallen below your limit of $" + fmt.Sprintf("%.2f", btcLimit)
				btcLimit = 0
			}
			url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

			response, err := http.Get(url)
			if err != nil {
				log.Println("Error fetching BTC price:", err)

				continue
			}

			var data dto.CoinGeckoResponse
			if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
				log.Println("Error decoding BTC price data:", err)

				continue
			}

			btcPrice = data.Bitcoin.USD
			err = response.Body.Close()
			if err != nil {
				return
			}

			time.Sleep(10 * time.Second)
		}
	}()
}

var word string
var hiddenWord string
var guesses int
var maxGuesses int

// handleHangmanGame function handles hangman game commands
// start command starts a new game
// single letter is counted as a guess
func (h *handler) handleHangmanGame(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	// Initialize game state
	if args[2] == "start" {
		word = utils.PickRandomWord()
		hiddenWord = strings.Repeat("-", len(word))
		guesses = 0
		maxGuesses = 6

		// Send initial game message
		embed := discordgo.MessageEmbed{
			Title: "Hangman Game! Guess a letter to reveal the word. You have " +
				"" + strconv.Itoa(maxGuesses) + " guesses.\n Your Word is: \n\n" + hiddenWord +
				"\n" + utils.HangmanAscii[0],
		}
		_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		if guesses == maxGuesses {
			embed := discordgo.MessageEmbed{
				Title: "You ran out of guesses! The word was: " + word +
					"\n Please enter (/bot hangman start) to start a new game",
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			if err != nil {
				log.Fatal(err)
				return
			}
			return
		}
		guess := strings.ToLower(args[2])
		if len(guess) != 1 || !utils.IsLetter(guess) {
			embed := discordgo.MessageEmbed{
				Title: "Invalid guess. Please enter a single letter.",
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		if strings.Contains(word, guess) {
			for i, char := range word {
				if string(char) == guess {
					hiddenWord = hiddenWord[:i] + guess + hiddenWord[i+1:]
				}
			}
			embed := discordgo.MessageEmbed{
				Title: "Correct! The word is now: \n" + hiddenWord + "\n" + utils.HangmanAscii[guesses],
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			guesses++
			embed := discordgo.MessageEmbed{
				Title: "Incorrect. You have " + strconv.Itoa(maxGuesses-guesses) + " guesses left.\n" +
					hiddenWord + "\n" + utils.HangmanAscii[guesses],
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			if err != nil {
				log.Fatal(err)
			}
		}
		// End game message
		if guesses == maxGuesses {
			embed := discordgo.MessageEmbed{
				Title: "You ran out of guesses! The word was: " + word,
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			if err != nil {
				log.Fatal(err)
			}
		} else if hiddenWord == word {
			embed := discordgo.MessageEmbed{
				Title: "Congratulations! You guessed the word: " + word,
			}
			_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

// handleHelp prints help message to user
func (h *handler) handleHelp(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	var helpMessage string
	if args[2] == "hangman" {
		helpMessage = utils.HelpMessageHangman
	} else if args[2] == "btc" {
		helpMessage = utils.HelpMessageBTC
	} else {
		helpMessage = utils.HelpMessage
	}
	embed := discordgo.MessageEmbed{
		Description: helpMessage,
	}
	_, err := s.ChannelMessageSendEmbed(m.ChannelID, &embed)
	if err != nil {
		log.Fatal(err)
	}
}
