import {SeverityColorValues, DEFAULT_SEVERITY_LEVELS} from 'src/logs/constants'

import {ColorScale} from 'src/types/histogram'

const DEFAULT_COLOR_VALUE = '#7A65F2'

export const colorForSeverity: ColorScale = (
  colorName,
  severityLevel
): string => {
  return (
    SeverityColorValues[colorName] ||
    DEFAULT_SEVERITY_LEVELS[severityLevel] ||
    DEFAULT_COLOR_VALUE
  )
}
