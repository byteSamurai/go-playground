package main

import (
	"fmt"
	"strings"
	"sync"
)

// Quelle Bundeslage Cybercrime 2015 ;)
// https://www.bka.de/SharedDocs/Downloads/DE/Publikationen/JahresberichteUndLagebilder/Cybercrime/cybercrimeBundeslagebild2015.html
var vorwort = strings.ToLower(
	`Cybercrime umfasst die Straftaten die sich gegen das Internet Datennetze informationstechnische Systeme oder deren Daten
	richten Cybercrime im engeren Sinne oder die mittels dieser Informationstechnik begangen werden Das Bundeslagebild
	Cybercrime bildet schwerpunktmäßig die im Jahr 2015 polizeilich erfassten Fälle von Cybercrime im engeren Sinne ab Es
	wird über die Entwicklungen im Jahr 2015 berichtet und das Gefahren und Schadenspotenzial von Cybercrime und dessen
	Bedeutung für die Kriminalitätslage in Deutschland beschrieben Grundlage für den statistischen Teil des Lagebildes
	sind die Daten der Polizeilichen Kriminalstatistik Diese beinhaltet alle Straftaten einschließlich der mit Strafe
	bedrohten Versuche die polizeilich bearbeitet und an die Staatsanwaltschaft abgegeben wurden Nicht berücksichtigt sind
	Cybercrime Straftaten bei denen von einer politischen oder nachrichtendienstlichen Motivation ausgegangen wird Seit dem
	Jahr 2014 findet eine PKS Erfassung von Delikten der Cybercrime bundeseinheitlich ausschließlich in Fällen statt in
	denen konkrete Anhaltspunkte für eine Tathandlung innerhalb Deutschlands vorliegen Dies zu erkennen gestaltet sich in
	diesem Phänomen jedoch häufig schwierig In Anbetracht der sehr großen Anzahl von CybercrimeStraftaten die der Polizei
	nicht zur Kenntnis gelangen bedarf es zur Einschätzung des Gefahrenpotenzials von Cybercrime auch der Einbeziehung
	nichtpolizeilicher Informationsquellen Studien von Antivirensoftware Herstellern oder behördlicher Einrichtungen die
	die polizeilich bekanntgewordenen Straftaten ergänzen Aussagen des Lagebildes zu den verschiedenen Erscheinungsformen
	von Cybercrime beruhen daher sowohl auf Erkenntnissen aus dem kriminalpolizeilichen Informationsaustausch zu
	Sachverhalten im Zusammenhang mit Cybercrime als auch auf polizeiexternen Quellen`)

var mobileBedrohungen string = strings.ToLower(
	`Mobile Endgeräte wie Smartphones und Tablets gewinnen weiterhin Marktanteile Gemäß einer repräsentativen Umfrage
	des Bundesverbandes Informationswirtschaft Telekommunikation und neue Medien nutzten Anfang des Jahres 2015 rund 44 Mio
	Bundesbürger ein Smartphone was einer Zunahme von rund zwei Mio innerhalb der zurückliegenden sechs Monate
	entspricht Neben klassischen Funktionen wie Telefonieren und der Verwendung als Foto oder Videokamera werden dabei
	meist folgende Anwendungen genutzt Surfen im Internet zusätzliche Apps soziale Netzwerke Das Abspeichern
	persönlicher Informationen und das Abwickeln sensibler Vorgänge über die mobilen Geräte macht sie zu einem lohnenden
	Angriffsziel für Kriminelle Hinzu kommt ein langsames Updateverhalten der Gerätehersteller so dass bekannte
	Sicherheitslücken oftmals monatelang ungeschlossen bleiben bzw auf zahlreichen Geräten sogar niemals geschlossen
	werden Die steigende Verbreitung sowie die teilweise immer noch mangelnde Sensibilität der Nutzer hinsichtlich der
	Gefahren im Umgang mit mobilen Endgeräten sorgen für eine weiterhin hohe Attraktivität für die Täterseite Ein
	wesentlicher Aspekt ist dass mobile Endgeräte im Gegensatz zum klassischen PC in der Regel ständig online sind und die
	jeweiligen Nutzer mittlerweile einen Großteil ihrer digitalen Aktivitäten über diese Geräte abwickeln wie
	Transaktionen im Onlinebanking Zugriff auf Email Konten und Soziale Netzwerke oder auch Aktivitäten im Bereich des
	Ecommerce oft über entsprechende Apps Dieser Trend steigert die Bedeutung und Attraktivität mobiler Endgeräte für
	Cyberkriminelle was insbesondere durch die Zunahme von Malwareentwicklungen im Bereich mobiler Betriebssysteme
	unterstrichen wird Die Anzahl der Varianten von Schadsoftware für mobile Plattformen nimmt laut BSI weiterhin rasant zu
	Rund 96% der Schadsoftware trifft aufgrund seines Verbreitungsgrads das Betriebssystem Android3.6 Insgesamt 59% der
	laut BSI bis September 2015 von AntivirusProdukten detektierten schadhaften Android Apps sind Trojanische Pferde 2014
	lag ihr Anteil bei 51%`)

