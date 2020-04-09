package main

import (
	"context"
	"crypto/sha256"
	"crypto/tls"
	"encoding/hex"
	"encoding/json"
	"flag"
	"log"
	"net/http"

	awvs "github.com/chennqqi/awvs-go-sdk/go"
)

func pretty(v interface{}) string {
	txt, _ := json.Marshal(v)
	return string(txt)
}

func main() {
	var host, user, pass string
	var https, debug bool
	flag.StringVar(&host, "h", "", "set host")
	flag.StringVar(&user, "u", "", "set email")
	flag.StringVar(&pass, "p", "", "set password")
	flag.BoolVar(&https, "s", true, "set http/https")
	flag.BoolVar(&debug, "d", false, "set debug mode")
	flag.Parse()

	if user == "" || pass == "" {
		log.Println("Both user and pass are required")
		return
	}

	cfg := awvs.NewConfiguration()
	cfg.Host = host
	cfg.Debug = debug
	if !https {
		cfg.Scheme = "http"
	}

	cfg.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := awvs.NewAPIClient(cfg)
	api := client.DefaultApi

	// sha256(password)
	h := sha256.New()
	h.Write([]byte(pass))
	hexPassword := hex.EncodeToString(h.Sum(nil))

	resp, err := api.Login(context.Background(), awvs.LoginReq{
		Email:          user,
		Password:       hexPassword,
		RememberMe:     true,
		LogoutPrevious: true,
	})
	if err != nil {
		log.Println("Login failed", err)
		return
	}

	ctx := context.WithValue(context.Background(), awvs.ContextAPIKey,
		awvs.APIKey{
			Prefix: "",
			Key:    resp.Header.Get("X-Auth"),
		})

	m, _, err := api.GetMe(ctx)
	if err != nil {
		log.Println("GetMe failed", err)
		return
	}

	log.Println("GetMe:", pretty(m))
	info, _, err := api.GetInfo(ctx)
	if err != nil {
		log.Println("GetInfo", err)
		return
	}
	log.Println("GetInfo:", pretty(info))

	targets, _, err := api.GetTargets(ctx)
	if err != nil {
		log.Println("GetTargets", err)
		return
	}
	log.Println("GetTargets:", pretty(targets))
}
