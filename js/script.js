function onButtonClick(event) {
  console.log('pogi')
  const input = document.querySelector('#points-input')
  const warning = document.querySelector('#max-points-warning')
  const instruction = document.querySelector('#instruction-text')
  const pointsText = document.querySelector('#points')
  const grades = document.querySelector('#grades')
  const maxPointsText = input.value
  // if jest

  const maxPoints = parseInt(maxPointsText, 10)

  if (maxPoints < 6) {
    warning.style.display = 'block'
    instruction.style.display = 'block'
    grades.style.display = 'none'
    return
  }

  warning.style.display = 'none'
  instruction.style.display = 'none'

  const percentages = [99, 88, 73, 58, 45, 0]
  let prevPoints = null
  let points = []
  for (const percentage of percentages) {
    points.push(computePoints(percentage, maxPoints, prevPoints))
    prevPoints = points[points.length-1]
  }
  
  for (const {max, min} of points) {
    const element = document.createElement('div')
    element.innerText = min + ' - ' + max
    pointsText.appendChild(element)
  }
  grades.style.display = 'flex'
}


function computePoints(minPercentage, maxPoints, prevGrade) {
	if (prevGrade == null) {
		return {
		 max: maxPoints,
		 min: Math.round((minPercentage * maxPoints) / 100)
		}
	}

	return {
		max: prevGrade.min - 1,
		min: Math.round((minPercentage * maxPoints) / 100)
	}
}


window.onload = () => {
  const button = document.querySelector('#submit-button')
  button.addEventListener('click', onButtonClick)
  console.log(button)
}

