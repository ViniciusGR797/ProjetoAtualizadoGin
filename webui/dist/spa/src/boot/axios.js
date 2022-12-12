import { boot } from 'quasar/wrappers'
import axios from 'axios'

// Default API Base Path
const apiBasePath = '/' // Prod

// const apiBasePath = 'http://localhost:8081/' // Dev

// Be careful when using SSR for cross-request state pollution
// due to creating a Singleton instance here;
// If any client changes this (global) instance, it might be a
// good idea to move this instance creation inside of the
// "export default () => {}" function below (which runs individually
// for each client)
const api = axios.create({ baseURL: apiBasePath })

/**
 * Devolte os Headers HTTP padrão e customizados
 * @param {*} headers objeto ex: {'Authorization': 'Basic qwerqwerqwer'}
 */
const setHeaders = (headers = {}) => {
  let localHeaders = {}

  if (window.localStorage.autenticated) {
    localHeaders.Authorization = 'Bearer ' + JSON.parse(window.localStorage.getItem('autenticated')).token
  }

  if (headers) {
    localHeaders = { ...localHeaders, ...headers }
  }

  return localHeaders
}

/**
 * Metodos HTTP para interações Ajax
 */

const getData = (url, headers) => {
  return api.get(url, { headers: setHeaders(headers) })
}

const postData = (url, data, headers) => {
  return api.post(url, data, { headers: setHeaders(headers) })
}

const putData = (url, data, headers) => {
  return api.put(url, data, { headers: setHeaders(headers) })
}

const deleteData = (url, headers) => {
  return api.delete(url, { headers: setHeaders(headers) })
}

export default boot(({ app }) => {
  // for use inside Vue files (Options API) through this.$axios and this.$api

  app.config.globalProperties.$axios = axios
  // ^ ^ ^ this will allow you to use this.$axios (for Vue Options API form)
  //       so you won't necessarily have to import axios in each vue file

  app.config.globalProperties.$api = api
  // ^ ^ ^ this will allow you to use this.$api (for Vue Options API form)
  //       so you can easily perform requests against your app's API
})

export {
  api,
  getData,
  postData,
  putData,
  deleteData
}
