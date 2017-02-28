# Software Lab

## What is it?
Software Lab is a local mirror that lets users at hackathons quickly and easily download commonly used software packages at hackathons.  The software is designed to be easy to both setup and scale with a one button setup process.  Overall, hackers can download software like Android Studio and Unity in seconds / minutes rather than hours through an easy to use web interface, rather than SAMBA or FTP.  Software Lab is currently implemented at the University of Akron and is scheduled to roll out at other MLH Hackathons for the 2017 Spring Season.


## Why Use it?
Software Lab gives hackers substantially faster download speeds that are consistent.  In addition, software downloads don't take up bandwidth on the internet, as everything is local.  Hackers will spend less time downloading software and more time working on their projects.  In addition, the network is under less stress and the internet will be faster overall for the hackathon.


## How does it work?
The application consists of two parts.  The first is a local server, which mirrors the software packages locally.  The system is scalable, so several local servers can be set up.  The remote server, the second part, tracks what local servers exist and point users to the local servers.  This allows users to go to a simple web site, rather than some internal IP address which could change, and makes the download process as easy as downloading it normally from the web.
