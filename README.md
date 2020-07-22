# desearch
A wrapper around [sci-hub.tw](https://sci-hub.tw) for downloading Research Papers.

## Installation
If you are a GO user, 
```shell
go get github.com/krAshwin/desearch
```
For linux & windows users, they can use by directly downloading the [portable executables](https://github.com/krAshwin/desearch/releases). 

## Usage
If you are a **linux** user, download the executable from [linux/desearch](https://github.com/krAshwin/desearch/releases/download/v2.1_linux/desearch) and traverse to the directory where the downloaded file is and run with the following command,
```shell
./desearch [url]
```

If you are a **windows** user, download the executable from [windows/desearch.exe](https://github.com/krAshwin/desearch/releases/download/v2.1_windows/desearch.exe) and open Command Prompt(cmd) and go to the folder where desearch.exe is and run with the following command,
```cmd
desearch.exe [url]
```

To run the code globally, <br>
For Linux: Move the executable to `/bin` directory and then it can be run globally.
```shell
desearch [url]
```

# Output
May spit error if unable to download the file.<br>
Otherwise, 
```shell
Downloading Research Paper: {filename}
Done
```

## Known url sources
- [IEEE](https://ieeexplore.ieee.org/Xplore/)
- [ACM](https://dl.acm.org/)
- [Science Direct](https://www.sciencedirect.com/)
- [Springer](https://www.springer.com/in)
- [Nature](https://www.nature.com)

## Author
**Kumar Ashwin** <br>
**Email - krahwin00@gmail.com**
