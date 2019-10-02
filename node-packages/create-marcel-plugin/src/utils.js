const { execSync } = require('child_process')

const shouldUseYarn = () => {
  if (process.argv.includes('--use-npm')) return false
  try {
    execSync('yarnpkg --version', { stdio: 'ignore' });
    return true;
  } catch (e) {
    return false;
  }
}

module.exports = {
  shouldUseYarn
}