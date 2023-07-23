package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/Delta456/box-cli-maker/v2"
	"github.com/common-nighthawk/go-figure"
	"github.com/ttacon/chalk"
)

var openPorts []string

func sliceToString(slice []string) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(slice)), ","), "[]")
}

func saveToFile(filename, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	return err
}

func makeReq(targetIp string, port int) {
	url := fmt.Sprintf("%s:%s", targetIp, strconv.Itoa(port))
	conn, err := net.DialTimeout("tcp", url, time.Second*5)
	if err != nil {
		fmt.Println(url, chalk.Red.Color("is closed"))
	} else {
		defer conn.Close()
		fmt.Println(url, chalk.Blue.Color("is Open"))
		openPorts = append(openPorts, url)
	}
}

func main() {
	var targetIp string
	var startPort int
	var endPort int
	var ResultfileName string
	// other ways to enter did not work, I had to do it this way :/
	portScanLogo := figure.NewColorFigure("SIMPLE PORT", "", "green", true)
	portScanLogo2 := figure.NewColorFigure("SCANNER", "", "green", true)

	Box := box.New(box.Config{Px: 2, Py: 1, Type: "Single", Color: "Cyan"})
	portScanLogo.Print()
	portScanLogo2.Print()

	Box.Print("Created By Grewather", "github.com/Grewather/simple-port-scanner")
	fmt.Println(strings.Repeat("=", 50))
	fmt.Print("Target ip adress: ")
	fmt.Scanln(&targetIp)

	fmt.Print("Select port range = first port: ")
	fmt.Scanln(&startPort)
	if startPort < 1 || startPort > 65535 {
		fmt.Println("Invalid starting port. Please enter a valid port number (1-65535).")
		return
	}

	fmt.Print("Select port range = last port: ")
	fmt.Scanln(&endPort)
	if endPort < 1 || endPort > 65535 {
		fmt.Println("Invalid ending port. Please enter a valid port number (1-65535).")
		return
	}

	if startPort > endPort {
		fmt.Println("Invalid port range. The first port should be less than or equal to the last port.")
		return
	}
	fmt.Print("Select result filename: ")
	fmt.Scanln(&ResultfileName)
	for port := startPort; port <= endPort; port++ {
		makeReq(targetIp, port)
	}
	data := sliceToString(openPorts)
	filename := ResultfileName + ".txt"
	err := saveToFile(filename, data)
	if err != nil {
		fmt.Println("Error while saving to file:", err)
		return
	}

	fmt.Println("The open ports were saved to a file: ", filename)

}
