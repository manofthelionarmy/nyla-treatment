package google

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"nylatreatment/internal/model/medicine"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

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

// DefaultConfig stores default values
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

	// TODO: override defaults by detecting config.json in .config/nyla-treatment/config.json

	svc := CalendarService{
		cfg: DefaultConfig,
	}

	for _, f := range opts {
		f(&svc.cfg)
	}
	return &svc
}

// AddToCalendar adds a calendar event to google calendar
func (svc CalendarService) AddToCalendar(medicineRecord medicine.MedicineRecord) error {
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

	calendarSvc, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		return fmt.Errorf("Unable to retrieve Calendar client: %v", err)
	}

	// because the db has a different timezone, I have to make sure I get it set up correctly in the application
	startTime := convertToRFC339(medicineRecord.TimeTaken)
	endTime := convertToRFC339(medicineRecord.TimeTaken.Add(15 * time.Minute))
	event := calendar.Event{
		Summary:     "Time for Medicine",
		Description: medicineRecord.Name + " Treatment",
		Start: &calendar.EventDateTime{
			DateTime: startTime,
			TimeZone: time.Now().Local().String(),
		},
		End: &calendar.EventDateTime{
			DateTime: endTime,
			TimeZone: time.Now().Local().String(),
		},
		Reminders: &calendar.EventReminders{
			UseDefault: false,
			Overrides: []*calendar.EventReminder{
				&calendar.EventReminder{
					Method:  "email",
					Minutes: 30,
				},
				&calendar.EventReminder{
					Method:  "email",
					Minutes: 15,
				},
			},
			// https://stackoverflow.com/a/65441406
			ForceSendFields: []string{"UseDefault"},
		},
	}
	generatedEvent, err := calendarSvc.Events.Insert("primary", &event).Do()
	if err != nil {
		return err
	}

	sb := strings.Builder{}
	sb.WriteString("Calendar Event Created!" + "\n")
	sb.WriteString("Summary: " + generatedEvent.Summary + "\n")
	sb.WriteString("Start Time: " + medicineRecord.TimeTaken.Format("Mon, Jan _2 3:04PM") + "\n")
	sb.WriteString("calendar link: " + generatedEvent.HtmlLink + "\n")
	fmt.Println(sb.String())
	return nil
}

func convertToRFC339(t time.Time) string {
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()
	newTime := time.Date(year, month, day, hour, minute, 0, 0, time.Now().Location())
	return newTime.Format(time.RFC3339)
}

func getAuthCode(config *oauth2.Config, tokCh chan *oauth2.Token) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		values := r.URL.Query()
		authCode := values.Get("code")

		tok, err := config.Exchange(context.TODO(), authCode)
		if err != nil {
			log.Printf("Unable to retrieve token from web: %v", err)
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

func printAuthURL(config *oauth2.Config) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	printAuthURL(config)

	serverMux := http.NewServeMux()
	tokCh := make(chan *oauth2.Token)
	serverMux.HandleFunc("/auth", getAuthCode(config, tokCh))
	// TODO: make configurable
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
