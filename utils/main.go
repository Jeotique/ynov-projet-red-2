package utils

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"projet/structs"
	"projet/values"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/mattn/go-tty"
)

func ListenInput() string {
	println("listening for input")
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		println("aaa")
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		return string(r)
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func Write(text string, args ...any) {
	if args != nil && len(args) > 0 {
		fmt.Printf(text, args)
	} else {
		fmt.Print(text)
	}
}

func Writeln(text string, args ...any) {
	if args != nil && len(args) > 0 {
		fmt.Printf(text, args)
		fmt.Println()
	} else {
		fmt.Println(text)
	}
}

func Writeanim(text string, timee time.Duration, args ...any) {
	for _, character := range text {
		fmt.Print(string(rune(character)))
		time.Sleep(timee * time.Millisecond)
	}
}

/*func WaitForInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	return scanner.Text()
}*/

func WaitForInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func RandomNumber(min int, max int) int {
	/*rand.Seed(time.Now().UnixNano())
	return rand.Intn((max - min + 1) + min)*/
	rand.Seed(time.Now().UnixNano())
	return rand.Intn((max - min + 1)) + min
}

func WaitForNumberInput() (bool, int) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatal(err)
	}
	if !IsNumeric(scanner.Text()) {
		return false, -1
	}
	number, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
		return false, -1
	}
	return true, number
}

func IsNumeric(s string) bool {
	for _, char := range s {
		if !('0' <= char && char <= '9') {
			return false
		}
	}
	return true
}

func PlaySound(path string, rate ...int) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	if len(rate) < 1 {
		rate = []int{1}
	}

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	defer streamer.Close()

	//speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Init(format.SampleRate*beep.SampleRate(rate[0]), format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}

func CenterText(text string, width int) string {
	return strings.Repeat(" ", width) + text
}

func WriteColorLn(text string, color string) {
	switch color {
	case "red":
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case "green":
		fmt.Printf("\033[32m%s\033[0m\n", text)
	case "yellow":
		fmt.Printf("\033[33m%s\033[0m\n", text)
	case "blue":
		fmt.Printf("\033[34m%s\033[0m\n", text)
	case "magenta":
		fmt.Printf("\033[35m%s\033[0m\n", text)
	case "cyan":
		fmt.Printf("\033[36m%s\033[0m\n", text)
	case "white":
		fmt.Printf("\033[37m%s\033[0m\n", text)
	default:
		fmt.Printf("\033[37m%s\033[0m\n", text)
	}
	fmt.Println()
}

func WriteanimColor(text string, timee time.Duration, color string, args ...any) {
	for _, character := range text {
		switch color {
		case "red":
			fmt.Printf("\033[31m%s\033[0m", string(character))
		case "green":
			fmt.Printf("\033[32m%s\033[0m", string(character))
		case "yellow":
			fmt.Printf("\033[33m%s\033[0m", string(character))
		case "blue":
			fmt.Printf("\033[34m%s\033[0m", string(character))
		case "magenta":
			fmt.Printf("\033[35m%s\033[0m", string(character))
		case "cyan":
			fmt.Printf("\033[36m%s\033[0m", string(character))
		case "white":
			fmt.Printf("\033[37m%s\033[0m", string(character))
		default:
			fmt.Printf("\033[37m%s\033[0m", string(character))
		}
		time.Sleep(timee * time.Millisecond)
	}
}

func WriteanimColorln(text string, timee time.Duration, color string, args ...any) {
	for _, character := range text {
		switch color {
		case "red":
			fmt.Printf("\033[31m%s\033[0m", string(character))
		case "green":
			fmt.Printf("\033[32m%s\033[0m", string(character))
		case "yellow":
			fmt.Printf("\033[33m%s\033[0m", string(character))
		case "blue":
			fmt.Printf("\033[34m%s\033[0m", string(character))
		case "magenta":
			fmt.Printf("\033[35m%s\033[0m", string(character))
		case "cyan":
			fmt.Printf("\033[36m%s\033[0m", string(character))
		case "white":
			fmt.Printf("\033[37m%s\033[0m", string(character))
		default:
			fmt.Printf("\033[37m%s\033[0m", string(character))
		}
		time.Sleep(timee * time.Millisecond)
	}
	fmt.Println()
}

func WriteColor(text string, color string) {
	switch color {
	case "red":
		fmt.Printf("\033[31m%s\033[0m\n", text)
	case "green":
		fmt.Printf("\033[32m%s\033[0m\n", text)
	case "yellow":
		fmt.Printf("\033[33m%s\033[0m\n", text)
	case "blue":
		fmt.Printf("\033[34m%s\033[0m\n", text)
	case "magenta":
		fmt.Printf("\033[35m%s\033[0m\n", text)
	case "cyan":
		fmt.Printf("\033[36m%s\033[0m\n", text)
	case "white":
		fmt.Printf("\033[37m%s\033[0m\n", text)
	default:
		fmt.Printf("\033[37m%s\033[0m\n", text)
	}
}

