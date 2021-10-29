package examples

import (
	"fmt"
	"regexp"
	"strings"

	ac "github.com/grokify/go-adaptivecards"
)

func ExampleCardBullets(linePrefix string) ac.AdaptiveCard {
	strs := LorusIpsumStrings(regexp.MustCompile(`\.`))
	strs = strs[:10]
	if len(linePrefix) > 0 {
		for i := range strs {
			strs[i] = linePrefix + strings.TrimSpace(strs[i])
		}
	}
	card := ac.NewAdaptiveCard()
	card.Body = append(card.Body,
		ac.ElementTextBlock{
			Type:      ac.ElementTypeTextBlock,
			Text:      strings.Join(strs, "\n"),
			Wrap:      true,
			IsVisible: true,
		})
	card.Inflate(true)
	return card
}

func ExampleCardBulletsMulti() ac.AdaptiveCard {
	card := ac.AdaptiveCard{
		Body: []ac.Element{},
	}
	blocks := TestTextBlocks()
	strs := LorusIpsumStrings(nil)
	strInc := 0
	bullet := "1."
	for _, b := range blocks {
		b.Text = FontString(b)
		b.Wrap = true
		card.Body = append(card.Body, b)
		textParts := []string{}
		for i := 0; i < 3; i++ {
			strIdx := modInt(strInc, len(strs))
			textParts = append(textParts, fmt.Sprintf(
				"%s %s", bullet, strings.TrimSpace(strs[strIdx])))
			strInc++
		}
		card.Body = append(card.Body,
			ac.ElementTextBlock{
				Type: ac.ElementTypeTextBlock,
				Wrap: true,
				Text: strings.Join(textParts, "\n"),
			})
		if bullet == "1." {
			bullet = "*"
		} else {
			bullet = "1."
		}
	}
	card.Inflate(true)
	return card
}

var rxPunct = regexp.MustCompile(`[[:punct:]]`)

func LorusIpsumStrings(rxSplit *regexp.Regexp) []string {
	if rxSplit == nil {
		rxSplit = rxPunct
	}
	parts := []string{}
	m := rxSplit.Split(LorusIpsum(), -1)
	for i := range m {
		parts = append(parts, m[i])
	}
	return parts
}

