const SET_PRODUCT_MUTATION = (state, obj) => {
  state.product = obj.dados
}

const SET_LIST_PRODUCTS_MUTATION = (state, obj) => {
  state.list_products = obj.lista
}

const SET_PRODUCTS_IN_LIST_MUTATION = (state, obj) => {
  state.list_products.push(obj.dados)
}

export default {
  SET_PRODUCT_MUTATION,
  SET_LIST_PRODUCTS_MUTATION,
  SET_PRODUCTS_IN_LIST_MUTATION
}
