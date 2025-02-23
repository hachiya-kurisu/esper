# esper

makes sure gemini/spartan/nex/https services are up and running

reads and writes gemtext

## usage

the following command reads services.gmi, checks each link in the document and prepends a pair of emoji representing its current status:

* success: the two emojis will match (🐝🐝)
* failure: the two emojis will not match (🐐🐒)

```
$ esper -c services.gmi
# my services

## blekksprut.net

=> gemini://blekksprut.net/ 🐙🐙 gemini
=> spartan://blekksprut.net/ 🦍🦍 spartan
=> gemini://blekksprut.net/notfound 🦋🦉 404 😱

## manatsu.town

=> nex://manatsu.town/ 🐋🐋 真夏駅

last updated 2025.02.22 01:31
```
