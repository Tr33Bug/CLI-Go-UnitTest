
<h1 style="text-align: center"> Funktionsweise des GO-Testing-Tools </h1>

<h4 style="text-align: center">Götz-Henrik Wiegand | Matr.</h4>



<div style="text-align: center"><tiny>Projektbericht im Rahmen des 5. Semesters des Informatik Bachelor Studiums</tiny></div>
<div style="text-align: center"><small>Betreuer: Prof. Dr. Martin Sulzmann</small></div>
<div style="text-align: center"><small>Hochschule Karlsruhe – Technik und Wirtschaft</small></div>
<div style="text-align: center"><small>an der</small></div>
<div style="text-align: center"><small>Fakultät für Informatik und Wirtschaftsinformatik</small></div>





___
##### Zusammenfassung

Der gezielte Einsatz von Tests in der Entwicklung kann viele Probleme und Stunden sparen. 
	
Die Programmiersprache Go wird 

% Worum geht es?
% Wie bist du vorgegangen?
% Was sind deine wichtigsten Ergebnisse?
% Was bedeuten deine Ergebnisse?

---
## Einleitung
% ca. 1 Seite
% Was wird in der Arbeit behandelt?
% Was der Kontext/Der Rahmen der Arbeit

Im Rahmen der Projektarbeit wird ein Programm in Go geschrieben, an dem das GO-test-Tool vorgestellt werden soll. Hierfür werden gängige Testmehtoden wie z.b. Unit-Tests gezeigt und wie diese in GO umgesetzt werden können. 

% Hintergrund und Einordnung
% Motivation
% Relevanz der Fragestellung
% Abgrenzung

% Haubteil
## Material und Methoden
% Wie bin ich vorgegangen? 
% Was habe ich genutzt? (Prgramme und co.)
% Welche Methoden werden getestet bzw. welches Tool wird gezeigt.

Die Programmiersprache Go wurde vor allem für das Entwickeln von skalierbaren Webanwendungen und Cloud Computing entwickelt. Neben diesen Aufgaben, hat sich GO auch als eine gängige Sprache für das Entwickeln von Comand Line Interfaces (CLI) herausgestellt. 
Mit Hilfe eines selbstgeschriebenen CLIs[^GitHub_Tr33Bug] werden verschiedene Beispiele dargestellt und Möglichkeiten aufgezeigt, die das Go-Testing Tool bereitstellt. Die Entwicklung des Programms sowie die Ausführung der Tests werden mithilfe von Visual Studio Code und der GO-Erweiterung[^VS Code] durchgeführt. Für die Programmierung des CLIs werden das Cobra-Framework[^GitHub_Cobra] sowie die GO eigenen Pakete[^GO-Packages] "fmt", "testing" und "strings" genutzt. Alle Testbefehle und Programmausführungen werden über den Terminal ausgeführt und beschrieben. Die Einbettung und Umsetzung der Ausführung von Tests in Grafischen Oberflächen wird hier nicht weiter berücksichtigt. 

### Tests ausführen
Der erste Tests soll die Funktion `printHelloWorld()` in `./cmd/bsp1.go` testen. Hierfür wurde eine weitere Datei im gleichen Ordner und Paket mit dem Namen `bsp1_test.go` angelegt. Die Funktion `printHelloWorld()` gibt als Rückgabewert einen String mit dem Inhalt "*Hello World*" zurück.
In der Comandozeile wird die Funktion nach dem Installieren mit `go install` durch das eingeben von `myCli bsp1` aufgerufen und gibt den Rückgabewert der Funktion `printHelloWorld()` auf dem Terminal aus.

Um die Funktion zu testen wird in der Test-Datei (`./cmd/bsp1_test.go`) eine Funktion angelegt , die überprüfen soll, ob der zurückgegebene Wert mit dem String "*Hello World*" übereinstimmt. Sollte der Test fehlschlagen, so wird als Fehler-Mitteilung der Text " *'Hello World' test failed*" zurückgegeben. 

Um diesen Test aus zu führen benötigen wir den `go test`-Befehl. Wird dieser ohne weitere Argumente ausgeführt, so kompiliert GO das aktuelle Verzeichnis und führt alle gefundenen Tests darin aus.
Um eine spezifische Datei zu testen reicht der Befehl `go test` mit dem entsprechenden Pfad der Testdatei dahinter. In dem hier gezeigten Beispiel also `go test ./cmd/`. 

