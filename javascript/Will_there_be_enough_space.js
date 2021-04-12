const enough = (cap, on, wait) => {
  const less = cap - on
  const difference = wait - less
  return difference < 0 
    ? 0
    : difference
}