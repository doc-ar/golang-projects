# Introduction

When I started learning GO language I did not understand where to start. However, I quickly understood that the best way to learn is to build stuff I am interested in. So This repository contains all the projects I have created while learning GO language. I will be adding seperate README files for each project. They are listed below. The resources I am currently using or plan to use for learning go are listed at the bottom of this page.

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
- [Standard Library Webserver](go-webserver/README.md)
