


run command

`
DEVOTIONAL_BACKEND_SCRAPER_FREQ='@every 30s' DEVOTIONAL_BACKEND_SCRAPER_LINK='http://127.0.0.1:8081' DEVOTIONAL_BACKEND_MONGO_URI="mongodb://127.0.0.1:27017/" PORT=8080 ./backend
`


docker command 

`
docker run -e DEVOTIONAL_BACKEND_SCRAPER_FREQ='@every 30s' -e DEVOTIONAL_BACKEND_SCRAPER_LINK='http://host.docker.internal:8081' -e DEVOTIONAL_BACKEND_MONGO_URI="mongodb://host.docker.internal:27017/" -e PORT=8080 64ea5e
`


non error log format

`
[repo/project/type] action | result | timestamp
`



## ⚙️ Environment Variables

| ENV                                    | Default                      |
| -------------------------------------- | ---------------------------- |
| PORT                                   | 8080                         |
| DEVOTIONAL_BACKEND_SCRAPER_FREQ        | @every 30s                   |
| DEVOTIONAL_BACKEND_SCRAPER_LINK        | http://localhost:8081        |
| DEVOTIONAL_BACKEND_MONGO_URI           | mongodb://127.0.0.1:27017/   |
