package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"context"
	"github.com/alohihq/signplus-go"
	"github.com/alohihq/signplus-go/signplus1"
)

func main() {
	loadEnv()

	config := signplus.NewConfig()
	client := signplus.NewSignplus(config)

	envelopeLegalityLevel := signplus1.EnvelopeLegalityLevelSes

	request := signplus1.CreateEnvelopeRequest{
		Name:          "name",
		LegalityLevel: envelopeLegalityLevel,
		ExpiresAt:     signplus.Ptr(int64(6)),
		Comment:       signplus.Ptr("comment"),
		Sandbox:       signplus.Ptr(true),
	}

	response, err := client.Signplus1.CreateEnvelope(context.Background(), request)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", response)
}

func loadEnv() error {
	file, err := os.Open(".env")
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
