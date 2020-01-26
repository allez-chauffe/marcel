export const withoutNil = obj => {
  if (!obj || typeof obj !== 'object') return obj
  const cleanObj = {}

  Object.keys(obj).forEach(key => {
    const value = obj[key]
    if (value !== null && value !== undefined) cleanObj[key] = value
  })

  return cleanObj
}
