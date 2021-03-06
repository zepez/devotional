package utils

import (
	"strings"
	"testing"
)

func TestScrape(t *testing.T) {
	link := "https://www.lhm.org/dailydevotions/default.asp?date=20211109"
	html, plain, name, image := Scrape(link)

	expectName := "\"No Crisis Too Great\""
	expectHtml := "\u003cp\u003e\u003ci\u003eDaniel 12:1-3 - At that time shall arise Michael, the great prince who has charge of your people. And there shall be a time of trouble, such as never has been since there was a nation till that time. But at that time your people shall be delivered, everyone whose name shall be found written in the book. And many of those who sleep in the dust of the earth shall awake, some to everlasting life, and some to shame and everlasting contempt. And those who are wise shall shine like the brightness of the sky above; and those who turn many to righteousness, like the stars forever and ever.\u003c/i\u003e\u003c/p\u003e\u003cp\u003eThe book of Daniel is a fascinating book. In it we read of Daniel. Exiled from Judah to Babylonia, he came from primo Israelite stock: noble, well-educated, handsome, full of knowledge and insight. Yes, he was a man competent on all fronts (see Daniel 1:3-7). As such, Daniel was selected to receive a first-rate Babylonian education so that at the end of three years he and his three friends, Shadrach, Meshach, and Abednego, could serve the king of Babylon in his court.\u003cbr/\u003e\u003cbr/\u003eDaniel was also one to whom God gave profound messages like today\u0026#39;s text about the people Israel. In it is a biblical reference to the resurrection, a final judgment, and an afterlife. It also speaks about a period of great anguish and deliverance for all those whose names are \u0026#34;found written in the book.\u0026#34; How in keeping with God\u0026#39;s character this is! Even in the middle of all the upheaval taking place in these long-ago centuries, God does not forsake His people; He will not forsake His people.\u003cbr/\u003e\u003cbr/\u003eAnd so, this message speaks to us today—as clearly as it did to Daniel and those who would listen. When times are bleak and nations are warring against each other, when it\u0026#39;s hard to discern just who our allies and enemies are, God is here with us. When persecution increases, when hatred and malice become commonplace, God shows that He watches over our lives \u003ci\u003eand\u003c/i\u003e our eternal futures.\u003cbr/\u003e\u003cbr/\u003eHow? Through the sacrificial life, death, and resurrection of His Son, Jesus Christ. Now, in devil-trampling victory, Jesus reigns over all in heaven and earth (see Revelation 5:12-13). Nothing we experience—no matter how awful or upsetting—is beyond His concern for us. In faith, we are His, and He will lead us through whatever would test that faith—be it a fiery furnace, a den of lions, or the crisis you\u0026#39;re facing today.\u003cbr/\u003e\u003cbr/\u003eNo stranger to life\u0026#39;s difficult twists and turns, like Daniel, the apostle Peter knew God would be there for him, too, protecting him unto eternal life. \u0026#34;And after you have suffered a little while, the God of all grace, who has called you to His eternal glory in Christ, will Himself restore, confirm, strengthen, and establish you\u0026#34; (1 Peter 5:10). What a wonderful encouragement!\u003cbr/\u003e\u003cbr/\u003eYou see, that\u0026#39;s what God does: He keeps our names safely in the book of life.\u003cbr/\u003e\u003cbr/\u003e\u003cb\u003eTHE PRAYER\u003c/b\u003e: Heavenly Father, strengthen us to face life\u0026#39;s ordeals with the eyes of faith. In Jesus\u0026#39; Name we pray. Amen.\u003cbr/\u003e\u003cbr/\u003eThis Daily Devotion was written by Paul Schreiber.\u003cbr/\u003e\u003cbr/\u003e\u003cb\u003eReflection Questions\u003c/b\u003e:\u003cbr/\u003e\u003cbr/\u003e1.\tWhat\u0026#39;s your favorite story from the book of Daniel?\u003cbr/\u003e\u003cbr/\u003e2.\tDo you have an example of how God strengthened you after you had \u0026#34;suffered a little while\u0026#34;?\u003cbr/\u003e\u003cbr/\u003e3.\tWhen did you first understand that your name was written in God\u0026#39;s book of life?                    \n                        \n                       \n                        \u003cbr/\u003e\n                      \u003c/p\u003e"
	expectPlain := `Daniel 12:1-3 - At that time shall arise Michael`
	expectImage := "https://www.lhm.org/uploads/devo_1200_goodnews.jpg"

	if name != expectName {
		t.Fatalf(`ERROR: Scrape(%s) returned name as %s, expected %s`, link, name, expectName)
	}

	if image != expectImage {
		t.Fatalf(`ERROR: Scrape(%s) returned image as %s, expected %s`, link, image, expectImage)
	}

	if html != expectHtml {
		t.Fatalf(`
ERROR: Scrape(%s) returned html as: 

%s

expected: 

%s
		`, link, html, expectHtml)
	}

	if !strings.Contains(plain, expectPlain) {
		t.Fatalf(`
ERROR: Scrape(%s) returned plain as:

%s

expected to contain:

%s
		`, link, plain, expectPlain)
	}
}
