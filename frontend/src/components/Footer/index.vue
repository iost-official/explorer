<template>
  <div class="footer-box">
    <div class="my-container">
      <div class="my-footer-top">
        <div class="my-footer-left">
          <img src="../../assets/iostFoot.png">
          <h2>IOST</h2>
        </div>
        <ul class="my-footer-middle">
          <li><a href="https://iost.io/tech">Technology</a></li>
          <li><a href="https://iost.io/team">Team</a></li>
          <li><a href="https://iost.io/resource">Resources</a></li>
          <li><a href="https://iost.io/career">Career</a></li>
          <li><a href="https://iost.io/community">Community</a></li>
        </ul>
        <ul class="my-footer-right">
          <li><a href="https://twitter.com/iostoken"><i class="fab fa-twitter"></i> </a></li>
          <li><a href="https://t.me/officialios"><i class="fab fa-telegram-plane"></i></a></li>
          <li><a href="https://www.reddit.com/r/IOStoken"><i class="fab fa-reddit-alien"></i></a></li>
          <li><a href="https://medium.com/@iostoken"><i class="fab fa-medium-m"></i></a></li>
        </ul>
      </div>
      <p>Our mission is to be the underlying architecture for the future of online service providers globally. We are laying the foundation for an ultra-high TPS blockchain infrastructure, one that meets the security and scalability needs of a decentralized economy.</p>
    </div>
  </div>
</template>

<script>
  export default {
    name: "Foot"
  }
</script>

<style lang="less" rel="stylesheet/less">
  .footer-box {
    text-align: left;
    background: #2C2E31;
    overflow: hidden;
    height: 180px;
    .my-container {
      width: 1000px;
      margin: 0 auto;
      .my-footer-top {
        display: flex;
        justify-content: space-between;
        ul {
          margin-bottom: 0;
        }
        .my-footer-left {
          display: flex;
          align-items: center;
          margin-top: 37px;
          img {
            width: 32px;
            height: 37px;
          }
          > h2 {
            margin: 0;
            margin-left: 14px;
            color: #FFFFFF;
          }
        }
        .my-footer-middle {
          margin-top: 48px;
          display: flex;
          > li {
            list-style: none;
            margin-left: 30px;
            &:first-child {
              margin-left: 0;
            }
            a {
              color: #FFFFFF;
              text-decoration: none;
              font-size: 12px;
              line-height: 15px;
            }
          }
        }
        .my-footer-right {
          margin-top: 48px;
          display: flex;
          > li {
            list-style: none;
            margin-left: 20px;
            a {
              color: #FFFFFF;
            }
            &:first-child {
              margin-left: 0;
            }
          }
        }
      }
      > p {
        margin-top: 38px;
        color: #FFFFFF;
        font-size: 12px;
        line-height: 18px;
        font-weight: 300;
        text-align: center;
        padding: 0 64px;
      }
    }
  }

  @media screen and (max-width:480px) {
    .footer-box {
      height: auto;
      .my-container {
        width: 100%;
        .my-footer-top {
          flex-direction: column;
          align-items: center;
          ul {
            margin-bottom: 0;
          }
          .my-footer-left {
            display: flex;
            align-items: center;
            margin-top: 25px;
            img {
              width: 32px;
              height: 37px;
            }
            > h2 {
              margin: 0;
              margin-left: 14px;
              color: #FFFFFF;
            }
          }
          .my-footer-middle {
            margin-top: 25px;
            display: flex;
            > li {
              list-style: none;
              margin-left: 30px;
              &:first-child {
                margin-left: 0;
              }
              a {
                color: #FFFFFF;
                text-decoration: none;
                font-size: 12px;
                line-height: 15px;
              }
            }
          }
          .my-footer-right {
            margin-top: 25px;
            display: flex;
            > li {
              list-style: none;
              margin-left: 20px;
              a {
                color: #FFFFFF;
              }
              &:first-child {
                margin-left: 0;
              }
            }
          }
        }
        > p {
          margin-top: 25px;
          margin-bottom: 20px;
          padding: 0 25px;
        }
      }
    }
  }

  @media screen and (max-width:480px) {
    .footer-box {
      .my-container {
        .my-footer-top {
          .my-footer-middle {
            > li {
              margin-left: 20px;
            }
          }
        }
      }
    }
  }
</style>