func LorusIpsum() string {
	return `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Aliquam ac erat sollicitudin, fermentum felis id, hendrerit dui. In hac habitasse platea dictumst. In pulvinar pharetra dui vel aliquet. Vivamus convallis libero ut sapien gravida efficitur. Ut laoreet malesuada bibendum. Duis et ultricies turpis. Phasellus ac lacinia massa, eu fringilla massa. Cras non est mi. Etiam eget scelerisque arcu, id pharetra velit.

    In sit amet volutpat nisl. Aliquam lacinia efficitur dolor sed interdum. Quisque interdum velit velit, in feugiat orci dictum non. Aliquam in purus magna. Maecenas porttitor efficitur diam sodales facilisis. Sed vel auctor libero, in faucibus eros. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia curae; Curabitur id enim lacinia, cursus sem non, tincidunt dolor. Nam lacinia in tortor ac tincidunt. Etiam diam ipsum, sollicitudin et augue eget, aliquam imperdiet est. Duis dui quam, consequat sed lectus semper, posuere imperdiet est. Suspendisse imperdiet nibh justo, eget tincidunt est euismod gravida. Quisque sagittis tempor facilisis.
    
    Nam gravida nec nibh at sollicitudin. Suspendisse potenti. Mauris pulvinar arcu sed feugiat lacinia. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Maecenas quis magna sapien. Suspendisse elementum pellentesque ipsum, maximus aliquam sem. Nam sit amet hendrerit massa, nec volutpat felis. Ut at maximus eros, id porttitor ex. Nunc in ullamcorper nulla, vel pellentesque ipsum. Lorem ipsum dolor sit amet, consectetur adipiscing elit. Mauris eu laoreet leo. Sed aliquet justo mi, suscipit convallis tortor commodo sed. Aenean condimentum iaculis elementum. Donec risus lorem, accumsan nec nisi sed, placerat mattis lacus.
    
    Vestibulum consequat mi id libero semper condimentum. Maecenas efficitur dolor vel pretium tincidunt. Suspendisse et ante vel erat facilisis consectetur vel et ex. Integer eu varius lectus, vel porttitor mauris. Sed bibendum rhoncus egestas. Mauris porttitor enim ut nisi malesuada molestie. Integer mi augue, sodales efficitur tellus quis, placerat condimentum orci. Vivamus feugiat elit non neque varius faucibus. Phasellus feugiat molestie dolor interdum bibendum. Vestibulum eu semper sapien. Sed ornare, odio imperdiet sagittis pharetra, augue massa dictum lectus, quis dapibus metus sem id nunc. Quisque vel interdum neque. Praesent auctor sollicitudin lacus, aliquet interdum dolor hendrerit at. Curabitur ut erat luctus, pulvinar odio non, dictum felis.
    
    Nunc dolor nisi, tempus vitae nisi vitae, cursus pellentesque felis. Sed pharetra nisi tortor, vitae porttitor orci ultricies id. Cras id sollicitudin augue, sit amet posuere tortor. Donec eget hendrerit mi. Donec eget urna nec arcu lobortis iaculis. In sodales consequat ornare. Ut suscipit orci at sem tempor, ut interdum arcu mollis. Nullam et pretium risus. Nunc dignissim iaculis euismod. In mattis, velit id dapibus tristique, lacus est finibus risus, a luctus magna nisl id lacus. Etiam sodales hendrerit leo, id convallis arcu gravida eget. Suspendisse commodo, nunc id pharetra porta, libero enim fermentum nibh, in consectetur leo sem sed mi. Suspendisse iaculis justo eget mi mattis, quis dapibus urna auctor. Vestibulum ut dictum tellus, volutpat malesuada ante. Praesent dapibus condimentum egestas. Etiam sollicitudin justo sed vestibulum blandit.
    
    Maecenas suscipit efficitur vehicula. Phasellus auctor, turpis sed blandit rutrum, nunc sem viverra ante, ut facilisis eros elit a lacus. Sed sit amet nibh non neque efficitur placerat. Pellentesque gravida leo sed nisl gravida, ut fermentum libero fermentum. Integer id velit ornare enim pharetra scelerisque. Phasellus sagittis leo a condimentum tristique. Fusce sed lacus vulputate augue fringilla tristique. Pellentesque ac augue id ante fermentum faucibus. Duis ornare tellus ut molestie efficitur.
    
    In laoreet, ex sagittis efficitur aliquam, orci sapien laoreet dolor, malesuada iaculis augue nibh vel risus. Duis condimentum ut quam sed tristique. Cras congue velit quis nisl aliquam, et hendrerit lorem iaculis. Quisque mattis accumsan urna, sed maximus sapien pulvinar vel. Fusce vitae vehicula lacus, sed dignissim massa. Suspendisse dictum ligula eu suscipit scelerisque. Ut porta felis vel enim mollis suscipit. Phasellus sit amet ipsum iaculis, bibendum dolor vel, rhoncus libero. Nam neque ante, commodo vitae ornare nec, congue vel tortor. Mauris vitae urna diam. Maecenas finibus nibh id tortor porttitor, non congue tortor semper. Morbi ullamcorper leo ut sodales posuere.
    
    In hac habitasse platea dictumst. Proin in fringilla tellus. Duis sit amet neque eu diam consequat tincidunt et quis massa. Donec at sapien sed metus tempus hendrerit. Pellentesque convallis mauris nibh, ornare vehicula massa semper at. Duis iaculis elementum neque quis ullamcorper. Proin faucibus non eros cursus condimentum. Maecenas euismod sit amet leo a placerat. Nunc lobortis justo id ligula maximus, vel convallis ligula viverra. Nam eu facilisis velit, nec porttitor ipsum. Suspendisse potenti. Suspendisse vestibulum est sed pulvinar sagittis. Nullam elementum quam sit amet placerat maximus.
    
    Aliquam pretium ex id dolor ultrices imperdiet. Donec eu massa vitae felis convallis dignissim ut vitae nunc. Praesent iaculis vestibulum odio egestas sollicitudin. Praesent efficitur porttitor est eget semper. Sed dignissim porttitor egestas. Nullam blandit vulputate nibh in iaculis. Praesent vel elit non massa scelerisque vestibulum at vitae dolor. Morbi et tellus felis. Morbi quis dui eu nibh vulputate dictum ornare et felis. Duis sed dui varius, viverra velit et, accumsan arcu. Duis volutpat ipsum vitae arcu cursus tempus. Maecenas dignissim venenatis sapien suscipit aliquam. Aenean quis ex lacinia, cursus odio vel, vehicula nisl. Nam molestie dui eget massa tristique accumsan. Pellentesque bibendum enim dolor.
    
    Sed semper leo gravida ante sollicitudin pharetra. Integer elementum malesuada augue ut mattis. Pellentesque tempus sapien eget velit aliquam molestie. Vestibulum vehicula fringilla risus vitae ornare. Nulla iaculis ex in ligula fringilla fringilla. Quisque scelerisque, mi at lacinia volutpat, erat justo egestas quam, sed posuere velit lorem rutrum nulla. Nunc dictum augue consequat massa egestas tincidunt. Integer in nisi velit. Integer in justo et lacus molestie maximus quis ut dui. Etiam molestie libero a ullamcorper pretium. Suspendisse non ornare odio. Nunc non tortor ultricies, pulvinar purus in, porta mi. Suspendisse tellus felis, lacinia et placerat ac, convallis non justo.`
}

func modInt(a, b int) int { return (a%b + b) % b }
