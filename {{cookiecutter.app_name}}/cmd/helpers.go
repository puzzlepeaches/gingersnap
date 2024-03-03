package cmd

import (
	valid "github.com/asaskevich/govalidator"
	"github.com/sirupsen/logrus"
)

func validateEmail(email string) bool {
	return valid.IsEmail(email)
}

func validateURL(url string) bool {
	return valid.IsURL(url)
}

func validateIP(ip string) bool {
	return valid.IsIP(ip)
}

func setLogging(verbose bool) {
	var level logrus.Level
	if verbose {
		level = logrus.DebugLevel
		logrus.Debug("Verbose logging enabled!")
	} else {
		level = logrus.InfoLevel
	}

	logrus.SetLevel(level)     // set level for global logger in logrus
	log.SetDefaultLevel(level) // set level for defaultLogger in log package
}

func checkInput(inputFile string) {
	// Expand the path and check if input file exists
	inputFile, _ = expandPath(inputFile)
	_, err := os.Stat(inputFile)
	if os.IsNotExist(err) {
		logrus.Fatalf("Could not find input file at %s", inputFile)
	} else if err != nil {
		logrus.Fatalf("Could not open input file at %s", inputFile)
	}

	// Open the input file
	file, err := os.Open(inputFile)
	if err != nil {
		logrus.Fatalf("Could not open input file at %s", inputFile)
	}
	defer file.Close()

}

func expandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(homeDir, path[2:])
	}
	return path, nil
}

func writeToFile(filePath string, username string) error {
	// Check if the file exists
	_, err := os.Stat(filePath)

	// Use the appropriate flag based on whether the file exists
	flag := os.O_CREATE | os.O_WRONLY
	if err == nil {
		// If the file exists, append to it
		flag = flag | os.O_APPEND
	} else if !os.IsNotExist(err) {
		// Return any error other than "not exists"
		return err
	}

	// Open the file with the determined flag and permissions
	file, err := os.OpenFile(filePath, flag, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the username to the file
	_, err = file.WriteString(username + "\n")
	if err != nil {
		return err
	}

	return nil
}

func randomUserAgent() string {
	userAgents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_6) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Safari/604.1.38",
		"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",
		"Mozilla/5.0 (Windows NT 6.3; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/63.0.3239.132 Safari/537.36",
		"Mozilla/5.0 (Windows NT 5.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.112 Safari/537.36",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/66.0.3359.181 Safari/537.36",
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.36",
		"Mozilla/5.0 (iPad; CPU OS 8_4 like Mac OS X) AppleWebKit/600.1.4 (KHTML, like Gecko) Version/8.0 Mobile/12H143 Safari/600.1.4",
		"Mozilla/5.0 (Windows NT 6.1; WOW64; rv:54.0) Gecko/20100101 Firefox/54.0",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:57.0) Gecko/20100101 Firefox/57.0",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(userAgents))
	return userAgents[randomIndex]
}

func loadFile(filePath string) ([]string, error) {
	// Expand the path
	expandedFilePath, err := expandPath(filePath)
	if err != nil {
		return nil, err
	}

	// Read the file and return its lines
	content, err := os.ReadFile(expandedFilePath)
	if err != nil {
		return nil, err
	}

	// Split the content into lines, trimming the last empty line if it exists
	lines := strings.Split(strings.TrimSuffix(string(content), "\n"), "\n")
	return lines, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func handleError(err error, message string, exit bool) {
	if err != nil {
		logrus.Errorf("%s: %v", message, err)
		if exit {
			os.Exit(1)
		}
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ensureDir(dirName string) error {
	if !fileExists(dirName) {
		return os.MkdirAll(dirName, 0755)
	}
	return nil
}

func toJSON(data interface{}) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func fromJSON(jsonStr string, target interface{}) error {
	return json.Unmarshal([]byte(jsonStr), target)
}
