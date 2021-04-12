function updateLight(current) {
  return  {
    green: 'yellow',
    yellow: 'red',
    red: 'green'
  }[current]
}