package google

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type CalendarService struct {
	cfg Config
}

type Config struct {
	credentialsFile string
	tokenFile       string
}

// DefaultConfig is this
var DefaultConfig = Config{
	credentialsFile: ".config/nyla-treatment/credentials.json",
	tokenFile:       ".config/nyla-treatment/token.json",
}

type CalendarServiceOpt func(cfg *Config)

func NewCalendarService(opts ...CalendarServiceOpt) *CalendarService {
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	DefaultConfig.credentialsFile = filepath.Join(homedir, DefaultConfig.credentialsFile)
	DefaultConfig.tokenFile = filepath.Join(homedir, DefaultConfig.tokenFile)

	svc := CalendarService{
		cfg: DefaultConfig,
	}

	for _, f := range opts {
		f(&svc.cfg)
	}
	return &svc
}

func (svc CalendarService) AddToCalendar() error {
	ctx := context.Background()
	b, err := os.ReadFile(svc.cfg.credentialsFile)
	if err != nil {
		return fmt.Errorf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	scopes := []string{calendar.CalendarEventsScope}
	config, err := google.ConfigFromJSON(b, scopes...)
	if err != nil {
		return fmt.Errorf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config, svc.cfg.tokenFile)

	_, err = calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}

	// t := time.Now().Format(time.RFC3339)
	// events, err := srv.Events.List("primary").ShowDeleted(false).
	// 	SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	// if err != nil {
	// 	return fmt.Errorf("Unable to retrieve next ten of the user's events: %v", err)
	// }
	// fmt.Println("Upcoming events:")
	// if len(events.Items) == 0 {
	// 	fmt.Println("No upcoming events found.")
	// } else {
	// 	for _, item := range events.Items {
	// 		date := item.Start.DateTime
	// 		if date == "" {
	// 			date = item.Start.Date
	// 		}
	// 		fmt.Printf("%v (%v)\n", item.Summary, date)
	// 	}
	// }
	return nil
}

func getAuthCode(config *oauth2.Config, tokCh chan *oauth2.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		authCode := values.Get("code")

		tok, err := config.Exchange(context.TODO(), authCode)
		if err != nil {
			log.Println("Unable to retrieve token from web: %v", err)
			tokCh <- nil
			return
		}
		tokCh <- tok
	}
}

func WithCredentialsFile(pth string) CalendarServiceOpt {
	return func(cfg *Config) {
		cfg.credentialsFile = pth
	}
}

func WithTokenFile(pth string) CalendarServiceOpt {
	return func(cfg *Config) {
		cfg.tokenFile = pth
	}
}

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config, tokenFileDestination string) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := tokenFileDestination
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

func printAuthUrl(config *oauth2.Config) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	printAuthUrl(config)

	serverMux := http.NewServeMux()
	tokCh := make(chan *oauth2.Token)
	serverMux.HandleFunc("/auth", getAuthCode(config, tokCh))
	server := http.Server{
		Addr:    ":8081",
		Handler: serverMux,
	}
	// how to close this properly?
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		server.ListenAndServe()
	}()
	server.RegisterOnShutdown(func() {
		wg.Wait()
	})
	// if handler fails, we return nil
	tok := <-tokCh
	server.Shutdown(context.Background())
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}
