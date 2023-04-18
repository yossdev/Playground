import {
  copyFileSync,
  existsSync,
  mkdirSync,
  readdirSync,
  rmSync,
  statSync,
} from "fs";
import path from "path";

function main() {
  console.time("Time took in nodejs");

  let cmdArgs: string = "";
  if (process.argv.length > 1) {
    cmdArgs = process.argv[2];
  }

  const root = "D:/Code/Playground";
  const cwd = "nodejs/copy-file";
  const dir = "dst";

  const pathDir = `${root}/${cwd}/${dir}`;
  if (!existsSync(pathDir)) {
    mkdirSync(pathDir);
  }

  if (cmdArgs === "clean") {
    cleanUp(pathDir);
  } else {
    const count: number = 10_000;
    const src = `${root}/common/src/test.md`;
    copyFile(count, src, dir);
  }

  console.timeEnd("Time took in nodejs");
}

function cleanUp(pathDir: string) {
  const stats = readdirSync(pathDir).map((file) =>
    statSync(path.join(pathDir, file))
  );

  const sumNBytes = stats.reduce((acc, { size }) => acc + size, 0);

  rmSync(pathDir, { recursive: true, force: true });

  console.log(`Total deleted: ${sumNBytes} bytes`);
}

function copyFile(count: number, src: string, dir: string) {
  console.time("Time took in nodejs createFile");

  let sumNBytes: number = 0;

  for (let i = 0; i < count; i++) {
    const dst = `./${dir}/test${i}.md`;
    copyFileSync(src, dst);
    sumNBytes += statSync(dst).size;
  }

  console.timeEnd("Time took in nodejs createFile");
  console.log(`Total written: ${sumNBytes} bytes`);
}

main();
