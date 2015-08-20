# hubit copartying rest go

## Doporučené zdroje
* https://golang.org/
* https://github.com/go-martini/martini
* http://0value.com/build-a-restful-API-with-Martini

## Postup instalace
1. nainstalujte si go
  * standartní cesty vaše oblíbeného OS
  * https://golang.org/dl/
2. nastavte si GOPATH - https://golang.org/doc/install#testing
3. stáhněte si tento repozitář (to už možná máte) ideálně do nastavené GOPATH a přepěnte se do složky copartying_rest_go
4. nainstalujte si tyto závislosti
  * ```go get github.com/go-martini/martini```
  * ```go get github.com/martini-contrib/render```
5. spusťte server ```go run server.go```
5. jestli jsem na něco nezapomněl a měli jste alespoň trochu štěstí, můžete zadat to prohlížeče adresu http://localhost:3000
6. správná odpověď je ```{"hello":"world"}```

   Pokud jste dostali cokoliv jiného, zkuste kontaktovat někoho kdo tomu rozumí
