---
title: "gorilla"
date: 2023-01-21T18:27:27-06:00
draft: false
tags:
  - go
  - FOSS
---

{{<
figure src="https://media.tenor.com/sidMbBYqr0EAAAAC/meme-made.gif"
>}}


Gorilla maintainers have decided  to shutter the doors to the project and archive the [gorilla repositories.](https://github.com/gorilla#gorilla-toolkit) 
The maintainers are signaling that there will no future development on the projects, and are encouraging the community to make and maintain forks of the 
gorillia repositories. If a critical issue were to arise in one of the gorilla projects, a zero day or other bug, it is very unlikely that the fix for said issue will make it to
the github.com/gorilla repositories. 
The gorilla maintainer(s) will likely defer to the community to incorporate the fix into one of the forks instead of un-archiving the original repositories.
  
I personally would prefer if the gorilla projects didn't die, since they are important, widely used libraries in the go community. Particularly websockets has been the go to websocket library. So I decided to fork the gorilla repositories I was interested in. Specifically: Mux, websocket, sessions, and handlers. Feel free to also use them if you would like.  


## Maintaince / Support

These libraries are already battle tested and feature stable. Because of this I don't plan on much development on them, only if a critical issue were to arise that meant a patch would be required. If you are aware of a critical issue with any of the libraries please contact me at tyler@smantic.dev. 
If you have questions about how to use these repos kindly direct them to the community at large. Join either the [discord](https://discord.gg/golang) or the [slack](https://blog.gopheracademy.com/gophers-slack-community/) and ask your question accordingly.
Or use other [Go help resources](https://go.dev/help).

## License
I am licensing future versions under LGPL. For the most part this wont matter because these libraries almost always sit on a webserver somewhere, so are vulnerable to the ["Application Serivce Loophole"](https://fossa.com/blog/open-source-software-licenses-101-agpl-license/) of the GPL. But I do wish that if people were to make improvements upon these libraries that they open source them and share them with the wider community, like you would normally do with typical LGPL software.  

## Migration

Migrating to my fork is easy. 
```bash 
find . -type f \
    -name '*.go|go.mod' -o -name 'go.mod' \
    -exec sed -i -E 's,github.com/gorilla/(websocket|mux|handlers|schema),go.smantic.dev/\1,g' {} \;
```

If you're on a mac, I recommend trying with gsed.


# Repositories  

* mux - [go.smantic.dev/mux](https://github.com/smantic/mux)
* websocket - [go.smantic.dev/websocket](https://github.com/smantic/websocket)
* sessions - [go.smantic.dev/sessions](https://github.com/smantic/sessions)
* handlers- [go.smantic.dev/handlers](https://github.com/smantic/handlers)

