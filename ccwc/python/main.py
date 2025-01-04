import argparse
from io import TextIOWrapper


def byte_count(file_path: TextIOWrapper) -> int:
    bytes = 0
    with open(file_path.name, "rb") as file:
        while line := file.readline():
            bytes += len(line)
    return bytes


def line_count(file_path: TextIOWrapper) -> int:
    lines = 0
    while file_path.readline():
        lines += 1
    return lines


def character_count(file_path: TextIOWrapper) -> int:
    characters = 0
    with open(file_path.name, "rb") as file:
        while line := file.readline().decode():
            characters += len(line)
    return characters


def word_count(file_path: TextIOWrapper) -> int:
    words = 0
    with open(file_path.name, "r") as file:
        while line := file.readline():
            words += len(line.split())
    return words


def main(args: argparse.Namespace) -> int:
    count = ""
    if args.byte:
        count = f"{count} {byte_count(args.file_path)}"
    if args.line:
        count = f"{count} {line_count(args.file_path)}"
    if args.word:
        count = f"{count} {word_count(args.file_path)}"
    if args.character:
        count = f"{count} {character_count(args.file_path)}"

    print(count, args.file_path.name)
    return 0


if __name__ == "__main__":
    parser = argparse.ArgumentParser(prog="ccwc", description="wc clone")
    parser.add_argument("-c", dest="byte", action="store_true", help="bytes count")
    parser.add_argument("-l", dest="line", action="store_true", help="lines count")
    parser.add_argument("-w", dest="word", action="store_true", help="words count")
    parser.add_argument(
        "-m", dest="character", action="store_true", help="characters count"
    )
    parser.add_argument(
        "file_path",
        type=argparse.FileType("r"),
        help="file path from which to count",
    )

    main(parser.parse_args())
