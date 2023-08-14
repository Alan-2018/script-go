package main

import "log"

// var v struct{} = isyntax.TestISyntaxFuncsInit2()

func main() {

	/*
		golang syntax
	*/

	// isyntax.TestISyntaxFuncsPrints()

	// isyntax.TestISyntaxConditionalStatements()

	// isyntax.TestISyntaxTypesAndDefaults()

	// isyntax.TestISyntaxInitializationsFuncs()

	// isyntax.TestISyntaxTypesInterface()

	// isyntax.TestISyntaxStringsFuncs()

	// isyntax.TestISyntaxIOFuncs()

	// isyntax.TestISyntaxTimesFuncs()

	// isyntax.TestISyntaxErrors()

	// isyntax.TestISyntaxJsons()

	// isyntax.TestISyntaxStructs()

	// isyntax.TestISyntaxUnsafeFuncs()

	/*
		math
	*/

	// isyntax.TestISyntaxMathRandFuncs()

	// isyntax.TestISyntaxMathFuncs()

	/*
		algorithms
	*/

	// isyntax.TestMapsSort()
	// isyntax.TestStringSliceSort()

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

	// iextensions.TestIExtensionsJwtFuncs(
	// 	"eyJhbGciOiJSUzI1NiIsImtpZCI6ImxzcC1jZXJ0IiwidHlwIjoiSldUIn0.eyJvd25lciI6Imxha2Utc3VwZXJpb3ItdXNlcnMiLCJuYW1lIjoiY29uZ2Nhbi10ZXN0LTkwMDAwIiwiY3JlYXRlZFRpbWUiOiIyMDIzLTA0LTE4VDExOjE1OjA3WiIsInVwZGF0ZWRUaW1lIjoiIiwiaWQiOiI5YmQyMGY2NC02YWUxLTRmZjctYmI5Ni03N2NkMTRjOTQyYWMiLCJ0eXBlIjoibm9ybWFsLXVzZXIiLCJwYXNzd29yZCI6IiIsInBhc3N3b3JkU2FsdCI6IiIsImRpc3BsYXlOYW1lIjoiZ2l2ZW5fbmFtZSIsImZpcnN0TmFtZSI6IiIsImxhc3ROYW1lIjoiIiwiYXZhdGFyIjoiIiwicGVybWFuZW50QXZhdGFyIjoiIiwiZW1haWwiOiIxMzA2NDM2OTI1QHFxLmNvbSIsImVtYWlsVmVyaWZpZWQiOmZhbHNlLCJwaG9uZSI6IjEzNzE2ODUzMjg5IiwibG9jYXRpb24iOiIiLCJhZGRyZXNzIjpbXSwiYWZmaWxpYXRpb24iOiIiLCJ0aXRsZSI6IiIsImlkQ2FyZFR5cGUiOiIiLCJpZENhcmQiOiIiLCJob21lcGFnZSI6IiIsImJpbyI6IiIsInJlZ2lvbiI6IiIsImxhbmd1YWdlIjoiIiwiZ2VuZGVyIjoiIiwiYmlydGhkYXkiOiIiLCJlZHVjYXRpb24iOiIiLCJzY29yZSI6MjAwMCwia2FybWEiOjAsInJhbmtpbmciOjE2OSwiaXNEZWZhdWx0QXZhdGFyIjpmYWxzZSwiaXNPbmxpbmUiOmZhbHNlLCJpc0FkbWluIjpmYWxzZSwiaXNHbG9iYWxBZG1pbiI6ZmFsc2UsImlzRm9yYmlkZGVuIjpmYWxzZSwiaXNEZWxldGVkIjpmYWxzZSwic2lnbnVwQXBwbGljYXRpb24iOiJsYWtlLXN1cGVyaW9yLXNlcnZpY2UiLCJoYXNoIjoiIiwicHJlSGFzaCI6IiIsImNyZWF0ZWRJcCI6IiIsImxhc3RTaWduaW5UaW1lIjoiIiwibGFzdFNpZ25pbklwIjoiIiwiZ2l0aHViIjoiIiwiZ29vZ2xlIjoiIiwicXEiOiIiLCJ3ZWNoYXQiOiIiLCJ1bmlvbklkIjoiIiwiZmFjZWJvb2siOiIiLCJkaW5ndGFsayI6IiIsIndlaWJvIjoiIiwiZ2l0ZWUiOiIiLCJsaW5rZWRpbiI6IiIsIndlY29tIjoiIiwibGFyayI6IiIsImdpdGxhYiI6IiIsImFkZnMiOiIiLCJiYWlkdSI6IiIsImFsaXBheSI6IiIsImNhc2Rvb3IiOiIiLCJpbmZvZmxvdyI6IiIsImFwcGxlIjoiIiwiYXp1cmVhZCI6IiIsInNsYWNrIjoiIiwic3RlYW0iOiIiLCJiaWxpYmlsaSI6IiIsIm9rdGEiOiIiLCJkb3V5aW4iOiIiLCJjdXN0b20iOiIiLCJsZGFwIjoiIiwicHJvcGVydGllcyI6e30sImlzcyI6Imh0dHBzOi8vc3NvLnJkZS5iYXNlYml0Lm1lIiwic3ViIjoiOWJkMjBmNjQtNmFlMS00ZmY3LWJiOTYtNzdjZDE0Yzk0MmFjIiwiYXVkIjpbIjQ1Yjk2YzIzMDU1MjcyZmZjNGJhIl0sImV4cCI6MTY4MjI0MDY1NywibmJmIjoxNjgyMjI5ODU3LCJpYXQiOjE2ODIyMjk4NTd9.j3_fOzG5oCiVlaWVDDIHlV_vQAbHSU_YmVWMI-PtnWItqeAOfIJE0tM5BalnkFoD5Yf_M52-Qrx818CjPrbI5vQr4QmFdokGZ7JKnLvndeKIsgDM7DZ7Skjv8l2oPgPSnqw8na_gmqxCU-iPIrw2lklnWwtYGRi6-90bA8xKZUlSWHp4Ko0ElDpyP9gym11UqhKzc3Vv4TT-pQuvkwSPYIX00Vi5IRp1d5NqtMpQIMhyNrALAylvKCrbAT_5msclBQc1HXZKwHZJVDhpI9iZyqj950aCoo7BvidoUbWR-9VsX8aZ_pdxPNLMT1JigmK6u6txkkeKV0HUnf5ZGR-C7QHmAKVgFdVx3T5ybzdVVQzU4wWHpTnq8wbZW7bBX3A4F3kpvkzSyUQ9A_F0AqxtvQDkdNC-cc-afSd7w_fJnqf64XUkvoP6IT8kxlzCJSpbKNGDwLX-14gUyy3i-BBmIUmrW0HT0OvGEujlZyOEGTNiadVPCajEUNyU0ENwPndre_W4K9Feh3WkulEYJPqYcFpWG-Ds7RUWOIrAlgII4BYmX9CdgDJHF4dp4ocN6O4U8WCREqITG5Qx5xNByt7r7tji4Bq28_BzZpdWW8rpZ5GwunPZxSxqRVN0UjOkYfmFWjgfOUmw9n4M-oXb3TDclwwdUyv4PZoG0vuCfQKhDhI",
	// )

	/*
		bce sms
	*/
	// ? package.package
	// iextensions.baidubce.TestBaiduBceSms()
	// baidubce.TestBaiduBceSms()

	/*
		mysql
	*/
	// mysql.TestISqlsMysqls()

	// mysql.TestISqlsMysqlSqlx()

	// mysql.TestISqlsMysqlSqlxConcurrent()

	/*
		?
	*/

	/*
		temp
	*/
	Temp()

	// s := strings.HasPrefix("", "Bearer ")
	// log.Println(s)

	// s2 := strings.TrimLeft("Bearer xxx", "Bearer ")
	// log.Println(s2)

	// tokenStr := "Bearer eyJhbGciOiJSUzI1NiIsImtpZCI6ImxzcC1jZXJ0IiwidHlwIjoiSldUIn0.eyJvd25lciI6Imxha2Utc3VwZXJpb3ItdXNlcnMiLCJuYW1lIjoiY29uZ2Nhbi10ZXN0LTkwMDAwIiwiY3JlYXRlZFRpbWUiOiIyMDIzLTA0LTE4VDExOjE1OjA3WiIsInVwZGF0ZWRUaW1lIjoiIiwiaWQiOiI5YmQyMGY2NC02YWUxLTRmZjctYmI5Ni03N2NkMTRjOTQyYWMiLCJ0eXBlIjoibm9ybWFsLXVzZXIiLCJwYXNzd29yZCI6IiIsInBhc3N3b3JkU2FsdCI6IiIsImRpc3BsYXlOYW1lIjoiZ2l2ZW5fbmFtZSIsImZpcnN0TmFtZSI6IiIsImxhc3ROYW1lIjoiIiwiYXZhdGFyIjoiIiwicGVybWFuZW50QXZhdGFyIjoiIiwiZW1haWwiOiIxMzA2NDM2OTI1QHFxLmNvbSIsImVtYWlsVmVyaWZpZWQiOmZhbHNlLCJwaG9uZSI6IjEzNzE2ODUzMjg5IiwibG9jYXRpb24iOiIiLCJhZGRyZXNzIjpbXSwiYWZmaWxpYXRpb24iOiIiLCJ0aXRsZSI6IiIsImlkQ2FyZFR5cGUiOiIiLCJpZENhcmQiOiIiLCJob21lcGFnZSI6IiIsImJpbyI6IiIsInJlZ2lvbiI6IiIsImxhbmd1YWdlIjoiIiwiZ2VuZGVyIjoiIiwiYmlydGhkYXkiOiIiLCJlZHVjYXRpb24iOiIiLCJzY29yZSI6MjAwMCwia2FybWEiOjAsInJhbmtpbmciOjE2OSwiaXNEZWZhdWx0QXZhdGFyIjpmYWxzZSwiaXNPbmxpbmUiOmZhbHNlLCJpc0FkbWluIjpmYWxzZSwiaXNHbG9iYWxBZG1pbiI6ZmFsc2UsImlzRm9yYmlkZGVuIjpmYWxzZSwiaXNEZWxldGVkIjpmYWxzZSwic2lnbnVwQXBwbGljYXRpb24iOiJsYWtlLXN1cGVyaW9yLXNlcnZpY2UiLCJoYXNoIjoiIiwicHJlSGFzaCI6IiIsImNyZWF0ZWRJcCI6IiIsImxhc3RTaWduaW5UaW1lIjoiIiwibGFzdFNpZ25pbklwIjoiIiwiZ2l0aHViIjoiIiwiZ29vZ2xlIjoiIiwicXEiOiIiLCJ3ZWNoYXQiOiIiLCJ1bmlvbklkIjoiIiwiZmFjZWJvb2siOiIiLCJkaW5ndGFsayI6IiIsIndlaWJvIjoiIiwiZ2l0ZWUiOiIiLCJsaW5rZWRpbiI6IiIsIndlY29tIjoiIiwibGFyayI6IiIsImdpdGxhYiI6IiIsImFkZnMiOiIiLCJiYWlkdSI6IiIsImFsaXBheSI6IiIsImNhc2Rvb3IiOiIiLCJpbmZvZmxvdyI6IiIsImFwcGxlIjoiIiwiYXp1cmVhZCI6IiIsInNsYWNrIjoiIiwic3RlYW0iOiIiLCJiaWxpYmlsaSI6IiIsIm9rdGEiOiIiLCJkb3V5aW4iOiIiLCJjdXN0b20iOiIiLCJsZGFwIjoiIiwicHJvcGVydGllcyI6e30sImlzcyI6Imh0dHBzOi8vc3NvLnJkZS5iYXNlYml0Lm1lIiwic3ViIjoiOWJkMjBmNjQtNmFlMS00ZmY3LWJiOTYtNzdjZDE0Yzk0MmFjIiwiYXVkIjpbIjQ1Yjk2YzIzMDU1MjcyZmZjNGJhIl0sImV4cCI6MTY4MjI0MDY1NywibmJmIjoxNjgyMjI5ODU3LCJpYXQiOjE2ODIyMjk4NTd9.j3_fOzG5oCiVlaWVDDIHlV_vQAbHSU_YmVWMI-PtnWItqeAOfIJE0tM5BalnkFoD5Yf_M52-Qrx818CjPrbI5vQr4QmFdokGZ7JKnLvndeKIsgDM7DZ7Skjv8l2oPgPSnqw8na_gmqxCU-iPIrw2lklnWwtYGRi6-90bA8xKZUlSWHp4Ko0ElDpyP9gym11UqhKzc3Vv4TT-pQuvkwSPYIX00Vi5IRp1d5NqtMpQIMhyNrALAylvKCrbAT_5msclBQc1HXZKwHZJVDhpI9iZyqj950aCoo7BvidoUbWR-9VsX8aZ_pdxPNLMT1JigmK6u6txkkeKV0HUnf5ZGR-C7QHmAKVgFdVx3T5ybzdVVQzU4wWHpTnq8wbZW7bBX3A4F3kpvkzSyUQ9A_F0AqxtvQDkdNC-cc-afSd7w_fJnqf64XUkvoP6IT8kxlzCJSpbKNGDwLX-14gUyy3i-BBmIUmrW0HT0OvGEujlZyOEGTNiadVPCajEUNyU0ENwPndre_W4K9Feh3WkulEYJPqYcFpWG-Ds7RUWOIrAlgII4BYmX9CdgDJHF4dp4ocN6O4U8WCREqITG5Qx5xNByt7r7tji4Bq28_BzZpdWW8rpZ5GwunPZxSxqRVN0UjOkYfmFWjgfOUmw9n4M-oXb3TDclwwdUyv4PZoG0vuCfQKhDhI"
	// if strings.HasPrefix(tokenStr, "Bearer ") {
	// 	t := strings.TrimLeft(tokenStr, "Bearer ")
	// 	t2 := strings.TrimPrefix(tokenStr, "Bearer ")

	// 	log.Println(t == t2)

	// 	log.Println(t)
	// 	log.Println(t2)
	// }

	// log.Println(tokenStr)

}

type jarray = []interface{}
type jmap = map[string]interface{}

func Temp() {
	j := make(jmap)

	j["x"] = int(9)

	log.Println(j["x"].(int))

	// jarr := make(jarray, 1)
	// jarr = append(jarr, "id1")

	// log.Println(jarr)

	// s := uuid.NewString()[:8]
	// n, flag := new(big.Int).SetString(s, 16)
	// log.Println(s, n, n.String(), len(n.String()), flag)

	// code := "86758181"
	// for len(code) < 10 {
	// 	code = "0" + code
	// }

	// log.Println(code)

	// var phones []string
	// var phonesExt []string

	// i := "+85251049640"

	// if strings.HasPrefix(i, "+") && !strings.HasPrefix(i, "+86") {
	// 	phonesExt = append(phonesExt, i)
	// } else {
	// 	phones = append(phones, i)
	// }

	// log.Println(phonesExt)

}
