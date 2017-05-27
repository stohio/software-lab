<template>

  <div id="bundle-container">
    <h1>I want to build...</h1>
    <ul v-bind:class="{active: (openTab != '')}">
      <li 
        v-on:click="toggleActive('mobile')"
        v-bind:class="{active: (openTab === 'mobile'), hidden: (openTab != '' && openTab != 'mobile')}"
        class="bundle-card left rp-gradient">
        <div class="block"></div>
        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
        width="32px" height="32px" viewBox="0 0 32 32" enable-background="new 0 0 32 32" xml:space="preserve">
        <defs>
          <linearGradient id="rp-gradient" gradientTransform="rotate(60)">
            <stop offset="0%" stop-color="#d11d29"></stop>
            <stop offset="100%" stop-color="#542437"></stop>
          </linearGradient>
        </defs>
        <g>
          <path d="M1.5,32h14c0.276,0,0.5-0.224,0.5-0.5S15.776,31,15.5,31h-14C1.224,31,1,30.776,1,30.5v-29
          C1,1.224,1.224,1,1.5,1h21C22.776,1,23,1.224,23,1.5v5C23,6.776,23.224,7,23.5,7S24,6.776,24,6.5v-5C24,0.673,23.327,0,22.5,0h-21
          C0.673,0,0,0.673,0,1.5v29C0,31.327,0.673,32,1.5,32z"/>
          <path  d="M18,10.5v20c0,0.827,0.673,1.5,1.5,1.5h11c0.827,0,1.5-0.673,1.5-1.5v-20C32,9.673,31.327,9,30.5,9h-11
          C18.673,9,18,9.673,18,10.5z M31,10.5v20c0,0.276-0.224,0.5-0.5,0.5h-11c-0.276,0-0.5-0.224-0.5-0.5v-20c0-0.276,0.224-0.5,0.5-0.5
          h11C30.776,10,31,10.224,31,10.5z"/>
          <circle cx="12" cy="28" r="1"/>
          <circle cx="25" cy="28" r="1"/>
        </g>
        </svg>
        <h2>Mobile Apps</h2>
      </li>

      <li 
      v-on:click="toggleActive('java')"
        v-bind:class="{active: (openTab === 'java'), hidden: (openTab != '' && openTab != 'java')}"
      class="bundle-card right oy-gradient">
        <div class="block"></div>
        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" width="32.02199935913086" height="32" viewBox="0 0 32.02199935913086 32" enable-background="new 0 0 16 16" xml:space="preserve" fill="#000000">
        <defs>
          <linearGradient id="oy-gradient" gradientTransform="rotate(60)">
            <stop offset="0%" stop-color="#e16622"></stop>
            <stop offset="100%" stop-color="#f6c435"></stop>
          </linearGradient>
        </defs>
        <g><path d="M 25.89,10.004c 0.056-0.836, 0.090-1.514, 0.108-1.92c 0.022-0.544-0.178-1.074-0.554-1.468S 24.544,6, 24,6L 2,6 C 1.456,6, 0.934,6.222, 0.556,6.616S-0.020,7.54, 0.002,8.084c 0.11,2.642, 0.854,15.996, 4.664,19.406C 5.032,27.818, 5.508,28, 6,28l 14,0 c 0.492,0, 0.968-0.182, 1.334-0.51c 1.194-1.070, 2.088-3.116, 2.754-5.504c 0.99-0.004, 1.674-0.006, 1.754-0.006 c 3.402,0, 6.16-2.48, 6.16-5.99C 32,11.224, 27.88,10.022, 25.89,10.004z M 20,26L 6,26 c-3.374-3.020-4-18-4-18l 22,0 C 24,8, 23.374,22.98, 20,26z M 25.84,19.98l-0.63,0.002c-0.18,0-0.394,0.002-0.636,0.002c 0.596-2.738, 0.954-5.66, 1.16-7.984l 0.106,0 C 26.010,12, 30,12.046, 30,15.99 C 30,18.302, 28.25,19.98, 25.84,19.98z"></path></g></svg>
        <h2>Java Apps</h2>
      </li>

      <li 
      v-on:click="toggleActive('games')"
        v-bind:class="{ active: (openTab === 'games'), hidden: (openTab != '' && openTab != 'games')}"
        class="bundle-card left by-gradient">
        <div class="block"></div>
        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
        width="34px" height="32px" viewBox="0 0 34 32" enable-background="new 0 0 34 32" xml:space="preserve">
        <defs>
          <linearGradient id="by-gradient" gradientTransform="rotate(60)">
            <stop offset="0%" stop-color="#3a5255"></stop>
            <stop offset="100%" stop-color="#f6c435"></stop>
          </linearGradient>
        </defs>
        <g>
          <path d="M25.5,1C25.776,1,26,0.776,26,0.5S25.776,0,25.5,0h-17C8.224,0,8,0.224,8,0.5S8.224,1,8.5,1H25.5z"/>
          <path  d="M25.5,2h-17h-2C3.217,2,1.687,3.166,0.981,4.144c-0.864,1.2-1.026,2.803-0.456,4.514
          c0.882,2.646,3.846,5.825,5.668,7.237C6.285,15.966,6.393,16,6.5,16c0.149,0,0.297-0.066,0.396-0.194
          c0.169-0.218,0.129-0.532-0.089-0.702c-1.674-1.296-4.525-4.342-5.332-6.763C1.008,6.944,1.121,5.661,1.793,4.729
          C2.608,3.598,4.235,3,6.5,3H8v9.5c0,6.726,6.673,10.601,8.036,11.322C16.015,23.878,16,23.937,16,24v2.5c0,0.276,0.224,0.5,0.5,0.5
          s0.5-0.224,0.5-0.5V24c0-0.058-0.015-0.112-0.033-0.164C18.481,23.124,26,19.25,26,12.5V3h1.5c2.265,0,3.892,0.598,4.707,1.729
          c0.672,0.932,0.785,2.215,0.318,3.613c-0.63,1.891-3.769,5.553-5.332,6.763c-0.218,0.169-0.258,0.483-0.089,0.702
          C27.203,15.934,27.351,16,27.5,16c0.107,0,0.215-0.034,0.307-0.104c1.732-1.342,4.961-5.119,5.668-7.237
          c0.57-1.711,0.408-3.314-0.456-4.514C32.313,3.166,30.783,2,27.5,2H25.5z M25,12.5c0,6.093-7.152,9.804-8.486,10.443
          C15.286,22.278,9,18.576,9,12.5V3h16V12.5z"/>
          <path  d="M11,28c-1.141,0-2,0.86-2,2v1.5C9,31.776,9.224,32,9.5,32h14c0.276,0,0.5-0.224,0.5-0.5v-1.425
          C24,28.912,23.121,28,22,28H11z M23,30.075V31H10v-1c0-0.589,0.411-1,1-1h11C22.57,29,23,29.462,23,30.075z"/>
        </g>
      </svg>
        <h2>Games</h2>
      </li>

      <li v-on:click="toggleActive('os')"
        v-bind:class="{active: (openTab === 'os'), hidden: (openTab != '' && openTab != 'os')}"
        class="bundle-card right pb-gradient">
        <div class="block"></div>
        <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px"
        width="39px" height="32px" viewBox="0 0 39 32" enable-background="new 0 0 39 32" xml:space="preserve">
        <defs>
          <linearGradient id="pb-gradient" gradientTransform="rotate(60)">
            <stop offset="0%" stop-color="#542437"></stop>
            <stop offset="100%" stop-color="#3a5255"></stop>
          </linearGradient>
        </defs>
        <g>
          <path  d="M1.5,25H19v1.5c0,0.276,0.224,0.5,0.5,0.5s0.5-0.224,0.5-0.5V25h17.5c0.827,0,1.5-0.673,1.5-1.5v-22
          C39,0.673,38.327,0,37.5,0h-36C0.673,0,0,0.673,0,1.5v22C0,24.327,0.673,25,1.5,25z M1,1.5C1,1.224,1.224,1,1.5,1h36
          C37.776,1,38,1.224,38,1.5v22c0,0.276-0.224,0.5-0.5,0.5h-36C1.224,24,1,23.776,1,23.5V1.5z"/>
          <path  d="M3.5,22h32c0.276,0,0.5-0.224,0.5-0.5v-18C36,3.224,35.776,3,35.5,3h-32C3.224,3,3,3.224,3,3.5v18
          C3,21.776,3.224,22,3.5,22z M4,4h31v17H4V4z"/>
          <path  d="M14,28c-1.14,0-2,0.86-2,2v1.5c0,0.276,0.224,0.5,0.5,0.5h14c0.276,0,0.5-0.224,0.5-0.5v-1.425
          C27,28.912,26.122,28,25,28H14z M26,30.075V31H13v-1c0-0.589,0.411-1,1-1h11C25.57,29,26,29.462,26,30.075z"/>
        </g>
        </svg>
        <h2>Operating Systems</h2>
      </li>
    </ul>
  </div>

