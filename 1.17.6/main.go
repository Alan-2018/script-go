package main

import "github.com/flower/script-go/isyntax"

// var v struct{} = isyntax.TestISyntaxFuncsInit2()

func main() {

	/*
		golang syntax
	*/

	// isyntax.TestISyntaxFuncsPrints()

	// isyntax.TestISyntaxConditionalStatements()

	// isyntax.TestISyntaxTypesAndDefaults()

	// isyntax.TestISyntaxInitializationsFuncs()

	isyntax.TestISyntaxTypesInterface()

	// isyntax.TestISyntaxStringsFuncs()

	// isyntax.TestISyntaxIOFuncs()

	// isyntax.TestISyntaxTimesFuncs()

	// isyntax.TestISyntaxErrors()

	// isyntax.TestISyntaxJsons()

	// isyntax.TestISyntaxStructs()

	// isyntax.TestISyntaxUnsafeFuncs()

	/*
		algorithms
	*/

	// isyntax.TestMapsSort()

	/*
		utils
	*/

	// iutils.TestIUtils()

	// iutils.TestIUtilsLog()

	// iutils.GenerateKeysByEd25519()

	/*
		HTTP
	*/

	// ihttp.TestIHttpStatusCode()
	// ihttp.TestIHttpFuncs()

	/*
		redis
	*/

	// iredis.TestIRedisConcurrent()

	/*
		crypto
	*/

	// icrypto.TestICryptoBase64Funcs(
	// 	// []byte("bc3af9c93ef6cf2e5e18:f955ae9d3ee73e76a843fe7278a1998db68d94b2"),
	// 	[]byte("45b96c23055272ffc4ba:bf75a9dec90e68a4959534354ec8bb00422e20e9"),
	// )

	// icrypto.TestICryptoHex()

	// icrypto.TestICryptoUuid()

	/*
		jwt
	*/

	// iextensions.TestIExtensionsJwtFuncs(
	// 	"eyJhbGciOiJSUzI1NiIsImtpZCI6ImxzcC1jZXJ0IiwidHlwIjoiSldUIn0.eyJvd25lciI6Imxha2Utc3VwZXJpb3ItdXNlcnMiLCJuYW1lIjoiZmxvd2VyMiIsImNyZWF0ZWRUaW1lIjoiMjAyMy0wMy0xNVQwOToyMDo1MVoiLCJ1cGRhdGVkVGltZSI6IiIsImlkIjoiZjJlOGMxMjEtN2U0MC00ZmQ0LWFhMzgtMTQxZDI0YzFmMDEzIiwidHlwZSI6Im5vcm1hbC11c2VyIiwicGFzc3dvcmQiOiIiLCJwYXNzd29yZFNhbHQiOiIiLCJkaXNwbGF5TmFtZSI6ImdpdmVuX25hbWUiLCJmaXJzdE5hbWUiOiIiLCJsYXN0TmFtZSI6IiIsImF2YXRhciI6IiIsInBlcm1hbmVudEF2YXRhciI6IiIsImVtYWlsIjoiZmxvd2VyLmNvbmcyQGJhc2ViaXQuYWkiLCJlbWFpbFZlcmlmaWVkIjpmYWxzZSwicGhvbmUiOiIiLCJsb2NhdGlvbiI6IiIsImFkZHJlc3MiOltdLCJhZmZpbGlhdGlvbiI6IiIsInRpdGxlIjoiIiwiaWRDYXJkVHlwZSI6IiIsImlkQ2FyZCI6IiIsImhvbWVwYWdlIjoiIiwiYmlvIjoiIiwicmVnaW9uIjoiIiwibGFuZ3VhZ2UiOiIiLCJnZW5kZXIiOiIiLCJiaXJ0aGRheSI6IiIsImVkdWNhdGlvbiI6IiIsInNjb3JlIjoyMDAwLCJrYXJtYSI6MCwicmFua2luZyI6MjUsImlzRGVmYXVsdEF2YXRhciI6ZmFsc2UsImlzT25saW5lIjpmYWxzZSwiaXNBZG1pbiI6ZmFsc2UsImlzR2xvYmFsQWRtaW4iOmZhbHNlLCJpc0ZvcmJpZGRlbiI6ZmFsc2UsImlzRGVsZXRlZCI6ZmFsc2UsInNpZ251cEFwcGxpY2F0aW9uIjoibGFrZS1zdXBlcmlvci1zZXJ2aWNlIiwiaGFzaCI6IiIsInByZUhhc2giOiIiLCJjcmVhdGVkSXAiOiIiLCJsYXN0U2lnbmluVGltZSI6IiIsImxhc3RTaWduaW5JcCI6IiIsImdpdGh1YiI6IiIsImdvb2dsZSI6IiIsInFxIjoiIiwid2VjaGF0IjoiIiwidW5pb25JZCI6IiIsImZhY2Vib29rIjoiIiwiZGluZ3RhbGsiOiIiLCJ3ZWlibyI6IiIsImdpdGVlIjoiIiwibGlua2VkaW4iOiIiLCJ3ZWNvbSI6IiIsImxhcmsiOiIiLCJnaXRsYWIiOiIiLCJhZGZzIjoiIiwiYmFpZHUiOiIiLCJhbGlwYXkiOiIiLCJjYXNkb29yIjoiIiwiaW5mb2Zsb3ciOiIiLCJhcHBsZSI6IiIsImF6dXJlYWQiOiIiLCJzbGFjayI6IiIsInN0ZWFtIjoiIiwiYmlsaWJpbGkiOiIiLCJva3RhIjoiIiwiZG91eWluIjoiIiwiY3VzdG9tIjoiIiwibGRhcCI6IiIsInByb3BlcnRpZXMiOnt9LCJpc3MiOiJodHRwczovL3Nzby5hcmtzaG9wLnBhcml0eWJpdC5haSIsInN1YiI6ImYyZThjMTIxLTdlNDAtNGZkNC1hYTM4LTE0MWQyNGMxZjAxMyIsImF1ZCI6WyJiYzNhZjljOTNlZjZjZjJlNWUxOCJdLCJleHAiOjE2Nzg4ODY5MjcsIm5iZiI6MTY3ODg3NjEyNywiaWF0IjoxNjc4ODc2MTI3fQ.I7k7t23A0JTGIznMmAsXK03w0dj_2sdC4NSej1WeGd6Dcv3C_MJG9a1UFY1QlSAe_QtUMPi93tkstwGGRYCIr4uPfutYgjfKDFQB294N2tc1A9yAHbjbQYYLzwAvCwrY6d-u4iALzXnalsQ3RawQITtI7otell7BJecH-28AYNIwDteKyGOM9V1czjp40m0e80rbnVBRZtrZuX5uxzEoZKy_nSKagbLQyhg1ybGVOAyDvOOGP2KIr1RiHexWn1TvRT8qclk8NBfJchhBOknwQ6ZSvPwHp6NS45MKym738oRwsRe1p1bSViLeIdt0tJ4hDYFwer6SFhInFeRPoSnD7tB7P9r-HKMk-kkIgDiysGW383PIkU2OnUx0Twsd1Rpg9_EHGt8J5wvyJ9O7oVx7GBIC3lv1xfhzUf5ZeYZCgcofaWYFCLdEFzhJ-2iWwJcMSyEGVOtve3pQk32dYjD-FyT1WPdCKrQU2jSAk0Y31lKcFWXT_Mh1rZLB9a6YY_c6YuDFJSl5bwRJSlLzGC7aCh2us9312Pavdr4-Cuw_NvgzVlAxt2MofUkCToiFNzxX8F7maWNyk3M1qtpGkXegfViSxL0RF4cn5eU1SITSC8OWybIDiQL1udSMBr-Njrd7onGOeWaSMixCvATW12U1zpv2dgDS9fBreWm6YGtCBI8",
	// )

	/*
		mysql
	*/
	// mysql.TestISqlsMysqls()

	// mysql.TestISqlsMysqlSqlx()

	// mysql.TestISqlsMysqlSqlxConcurrent()

	/*
		?
	*/

}
