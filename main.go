package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
	"time"

	"github.com/fatih/color"
	"gopkg.in/yaml.v2"
)

type responseModifier struct {
	http.ResponseWriter
	Config *Config
}

func (rm *responseModifier) WriteHeader(code int) {
	rm.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	rm.Header().Set("Access-Control-Allow-Credentials", "true")
	rm.rewriteCookieDomain()
	rm.ResponseWriter.WriteHeader(code)
}

func (rm *responseModifier) rewriteCookieDomain() {
	for _, cookie := range rm.ResponseWriter.Header().Values("Set-Cookie") {
		rewrittenCookie := strings.Replace(cookie, "Domain="+rm.Config.Cookie.Domain.From, "Domain="+rm.Config.Cookie.Domain.To, 1)
		rm.ResponseWriter.Header().Set("Set-Cookie", rewrittenCookie)
	}
}

func main() {

	configPath := flag.String("config", "config.yaml", "Path to the configuration file")
	flag.Parse()

	config, err := loadConfig(*configPath)
	if err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		return
	}

	describeProxyConfig(*config)

	proxy := httputil.NewSingleHostReverseProxy(config.URL)

	proxy.Director = func(req *http.Request) {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
		req.Host = config.URL.Host
		req.URL.Scheme = config.URL.Scheme
		req.URL.Host = config.URL.Host
	}

	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		startTime := time.Now()

		logger := getLogger(req)

		logger("Started %s %s", req.Method, req.URL.Path)

		for _, header := range config.AllowedHeaders {
			res.Header().Add("Access-Control-Allow-Headers", header)
		}
		rm := &responseModifier{res, config}
		proxy.ServeHTTP(rm, req)
		logger("Completed %s %s in %v", req.Method, req.URL.Path, time.Since(startTime))

	})

	http.ListenAndServe(":4000", nil)
}

func banner() {
	println(`
#######                                                    
#       #    #  ####   ####   ####  #    # # ######  ####  
#       #    # #    # #    # #    # #   #  # #      #      
#####   #    # #      #    # #    # ####   # #####   ####  
#       #    # #      #    # #    # #  #   # #           # 
#       #    # #    # #    # #    # #   #  # #      #    # 
#        ####   ####   ####   ####  #    # # ######  ####`)
}

func describeProxyConfig(config Config) {
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	// white := color.New(color.FgWhite).SprintFunc()
	banner()
	fmt.Println("\nProxy Configuration:")
	fmt.Printf("Target: %s\n", cyan(config.Target))
	fmt.Printf("Port: %s\n", cyan(config.Port))
	fmt.Printf("Allowed Headers: %s\n", cyan(strings.Join(config.AllowedHeaders, ", ")))
	fmt.Printf("Cookie Domain: From '%s' to '%s'\n", magenta(config.Cookie.Domain.From), magenta(config.Cookie.Domain.To))
}

type Config struct {
	Target         string   `yaml:"target"`
	Port           int      `yaml:"port"`
	AllowedHeaders []string `yaml:"allowed_headers"`
	Cookie         Cookie   `yaml:"cookie"`
	URL            *url.URL `yaml:"-"`
}

type Cookie struct {
	Domain struct {
		From string `yaml:"from"`
		To   string `yaml:"to"`
	} `yaml:"domain"`
}

func loadConfig(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}
	proxyTarget, err := url.Parse(config.Target)
	if err != nil {
		return nil, fmt.Errorf("failed to parse target URL: %w", err)
	}
	config.URL = proxyTarget

	return config, nil
}

func getLogger(req *http.Request) func(format string, a ...interface{}) {
	cyan := color.New(color.FgCyan).SprintFunc()
	magenta := color.New(color.FgMagenta).SprintFunc()
	white := color.New(color.FgWhite).SprintFunc()

	return func(format string, a ...interface{}) {
		msg := fmt.Sprintf(format, a...)
		logMsg := fmt.Sprintf("[%s] %s %s %s", cyan(time.Now().Format("2006-01-02 15:04:05")), magenta(req.Method), white(req.URL.Path), req.Proto)
		fmt.Printf("%s %s\n", logMsg, msg)
	}
}