</template>

<script>
export default {
  name: 'bundle-list',
  data () {
    return {
      msg: 'The Software Lab',
    }
  },

  props: ['openTab'],

  mounted() {
    
  },

  methods: {
    toggleActive: function(name) {

      if (this.openTab === name) {
        var el = document.querySelector('.active')
        console.log(el);
        el.classList.remove('active');
        setTimeout(function () { this.openTab = '' }.bind(this), 600)
        this.$emit('changeBundle', '');

      } else {
        this.openTab = name;
        this.$emit('changeBundle', name);
      }
    }
  }

}
</script>

<style lang="scss">
#bundle-container {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #ddd5bf;

  width: 100%;
  max-width: 940px;

  margin: 0 auto;
  margin-top: 40px;

  h1 {
    text-align: left;
  }

  ul {
    margin-top: 40px;
    display: flex;
    flex-wrap: wrap;
    justify-content: row-end;
    max-height: auto;
    height: auto;
  }

  ul.active {
    transition: 0.5s height ease;
    max-height: 280px;
    height: 280px;
  }
}

// a {
//   color: #42b983;
// }

.bundle-card.see-all {
  width: 100%;
  min-height: 150px;
  max-height: 150px;

  display: flex;

  .block {
    transform: rotate(60deg);
    height: 230%;
  }

  h2 {
    text-align: right;
    top: 40%;
  }

  svg {
    margin-left: 20px;
    height: 180px;
    transform: translate(-200%, -30px);
    transition: 0.5s transform ease;
  }

  &:hover {
    svg {
      transform: translate(0, -30px);
    }
  }
}

