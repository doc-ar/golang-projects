#### Table of Contents

- [Introduction](#introduction)<br>
- [Prerequisites](#prerequisites)<br>
  - [Linux](#linux)
  - [Windows](#windows)
- [Project Docs](#project-docs)<br>
- [Learning Resources](#learning-resources)<br>

# Introduction

When I started learning GO language I did not understand where to start. However, I quickly understood that the best way to learn is to build stuff I am interested in. So This repository contains all the projects I have created while learning GO language. I will be adding seperate README files for each project. They are listed below.

# Prerequisites

For running and building any of these projects, you will need GO installed.

> Project specific dependancies will be mentioned in their respective README files

## Linux

For linux distributions its a simple process of installing go using your package manager.

```
# Arch Linux
sudo pacman -S go
```

```
# Debian / Ubuntu bases Distros
sudo apt-get install go
```

## Windows

Installing any language in Windows tends to create many problems (Mostly $PATH problems). So I've come up with a good solution. We will use [Windows Utility](https://github.com/ChrisTitusTech/winutil) by [Chris Titus Tech](https://www.youtube.com/@ChrisTitusTech) on Youtube

> This utility uses winget for installation.

1. First open a powershell window in administrator mode
2. Enter the following command:

```
irm "https://christitus.com/win" | iex
```

3. A GUI window will open up.
4. Go to `Install` tab (if not already there) and select `Go` from the list
5. Click `Install Selected` button above the list
6. You can check the installation progress in the original powershell window

# Project Docs

- [Tic Tac Toe](TicTacToe/README.md)
- [Basic Webserver](go-webserver/README.md) <sup>This project does not do anything useful other than creating a basic webserver</sup>

# Learning Resources

These are the learning resources I've used or I am planning to use while learning go language

- [Go by Example](https://gobyexample.com/) is my go to website for learning go syntax
- [Golang Projects Playlist](https://youtube.com/playlist?list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&si=H8EVRUZRGPX6UqxJ) by [Akhil Sharma](https://www.youtube.com/@AkhilSharmaTech) on Youtube
- [Gophercises](https://gophercises.com/) is a website which with great project based exercises
- [Let's Go Further](https://lets-go-further.alexedwards.net/) is a great website for learning advanced patterns for learning to create APIs and web apps
- [Exercism](https://www.exercism.org) is a free learning website for a lot of languages. It uses interactive exercises to teach you.
