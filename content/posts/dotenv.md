---
title: "dotenv"
date: 2023-10-29T21:10:51-05:00
draft: false
---

I recently submitted a PR to an open source project, a CLI tool.
I stumbled upon a few requests for the maintainer to add support of *parsing an .env file* while invoking the CLI. 

I find this desire ridiculous. And there is no shortage of this desire in other projects.  
Should every application parse a .env file for configuration? Obviously, no. 
We all ready have plenty of convenient ways of automatically setting env vars.
  
Namely 
- `KEY=value ./cli` use a script if you must
- `set -a; source .env; set +a` (create a custom function in your profile!)
- direnv -> https://direnv.net/
    - automatically sets env vars based on a CWD's .envrc file. 

Note that these approaches have specific caveats, such as how you specify the env vars, whether the env vars are exported, or what the name of file is.

Parsing, and exporting env vars in cross platform compatible way, and in a way that most people will expect is not a trival amount of complexity to add to your application. 
Not to mention the complexity it may add to your user's experience. As in: 
- What happens when I have the same key in a .env file and the environment.
- How are the values expected? 
    - key="value"
    - key=no quotes value   
    - key = value
    - export key=value

If a key, value pair never hits the environment, should it really be called an env var? 
Calling it an env var I might expect it to appear in the output of env. 

What is the benefit of including this complexity? 

Since we already have extremely convenient ways of setting env vars, why should we enforce that every application contains the complexity that is parsing some non-standard key value paring? 
Instead we can keep our programs simple, and parse configuration one way. 
Safe to say I wont be adding dotenv support to any of my apps that utilize env vars anytime soon. And I wouldnt expect other open source developers to do so either.
