import os
import shutil
import sys
import time


def main():
    start = time.time()

    cmd = ""
    if len(sys.argv) > 1:
        cmd = sys.argv[1]

    dir = "dst"

    if cmd == "clean":
        pathDir = f"./{dir}"
        cleanUp(pathDir)
    else:
        count = 10000
        createFile(count, dir)

    duration = time.time() - start
    print(
        f"Time took in python: {duration * 1000:.04f}ms")


def cleanUp(pathDir):
    size = 0

    if os.path.exists(pathDir):
        for path, _, files in os.walk(pathDir):
            for f in files:
                fp = os.path.join(path, f)
                size += os.path.getsize(fp)

        shutil.rmtree(pathDir)

    os.mkdir(pathDir)
    print(f"Total deleted: {size} bytes")


def createFile(count, dir):
    start = time.time()

    sumNBytes: int = 0
    src = "./src/test.md"

    for i in range(count):
        dst = f"./{dir}/test{i}.md"

        sumNBytes += os.path.getsize(src)

        shutil.copy(src, dst)

    duration = time.time() - start
    print(
        f"Time took in python createFile: {duration * 1000:.04f}ms")

    print(f"Total written: {sumNBytes} bytes")


main()