var iot string = strings.ToLower(
	`Der Begriff Internet der Dinge beschreibt den Trend dass neben den standardmäßig genutzten Geräten zunehmend
	auch sogenannte intelligente Endgeräte an das Internet angeschlossen und durchgängig online sind Dazu zählen
	beispielsweise Kühlschränke Fernseher oder Router aber auch Sensoren über die andere Geräte via Internet per
	Smartphone oder Tablet gesteuert werden Diese Geräte verfügen in der Regel über eine beachtliche Rechenleistung und
	sind mit entsprechenden Betriebssystemen ausgestattet welche oftmals eigens für die Geräte durch den Hersteller auf
	Basis von Open Source Code entwickelt werden Oftmals verfügen diese sogenannten intelligenten Endgeräte über
	keine oder nur unzureichende Schutzmechanismen und nutzen häufig veraltete Betriebssysteme Software mit
	Sicherheitslücken Für Cyberkriminelle sind solche Geräte leicht angreifbar wobei Infektionen für die Benutzer kaum
	feststellbar sind Der Trend zum sogenannten Smart Home die Vernetzung von Haustechnik und Haushaltsgeräten
	und die gezielte Fernsteuerung der Funktionen verbreitet sich ebenfalls fortwährend Hierdurch eröffnen sich
	vielfältige neue Möglichkeiten zur Begehung von Straftaten Auch die fortschreitende Vernetzung in und von
	Kraftfahrzeugen eröffnet neue Angriffsgelegenheiten für Cyberkriminelle durch die Manipulation interner Steuerbefehle
	von Kraftfahrzeugen Erpressungen könnten ein Ziel solcher Manipulationen sein indem Kraftfahrzeuge beispielsweise
	verschlossen oder sogar gestoppt werden und erst nach Zahlung von Lösegeld wieder freigegeben werden Immer mehr
	Kraftfahrzeuge sind mittlerweile auch internetfähig und verfügen über handelsübliche Internetbrowser die oft analog
	zu den intelligenten Endgeräten über keine oder nur unzureichende Schutzmechanismen verfügen und häufig veraltete
	Betriebssysteme Software mit Sicherheitslücken nutzen Eine Studie geht davon aus dass bis zum Jahr 2020 mehr als eine
	Billion internetfähiger Endgeräte weltweit mit dem Internet verbunden sein werden `)

var industrie4_0 string = strings.ToLower(
	`Die Entwicklung hin zum Internet der Dinge beeinflusst auch die Entwicklungen im Unternehmenssektor Die Nutzung
	privater mobiler Endgeräte und Sozialer Netzwerke im Arbeitskontext nimmt stetig zu Der Trend des Bring your own device
	birgt erhebliche Risiken Die Vereinigung von privaten und beruflichen Internet und Computeraktivitäten auf einem
	privaten Endgerät erleichtert es Cyberkriminellen aufgrund der teilweise schwächeren Absicherung dieser Geräte auch
	auf Unternehmensdaten zuzugreifen Hier werden Einfallstore für den Diebstahl geistigen Eigentums oder
	Wirtschaftsspionage geöffnet Ebenso gewinnt die elektronische und webbasierte Steuerung von Prozessen in Unternehmen
	weiter an Bedeutung Die zunehmende Vernetzung die Abhängigkeit vernetzter sich selbst steuernder Produktionsprozesse
	und Logistikketten von der Verfügbarkeit der Netze und die Problematik der sicheren Trennung und Abschottung dieser
	Netze zum Internet stellen dabei eine große Herausforderung dar Die Folge all dieser Entwicklungen ist eine steigende
	Abhängigkeit der Unternehmen von der Informationstechnik einhergehend mit einem sehr hohen Gefährdungspotenzial
	Angriffe auf die IT Infrastruktur von Unternehmen führen mittlerweile nicht mehr alleine zur Störung der Kommunikation
	sondern bergen vielmehr die Gefahr eines kompletten Produktionsstillstands mit allen damit verbundenen Folgen
	Insbesondere die Gefahr der digitalen Erpressung dürfte für Unternehmen kontinuierlich zunehmen`)

func main() {
	var wg sync.WaitGroup
	texts := [4]string{vorwort, mobileBedrohungen, iot, industrie4_0}

	wordCountMap := make(WordCountMap)
	wordCountChan := make(chan WordCount, 10)

	go WordCountReceiver(&wordCountMap, wordCountChan)

	for _, text := range texts {
		lines := strings.Split(text, "\n")
		for _, line := range lines {
			wg.Add(1)
			go LineToWordCount(line, wordCountChan, &wg)
		}
	}

	wg.Wait()
	close(wordCountChan)
	fmt.Println("\"cybercrime\" appears", wordCountMap["cybercrime"], "times")
}

// Should receive every WordCount from the channel and pass it into the map
func WordCountReceiver(countMap *WordCountMap, counts <-chan WordCount) {
	for count := range counts {
		countMap.IncreaseCounter(count)
	}
}

type WordCountMapping interface {
	IncreaseCounter(wc WordCount)
}

type WordCountMap map[string]int

type WordCount struct {
	word  string
	count int
}

type WordCountList []WordCount

// Produces for every passed string a WordCount with a count of 1
func SimpleCounter(s string) WordCount {
	return WordCount{word: s, count: 1}
}

// Returns a new slice containing the results of applying
// the function `f` to each string in the original slice.
func MapStringToWordCount(vs []string, f func(string) WordCount) WordCountList {
	vsm := make(WordCountList, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// Every line should be split into an array of words. Using the MapStringToWordCount function and the words
// array, every WordCount should be send to the passed channel
func LineToWordCount(line string, wcmChan chan<- WordCount, wg *sync.WaitGroup) {
	defer wg.Done()

	words := strings.Split(line, " ")

	counts := MapStringToWordCount(words, SimpleCounter)

	for _, count := range counts {
		wcmChan <- count
	}
}

// This method should increase every WordCounting type in the given map
func (wcm WordCountMap) IncreaseCounter(wc WordCount) {
	wcm[wc.word] += 1
}