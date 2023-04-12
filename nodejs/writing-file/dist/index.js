import { copyFileSync, mkdirSync, readdirSync, rmSync, statSync } from "fs";
import path from "path";
function main() {
    console.time("Time took in nodejs");
    let cmd = "";
    if (process.argv.length > 1) {
        cmd = process.argv[2];
    }
    const dir = "dst";
    if (cmd === "clean") {
        const pathDir = `./${dir}`;
        cleanUp(pathDir);
    }
    else {
        const count = 10000;
        createFile(count, dir);
    }
    console.timeEnd("Time took in nodejs");
}
function cleanUp(pathDir) {
    const stats = readdirSync(pathDir).map((file) => statSync(path.join(pathDir, file)));
    const sumNBytes = stats.reduce((acc, { size }) => acc + size, 0);
    rmSync(pathDir, { recursive: true, force: true });
    mkdirSync(pathDir);
    console.log(`Total deleted: ${sumNBytes} bytes`);
}
function createFile(count, dir) {
    console.time("Time took in nodejs createFile");
    let sumNBytes = 0;
    const src = "./src/test.md";
    for (let i = 0; i < count; i++) {
        const dst = `./${dir}/test${i}.md`;
        copyFileSync(src, dst);
        sumNBytes += statSync(dst).size;
    }
    console.timeEnd("Time took in nodejs createFile");
    console.log(`Total written: ${sumNBytes} bytes`);
}
main();
//# sourceMappingURL=index.js.map