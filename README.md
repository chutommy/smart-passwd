# SmartPasswd

SmartPassword is a generator of extremely strong passwords that can be easily remembered thanks to the reminders with which passwords are being generated.
The generator allows the user to select the length of the generated password and the number of extra symbols that will be randomly added. However, the best feature of this app is that the user can generate a password with a custom reminder.

Under the hood, the application collects words from the English dictionary (all words are parsed into the database from <a href="https://wordnet.princeton.edu/">WordNet</a> - a lexical database for English - hosted by Princeton University) and composes them according to the random number generator into a certain structure, which very often sounds laughingly but allows the user to quickly recall the generated password. Then the generated phrase goes through an algorithm that may or may not (again depends on RNG results) replace each of the letters with another one. But in order to preserve to a certain memorability, each character can be replaced by a character, which sounds or looks similarly.

### Screenshots
<img src="https://raw.githubusercontent.com/chutified/smart-passwd/master/imgs/1.jpg">
<img src="https://raw.githubusercontent.com/chutified/smart-passwd/master/imgs/2.jpg">
<img src="https://raw.githubusercontent.com/chutified/smart-passwd/master/imgs/3.jpg">

#### server logs
<img src="https://raw.githubusercontent.com/chutified/smart-passwd/master/imgs/4.jpg">

### Used languages and tools
  - Go (Gin framework, Postgres driver)
  - Postgres SQL
  - Javascript + jQuery
  - HTML (HTML5 UP design), CSS (SCSS) + Bootstrap
  - Docker + Docker-compose
