---
title: "retro-gaming-devices"
date: 2021-04-07T17:37:50-05:00
draft: false 
---

I loved playing pokemon games as kid. Problem is that I still love playing them. 
I've played pokemon games on emulators at my PC, but sitting at computer after I just sat there all day for my job doesn't really capture the joys of my childhood.  
So I've started to trying to find a handheld device that will let me rexperiance the glory days of the game boy and the nintendo DS.  

However, it's been a decade plus since I last used any of my nintendo handhelds, 
and even if I knew where they were I doubt they'd be very useable at this point. Not to mention that I have new expectations of the hardware I use. 
The only charging adapters I have now are usbc, and the only SD cards I own are 256 gb or more. Both of which didn't exist in 2006. 
  
Ideally I'd like to use my switch to play the classic pokemon games that I loved. It has everything I want in a handheld console in 2021. 
Support for large capacity sd cards, usbc connection and charging, as well as detachable controller and the ability drive a widescreen tv. 
However, the ability to play the old pokemon games that I love is virtually non-existant on the switch. 
I've explored what homebrew is like on the switch, with moderate success but immense disapointment.

  
## Homebrew on the switch
  
The current state of homebrew on the switch is inconvient to say the least. Homebrew currently works through an exploit called Fusée Gelée.
Which works by exploiting the boot process of vulnerable Tegra SoC (system on a chip) CPU. Which are dual core ARM cpus and happen to be used in the switch.
Since we are exploiting the boot process of the processor there is no presisting of the exploit after a reboot. 
So if your switch ever turns off you'll` need to do the exploit over again. 
Which consists of putting your switch into recovery mode by grounding specific pins on switch's joycon rail 
and then holding the volume up button and the power button while the switch is turned off.    
  
After doing this you use a Fusée Gelée injector on a computer connected to the switch's usbc port. 
Using the injector you inject the Fusée Gelée 
payload along with any code that you want executed after the exploitable boot process, such as boot loader or straight to custom firmware.  
  
I got as far as being able to boot into atmosphere (a popular homebrew OS for switch), and attempted to download the melon DS emulator from the homebrew store.   
But for Melon DS to work you need a dump of a nintendo DS's BIOS/firmware, which doesn't come with the homebrew download (because it would be illegal to do so).  
Even after all of that, the emulator isn't all that great. And it's commly known that melon DS and other emulators don't run well on the switch. 

>Keep in mind that all of the aforementioned emulators do not currently support a hardware renderer and many games run at slow speeds. 
You can however overclock the Switch's CPU up to 1.75 GHz to achieve somewhat acceptable framerates.

Oh, I forgot to mention that if your switch comes into contact with nintendo servers while doing homebrew stuff you run the risk of banning your switch.
Leaving your switch unable to access any nintendo online features or the estore. 
And the only way to recover from that is to back up your switch's eMMC (embeded multi-media controller) and some authentication keys. 
Which would be rather inconveint for me since I want to grab the pokemon snap remake coming out this month. 

### Homebrew on NDS 

If I had a NDS I could play gameboy / NDS cartrages directly on the device and get homebrew software to play ROM hacks on the SD card.  
Which would be nice, and there should be no problem running the games since they were designed to run on that hardware.
But I don't have a NDS to try to homebrew on at the moment, 
and also the NDS wouldn't exactly be my first choice since it's not usbc and doesn't support SD cards larger than 2 GB (or 32GB for SDHC).  


## RG350 retro handheld 

My first attempt at finding a device that would allow me to re-experiance the golden age of pokemon games 
is a rasberry pi with a 3.5" ips screen, plastic/metal case, and 3D printed buttons. 
It has good support for expanding storage, and has usbc charging/data transfer. 
However, it's incapable of charging from anything that outputs more than 5 amps, and in my experiance suffers other battery issues.
To resolve I had to unplug and replug the battery back in. All in all it's not the worst, and is somewhat useful. 
But it won't play DS games well, so I continued to search for a device that could do gameboy games and NDS games. 


### Analouge Pocket

Im looking forward to Analouge's 'pocket' device, and I hope that I can get my hands on one when they do the next round of production
(preoders for the batch shipping in October 2021 sold out in 16 minutes).
Its got usbc charging, a micro-sd slot, and HDMI output all in a slick designed Gameboy like handheld console. 
It won't play DS games of course but I'd definitely get it just so I can play old Gameboy games on it. 


## Simple Solution 

The solution that I've ended up with using the most for now is just hooking up my PC up to the TV in my living room through a long HDMI cable I happened to have. 
I play from the couch using a wireless xbox controller that I got from my Oculus, and it's as simple as that. 
I can play DS games, gameboy, or even gamecube games, including any hacked variants of those games.
But I can't exactly take my TV to the park to play pokemon like I did as kid (you know after they added backlight screens).
This will work for now but I will continue to search for a better solution. 
