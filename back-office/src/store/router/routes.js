//@flow

export default {
  '/': {
    title: 'HOME',

    '/medias': {
      title: 'MEDIAS',
      '/:mediaID': { title: 'MEDIA_DETAILS' },
    },
  },
}
