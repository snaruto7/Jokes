package main

type Joke struct {
	ID    int    `json:"id" binding:"required"`
	Likes int    `json:"Likes"`
	Joke  string `json:"joke" binding:"required"`
}

var jokeList = []Joke{
	{1, 0, "Did you hear about the restaurant on the moon? Great food, no atmosphere."},
	{2, 0, "What do you call a fake noodle? An Impasta."},
	{3, 0, "How many apples grow on a tree? All of them."},
	{4, 0, "Want to hear a joke about paper? Nevermind it's tearable."},
	{5, 0, "I just watched a program about beavers. It was the best dam program I've ever seen."},
	{6, 0, "Why did the coffee file a police report? It got mugged."},
	{7, 0, "How does a penguin build it's house? Igloos it together."},
	{8, 0, "Dad, did you get a haircut? No I got them all cut."},
	{9, 0, "What do you call a Mexican who has lost his car? Carlos."},
	{10, 0, "Dad, can you put my shoes on? No, I don't think they'll fit me."},
	{11, 0, "Why did the scarecrow win an award? Because he was outstanding in his field."},
	{12, 0, "Why don't skeletons ever go trick or treating? Because they have no body to go with."},
}

func ExportedJokes() []Joke {
	return jokeList
}
