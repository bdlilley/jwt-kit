package idp

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
)

type Provider struct {
	Issuer        string
	JWKSUrl       string
	KID           string
	RsaPrivateKey *rsa.PrivateKey
	PrivateKey    string
	RsaPublicKey  *rsa.PublicKey
	PublicKey     string
}

var Provider1 *Provider = &Provider{
	Issuer:     "https://bdlilley.github.io/jwt-kit/provider1",
	JWKSUrl:    "https://bdlilley.github.io/jwt-kit/provider1/.well-known/jwks.json",
	KID:        "fOpWQOJ_sSkAzzv5NU1jfy9P6nOwz9udUew26ferYMs",
	PrivateKey: "eyJOIjoxNzYyOTMxODI1NjUwMzI0OTU1MzM0NTM5MDIwMzI0MDk0MzgzNjY1ODYzMDQ1MTY5ODA4NzE2OTg5OTcwMTY2MTY4Njk5MzAwODE4Mzc4NDMyMzM3MjcxNDI5NjcxOTQ3NDM1Mjg4NTY4ODk1NjkzMzgzOTg5MDk3OTgxNzk2MTAyNjIzMzkxNTM1NjQyMTU5MDM5MTMyNDEwNzQ5NzI5NTU0NTgzMzUyMjg1Nzg1Njk4ODY2NzAyMDU2NDA1Nzk4MDE3OTIzNzc2MDQ1MTg2NDYyNjIwNzc2MTMyMTU5NTUyMTMwNjI0Njk3MzQ4NTg3MjkzMjYzNTEzNDk2NDAxNzAxNjgzMTE2MTY0MDcxNjU5NzM3Mzc0NDgyMzE1MDkwMjg5MTYwNzQyNTgzNjU5NDY5NTg1MjcyMTIxOTMxOTAzNTA5MDgxNTAzODQ0MzA5MTYzNzM0MTI2MzIwMjI2Mjk4MzI2MjI4NjA2Mzc4Njk3NzU3MDYwMTUyMDA0ODMyNTM3Mjc1OTc3NzY2MTM4NjM0Mzc1NjkwMzYzNzAzOTEzNjE3NDUzMTA1MTUzMzQxMTI4NTMxMDIwNzkwODQ5OTM0MjAyODg0NDk1NjMxNDcwMjUzOTI5OTA3MzI1MjI4OTE0MTg2MzQ5MDgzNTI0NTQwMjMyNjY3Mjk3NzYxNDkzOTQ1NzYzNjAyMzUzNTc2NTUyOTA1OTE5MjkxNDk4Mjg4MjA3MTU4Mjc2MDY3OTI4MTUzMTI1NjU3NzU5Nzk2MjA3MDE5NDQyMDU0Njg1NzYwMTI2NTM4MTA1MzcwODkxODU4MDk5NywiRSI6NjU1MzcsIkQiOjE2NjQ4MjgyNzU0NzAzMjM4MDkyMzIyNTk5NDQyNzM1ODg5ODY0NTE2MjY3NzM2NjMxMTMzNzg2MTgzMjYzNDM2NTU5NjIyNzY2OTMzMDk3NTEzOTc3NDA2NDY5OTkzNTY0OTczOTg1NjE1NTk5NTAzMTEwNDY5NzAwMTgzNjE1NDgzMDY0Nzg4MTU2NDQ0OTQyOTk4OTAyMTMxODE4NzgzNjA4MjE1NzY3MTE4OTIxNjYzMDA4NzExMTUxMDk4NTkxNTMyOTIwNzE2NDU1ODMyNjA5MzA0NjQyMzg4MTcwNzU1NjU5MzAxMjM3MzAxNjg3MjI3MDc0NjA2NzU1MzEzOTkzMTg5ODEwMjQzNzM5NDE3NjAzMjAyMjY5MDIwNjk2Mzk5OTIwNDA1MDM4Nzg0ODk1NzA1MTkxOTcyOTI0OTEyMTMzNTE0ODA2ODA3NTgyNDU2NDQwNjMzMDYyMTI4NDkwNDE3NDUxNDMzOTc4NzM5MDA1MjkyNTkxNDQ1MzAyODA3ODQxNzIxMTczNzUzMzA0MzU4MjQ0NDU3MjI2OTYzMTQ1OTgzMzA5MjMxOTMyMjc2OTAxOTQyMjYzMTQ5ODc4MTUxMDU1MDU5NDg4NjM0MDQyMDY3OTM2Mjk5NzQzMTU5NzExNTk5NzIxMTEwMjQ0ODQzNzE4NTA1OTMzNTY0ODIyOTUyNjk5MzU5ODkyMjYyODAzODYxMzM1MDI4NTg4Mzg4MDMxNjQ3OTA2NTQyMjM5MjAxODc1ODYxMzA5MDQ4NzU2MjcwMTEzOTE1NDUxMzAzNzc1Mzc1NDkwNDk2MjMyMzMsIlByaW1lcyI6WzE0NDQ5MjU5MDU4MTg2Nzk2OTM5NzEyMTU3MzM3MzU5NjM5Mzg2MzIwNDcwMzc0MzI1OTgzNjMzMTkwMjQyMjMxODcwOTU4NjM0Mzk3NzA2MjkyMTA0NzEzNjA1MzIzNDQ1MjE3NzIyNjE2NzMzNzg0ODk4NTQ2Nzc3Mzc0NjU3NDkwOTYxOTkzMjQyMzIyMDM2OTM1OTk4MTYxODIxMDAxOTY1MzU1NjQyMTg1MjczOTkwMzI1NTM5NzU4MDU2MTczMjU3MTA1MDkyNDcxMDc5NTU2MTg4NTcyNzkxMTg3NjAzNTEzMTY1MTIwMjkwMjc5NjcwNzYwNzI2NjA1NTgxMTIwMjA0MTA2NTg4ODk3MzYzNzk0ODE4MDY4MzI1MzQwNjUzMzE1NDQ2NjUxNjkzMzc4MSwxMjIwMDg0NTg2MDMzODE5MTM4ODU4NDU1NTM2OTcxNzQxMjMxOTk0OTY3ODgyNzczODM0NzIyNzc4NzQ3ODYwMTMwODMzOTMzMDEyMDA2MjY1Njk2NTgyOTEzMDA2MjczNTcyOTkyNDk4ODgwNjEzMTMxMzE0OTIwMzAyNTUzNDQ4OTIwNjMzMzM5NjgyOTIyMTM5MjUxODQyMDE5NjAxMzYxNDI0MjUxOTU2NjA3MTE4OTY5MTE2MjI0NzAxMDYxMjQ1MTQyNzU2NDIxMTA0MjMyMzM0MzU2MjczODQxMjI3NzA3NjI1MDI4NTMwMzQ3NDEyNDQwMzE5MjU4OTI0MDcyMzIyMTcyODMwNDA3NjI5MjY4MTIwOTA1OTM4NTEyOTg4OTQzMjM4MDk4NDk1MTMxMzddLCJQcmVjb21wdXRlZCI6eyJEcCI6MTE0OTE1ODkxODgyNTcyMDE3MzQ3NzA4NDc5OTAyNjI4OTE1NTg4NzIwMTk3MjcwMzM4NzI3NTc4MTgzNTkxODA0NTY1MDcxMDE5NzQxMTAwMDQzODA0NTUwNDg4NTI4NDA1NTc4MjQ0MDgwMzUxMzAzMzA2ODQyNzE5NDI5MDA5MjIyODY5Njc4NDY4MDg4MTM0ODIxMjYwNjkwOTczNjg4ODIyODQzMDkzNTE2NzY5Mjk0MjUzMjg5NDgwNjYwMzY5MzM0Mzk2MDg2MTQ2MzkxOTA0OTc4OTg3NTk4MTkzNDI4MTg3NjE4MDQ3Nzg1MjE0MDAxMTU4MjE0NzY5Njg3MjIyMDY5NzM4MjU4NzcxMTM2MjYwOTY4MjA5OTA0ODQ4NDg3NDM1NzU1NDMyNzQyMTUzLCJEcSI6NzgzMzU0NzM0MTM2OTE1NjYxNzYxODQ1ODU5MzU3MjYyNzMwMzY0Mjg2NzE2OTg5NzUyNjIwMTI0MjY4MDA4Mjc5MDY3MjQ4MDE2ODMzMjMzODY3NTk4Njk3MTI0OTUyMDAyNzUyMzEzNDcwMjI5NjMxMTkyNTk2ODAwMTQxMDUwNDU1MTk5NTAwNTQ0Mzk0NzM1NDIzMzMzNTEzOTA0OTA5OTg5OTI0Mzc2MDAzMDg3NTk5MTAzOTAzMTg0MDIyMDgwMjQ2NTMxMDM4NzYwMTQ5MDQ4MDk1NjI5ODA3NDU1MDE3NDkzMTEwMDA2NTYwNTQ0NzQwNTg1NTI4NDM0NDI4MTExOTQyNjk0MzIzNjk4NDM1MTQzMzc2NzE5NzI3MDE3NTQzNTc5NTQ2MDM0NzMwNTcsIlFpbnYiOjEzNjMxMzY1ODk2NDcyOTc5NzU1MjUxNjc1OTkxMjQzMDE3MTE1NTI0NDcxMjE3MTk3Nzg4NjMwNzY5OTE5MTg2Mjg0NTI4MTc3OTQxMDI3Njk1NDUwMzIxMTMzOTM0OTU1MjIwMDE1OTc2OTA1MzkwNDg3OTQzMDU5NTgzOTI1MDQxMTc2Mjc4MzU0MjI0MTc2Mzc0OTQ1ODQxOTU5MDczMzI3NDQ0MDA3NTkzNzA4OTI4ODQxMTEzODc3Nzc1NDc2MTYxMzk4MjI4OTA2Njk5NjkzMjg4MzU2ODcwODcwOTI1MTQ5ODE5MTkyNjk5MjQ2MzA2NTU4MTgzMjE0OTc2ODY0ODQxNjg5Mzg3NTQ0NTgzOTMxNTAxMTYwODA5MTAwOTQ3NDcxODcyOTY0ODc5NTg4NiwiQ1JUVmFsdWVzIjpbXX19",
	PublicKey:  "eyJOIjoxNzYyOTMxODI1NjUwMzI0OTU1MzM0NTM5MDIwMzI0MDk0MzgzNjY1ODYzMDQ1MTY5ODA4NzE2OTg5OTcwMTY2MTY4Njk5MzAwODE4Mzc4NDMyMzM3MjcxNDI5NjcxOTQ3NDM1Mjg4NTY4ODk1NjkzMzgzOTg5MDk3OTgxNzk2MTAyNjIzMzkxNTM1NjQyMTU5MDM5MTMyNDEwNzQ5NzI5NTU0NTgzMzUyMjg1Nzg1Njk4ODY2NzAyMDU2NDA1Nzk4MDE3OTIzNzc2MDQ1MTg2NDYyNjIwNzc2MTMyMTU5NTUyMTMwNjI0Njk3MzQ4NTg3MjkzMjYzNTEzNDk2NDAxNzAxNjgzMTE2MTY0MDcxNjU5NzM3Mzc0NDgyMzE1MDkwMjg5MTYwNzQyNTgzNjU5NDY5NTg1MjcyMTIxOTMxOTAzNTA5MDgxNTAzODQ0MzA5MTYzNzM0MTI2MzIwMjI2Mjk4MzI2MjI4NjA2Mzc4Njk3NzU3MDYwMTUyMDA0ODMyNTM3Mjc1OTc3NzY2MTM4NjM0Mzc1NjkwMzYzNzAzOTEzNjE3NDUzMTA1MTUzMzQxMTI4NTMxMDIwNzkwODQ5OTM0MjAyODg0NDk1NjMxNDcwMjUzOTI5OTA3MzI1MjI4OTE0MTg2MzQ5MDgzNTI0NTQwMjMyNjY3Mjk3NzYxNDkzOTQ1NzYzNjAyMzUzNTc2NTUyOTA1OTE5MjkxNDk4Mjg4MjA3MTU4Mjc2MDY3OTI4MTUzMTI1NjU3NzU5Nzk2MjA3MDE5NDQyMDU0Njg1NzYwMTI2NTM4MTA1MzcwODkxODU4MDk5NywiRSI6NjU1Mzd9",
}