Bei einer erfolgreichen Ausführung bekommen wir in der Ausgabe eine kurze Übersicht über den Status, den Paketpfad und die Dauer der Ausführung des Testes. Um eine Auflistung der durchgeführten Tests zu bekommen benötigt man das *verbose*-Flagg. Test die mit diesem Flagg ausgeführt werden, zeigen für jeden durchgeführten Tests ein Startpunkt, der mit *=== RUN* gekennzeichnet wird und einen Endpunkt, der bei erfolgreichem Test mit *--- PASS* oder wenn ein Fehler auftritt mit *--- FAIL* beendet wird. Logs werden bei erfolgreichem Testdurchlauf auch nur angezeigt, wenn das *verbose*-Flagg gesetzt wurde.  

Das *verbose*-Flagg kann man als `go test -v` oder als `go test -verbose` einsetzen. Im weiteren Verlauf des Berichtes wird nur die kürzere Schreibweise (`-v`)verwendet, alle Befehle sind analog auch mit der vollen Schreibweise(`-verbose`) ausführbar. 





##### Code

```go
// printHelloWorld() aus ./cmd/bsp1.go
func printHelloWorld() string {
	return ("Hello World")
}

// TestPrintHelloWorld() aus ./cmd/bsp1_test.go
func TestPrintHelloWorld(t *testing.T) {
	if "Hello World" != printHelloWorld() {
		t.Error("'Hello World' test failed!")
	}
}
```

##### Ausgabe

```bash
tr33@bug:~$ myCly bsp1
Hello World
tr33@bug:~$ go test ./cmd/bsp1
ok      github.com/Tr33Bug/myCli/cmd    0.164s
tr33@bug:~$ go test -v ./cmd/bsp1
=== RUN   TestPrintHelloWorld
--- PASS: TestPrintHelloWorld (0.00s)
PASS
ok      github.com/Tr33Bug/myCli/cmd    0.162s

```



### Spezifische Tests ausführen

Werden nun Änderungen an einer Funktion vorgenommen und man möchte diese anhand der geschriebenen Tests prüfen, ohne alle Tests des gesamten Projektes aus zu führen, so kann man dies über `go test` verknüpft mit dem *run*-Flagg und dem Namen der Testfunktion erreichen. Damit kann man eine oder mehrere Testfunktionen nach ihrem Namen aufrufen.

Es soll aus dem `./cmd/bsp2_test.go` nur die Funktion `TestPrintHello()` ausgeführt werden. Hierfür wird in dem Terminal `go test -run=TestPrintHello$ ./cmd/` verwendet. Das Dollarzeichen am Ende des Namens legt fest, dass GO nach genau diesem Test suchen soll (direkter Aufruf). Lassen wir das Dollarzeichen weg und führen `go test -run=TestPrint ./cmd/` aus, so werden nach allen Tests gesucht, die mit "*TestPrint*" anfangen und dann ausgeführt("beginnt mit" Aufruf). In unserem Beispiel sind in der Ausgabe dann die Testergebnisse beider Tests und des Tests aus dem `./cmd/bsp1_test.go` zu sehen (siehe Ausgabe).

In der `Test Print World()` Funktion wird mit einem Subtest gearbeitet. Diese sind besonders sinnvoll, um komplexere Funktionen zu testen und abzudecken. In diesem Beispiel zu Demonstrationszwecken stark vereinfacht. Auch diese Subtests können über die *run*-Flagg aufgerufen werden. Hierfür wird zuerst der übergeordnete Test und dann mit einem Schrägstrich der Subtest genannt. Auch hier sind wieder "direkte" oder "beginnen mit"-Aufrufe möglich. 

##### Code

```go
// printHello() und printWorld() aus ./cmd/bsp2.go
func printHello() string {
	return ("Hello")
}

func printWorld() string {
	return ("World")
}

// () aus ./cmd/bsp2_test.go
func TestPrintHello(t *testing.T) {
	if "Hello" != printHello() {
		t.Error("Printing Hello has failed!")
	}

}

func TestPrintWorld(t *testing.T) {
	t.Run("subtestPrintWorld", func(t *testing.T) {
		if "World" != printWorld() {
			t.Error("Printing World has failed!")
		}
	})
}
```

##### Ausgabe

```bash
tr33@bug:~$go test -v -run=TestPrintHello$ ./cmd/
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
PASS
ok      github.com/Tr33Bug/myCli/cmd    0.161s
tr33@bug:~$go test -v -run=TestPrint ./cmd/
=== RUN   TestPrintHelloWorld
    bsp1_test.go:11: This Log should only appear in verbose mode or if something went wrong
--- PASS: TestPrintHelloWorld (0.00s)
=== RUN   TestPrintHello
--- PASS: TestPrintHello (0.00s)
=== RUN   TestPrintWorld
=== RUN   TestPrintWorld/subtestPrintWorld
--- PASS: TestPrintWorld (0.00s)
    --- PASS: TestPrintWorld/subtestPrintWorld (0.00s)
PASS
ok      github.com/Tr33Bug/myCli/cmd    0.160s
tr33@bug:~$go test -v -run=TestPrintWorld$/subtestPrint ./cmd/
=== RUN   TestPrintWorld
=== RUN   TestPrintWorld/subtestPrintWorld
--- PASS: TestPrintWorld (0.00s)
    --- PASS: TestPrintWorld/subtestPrintWorld (0.00s)
PASS
ok      github.com/Tr33Bug/myCli/cmd    0.160s
```



