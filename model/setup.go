package model

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_CONNECTION_STRING")
	if len(dsn) == 0 {

		panic("Connection url not provided please provide it in the DB_CONNECTION_STRING env")
	}
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {

		panic("Failed to connect to database!")

	}

	var books []Book
	var count int64
	database.Find(&books).Count(&count)
	if count == 0 {

		err = seedData(err, database)
	}

	if err != nil {

		fmt.Errorf("Failed to seed the data")
	}

	DB = database
}

func seedData(err error, database *gorm.DB) error {
	err = database.AutoMigrate(&Book{})
	book := &Book{1, "Beyond Good And Evil", "Friedrich Nietzsche", "Beyond Good And Evil asks, is truth absolute? Do humans invent ways to fortify already held views or truly seek the truth? Are the powerful more ‘right’ than the weak? Or is Nietzsche writing down page after page to hear himself talk?", 1}
	database.Create(&book)

	book = &Book{6, "Culture Map: Decoding How People Think, Lead, and Get Things Done Across Cultures", "Erin Meyer", "Whether you work in a home office or abroad, business success in our ever more globalized and virtual world requires the skills to navigate through cultural differences and decode cultures foreign to your own. Renowned expert Erin Meyer is your guide through this subtle, sometimes treacherous terrain where people from starkly different backgrounds are expected to work harmoniously together.", 6}
	database.Create(&book)

	book = &Book{2, "The Path\nA New Way to Think About Everything", " Professor Michael Puett, Christine Gross-Loh", "In the first book of its kind, The Path draws on the work of six of the great - but largely unknown - Chinese philosophers to offer a fresh and revolutionary guide to human flourishing.\n\nBy examining the teachings of Chinese thinkers and explaining what they reveal about our daily lives - from greeting others to raising children - The Path challenges some of our deepest held assumptions. It shows that the way to live well is not to slavishly follow a grand plan, as so much of Western thought would have us believe, but rather to follow a path - one of self-cultivation and self-discovery.", 2}
	database.Create(&book)

	book = &Book{3, "I Will Never See the World Again", "Ahmet Altan", "The destiny I put down in my novel has become mine. I am now under arrest like the hero I created years ago. I await the decision that will determine my future, just as he awaited his. I am unaware of my destiny, which has perhaps already been decided, just as he was unaware of his. I suffer the pathetic torment of profound helplessness, just as he did. Like a cursed oracle, I foresaw my future years ago not knowing that it was my own. Confined in a cell four metres long, imprisoned on absurd, Kafkaesque charges, novelist Ahmet Altan is one of many writers persecuted by Recep Tayyip Erdogan's oppressive regime. In this extraordinary memoir, written from his prison cell, Altan reflects upon his sentence, on a life whittled down to a courtyard covered by bars, and on the hope and solace a writer's mind can provide, even in the darkest places.", 3}
	database.Create(&book)

	book = &Book{4, "101 Essays That Will Change the Way You Think", "Brianna Wiest", "Over the past few years, Brianna Wiest has gained renown for her deeply moving, philosophical writing. This new compilation of her published work features pieces on why you should pursue purpose over passion, embrace negative thinking, see the wisdom in daily routine, and become aware of the cognitive biases that are creating the way you see your life.\n\nSome of these pieces have never been seen; others have been read by millions of people around the world. Regardless, each will leave you thinking: This idea changed my life.", 4}
	database.Create(&book)

	book = &Book{5, "The Nordic Theory of Everything: In Search of a Better Life", "Anu Partanen", "Partanen is a careful, judicious writer and she makes a careful, judicious case.. it s useful to know what the outsider knows: there are other ways of organizing humanity.", 5}
	database.Create(&book)
	return err
}