var Provider2 *Provider = &Provider{
	Issuer:     "https://bdlilley.github.io/jwt-kit/provider2",
	JWKSUrl:    "https://bdlilley.github.io/jwt-kit/provider2/.well-known/jwks.json",
	KID:        "WTwPMNG_FBQTM9VNl_cAnFIoHuZx4JFwrXqVDA48CFE",
	PrivateKey: "eyJOIjoxODc4ODg3NTAwNjA0NzM2ODUyNDc2ODA0MTcyODExNDU0ODgzOTI4MjI1OTY0OTIzODk5NDczOTQyMTgxOTc4ODAyOTc3OTI0NzQ0NzMwOTQ4NjE3Nzc1MTUxNjM3NDEzMjkyNzA1Mzk0MTg5ODIyMTE4NjIwMzk4MTgxNzA1MDU2OTYxOTA4MjYwNzAzMTU5NDQzMDcwMDM3NjM5MDQzNTcyNjI0MDQzODUxMzI3MzkxMTE4Mjk5Mzk1MTg5Mzk5NjkzMzI5NTk1MDUwNzk2Nzk3MDA1NDY2MjAwNTgzNDczODY4MDQwNjI2NzE1Mzc4MDA0MzQwMDk4MTMxMDI3OTEyMDg1OTY2MjI4MjAxNTUzNzEyMTg5MTYzNDU4Nzk0NDAxMDgyOTMwNzc1NDI5ODcwMDA5MTA4NTk4NzEwODU4MTE3MjgxNzYxMjM1Njc0NDY3NjU1MzUyMjE2NTc0NzIyNzU5NzM0OTc1MTg2ODA2NTU3NjgwMDk2ODcxMTQ0MzAxNDIzMzU2MzQ1MTA4NjM1MzY0NTI3ODI1NTczMTUxOTI0MTc1ODUzMTU1OTk5ODM0NzMzNjMwOTc0MTA3MDE5MjA3MDkwNjg1NzY5NTE1ODgzODQxNzA5ODgwODEwNTIyMDQ2OTgxNDk4OTYwNDM0NDQwMjc4OTEzMTkzNjYyMjAyNTQ2NTg2MjI2OTk0NzU3OTg0ODQ3MzE0MTQ1MjUzOTA5Mjc5MTQzNTczMjMxNzAxOTczNDE4MTAxMjg5MTIwMTU1MjM1MjMzMjY4Mjg4NzQxNzYxNjk3NjMyMDk0NDg1ODUzMSwiRSI6NjU1MzcsIkQiOjE4NTUyOTI4MjI1OTA4MjU2NTA3MTkxOTUzNDM2ODg2MTA4ODE3Njk1NTMzMDcyMDE4MTMxODI3MzE4MDU5MTk5NTc1MTg4NTgzODQyNzkzOTM0NTQ4NTMwMDE1NTczNDI2MjgxOTY4NDg3OTY4NjUxMDgwMjE0Mjk4NDc3MzI3NDcyNDU1NzAxNTQwMDY4MDkyODYzMzk1Mzk3Mzc0NDcwMjYyNDE1NDg2NDg3NzU1MTI4OTg1MDk2NTE5NTY3MDg2MDM1OTk5NzI3NTAwMDc4NDE2OTg3NjE2MTEwNjc5Mjk5OTIxNjI5MTczOTQxMjgwODI5NTg3MTIwMzIxNTQ5NzA1MTYzNjgxMzU0MzI0MzI3NzczNzE3MDY0MjY0ODMwNjg2NzM4OTA4MTI4NDExMjE5MDAxNDAwMDc1OTczMTEyOTAyMjMxMTExNjM2OTQ1MjY2MDEwNzUyNDE4Nzg0MDAzNDkzNDY5MzMxMDc2MTM5MjAyMDYyOTcyNTYyOTA2MjUzNzYzNDM0NTM3OTA0NjA0ODA1NDc3MTUxNzU1ODkzMzYzNjE2NDg2MjMxNzM2MDEwNzIxMzU1MTY4OTgyNzg3MTYwNDMyMzIzODA3MDc2OTQ0NTI5MzA3NDY1MDIzNjY3NDcwNTAzNTk3MTIxODU1ODk1NTI1MTE3NTAyNDM0MzU0Nzk1MzA1NDEzMDM3MDU3ODM3NTU2NjM2ODUzOTI4OTcxODkzMjU4ODYzNTMwNjUzMzQ3NjI3MDgxOTYwOTU3OTM2MDc2Nzc3MDg4NDU3MDg4NTAwNTUyOTk3MDA4NjQ5NDczLCJQcmltZXMiOlsxNjI1NjgyMzkzNDU3NzIyMTE3ODY0ODU0MDMxNzcxMDkyMDM3ODk1Mjc4MzYwNTI4OTUxMjIzOTU2MTQ4MDExMzQ2NTM3OTE2NDM0NTA4OTcyODcwMTUxMzc2OTkxOTg1NDA3NDYzODE4Nzg5MTk0NTM0MjY2NTM0NzI2NjcyNzIyOTQyMzQxMjQ4MzIxODU3NTc2OTg5OTAzNzA5MTA0NjI2NDQ3MTQwMDI4MTM0MjEyMzg3ODQwODAyNzk4NjIzMjYxMTM3ODkyMzQ4NzA3MDY2NTA3NTQ3ODg5OTY4OTE4NjQwODUwNTc1ODQwMzk5NDk1MjgxMjMxNTgxOTc1MDEyOTQyNTYxMDk2NDgxMzU3NzYzNzAzNzAwMDkzMjkyNDUwOTYxOTg1OTA5ODUyMjQ5MzEsMTE1NTc1MzEyMTM3NTk4MDI4NzIwOTU2NjE0MDY2MzQ4MTYxNTA5MzU5MDI0ODg2Nzc0NTUzMjg5NTc5MzMwOTA1NTc3MDIxNzI1MTczMTQxODczMjM3OTEyMDcyODQ3NzY1NTIzOTkxNTYwODg1NzMzMDgzNjUyNjkyOTkxNjI4NjM3MTk0ODUzMDM5NTQ5Njg2NjE2MTk5Mjg3Mjc5NTMyNjUzMTY0NDkyMjk1MjkwMDUxOTYwMDA4OTg0OTA5OTYyMDYyODk5ODY0Njk4NDM1MjAwNzQzODA2MTAxMjAxNjM2NzY3MDIzMzYzODc5Njk0MTMzNzgxMDAyMjkyNTgxMDE2MTMwNDIxODI1MjA3MTk0MjczNzczNzE4MzY3NDUzMjE2NDQyNjI4MzU5NDI1NjAxXSwiUHJlY29tcHV0ZWQiOnsiRHAiOjE1NTk4OTgwMzE4MzgzMzMzOTAwMjI5MDgzNjkxMzM4NDk5MjkwOTQxNjkzOTU5NDIxODY4MjE0NjY5OTM3MjQwNTcwNTIzMDQ0MjMyMTI3OTIxNDcwMjMzNTA4MTE2MTc5MDA1NTAyNTc3ODY1NzA5MTg1MjQ2NjYwMDk4Mzg5MzM0NjA5MzI0NDEyMjczOTg0NzMwNjExNzI5OTc2NTA4NjAzNDAzOTM5ODYxMzMzMjg0NDA1NjU5MjI4MjIwOTE3MDA1MTUwNzMzMjI2Nzk0NjE2MzczNTIxNjg0MDQwMzgxNTc1Mjc2MzI2NjEyOTg1Mzc2MzE1NzA2ODU3NTc2NDM2MDQyNjg2NTA1NjczMTU5MTI2OTgyODAwNjExMzYzOTI4NTgxNDg1ODY5MjEyNjEyMywiRHEiOjU3MDg0ODk2Mzc3NTI3OTMzNjgxNjk2ODM2ODYwNTE2ODA3MTE3NDc0ODg2NDg1MjY2MjIwNDU1MzcxNTE0NDMzMjc0NzYzNzcwNzUzMjMyNTYyMzE5MTY2NDgzMDI2MTA5OTgzODUwNDQ4MjMzMzgyMzYyOTA3MDI0NjE0MTcyNDM2NzMwMzU2Nzg5MTQ1NDE5NDY5NDA0NjI1MzE0NTMxODIxNDU4NjM1ODE3OTEyNjEzOTcyOTU0MDY5MzI3NzkxNTA2NzIyNDQxMDY4MjI2MzA2NDg0NTM1NTA2NTk3NzUzNjY4MTM0NzM3NDU4MDExNDkxMzc4Nzc5MDc0NTgxNDk1ODI4OTQ3ODM4MzUwMTk3MjcyNDExODQ3NzEyODEwNDgyODc2MDUyOTE2NTkwNzMsIlFpbnYiOjg2NDE1NjA0ODcwMTIwODM2OTc5Mzk3Mjk5NzU1NDcwNzE3NDg5NTY0NjExODY4MzA4NjQwMTc0NjIyMDY5NjcyNjM1NTgyMjY0MzE3MTgwMTgwNjI0NDU3MzM3OTc4NDgxODY5OTc0MDkxMTM0ODA0NjE4NTUyNDQwNDI1NDM5MjgxMDk5NjY1MTI0NDA5Mjk1MjkwNDc3NDA5Njc2MzI4OTUzNTE0MTYwNzk3MzExNTA5NTI4NDIzMjgzMjA5Mzk5MDgwMzUzMzM5MzgxMDI5NTcyNjY3NDIzMTQ2NTAxNjM5NzU5NjQ0OTA2NTEzNDkwNzU5OTk0OTY3MzI4ODgzNjc4OTczMDM3MDUwNTA3MTE0Nzg3NTEwMDM1MDc3NDc3NDMwOTY4NjY4OTExMzg1NDExLCJDUlRWYWx1ZXMiOltdfX0=",
	PublicKey:  "eyJOIjoxODc4ODg3NTAwNjA0NzM2ODUyNDc2ODA0MTcyODExNDU0ODgzOTI4MjI1OTY0OTIzODk5NDczOTQyMTgxOTc4ODAyOTc3OTI0NzQ0NzMwOTQ4NjE3Nzc1MTUxNjM3NDEzMjkyNzA1Mzk0MTg5ODIyMTE4NjIwMzk4MTgxNzA1MDU2OTYxOTA4MjYwNzAzMTU5NDQzMDcwMDM3NjM5MDQzNTcyNjI0MDQzODUxMzI3MzkxMTE4Mjk5Mzk1MTg5Mzk5NjkzMzI5NTk1MDUwNzk2Nzk3MDA1NDY2MjAwNTgzNDczODY4MDQwNjI2NzE1Mzc4MDA0MzQwMDk4MTMxMDI3OTEyMDg1OTY2MjI4MjAxNTUzNzEyMTg5MTYzNDU4Nzk0NDAxMDgyOTMwNzc1NDI5ODcwMDA5MTA4NTk4NzEwODU4MTE3MjgxNzYxMjM1Njc0NDY3NjU1MzUyMjE2NTc0NzIyNzU5NzM0OTc1MTg2ODA2NTU3NjgwMDk2ODcxMTQ0MzAxNDIzMzU2MzQ1MTA4NjM1MzY0NTI3ODI1NTczMTUxOTI0MTc1ODUzMTU1OTk5ODM0NzMzNjMwOTc0MTA3MDE5MjA3MDkwNjg1NzY5NTE1ODgzODQxNzA5ODgwODEwNTIyMDQ2OTgxNDk4OTYwNDM0NDQwMjc4OTEzMTkzNjYyMjAyNTQ2NTg2MjI2OTk0NzU3OTg0ODQ3MzE0MTQ1MjUzOTA5Mjc5MTQzNTczMjMxNzAxOTczNDE4MTAxMjg5MTIwMTU1MjM1MjMzMjY4Mjg4NzQxNzYxNjk3NjMyMDk0NDg1ODUzMSwiRSI6NjU1Mzd9",
}

