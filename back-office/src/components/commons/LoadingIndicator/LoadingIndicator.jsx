//@flow
import React from 'react'
import ProgressBar from 'react-toolbox/lib/progress_bar/ProgressBar'

const LoadingIndicator = () => (
  <ProgressBar type="circular" mode="indeterminate" className="loader" />
)

export default LoadingIndicator
