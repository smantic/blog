---
title: "dual-booting"
date: 2021-04-06T22:32:51-05:00
draft: false
toc: true
tags: 
  - linux
---

---

This is a guide to a bare metal dual boot of Windows 10 and Ubuntu. No emulation, or virtualization.
Most of this guide comes from my personal experiance over several years of having a dual boot install, 
any pain points that I highlight are because I have experianced them myself. Obiviously, this is not 
a complete resource of all the information you might need to do this, but this post does have all the information you mighy need.   

Notice that I say that it is a dual boot of Windows 10, and Ubuntu (windows first). 
Unfortuantly the Windows installer has a [terrible habit](#1---windows-installer) of destroying any other operating system
that may be existing peacefully on the computer that it is being installed on. 
The solution to this? Simple, just install windows first.  

## Installing Windows

Thankfully, thats not an issue for you because you are an intelligent, responsible human being that will avoid uneccessary complications when possible. 
Here is what you will need to install windows.
A flash drive with atleast 8gb of space and a windows computer with atleast 8gb of space on the C:/ drive. 
Yes you need a windows computer to install windows. 
Once you have those go ahead and download and run window's installation media.
Again make sure there is plenty of storage both on the flashdrive and computer that you are running the installation media on because other wise the installation media will fail!!  

Depending on what you are installing onto you may need to familarlize yourself with the partions on your system. 
When are doing the windows installation, the installer is not going to be helpful to identify what partition is which. 
If you are on windows open up the Disk Management program, and if you are linux open up Gparted or whatever your distro uses to do Partition Management. 
You will want to identify what drive you want to install onto, and what partions you will need or want to delete ( format ).
If you are installing onto brand new drives with no existing partions then dont worry about it.   

After closing your 5 tabs of reddit that you opened after getting bored waiting for the installation media; go ahead and plug that flash drive into the computer you want to dual boot on. 
Restart your computer now. But before you do that look up what your bios key is ---  too late? Quick! Spam f2 and delete vigoriously. 
Nice. Now click around your bios interface randomly for 10 minutes until you find where you can select what device to boot into. 
Select the one that has the brand name of your flash drive and UEFI in the name of the partition. 

Select the advanced installation option when you get to it because we don't trust windows to anything right. 
This will bring up a view of the different partions on your system. 
Hopefully you have [identified](#7---identifing-your-disks) which drive you wanted to install on before hand.
Go ahead and delete any partitions you don't need anymore (such as an old windows partion, or old linux partions ). 
After you have created some free space on the drive you want to install on (i.e. your SSD, not the hard disk), 
select the free space to install into. 
After that Cortona will attempt to slide in your DMs 
and start ask you a serries of invasive questions such as your name, password, location, and your mother's maiden name. 
Make sure to select no, or skip whenever possible otherwise windows might server you adds. 

### Installing Windows and Linux on the Same Device

This part is for the really dense people out there that will inevitably want to have their linux partion and their windows on the same SSD for those fast boot times. 
Or you just only have 1 drive to install anything on. 
This idea is folish because the minute that Microsoft decides they want to change their boot process, 
or even decide it's a "good idea" to overwrite the master boot record again which will break your GRUB bootlooader.
Microsoft could change other things, and possibly even brick the linux kernal that happened to be on the same drive. 
My linux kernal went into a panic at the sight of windows attempting to update, and my windows update is stuck in a forever fail loop. 
But if you want to continue down this accursed path you will want to open window's Disk management program so that you can resize your windows partion to make room for linux.  

## Installing Linux

Now for the easy part. Grab that flash drive and download a software called Rufus to your windows machine. 
you'll also need to navigate to Ubuntu's download page and install the latest LTS version of Ubuntu. 
Open up rufus, select your flash drive, and hit the big "SELECT" button next to Boot selection. 
Select the Ubuntu .iso that you just downloaded. 
When it's ready select the big "START" button. 
If rufus asks you about Syslinux just say yes if you have acess to the internet. 
Choose the recommended selection ( Write in iso image mode ), and then Rufus will begin writing the Ubuntu installation media for you.     

After this boot into your bios again and select to boot into the flash drive.  
You should be greeted by a friendly GRUB interface, select the Ubuntu installation. 
You will get the option to to connect to the internet to download extra proprietary Software/Drivers, 
I'd recommend doing that as it's convenient and free software is a conversation for another day. 
I'd especially recommend doing this if you are running a computer with a Nividia graphics card 
or if your on a laptop that has just recently came into the market.
                    
When you get to the point where the Installer asks how you want to install, select "Something else", instead of the other options. We are going to manually create the partitions to ensure the partitions are set up how we want them to be.
  
These are the partitions that we will want to have, enter them in this order.  
* A swap partition for temporary files / hibernation. 
* Boot partition (/boot) for GRUB 
* Root Partition (/) for the Kernel and system files, etc. 
* (optional) Home partition (/home) for day to day files and programs. This one isn't necessary and you may just want a large combined root and home partition.
See [this](#3---order-of-partions) for how large to make each of these partitions.    


The Ubuntu installer can do this for you, but depending on what you are installing on top of, 
it might not be an ideal set up, or it might install along side your windows partition even if room on that drive is tight. 
I should also mention that you can elect to have root and /home in the same partition, 
but some separation between your files and the operating system can be pragmatic 
for when it comes time to upgrade Ubuntu your /home partition wont be touched.   

Once you're finished go ahead and reboot into your bios, ensure that the drive you installed Linux on is listed first in the boot order. 
After that, select to boot into windows, and then boot into Linux again. 
This is to just make sure your both of your installations work, and that they aren't going to brick each other after booting in one of them. 
After booting into windows your GRUB should recognize windows as an option to boot into now, so you won't need to go into the BIOS every time you want to boot into windows.

And that's it! All done! Hopefully, everything went smoothly for you, otherwise don't be afraid to ask for help. 
Do your best to identify what the problem may be and other people will do their best to help you solve it! 

Notes 
---

### 1 - Windows installer

I may have over exaggerated a bit.
Windows will just overwrite the master boot record on the drive, because the master boot record should show how to boot onto windows. 
Unfortunately, this makes it making much more difficult to boot into your ubuntu partition again.  
To do so you will have to do a boot-repair, which can be difficult if you're not sure what you're doing, and also not really guaranteed to work. 
So just install Linux last! To start the boot-repair process you will need to execute these commands:   
``` 
sudo add-apt-repository ppa:yannubuntu/boot-repair 
sudo apt-get update
sudo apt-get install boot-repair
```  

### 2 - Partition Sizes   

When creating custom partitions like this it's important to understand how large the partitions need to be in order to 
conserve space by not making extra large partitions, as well as making sure the partitions have enough space for their intended use.

 *  `Swap` - 2x your computer's memory. Swap is basically overflow for when you run out of memory. 
You can tell when your computer is using it because your computer slows down dramatically since it is reading / writing to disk instead of the low latency RAM. 
Having 2x your memory is nice such that your computer can go into hibernation (all of memory goes to swap partition), and still have room in swap for things that might have already been there.
 *  `/boot` - 200-500 MB is fine, GRUB should only use a fraction of that, but room for any updates is nice. 
 *  `/` - this is where your system files, and programs will live. How large this is depends on how much space you can spare, and whether you think you
will be installing a lot of applications to your computer. At the minimum however you will want to have 8-20 GB. 
 * `/home` - this is where you will store all of your personal files, pictures, and etc. How much space you allocate to this partition depends on how much space you have,
and if you plan on storing a lot of pictures or documents.
Note that if you are like me and you have a minimal amount of data you need to store in /home you can just go ahead and have a combined root and /home partition. 
  
Regarding root partition size - On my system I've installed all the extra software, and drivers and all together my system files take up 6 GB of space.
Which is much better than the 60 GB of system files needed for windows... 
 
### 3 - Order of Partitions 

Interestingly enough the order of the partitions that we put in might matter. 
I didn't know about this before writing this guide, so I've included some interesting findings.  

> Yes. Data located at the outer edge of a traditional hard disk will be sequentially read faster than data closer to 
> the center of the platter. This is just physics. The tangential velocity of the outer tracks is faster than inner tracks so the rotational latency is lower.
 
From a discussion on superuser forums discussing how reading on the edge of the hard disk is faster. 
I'm too lazy and not smart enough to be bothered to find the actual physics behind this. So if this wrong be sure to send me hate mail. 
 
> A partition shown on the left side would probably be actually located on the outer cylinders. GParted has property boxes that provide the actual disk addresses 
> (by sector numbers) to verify these relationships.
 
So such we know that partitions are (generally) listed left to right, indicating outer rings to inner rings on the right. We will want the swap partition
on the outermost ring (first partition), so we can quickly restore after a hibernation, or quickly write to swap when we have overflow. 
Next we will want our boot partition.
There seem to be plenty of interesting articles to read on this topic. 
But of course if you installing onto an SSD none of this matters anyway.  
  
### 4 - Separate Home directory

If you plan on storing a lot of things in your /home directory (photo albums or whatever), 
it may be prudent to have a separate /home partition so you can safely upgrade your system (or even change distros) without worrying about losing your data stored in /home. 
Also, I believe it's possible to have 2 distros on one system, that share a /home partition, so you can swap between your arch distro, and ubuntu or whatever 
and have access to the same files, altho I can't guarantee this will work perfectly.   
  
### 5 - What version to download?

In most cases you will want to want to use the latest long term support (LTS)
version as version provide great stability, infrequent updates, you know... plenty of support. 
There are some cases in which you may want to install the newer version of Ubuntu however. 
One of those cases is when you have purchased newer hardware that only newer versions of the Linux Kernel or
newer versions of the distro contains fixes needed for that piece of hardware to operate correctly.
Specifically this should only be a problem for laptops. Consult [this](https://certification.ubuntu.com/desktop) or
any up to date compatibility list to see if your hardware is compatible. 

### 6 - Windows Update Forever Fail Loop  

For some reason it's possible for your windows update to corrupt, and fail every time you attempt to update. 
This wouldn't be so bad if Microsoft wasn't so aggressive about updating windows, 
some older versions of windows 10 you couldn't turn your computer off without going through the update process if there is an update pending. 
If you do install windows and Linux on the same drive save you self some grief and disable updates completely. 
You can accomplish this by opening up the "Services" program, finding windows update, select the properties of the service, and change startup type to "disabled".
Although it's possible by the time you are reading this, Microsoft has changed how this is done, and this is no help to you.

### 7 - Identifying your disks

If you are in the middle of an installation of windows and you're not sure which drive is which it's still possible to figure out.
The easiest way to go about it if you are like me is to add up all the space on a particular drive to identify what drive it is.
The space is identified as bytes instead of Gigabytes, so you will need to divide the total by 10 9 in order to get what the size is in gigabytes.
Hopefully this should help you identify what drive is which, assuming you remember what the size of your drives are. 
