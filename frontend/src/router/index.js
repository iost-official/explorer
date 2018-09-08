import Vue from 'vue'
import Router from 'vue-router'
import Head from '../components/Header'
import Foot from '../components/Footer'
import Index from '../pages/Home'
import Blocks from '../pages/Blocks'
import Block from '../pages/BlockDetail'
import Txs from '../pages/Txns'
import Tx from '../pages/TxnsDetail'
import Accounts from '../pages/Accounts'
import AccountDetail from '../pages/AccountDetail'
import Search from '../pages/Search'
import ApplyIOST from '../pages/ApplyIOST'
import ApplyIOSTSuccess from '../pages/ApplyIOSTSuccess'
import LuckyBet from '../pages/LuckyBet'
import NotFound404 from '../pages/404'
import Feedback from '../pages/Feedback'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'Index',
      components: {
        Head: Head,
        Main: Index,
        Foot: Foot
      }
    },
    {
      path: '/blocks',
      name: 'Blocks',
      components: {
        Head: Head,
        Main: Blocks,
        Foot: Foot
      }
    },
    {
      path: '/block/:id',
      name: 'Block',
      components: {
        Head: Head,
        Main: Block,
        Foot: Foot
      }
    },
    {
      path: '/txs',
      name: 'Txs',
      components: {
        Head: Head,
        Main: Txs,
        Foot: Foot
      }
    },
    {
      path: '/tx/:id',
      name: 'Tx',
      components: {
        Head: Head,
        Main: Tx,
        Foot: Foot
      }
    },
    {
      path: '/Accounts',
      name: 'Accounts',
      components: {
        Head: Head,
        Main: Accounts,
        Foot: Foot
      }
    },
    {
      path: '/account/:id',
      name: 'AccountDetail',
      components: {
        Head: Head,
        Main: AccountDetail,
        Foot: Foot
      }
    },
    {
      path: '/search/:id',
      name: 'ExplorerSearch',
      components: {
        Head: Head,
        Main: Search,
        Foot: Foot
      }
    },
    {
      path: '/applyIOST',
      name: 'ApplyIOST',
      components: {
        Head: Head,
        Main: ApplyIOST,
        Foot: Foot
      }
    },
    {
      path: '/applyIOST/success',
      name: 'ApplyIOSTSuccess',
      components: {
        Head: Head,
        Main: ApplyIOSTSuccess,
        Foot: Foot
      }
    },
    {
      path: '/luckyBet',
      name: 'LuckyBet',
      components: {
        Head: Head,
        Main: LuckyBet,
        Foot: Foot
      }
    },
    {
      path: '/feedback',
      name: 'Feedback',
      components: {
        Head: Head,
        Main: Feedback,
        Foot: Foot
      }
    },
    {
      path: '*',
      name: '404',
      components: {
        Head: Head,
        Main: NotFound404,
        Foot: Foot
      }
    }
  ]
})
