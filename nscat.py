"""
nscat: displays contents of a file substituting passwords and api keys with Xs
"""
import click

secret_search = ["pw", "password", "passwd", "dbpass", "_key", "_key_", "credential"]


def check_line(line: str):
    for item in secret_search:
        if item in line.lower():
            data = line.strip().split("=")
            return f"{data[0]}={'X' * len(data[1])}"
    return line.strip()


@click.command()
@click.argument("filename", type=click.Path(exists=True))
def nscat(filename):
    with open(filename) as fn:
        data = fn.readlines()

    for item in data:
        print(check_line(item))


if __name__ == "__main__":
    nscat()
