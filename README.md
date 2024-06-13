# GoBitLy


### STEP-1

To access the UrlShorting Link 
```
go to http://localhost:8080/shorten
```
### STEP-2

Give the body param in .json (In POST method)
- you can change the url
```
{
  "url": "https://abhinandpn.netlify.app/"
}
```
### STEP-3
Then you got a body responce 
- eg
```
{
    "short_url": "1718290128702514889"
}
```
### STEP-4

Copy the responce id .

### STEP-5
Peste the id like this 
```
http://localhost:8080/short/{shortURL}

- in this example ...

http://localhost:8080/short/1718290128702514889
```
### Here you with the link...!
