#!/usr/bin/env zx
$.shell = '/usr/local/bin/bash'

require('zx/globals');

void async function () {
  console.log(chalk.bold.yellowBright('Advent 2021'))
  let day = parseInt(argv.day, 10);
  let inputFile = argv.input || "input.txt";
  console.log(chalk.blue("Day".padEnd(10, " ")), chalk.bold.redBright(`${day}`))
  console.log(chalk.blue("File".padEnd(10, " ")), chalk.bold.redBright(`${inputFile}`))

  await $`cd day_${day} && \
  cat ${inputFile} | go run .`
}()
