import Vue from 'vue'
import Router from 'vue-router'
import Head from '@/components/Head'
import Foot from '@/components/Foot'
import Index from '@/components/Index/Index'
import Blocks from '@/components/Blocks/Blocks'
import Block from '@/components/Blocks/Detail'
import Txs from '@/components/Txs/Txs'
import Tx from '@/components/Txs/Detail'
import Accounts from '@/components/Accounts/Accounts'
import AccountDetail from '@/components/Accounts/Detail'
import Search from '@/components/Search/Search'
import ApplyIOST from '@/components/ApplyIOST/ApplyIOST'
import ApplyIOSTSuccess from '@/components/ApplyIOST/ApplyIOSTSuccess'
import LuckyBet from '@/components/Dapp/LuckyBet/LuckyBet'
import NotFound404 from '@/components/404'
import Feedback from '@/components/Feedback/Feedback'

Vue.use(Router)

export default new Router({
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
      path: '/accounts',
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
