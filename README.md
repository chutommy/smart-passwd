![SmartPasswd logo](public/assets/images/logo.svg)

# Smart Passwd

Smart Passwd is a web application which helps people to generate strong passwords that are easy to remember.

## Features

### Strength level

Users can adjust the strength level of the password which adds an extra layer of security by inserting random
numbers or special characters into the final form of the password.

### Reminders

The generator creates each password along with its `reminder` which is a word or phrase that either graphically
or phonetically resembles the real password.

#### Custom reminders

Users can run the generator with a custom reminder which creates a new unique password based on the given text.

### Hide button

Hide button is placed on the left side of the `generated password` field. Three seconds after a password is generated
the values are hidden and its visibility can be easily toggled on and off.

### Copy buttons

Both `generated password` and `reminder` have copy buttons to easily copy their values to clipboard. After pressing
either of them a green notification appears to confirm it was successfully stored in the clipboard.

## Examples with random reminders

![Screenshot 1](imgs/screenshot_1.png)

### Fort short passwords

* Length:   6
* Strength: 0

#### outputs:

* `v&gon$` **[ wagons ]**
* `Sv&mi$` **[ swamis ]**
* `trcnDS` **[ trends ]**

### For long passwords

* Length:   20
* Strength: 2

#### outputs:

* `Or4PHanqeEIIn9$Gcrk8` **[ orphan peelings berk ]**
* `IO@fS@R1eHcolOGy$tS6` **[ loafs archeologists ]**
* `incONS0cqucntl44I@ak` **[ inconsequential oak ]**

## Examples with custom reminders

![Screenshot 2](imgs/screenshot_2.png)

### Smart Password

* Strength: 0

#### outputs:

* `$Martq4$$vOrb`
* `sm@rTP&S$v0rD`
* `sMArTq&SsvORD`

### I like to eat pineapple

* Strength: 3

#### outputs:

* `yL2ikEt0!!Atpi2nc4PPle`
* `l9like))OE@8tplNc4PPlE`
* `lIyketOe4tqI7nc&^^qIe6`

## Support for mobile devices

Tablet / iPad view | Smartphone view
| :---: | :---: |
| ![Screenshot 1](imgs/screenshot_3.png) | ![Screenshot 1](imgs/screenshot_4.png) |

## License

The Aerial design is provided by HTML5 UP (cc 3.0), the WordNet 1.6 database by Princeton University
and the slider component by Brandon McConnell.

The project is under the MIT license open source software. I welcome contributions both big
and small! Please, take a look at the community [contributing notes](https://github.com/chutommy/smart-passwd/blob/master/CONTRIBUTING.md).
