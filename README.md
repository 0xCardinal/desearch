# desearch
A wrapper around [sci-hub.tw](https://sci-hub.tw) for downloading Research Papers.

## Installation
If you are a GO user, 
```shell
go get github.com/krAshwin/desearch
```
## Usage
Run the following command from inside the desearch directory,
```shell
./desearch {url}
```

To get the usage & credit direction,
```shell 
desearch -h
```

To run the code globally, <br>
For Linux: Move the executable to `/bin` directory and then it can be run globally.
```shell
desearch {url}
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
