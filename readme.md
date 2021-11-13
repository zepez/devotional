

# ✨ Devotional Application

This is the respository for my devotionals application.


## 🤔 Structure

 - frontend - likely a server-side rendered application (nuxt maybe?)
 - [backend](https://github.com/zepez/devotional/tree/main/backend)
   - routinly grabs data from scraper
   - interfaces with mongodb to store scraped data
   - allows for retrieval of scraped data by id
   - allows for paginated view of scraped data
 - [scraper](https://github.com/zepez/devotional/tree/main/scraper)
   - scrapes data from [lhm.org](https://www.lhm.org/)
   - caches data to prevent rescrape within 1 day

## 🖊 Notes

 - this project was started for the intention of leaning go and a new front-end framework
 - all data is scraped from [lhm.org](https://www.lhm.org/)

## ✅ To do

- [ ] Scraper
  - [ ] testing `package/handlers`
  - [ ] add rate-limiting
- [ ] Backend
  - [ ] testing
  - [ ] add rate-limiting
- [ ] Frontend
  - [ ] start





  