.bundle-card.active {
  width: 100%;

  //transition: 1s width ease 4s;
  flex-grow: 1;
  transition: 0.5s flex-grow ease 0.5s, 0.5s width ease 0.5s;
  margin: 0%;
  margin-bottom: 1%;
}

.bundle-card.hidden {
  animation: 0.5s hide-card linear;
  animation-fill-mode: both;

  flex-shrink: 1;

  max-height: 0;
  height: 0;

  transition: 0.5s max-height ease 0.5s, 0.5s height ease 0.5s, 0.5s min-height ease 0.5s, max-width ease;
}

@keyframes hide-card {
  0% {
    opacity: 1;
    max-height: 100%;
  }

  50% {
    opacity: 0;
    min-height: 0;
    max-height: 0px;
    padding: 0;
    margin: 0;
    // min-height: 0;
    // max-height: 0;
    // padding: 0;
    // margin: 0;
    // max-width: 1px;
    // width: 1px;
    // min-width: 1px;
  }

  100% {
    min-width: 0;
    max-width: 0px;
    padding: 0;
    margin: 0;
    // min-height: 0;
    // max-height: 0;
    // padding: 0;
    // margin: 0;
    // max-width: 1px;
    // width: 1px;
    // min-width: 1px;
  }
}

.bundle-card {
  position: relative;
  width: 49%;
  min-height: 280px;
  background-color: rgba(10, 13, 16, 0.7);
  color: #ddd5bf;

  overflow: hidden;

  cursor: pointer;

  text-align: left;

  padding: 20px 40px;
  margin-bottom: 2%;

  border-radius: 5px;
  box-shadow: 0 1px 4px 0 rgba(0,0,0,0.16);

  transition: 0.5s max-height ease 0.5s, 0.5s height ease 0.5s, 0.5s min-height ease 0.5s, max-width ease 0.5s;

    &.left{
        margin-right: 0%;
    }
    &.right{
        margin-left: 0%;
    }
  &:hover, &.active {
    
    h2 {
      transform: translateY(-3px);
    }

    .block {
      top: 0%;
      left: 50%;
      
      &:after {
        left: -10%;
      }
    }
  }

  &.rp-gradient {

    svg {
      fill: url(#rp-gradient);
    }

    #rp-gradient {
      stop {
        stop-color: #000;
        transition: 0.3s stop-color ease;
      }

      stop:last-of-type {
        transition-delay: 0.1s;
      }
    }

    &:hover, &.active {
      #rp-gradient {
        stop:first-of-type {
          stop-color: #d11d29;
        }
        stop:last-of-type {
          stop-color: #542437;
        }
      }
    }
  }

  &.oy-gradient {

    svg {
      fill: url(#oy-gradient);
    }

    #oy-gradient {
      stop {
        stop-color: #000;
        transition: 0.3s stop-color ease;
      }

      stop:last-of-type {
        transition-delay: 0.1s;
      }
    }

    &:hover, &.active {
      #oy-gradient {
        stop:first-of-type {
          stop-color: #e16622;
        }
        stop:last-of-type {
          stop-color: #f6c435;
        }
      }
    }
  }

  &.pb-gradient {

    svg {
      fill: url(#pb-gradient);
    }

    #pb-gradient {
      stop {
        stop-color: #000;
        transition: 0.3s stop-color ease;
      }

      stop:last-of-type {
        transition-delay: 0.1s;
      }
    }

    &:hover, &.active {
      #pb-gradient {
        stop:first-of-type {
          stop-color: #542437;
        }
        stop:last-of-type {
          stop-color: #4B75BD;
        }
      }
    }
  }

  &.by-gradient {

    svg {
      fill: url(#by-gradient);
    }

    #by-gradient {
      stop {
        stop-color: #000;
        transition: 0.3s stop-color ease;
      }

      stop:last-of-type {
        transition-delay: 0.1s;
      }
    }

    &:hover, &.active {
      #by-gradient {
        stop:first-of-type {
          stop-color: #4B75BD;
        }
        stop:last-of-type {
          stop-color: #f6c435;
        }
      }
    }
  }

  h2 {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    font-size: 2rem;
    position: absolute;
    bottom: 20px;
    right: 30px;

    max-width: 50%;
    text-align: right;

    transition: 0.3s transform ease;
  }

  .block {
    position: absolute;
    top: 50%;
    left: 110%;
    width: 80%;
    height: 200%;
    background: rgba(0,0,0, 0.8);
    transition: 0.4s all ease;

    transform: rotate(30deg);

    overflow: visible;

    &:after {
      content: "";
      position: absolute;
      top: 0%;
      left: 0%;
      width: 110%;
      height: 100%;

      transition: 0.6s all ease;
      background: rgba(0,0,0, 0.6);
    }
  }

  svg {
    width: auto;
    height: 100%;
    padding: 10px 0;
    margin: 0 auto;
  }
}
</style>
