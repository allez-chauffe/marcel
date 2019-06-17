export default {
  '/': {
    title: 'HOME',

    '/medias': {
      title: 'MEDIAS',
      '/:mediaID': { title: 'MEDIA_DETAILS' },
    },
    '/users': {
      title: 'USERS',
    },
    '/profil': {
      title: 'PROFIL',
    },
  },
}
