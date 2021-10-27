---
title: "aws-kms"
date: 2021-10-13T18:33:39-05:00
draft: true
tags: 
  - aws
  - encryption
---

Amazon's Key Management Service is a service for managing encryption keys.
  
It's not meant to directly encrypt your data. That's what AWS Encryption SDK is for. 
[AWS Encryption SDK](https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/introduction.html) is a client side library that you can use to encrypt data. It is only available in some languages, not including go. 



Fear not, because in go we don't need AWS Encryption SDK; all of the tools that we need are included in go's standard library. 

Here's an example using an 
