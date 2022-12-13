import IndexProduc from '../components/products/IndexComponent.vue'
import FormProduct from '../components/products/FormComponent.vue'

const routes = [
  {
    path: '/',
    component: () => import('layouts/LoginLayout.vue'),
    children: [{
      path: '',
      name: 'Login',
      component: () => import('src/pages/LoginPage.vue')
    }]
  },

  {
    path: '/',
    component: () => import('layouts/MainLayout.vue'),
    children: [
      {
        name: 'Index',
        path: '',
        component: () => import('pages/IndexPage.vue')
      },
      {
        name: 'ListProducts',
        path: '/products',
        component: IndexProduc
      },
      {
        name: 'CreateProduct',
        path: '/create-product',
        component: FormProduct
      },
      {
        name: 'UpdateProduct',
        path: '/update-product/:id',
        component: FormProduct
      }
    ]
  },

  // Always leave this as last one,
  // but you can also remove it
  {
    path: '/:catchAll(.*)*',
    component: () => import('pages/ErrorNotFound.vue')
  }
]

export default routes
