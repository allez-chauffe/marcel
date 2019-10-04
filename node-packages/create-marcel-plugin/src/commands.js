const { shouldUseYarn } = require('./utils')

const installDepsYarn = [
  { command: () => 'yarnpkg', cwd: 'frontend' },
  { command: () => 'yarnpkg add marcel-plugin', cwd: 'frontend' },
  { command: () => 'yarnpkg add -D parcel-bundler', cwd: 'frontend' },
]

const installDepsNpm = [
  { command: () => 'npm i', cwd: 'frontend' },
  { command: () => 'npm add marcel-plugin', cwd: 'frontend' },
  { command: () => 'npm add --save-dev parcel-bundler', cwd: 'frontend' },
]

const commands = [
  ...(shouldUseYarn() ? installDepsYarn : installDepsNpm),
  { command: ({ eltName }) => `if [ -x "$(command -v git)" ] && [ ! "$(git rev-parse --is-inside-work-tree 2>/dev/null)" ]; then git init && git add . && git commit -m ":tada: Initialize ${eltName}"; fi` },
]

module.exports = commands
