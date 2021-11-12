# ✨ Devotional Web Scraper

## 🖊 Notes

 - Non-error log format: `[repo/project/type] action | result | timestamp`
  

## 🤔 Examples

 - GET `/devotional/:id` returns single devotional by id
 - GET `/devotionals/:page` returns paginated list of devotionals
 - GET `/devotionals/health` returns http status 200 if it is able to

## 🌟 Scripts

Assumes that your current directory is this project root (devotional/backend)

- Testing

  `go test -v ./...`

- Building

  `go build`

- Running

  ```
  DEVOTIONAL_BACKEND_SCRAPER_FREQ='@every 30s' DEVOTIONAL_BACKEND_SCRAPER_LINK='http://127.0.0.1:8081' DEVOTIONAL_BACKEND_MONGO_URI="mongodb://127.0.0.1:27017/" PORT=8080 ./backend
  ```

## 🐳 Docker 

- Building

  `docker build .`

- Running 
  
  ```
  docker run -e DEVOTIONAL_BACKEND_SCRAPER_FREQ='@every 30s' -e DEVOTIONAL_BACKEND_SCRAPER_LINK='http://host.docker.internal:8081' -e DEVOTIONAL_BACKEND_MONGO_URI="mongodb://host.docker.internal:27017/" -e PORT=8080 64ea5e
  ```

## ⚙️ Environment Variables

| ENV                                    | Default                      |
| -------------------------------------- | ---------------------------- |
| PORT                                   | 8080                         |
| DEVOTIONAL_BACKEND_SCRAPER_FREQ        | @every 30s                   |
| DEVOTIONAL_BACKEND_SCRAPER_LINK        | http://localhost:8081        |
| DEVOTIONAL_BACKEND_MONGO_URI           | mongodb://127.0.0.1:27017/   |
