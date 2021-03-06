package gratified

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

//Quotes struct with json
type Quotes struct {
	Author string `json:"Author"`
	Quote  string `json:"Quote"`
}

//ShowGratitudeQuote prints out a quote
func ShowGratitudeQuote() {

	var quotes = []byte(`
	[
	{
        "Author": "A.A. Milne",
        "Quote": "Piglet noticed that even though he had a Very Small Heart, it could hold a rather large amount of Gratitude."
    },
    {
        "Author": "Albert Schweitzer",
        "Quote": "At times, our own light goes out and is rekindled by a spark from another person. Each of us has cause to think with deep gratitude of those who have lighted the flame within us."
    },
    {
        "Author": "Alfred North Whitehead",
        "Quote": "No one who achieves success does so without the help of others. The wise and confident acknowledge this help with gratitude."
    },
    {
        "Author": "Anne Morrow Lindbergh",
        "Quote": "One can never pay in gratitude; one can only pay 'in kind' somewhere else in life."
    },
    {
        "Author": "Buddha",
        "Quote": "Let us rise up and be thankful, for if we didn't learn a lot today, at least we learned a little, and if we didn't learn a little, at least we didn't get sick, and if we got sick, at least we didn't die; so, let us all be thankful."
    },
    {
        "Author": "Buddhist proverb",
        "Quote": "'Enough' is a feast."
    },
    {
        "Author": "Charles Dickens",
        "Quote": "Reflect upon your present blessings, of which every man has plenty; not on your past misfortunes, of which all men have some."
    },
    {
        "Author": "Charles Schwab",
        "Quote": "The way to develop the best that is in a person is by appreciation and encouragement."
    },
    {
        "Author": "Dietrich Bonhoeffer",
        "Quote": "In ordinary life, we hardly realize that we receive a great deal more than we give, and that it is only with gratitude that life becomes rich."
    },
    {
        "Author": "Eckhart Tolle",
        "Quote": "Acknowledging the good that you already have in your life is the foundation for all abundance."
    },
    {
        "Author": "Epictetus",
        "Quote": "He is a wise man who does not grieve for the things which he has not, but rejoices for those which he has."
    },
    {
        "Author": "Frank A. Clark",
        "Quote": "If a fellow isn't thankful for what he's got, he isn't likely to be thankful for what he's going to get."
    },
    {
        "Author": "Fred De Witt Van Amburgh",
        "Quote": "Gratitude is a currency that we can mint for ourselves, and spend without fear of bankruptcy."
    },
    {
        "Author": "G.K. Chesterton",
        "Quote": "I would maintain that thanks are the highest form of thought; and that gratitude is happiness doubled by wonder."
    },
    {
        "Author": "Gerald Good",
        "Quote": "If you want to turn your life around, try thankfulness. It will change your life mightily."
    },
    {
        "Author": "Gertrude Stein",
        "Quote": "Silent gratitude isn't very much to anyone."
    },
    {
        "Author": "Henri Frederic Amiel",
        "Quote": "Thankfulness is the beginning of gratitude. Gratitude is the completion of thankfulness. Thankfulness may consist merely of words. Gratitude is shown in acts."
    },
    {
        "Author": "John E. Southard",
        "Quote": "The only people with whom you should try to get even are those who have helped you."
    },
    {
        "Author": "John F. Kennedy",
        "Quote": "As we express our gratitude, we must never forget that the highest appreciation is not to utter words but to live by them."
    },
    {
        "Author": "John Wooden",
        "Quote": "Things turn out best for people who make the best of the way things turn out."
    },
    {
        "Author": "Maya Angelou",
        "Quote": "This a wonderful day. I've never seen this one before."
    },
    {
        "Author": "Melody Beattie",
        "Quote": "Gratitude turns what we have into enough, and more. It turns denial into acceptance, chaos into order, confusion into clarity...it makes sense of our past, brings peace for today, and creates a vision for tomorrow."
    },
    {
        "Author": "Michael Josephson",
        "Quote": "The world has enough beautiful mountains and meadows, spectacular skies and serene lakes. It has enough lush forests, flowered fields, and sandy beaches. It has plenty of stars and the promise of a new sunrise and sunset every day. What the world needs more of is people to appreciate and enjoy it."
    },
    {
        "Author": "Mike Ericksen",
        "Quote": "I truly believe we can either see the connections, celebrate them, and express gratitude for our blessings, or we can see life as a string of coincidences that have no meaning or connection. For me, I'm going to believe in miracles, celebrate life, rejoice in the views of eternity, and hope my choices will create a positive ripple effect in the lives of others. This is my choice."
    },
    {
        "Author": "Naomi Williams",
        "Quote": "It is impossible to feel grateful and depressed in the same moment."
    },
    {
        "Author": "Neal A. Maxwell",
        "Quote": "We should certainly count our blessings, but we should also make our blessings count."
    },
    {
        "Author": "Oprah Winfrey",
        "Quote": "Be thankful for what you have; you'll end up having more. If you concentrate on what you don't have, you will never, ever have enough."
    },
    {
        "Author": "Raheel Farooq",
        "Quote": "Gratitude is more of a compliment to yourself than someone else."
    },
    {
        "Author": "Ralph Waldo Emerson",
        "Quote": "You cannot do a kindness too soon because you never know how soon it will be too late."
    },
    {
        "Author": "Rasheed Ogunlaru",
        "Quote": "In life, one has a choice to take one of two paths: to wait for some special day--or to celebrate each special day."
    },
    {
        "Author": "Robert Braathe",
        "Quote": "Gratitude and attitude are not challenges; they are choices."
    },
    {
        "Author": "Robert Brault",
        "Quote": "Enjoy the little things, for one day you may look back and realize they were the big things."
    },
    {
        "Author": "Robert Quillen",
        "Quote": "If you count all your assets, you always show a profit."
    },
    {
        "Author": "Stephen Richards",
        "Quote": "Gratitude also opens your eyes to the limitless potential of the universe, while dissatisfaction closes your eyes to it."
    },
    {
        "Author": "Steve Maraboli",
        "Quote": "Forget yesterday--it has already forgotten you. Don't sweat tomorrow--you haven't even met. Instead, open your eyes and your heart to a truly precious gift--today."
    },
    {
        "Author": "Tom Hopkins",
        "Quote": "Keep your eyes open and try to catch people in your company doing something right, then praise them for it."
    },
    {
        "Author": "Tom Perrotta",
        "Quote": "They both seemed to understand that describing it was beyond their powers, the gratitude that spreads through your body when a burden gets lifted, and the sense of homecoming that follows, when you suddenly remember what it feels like to be yourself."
    },
    {
        "Author": "William Arthur Ward",
        "Quote": "Feeling gratitude and not expressing it is like wrapping a present and not giving it."
    },
    {
        "Author": "William James",
        "Quote": "The deepest craving of human nature is the need to be appreciated."
    },
    {
        "Author": "Willie Nelson",
        "Quote": "When I started counting my blessings, my whole life turned around."
	}
	]
	`)

	var listquotes []Quotes

	err := json.Unmarshal(quotes, &listquotes)
	if err != nil {
		fmt.Println("error:", err)
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(listquotes)
	message := listquotes[n]
	fmt.Println(message.Quote, "-", message.Author)
}
