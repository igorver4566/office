import axios from 'axios'

export function WeatherLocation () {
  return new Promise((resolve, reject) => {
    navigator.geolocation.getCurrentPosition((position) => {
      axios.get('https://query.yahooapis.com/v1/public/yql?q=select * from weather.forecast where woeid in (SELECT woeid FROM geo.places WHERE text="(' + position.coords.latitude + ',' + position.coords.longitude + ')") and u="c"&format=json')
      .then(res => {
        return resolve(res.data.query.results.channel)
      })
      .catch(err => {
        return reject(err)
      })
    })
  })
}
