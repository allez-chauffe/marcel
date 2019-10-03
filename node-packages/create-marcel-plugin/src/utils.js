const path = require('path')
const { execSync } = require('child_process')
const toCamelCase = require('to-camel-case')

const shouldUseYarn = () => {
  if (process.argv.includes('--use-npm')) return false
  try {
    execSync('yarnpkg --version', { stdio: 'ignore' });
    return true;
  } catch (e) {
    return false;
  }
}

const getPluginInfo = eltName => {
  const [firstChar, ...endOfName] = toCamelCase(eltName)
  const name = [firstChar.toUpperCase(), ...endOfName].join('')

  return {
    eltName,
    name,
    path: path.resolve(eltName),
  }
}

const fatal = (...message) => {
  console.error(...message)
  process.exit(-1)
}

const toJSON = value => {
  const json = JSON.stringify(value, null, 2)
  return json.endsWith('\n') ? json : (json + '\n')
}

module.exports = {
  shouldUseYarn,
  getPluginInfo,
  fatal,
  toJSON
}