func init() {
	Provider1.SetRsaKeys()
	Provider2.SetRsaKeys()
}

func (p *Provider) SetRsaKeys() {
	p.SetPrivateRsaKey()
	p.SetPublicRsaKey()
}

func (p *Provider) SetPrivateRsaKey() {
	b, err := base64.StdEncoding.DecodeString(p.PrivateKey)
	if err != nil {
		panic(err)
	}
	var k rsa.PrivateKey
	if err = json.Unmarshal(b, &k); err != nil {
		panic(err)
	}
	p.RsaPrivateKey = &k
}

func (p *Provider) SetPublicRsaKey() {
	b, err := base64.StdEncoding.DecodeString(p.PublicKey)
	if err != nil {
		panic(err)
	}
	var k rsa.PublicKey
	if err = json.Unmarshal(b, &k); err != nil {
		panic(err)
	}
	p.RsaPublicKey = &k
}

/*
my very hacky method of embedding keys

* use https://mkjwk.org/ to generate key data, then add pems to files

	{
		b, _ := ioutil.ReadFile("certs/idp2/pri.pem")
		k, _ := jwt.ParseRSAPrivateKeyFromPEM(b)
		bb, _ := json.Marshal(k)
		fmt.Println(base64.StdEncoding.EncodeToString(bb))
	}
	{
		b, _ := ioutil.ReadFile("certs/idp2/pub.pem")
		k, _ := jwt.ParseRSAPublicKeyFromPEM(b)
		bb, _ := json.Marshal(k)
		fmt.Println(base64.StdEncoding.EncodeToString(bb))
	}
*/
