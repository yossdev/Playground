import { existsSync } from "https://deno.land/std@0.182.0/fs/exists.ts";
import * as path from "https://deno.land/std@0.183.0/path/mod.ts";

export default function main() {
  console.time("Time took in denojs");

  let cmdArgs = "";
  if (Deno.args.length > 0) {
    cmdArgs = Deno.args[0];
  }

  const root = "D:/Code/Playground";
  const cwd = "denojs/copy-file";
  const dir = "dst";

  const pathDir = `${root}/${cwd}/${dir}`;
  if (!existsSync(pathDir)) {
    Deno.mkdirSync(pathDir);
  }

  if (cmdArgs === "clean") {
    cleanUp(pathDir);
  } else {
    const count = 10_000;
    const src = `${root}/common/src/test.md`;
    copyFile(count, src, dir);
  }

  console.timeEnd("Time took in denojs");
}

function cleanUp(pathDir: string) {
  const files = Deno.readDirSync(pathDir);
  const stats = [];
  for (const file of files) {
    stats.push(Deno.statSync(path.join(pathDir, file.name)));
  }

  const sumNBytes = stats.reduce((acc, { size }) => acc + size, 0);

  Deno.removeSync(pathDir, { recursive: true });

  console.log(`Total deleted: ${sumNBytes} bytes`);
}

function copyFile(count: number, src: string, dir: string) {
  console.time("Time took in denojs createFile");

  let sumNBytes = 0;

  for (let i = 0; i < count; i++) {
    const dst = `./${dir}/test${i}.md`;
    Deno.copyFileSync(src, dst);
    sumNBytes += Deno.statSync(dst).size;
  }

  console.timeEnd("Time took in denojs createFile");
  console.log(`Total written: ${sumNBytes} bytes`);
}

// Learn more at https://deno.land/manual/examples/module_metadata#concepts
if (import.meta.main) {
  main();
}
