<template>
  <div class="q-pa-sm row q-gutter-sm justify-center">

    <div class="col-md-6">
      <q-card flat bordered>

        <q-card-section>
          <div v-if="local_product.id" class="text-h4">Editar Produto</div>
          <div v-else class="text-h4">Cadastrar Produto</div>
        </q-card-section>

        <q-separator />

        <q-card-section class="q-gutter-sm">

          <q-input
            filled
            v-model="local_product.name"
            label="Nome do produto"
          />

          <q-input
            filled
            v-model="local_product.code"
            label="Código"
          />

          <q-input
            filled
            type="number"
            v-model="local_product.price"
            label="Preço"
          />

        </q-card-section>

        <q-separator />

        <q-card-actions align="right">
          <q-btn flat icon="reply" color="info" :to="{name: 'ListProducts'}" label="Voltar" />
          <q-btn v-if="!local_product.id" flat icon="cleaning_services" color="warning" @click="resetProduct()" label="Limpar" />
          <q-btn v-if="local_product.id" flat icon="delete_forever" color="negative" @click="deleteProduct(item)" label="Remover" />
          <q-btn flat icon="send" color="positive" @click="saveProduct()" label="Salvar" />
        </q-card-actions>

      </q-card>
    </div>

  </div>
</template>

<script>
import {
  mapState,
  mapActions
} from 'vuex'

export default {

  data () {
    return {
      local_product: {
        id: null,
        name: '',
        code: '',
        price: 0
      }
    }
  },

  computed: {
    ...mapState('Products', [
      'product'
    ])
  },

  created () {
    if (this.$route.params.id) {
      (async () => {
        console.log(this.$route.params.id)
        await this.setProductAction({ id: this.$route.params.id })
        this.local_product = JSON.parse(JSON.stringify(this.product))
      })()
    } else {
      this.iniciarProductAction()
      this.local_product = JSON.parse(JSON.stringify(this.product))
    }
  },

  methods: {

    ...mapActions('Products', [
      'iniciarProductAction',
      'setProductAction',
      'setProductsInListAction',
      'deleteProductAction'
    ]),

    async deleteProduct () {
      if (confirm('Deseja realmente deletar este produto?') === true) {
        await this.deleteProductAction()
        this.$router.push({ name: 'ListProducts' })
      }
    },

    async saveProduct () {
      const tmpProduct = JSON.parse(JSON.stringify(this.local_product))
      tmpProduct.price = parseFloat(tmpProduct.price)
      await this.setProductsInListAction(tmpProduct)
      this.$router.push({ name: 'ListProducts' })
    },

    resetProduct () {
      this.iniciarProductAction()
      this.local_product = JSON.parse(JSON.stringify(this.product))
    }

  }
}
</script>
