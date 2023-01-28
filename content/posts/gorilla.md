---
title: "gorilla"
date: 2023-01-21T18:27:27-06:00
draft: true
---

{{<
figure src="https://media.tenor.com/sidMbBYqr0EAAAAC/meme-made.gif"
>}}


Gorilla maintainers have decided  to shutter the doors to the project and archive the [gorilla repositories.](https://github.com/gorilla#gorilla-toolkit) 
The maintainers are signaling that there will no future development on the projects, and are encouraging the community to make and maintain forks of the 
gorillia repositories. If a critical issue were to arise in one of the gorilla projects, a zero day or other bug, it is very unlikely that fix for said issue will make it to
the github.com/gorilla repositories. 
Instead of un-archiving the projects and merging in a fix the maintainers will likely defer to the community and discourage users from using github.com/gorrila repositories further.    


## Why Me

I personally know enough people that depend on gorilla libraries that It would make sense for me to offer to maintain a fork.
I hope to get more involved with open source software, and this is a great first step. Eventually I hope to introduce some open source projects of my own to the go community. 
I have gotten really good at maintaining other people's code at work, so I feel confident that I can do so in an open source environment as well. 
If you trust me as a maintainer feel free to also use this fork, so long as you comply with the license. 

## License

Even though the gorilla libraries are feature stable and battle tested the maintainers still gave up on maintaining them. 
To make sure that these libraries continue to be viable to actively maintain in a feature freeze state I will be licensing new versions of the gorilla repos under AGPL 3.0. [^1]
This will help ensure that when issues arise (if they ever do) that the fixes get published as open source software allowing me to easily integrate the fix.


## Maintaince / Support

Licensing under AGPL, having no github issues on the repository, and offering no support directly should make continued maintenance viable for me.  
If you have a patch that you would like to get merged you can get in touch with me at tyler@smantic.dev.
If you have questions about how to use these repos kindly direct them to the community at large. Join either the [discord](https://discord.gg/golang) or the [slack](https://blog.gopheracademy.com/gophers-slack-community/) and ask your question accordingly.
Or use other [Go help resources](https://go.dev/help).

## Migration

Migrating to my fork is easy. 
```bash 
find . -type f \
    -name '*.go|go.mod' -o -name 'go.mod' \
    -exec sed -i -E 's,github.com/gorilla/(websocket|mux|handlers|schema),go.smantic.dev/\1,g' {} \;
```

If you're on a mac, I recommend installing gsed.


# Repositories  

* mux - [go.smantic.dev/mux](https://github.com/smantic/mux)
* websocket - [go.smantic.dev/websocket](https://github.com/smantic/websocket)
* sessions - [go.smantic.dev/sessions](https://github.com/smantic/sessions)
* handlers- [go.smantic.dev/handlers](https://github.com/smantic/handlers)

---
[^1]: Yes I know that lots of corporations ban any use of AGPL licensed software in their products. 
