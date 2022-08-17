# Summer of Code Challenge: Word Guessing Game
I'm bored, let's play a word guessing game! We'll have to make it first...

## The Game

This game should be single-player. At the start of the game, the player should be shown a blank space for each letter in a randomly chosen word. The player then gets to guess one letter at a time. If they guess a letter that is in the word the corresponding empty spaces in the word should be filled with the letter. If they guess a letter that isn't in the word their guess counter should go up by 1. The player gets 10 guesses to get the full word or they lose. (This game is hangman if you're familiar with it.)
Rules

* Computer chooses a random word
* Player gets to make one guess each round
* For each correctly guessed letter, reveal each place the letter is at in the secret word
* If the player makes 10 incorrect guesses they lose
* If the player guesses every letter that is in the word they win

## Stretch Goal

Have a leaderboard of the highest scorers. The player gets 10 points for each correct guess and loses 5 points for each incorrect guess.