## bsp3

##### Code

```go
// waitAndPrint() aus ./cmd/bsp3.go
func waitAndPrint() string {
	time.Sleep(5 * time.Second)
	return "Hello World"
}

// TestWaitAndPrint() aus ./cmd/bsp3_test.go
func TestWaitAndPrint(t *testing.T) {
	// tell the testengine to skip the test in short mode
	if testing.Short() {
		t.Skip("This test is to long, skip in short mode")
	}
	// start test
	if "Hello World" != waitAndPrint() {
		t.Error("Waiting and printing Hello World has failed")
	}

}
```



##### Ausgabe

```bash
tr33@bug:~$
```



## bsp4

##### Code

```go
// sumUp() aus ./cmd/bsp4.go
func sumUp(x, y int) int {
	return x + y
}

// TestSumUp() aus ./cmd/bsp4_test.go
func TestSumUp(t *testing.T) {
	a, b := 5, 5
	result := sumUp(a, b)
	t.Log("The following test should fail. If you dont want it to fail, change the test in ./cmd/bsp4_test.go")
	if result != 11 {
		t.Errorf("This test should fail. Dont worry!")
	}
	
	// Replace with te following to fix the test
    /*
		if result != 10 {
			t.Errorf("Failed to calculate the sum of 5 and 5")
		}
	*/

}
```

##### Ausgabe

```bash
tr33@bug:~$
```

## bsp5

##### Code

```go
// TestStartRace() mit countUp() und countDown() aus ./cmd/bsp5_test.go
var superCount int

func TestStartRace(t *testing.T) {
	// Start subtest
	t.Run("raceDown", func(t *testing.T) {
		t.Parallel()
		countDown(t)
	})
	// Start second subtest
	t.Run("raceUp", func(t *testing.T) {
		t.Parallel()
		countUp(t)
	})

}

func countUp(t *testing.T) {
	superCount := 1
	t.Log(superCount)
}

func countDown(t *testing.T) {
	superCount := -1
	t.Log(superCount)
}

```

##### Ausgabe

```bash
tr33@bug:~$
```

## bsp6

##### Code

```go
// randomReturnHello() aus ./cmd/bsp6.go
func randomReturnHello() string {
	rand.Seed(time.Now().UnixNano())
	answers := []string{
		"...",
		"Hello",
		"Hello",
		"...",
		"...",
	}

	// return answers[rand.Intn(len(answers))]
	return answers[rand.Intn(len(answers))]

}


// TestRandomReturnHello() aus ./cmd/bsp6_test.go
func TestRandomReturnHello(t *testing.T) {
	answer := randomReturnHello()
	if answer != "..." {
		t.Errorf("The test failed, probably try out one more time...")
	}

}
```

##### Ausgabe

```bash
tr33@bug:~$
```

So bald das Projekt größer wird und damit auch der Umfang der Tests größer wird, sollte man nicht bei jedem Testdurchlauf alle 

## Ergebnisse
% Was ist durch die Methoden rausgekommen.
% Was sind die Ergebnisse des Tools

## Diskussion
% Was sind vergleichbare Tools? Vor und Nachteile? (JUnit)
% Ergebnisse diskutieren und kommentieren

## Fazit und Ausblick
% Weitere Möglichkeiten
% Gesamtdarstellung der Ergebnisse

[^VS Code]: Visual Studio Code Extension für GO: https://code.visualstudio.com/docs/languages/go

[^GO-Packages]: Die Dokumentation zum GO-Testing Paket: https://golang.org/pkg/testing/

[^GitHub_Tr33Bug]: Das Repository, in welchem das gesammte Comand Line Interface zu finden ist: https://github.com/Tr33Bug/myCli

[^GitHub_Cobra]: Das Repository, in welchem das Cobra-Tool zur Erstellung von Comand Line Interfaces zu finden ist: https://github.com/spf13/cobra

[^Blog_A.Edwards]: Alex Edwards Blogeintrag über die verschiedenen Tools, die GO mitbringt: https://www.alexedwards.net/blog/an-overview-of-go-tooling

artikel über go \\
https://entwickler.de/online/development/einfuehrung-programmierung-go-166821.html

​	