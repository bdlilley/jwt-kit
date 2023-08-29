package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bdlilley/jwt-kit/internal/idp"
	"github.com/brianvoe/gofakeit"
	"github.com/golang-jwt/jwt"
	"github.com/spf13/cobra"
)

type Config struct {
	Claims      []string
	Scopes      []string
	Audiences   []string
	Exp         string
	Sub         string
	Provider    string
	claimsMap   jwt.MapClaims
	expDuration time.Duration
	OutputJSON  bool
}

var config Config = Config{claimsMap: make(map[string]interface{})}

var rootCmd = &cobra.Command{
	Use:   "jwt-kit",
	Short: "jwt-kit - a simple CLI to generate JWTs using development IDPs",
	Long: fmt.Sprintf(`Jwt-kit contains embedded keypairs to sign JWTs as an IDP would.

Provider1:
  Public JWKS url: %s
  Issuer name: %s

Provider2:
  Public JWKS url: %s
  Issuer name: %s
`, idp.Provider1.JWKSUrl, idp.Provider1.Issuer, idp.Provider2.JWKSUrl, idp.Provider2.Issuer),
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var errs []string

		config.Provider = strings.ToLower(config.Provider)
		if config.Provider != "provider1" && config.Provider != "provider2" {
			errs = append(errs, fmt.Sprintf("invalid provider '%s', must be one of provider1,provider2", config.Provider))
		}

		for _, c := range config.Claims {
			parts := strings.Split(c, "=")
			if len(parts) != 2 {
				errs = append(errs, fmt.Sprintf("arg '%s' must be in format key=value", c))
				continue
			}
			config.claimsMap[parts[0]] = parts[1]
		}

		var err error
		config.expDuration, err = time.ParseDuration(config.Exp)
		if err != nil {
			errs = append(errs, fmt.Sprintf("invalid time duration '%s': %s", config.Exp, err))
		}

		if len(errs) > 0 {
			return fmt.Errorf("validation errors: %s", strings.Join(errs, "; "))
		}
		gofakeit.Seed(time.Now().UnixNano())
		config.claimsMap["beer_of_the_day"] = gofakeit.BeerName()
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		var provider *idp.Provider
		switch config.Provider {
		case "provider1":
			provider = idp.Provider1
		case "provider2":
			provider = idp.Provider2
		}

		theJwt := getUnsignedJwt(provider)

		tokenString, err := theJwt.SignedString(provider.RsaPrivateKey)
		if err != nil {
			return err
		}

		if config.OutputJSON {
			theToken, err := parseSignedJwtString(provider, tokenString)
			if err != nil {
				return err
			}
			formatted, err := json.MarshalIndent(theToken, "", "  ")
			if err != nil {
				return err
			}
			fmt.Printf("\n%s\n", formatted)
		} else {
			fmt.Println(tokenString)
		}

		return nil
	},
}

func Execute() {
	rootCmd.Flags().StringArrayVarP(&config.Claims, "claims", "c", []string{}, "add jwt claims")
	rootCmd.Flags().StringArrayVarP(&config.Scopes, "scopes", "s", []string{}, "add jwt scopes")
	rootCmd.Flags().StringArrayVarP(&config.Audiences, "audiences", "a", []string{"https://fake-resource.solo.io"}, "jwt audience")
	rootCmd.Flags().StringVarP(&config.Exp, "expires-in", "e", "8766h", "expires duration (uses https://pkg.go.dev/time#ParseDuration)")
	rootCmd.Flags().StringVarP(&config.Sub, "subject", "u", "glooey@solo.io", "jwt subject")
	rootCmd.Flags().BoolVarP(&config.OutputJSON, "json", "j", false, "output full token signed details as JSON")
	rootCmd.Flags().StringVarP(&config.Provider, "provider", "p", "provider1", "provider to use (provider1, provider2)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func getUnsignedJwt(provider *idp.Provider) *jwt.Token {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Header["kid"] = provider.KID

	now := time.Now().UTC()

	config.claimsMap["exp"] = now.Add(config.expDuration).Unix()
	config.claimsMap["iss"] = provider.Issuer
	config.claimsMap["aud"] = config.Audiences
	config.claimsMap["sub"] = config.Sub
	config.claimsMap["scope"] = strings.Join(config.Scopes, " ")
	token.Claims = config.claimsMap

	return token
}

func parseSignedJwtString(provider *idp.Provider, token string) (*jwt.Token, error) {
	t, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return provider.RsaPublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	return t, nil
}
