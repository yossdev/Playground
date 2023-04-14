import { copyFileSync, mkdirSync, readdirSync, rmSync, statSync } from "fs";
import path from "path";

function main() {
  console.time("Time took in nodejs");

  let cmdArgs: string = "";
  if (process.argv.length > 1) {
    cmdArgs = process.argv[2];
  }

  const root = "D:/Code/Playground";
  const cwd = "nodejs/writing-file";
  const dir = "dst";

  if (cmdArgs === "clean") {
    const pathDir = `${root}/${cwd}/${dir}`;
    cleanUp(pathDir);
  } else {
    const count: number = 10000;
    const src = `${root}/common/src/test.md`;
    createFile(count, src, dir);
  }

  console.timeEnd("Time took in nodejs");
}

function cleanUp(pathDir: string) {
  const stats = readdirSync(pathDir).map((file) =>
    statSync(path.join(pathDir, file))
  );

  const sumNBytes = stats.reduce((acc, { size }) => acc + size, 0);

  rmSync(pathDir, { recursive: true, force: true });
  mkdirSync(pathDir);

  console.log(`Total deleted: ${sumNBytes} bytes`);
}

function createFile(count: number, src: string, dir: string) {
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
