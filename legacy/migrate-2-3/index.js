const fs = require('fs')
const path = require('path')
const process = require('process')

const cwd = process.cwd()

const readJSON = file => JSON.parse(fs.readFileSync(file))

const usersLegacy = readJSON(path.resolve(cwd, 'users.json'))
const pluginsLegacy = readJSON(path.resolve(cwd, 'plugins.json'))
const mediasLegacy = readJSON(path.resolve(cwd, 'medias.json'))

const toDateString = ts => ts ? (new Date(ts * 1000)).toISOString() : null

const all = {
    users: usersLegacy.users.map(user => ({
        ...user,
        createdAt: toDateString(user.createdAt),
        lastDisconnection: toDateString(user.lastDisconnection),
    })),
    plugins: pluginsLegacy.plugins,
    medias: mediasLegacy.medias.map(media => ({
        ...media,
        plugins: media.plugins.map(({ backend, ...plugin }) => plugin)
    })),
}

fs.writeFileSync('all.json', JSON.stringify(all))
