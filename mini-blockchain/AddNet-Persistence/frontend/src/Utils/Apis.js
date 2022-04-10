//get port where react app is running

console.log(process.env.PORT);

export const apiBaseUrl = "http://localhost:" + process.env.REACT_APP_API_PORT;
