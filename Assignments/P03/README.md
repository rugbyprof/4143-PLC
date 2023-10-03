## Program 3 - Image Ascii Art
#### Due: 10-05-2023 (Thursday @ 9:30 a.m.) 

## Overview

This project is going to allow you to utilize the packaging strengths of Go, and make your own libs that can be used anywhere just by running the `go get ...` command. I want you to create four packages in one module or four different modules. I'm ok with both options. However, one or three will be placed in your github account (not in your course repo). 

#### Three modules
```
github.com/yourUserName/img_colors
├── Img_colors
    ├── main.go
    └── go.mod

github.com/yourUserName/img_text
├── Img_text
    ├── main.go
    └── go.mod

github.com/yourUserName/img_gray_scale
├── Img_gray_scale
    ├── main.go
    └── go.mod

github.com/yourUserName/img_get
├── Img_get
    ├── main.go
    └── go.mod
```


#### One Module 
```
github.com/yourUserName/img_mod
├── Colors
│   ├── go.mod
│   └── main.go
├── Grayscale
│   ├── go.mod
│   └── main.go
└── Text
│   ├── go.mod
│   └── main.go
└── GetPic
    ├── go.mod
    └── main.go
```

Each repository can then be imported to your projects as needed. I have provided the logic for each `package` in the form of a markdown file:

- [colors.md](./colors.md)
- [color_text.md](./color_text.md)
- [grahscale.md](./grayscale.md)
- [getPic](./getPic/)

These files have small explanations, but also the necessary code to implement each module. 

### Creating Folder and Repo(s)

The chat GPT discussion doc [HERE](../../Lectures/module_conversation.md) has a semi walkthrough on how to go about creating a module. I will write a seperate one here as another source of info.

1. Create a repository online in your github account. Name it something like I did above. 
2. Use vsCode to clone that repository and bring it down to your own computer. 
3. To clone:
   - Hit `ctrl`-`shift`-`p` to get your command bar at the top of vscode.
   - Type `git`, all the commands with `git` in them will show up.
   - Choose `clone from github`, then choose the repo you created. 
4. Depending on you using one or three mods, you can follow the directory structure from above.

 
### Adding Code 

- Depending on the directory structure you chose, you now need to add the code to your module (whether it contains 1 or 4 packages).
- For the `colors` module, here is what I've done/.
- After I added the subfolder `colors` (within the outer colors folder) I ran the following to initialize my color module.

```bash
$ go mod init colors     
go: creating new go.mod: module colors
go: to add module requirements and sums:
        go mod tidy
$ go mod tidy                                                 
```

I pasted the code from [colors.md](./colors.md) into a `main.go` file within the colors sub directory. However, the code I gave you is not very helpful as a package or module to be included to another project if there is only a main function. So you will need to place the logic into a function!

## More Coming ...
