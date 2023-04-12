import * as path from "https://deno.land/std@0.183.0/path/mod.ts";

export default function main() {
  console.time("Time took in denojs");

  let cmdArgs = "";
  if (Deno.args.length > 0) {
    cmdArgs = Deno.args[0];
  }

  const dir = "dst";

  if (cmdArgs === "clean") {
    const pathDir = `./${dir}`;
    cleanUp(pathDir);
  } else {
    const count = 10000;
    createFile(count, dir);
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
  Deno.mkdirSync(pathDir);

  console.log(`Total deleted: ${sumNBytes} bytes`);
}

function createFile(count: number, dir: string) {
  console.time("Time took in denojs createFile");

  let sumNBytes = 0;
  const src = "./src/test.md";

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
