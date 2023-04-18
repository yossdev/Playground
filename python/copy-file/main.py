import os
import shutil
import sys
import time


def main():
    start = time.time()

    cmd_args = ""
    if len(sys.argv) > 1:
        cmd_args = sys.argv[1]

    root = "D:/Code/Playground"
    cwd = "python/copy-file"
    dir = "dst"

    path_dir = f"{root}/{cwd}/{dir}"
    if not os.path.exists(path_dir):
        os.mkdir(path_dir)

    if cmd_args == "clean":
        clean_up(path_dir)
    else:
        count = 10_000
        src = f"{root}/common/src/test.md"
        copy_file(count, src, dir)

    duration = time.time() - start
    print(f"Time took in python: {duration * 1000:.04f}ms")


def clean_up(path_dir):
    size = 0

    if os.path.exists(path_dir):
        for path, _, files in os.walk(path_dir):
            for f in files:
                fp = os.path.join(path, f)
                size += os.path.getsize(fp)

        shutil.rmtree(path_dir)

    print(f"Total deleted: {size} bytes")


def copy_file(count, src, dir):
    start = time.time()

    sum_n_bytes: int = 0

    for i in range(count):
        dst = f"./{dir}/test{i}.md"

        sum_n_bytes += os.path.getsize(src)

        shutil.copy(src, dst)

    duration = time.time() - start
    print(f"Time took in python createFile: {duration * 1000:.04f}ms")

    print(f"Total written: {sum_n_bytes} bytes")


main()
