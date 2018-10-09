import axios from 'axios'

export const Settings = {
  host: 'http://localhost:8080/'
}

export function Registration (name, pass, mail) {
  return new Promise((resolve, reject) => {
    axios.post(Settings.host + 'app/register', {
      login: name,
      password: pass,
      email: mail
    })
    .then(function (response) {
      if (response.data.ok === 'true') {
        Authentication(response.data.token)
      } else {
        return resolve(response.data.token)
      }
    })
    .catch(err => {
      return reject(err)
    })
  })
}

function Authentication (token) {
  console.log(token)
}
