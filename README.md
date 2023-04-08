
# Subtitle Manager

Subtitleman is a command-line tool written in Go that finds all .srt files in a given directory and merges them into a single file named `out.srt`.

## Usage

To use Subtitleman, run the following command:

```
./subtitleman ./input/
```

The argument `./input/` is the path to the directory containing the .srt files that you want to merge.

## Example

Suppose you have the following .srt files in the `./input/` directory:

```
part0.srt
part1.srt
part2.srt
```

To merge these files using Subtitleman, run the following command:

```
./subtitleman ./input/
```

After running the command, a file named `out.srt` will be created in the `./input/` directory. This file will contain the merged contents of the `part0.srt`, `part1.srt`, and `part2.srt` files, ordered by filename.

## Contributing

If you find a bug or have an idea for a new feature, please open an issue on this GitHub repository. We welcome contributions from the community and appreciate your feedback!

## License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).

---

*This README.md file was generated by GPT 3.5.* 
