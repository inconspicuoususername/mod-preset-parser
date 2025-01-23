import { readFileSync } from "fs";

const fn = (htmlFile: string) => {
  const workshopIds: string[] = [];
  const regex = /(?<=id=)(\d+)/g;
  let match: RegExpExecArray | null;
  while ((match = regex.exec(htmlFile)) !== null) {
    workshopIds.push(match[0]);
  }
  return workshopIds;
};

const args = process.argv.slice(2);

if (args.length === 0) {
  console.error("Please provide a path to the HTML file");
  process.exit(1);
}

const htmlFile = readFileSync(args[0], "utf8");

console.log(fn(htmlFile).join(","));
