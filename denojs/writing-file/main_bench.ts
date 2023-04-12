import main from "./main.ts";

Deno.bench("Copy File", () => {
  main();
});
