

# ✨ Daily Devotional Web Scraper

## 🖊 Notes

 - all data is from [lhm.org](https://www.lhm.org/)
 - exmpale url pattern: [lhm.org/dailydevotions/default.asp?date=20211109](https://www.lhm.org/dailydevotions/default.asp?date=20211109)
 - [go-cache](https://github.com/patrickmn/go-cache) is used for in-memory caching. this prevents us from being rate limited or blocked for spamming them with requests
  

## 🤔 Examples

[Check out the examples dir for example requests and example responses.](https://github.com/zepez/devotional/tree/main/scraper/examples)

## 🌟 Scripts

Assumes that your current directory is this project root (devotional/scraper)

- Testing

  `go test -v ./...`

- Building

  `go build`

- Running

  `./scraper`


## 🐳 Docker 

- Building

  `docker build .`

- Running 
  
  `docker run -p 8080:8080 <image_hash>`



  