func ScrollTextAtLine(text string, width int, line int) {
	// Remplir avec des espaces pour faire entrer le texte de manière fluide
	displayText := CenterText(fmt.Sprintf("%-*s", len(text)+width, text), 50)

	// Faire défiler le texte
	for i := 0; i <= len(displayText); i++ {
		// Replace le curseur à la ligne spécifiée (sans affecter les autres lignes)
		fmt.Printf("\033[%d;1H", line)
		// Effacer la ligne où le texte doit défiler
		fmt.Print("\033[2K")
		// Affiche la partie défilante du texte
		if i+width < len(displayText) {
			fmt.Print(displayText[i : i+width])
		} else {
			fmt.Print(displayText[i:])
		}

		// Attendre avant de passer à la prochaine itération
		time.Sleep(100 * time.Millisecond)
	}
}

func DisplayTwoAsciiSideBySide(ascii1, ascii2 string, spacing int) {
	// Diviser chaque dessin ASCII en lignes
	lines1 := strings.Split(ascii1, "\n")
	lines2 := strings.Split(ascii2, "\n")

	// Trouver la ligne la plus longue pour chaque ASCII pour alignement
	maxLength1 := getMaxLineLength(lines1)
	maxLength2 := getMaxLineLength(lines2)

	// Ajuster la longueur des deux dessins pour qu'ils aient le même nombre de lignes
	maxLines := max(len(lines1), len(lines2))
	lines1 = adjustLineCount(lines1, maxLines)
	lines2 = adjustLineCount(lines2, maxLines)

	// Espacement entre les deux dessins
	space := strings.Repeat(" ", spacing)

	// Afficher les dessins côte à côte, centrés horizontalement
	for i := 0; i < maxLines; i++ {
		// Formater chaque ligne en centrant les deux dessins
		fmt.Printf("%s%s%s\n", padRight(lines1[i], maxLength1), space, padRight(lines2[i], maxLength2))
	}
}

// Fonction pour obtenir la longueur de la ligne la plus longue dans un dessin ASCII
func getMaxLineLength(lines []string) int {
	maxLength := 0
	for _, line := range lines {
		if len(line) > maxLength {
			maxLength = len(line)
		}
	}
	return maxLength
}

// Fonction pour ajuster le nombre de lignes du dessin ASCII
func adjustLineCount(lines []string, maxLines int) []string {
	for len(lines) < maxLines {
		lines = append(lines, "") // Ajouter des lignes vides si nécessaire
	}
	return lines
}

// Fonction pour ajouter des espaces à droite d'une ligne jusqu'à une longueur donnée
func padRight(str string, length int) string {
	return fmt.Sprintf("%-"+fmt.Sprint(length)+"s", str)
}

// Fonction pour obtenir le maximum entre deux entiers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func GetRandomDogClone(dogs map[string]structs.Dog) structs.Dog {

	if values.Training {
		return structs.NewDog(
			values.RobotDog.Name,
			values.RobotDog.Damage,
			values.RobotDog.MaxDamage,
			values.RobotDog.Health,
			values.RobotDog.MaxHealth,
			values.RobotDog.DodgeChance,
			values.RobotDog.Race,
		)
	}

	// Récupérer les clés de la map
	keys := make([]string, 0, len(dogs))
	for key := range dogs {
		keys = append(keys, key)
	}

	// Sélectionner une clé aléatoirement
	randomKey := keys[rand.Intn(len(keys))]

	// Récupérer le chien correspondant à la clé aléatoire
	selectedDog := dogs[randomKey]

	// Cloner le chien (créer une nouvelle instance indépendante)
	clonedDog := structs.NewDog(
		selectedDog.Name,
		selectedDog.Damage,
		selectedDog.MaxDamage,
		selectedDog.Health,
		selectedDog.MaxHealth,
		selectedDog.DodgeChance,
		selectedDog.Race,
	)

	return clonedDog
}

func GetClonedDog(dogs map[string]structs.Dog, key string) structs.Dog {

	// Récupérer le chien correspondant à la clé
	selectedDog := dogs[key]

	// Cloner le chien (créer une nouvelle instance indépendante)
	clonedDog := structs.NewDog(
		selectedDog.Name,
		selectedDog.Damage,
		selectedDog.MaxDamage,
		selectedDog.Health,
		selectedDog.MaxHealth,
		selectedDog.DodgeChance,
		selectedDog.Race,
	)
	return clonedDog
}
